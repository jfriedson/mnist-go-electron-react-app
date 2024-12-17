## Go service
Listens for an HTTP Post request at root "/" containing a flattened 28x28 single channel image and responds with the result of running the input through a neural network trained on MNIST for handwritten digit recognition from the [mnist-model generator module](../mnist-model-generator).

### Features
- Dynamic typing by means of reflection and type assertion in the neuralnet modules allows for easy to read and performant code that is simple to call using an interface method
- Modules accept pointers to their input so as to allow inplace operations to drastically improve memory use and minimize garbage to be collected

### Analysis of multiple Goroutine strategies
The following benchmarks were generated from the [modules](./neuralnet/module/) and [archived modules](./neuralnet/module/archive/) using tensor sizes in the MNIST model. The testing system has a 12 core AMD 7900x (64 MB L3 cache :shipit: ) and 64GB of RAM, so the minor scheduling and memory overhead that each goroutine needs is minimal on this system for these workloads. The MNIST model layers are very small in comparison to architectures of larger computer vision tasks - I would like to extend these tests to include much larger tensors.

Much work went into optimizing memory access and usage without sacrificing computational performance. Stuff like using accumulation vars and limiting the use of indirection for slices and the dereferencing of pointers were able to shave off noticeable chunks of time.

Additionally, the goroutine strategies will probably perform much better using a batched approach as opposed to the current goroutine-per-output method. The strategies with workerpools should mitigate the effect of the threading overhead; however, the channeling seems to add a decent amount of overhead itself as seen in the conv2d results.

#### Explanation of strategies
Sequential - no goroutines used\
Goroutine - one goroutine spawned per output element\
Workerpool - #cpus goroutine workers are spawned and assigned jobs by channel\
Static Workerpool - same as Workerpool strategy, but goroutines are kept alive for the duration of the service

The worker count of #cpus was reached through testing. Using 2*#cpus led to a drastic decrease in performance. Using an unbuffered job channel, a buffer size of #workers, a buffer size of 2*#workers, and a buffer size of #jobs/2 also led to abhorrent performance, so the channel buffer was set to match the total number of jobs (outputs) for a single pass of the layer; this is probably the optimal value due to the main routine being able to assign all jobs at once and allowing the workers to run uninterrupted.

#### Linear / Fully Connected
##### Input Dim (256) -> Output Dim (10)
| Strategy          | Iterations | Timing | Mem Used | Allocs |
| ----------------- | ---------: | -----: | -------: | -----: |
| Sequential        |    668530 | 1796 ns |     96 B |      3 |
| Goroutine         |    266930 | 4285 ns |   1256 B |     15 |
| Workerpool        |    167378 | 6798 ns |   3021 B |     30 |
| Static Workerpool |    281953 | 4078 ns |    125 B |      4 |

##### Input Dim (1024) -> Output Dim (256)
| Strategy          | Iterations | Timing | Mem Used | Allocs |
| ----------------- | ---------: | -----: | -------: | -----: |
| Sequential        |    8168 | 151364 ns |   1072 B |      3 |
| Goroutine         |   17389 |  69158 ns |  29786 B |    261 |
| Workerpool        |   15381 |  74663 ns |   6151 B |     30 |
| Static Workerpool |   13704 |  86828 ns |   1141 B |      4 |


#### Conv2d
##### Input Dim (1, 28, 28) -> Kernel Dim (5, 5) -> Output Dim (24, 24, 24)
| Strategy          | Iterations | Timing | Mem Used | Allocs |
| ----------------- | ---------: | -----: | -------: | -----: |
| Sequential        |   1815 |  629133 ns |  71349 B |    603 |
| Goroutine         |    541 | 2238759 ns | 1841204 B | 14429 |
| Workerpool        |    249 | 4683951 ns | 410733 B |    635 |
| Static Workerpool |    348 | 3276528 ns |  71508 B |    604 |

##### Input Dim (24, 24, 24) -> Kernel Dim (3, 3) -> Output Dim (32, 22, 22)
| Strategy          | Iterations | Timing | Mem Used | Allocs |
| ----------------- | ---------: | -----: | -------: | -----: |
| Sequential        |    196 | 5979858 ns |  86986 B |    739 |
| Goroutine         |    292 | 4406027 ns | 2074652 B | 16239 |
| Workerpool        |    168 | 7081673 ns | 467657 B |    773 |
| Static Workerpool |    186 | 6270808 ns |  87057 B |    740 |


#### Results Analysis
For smaller workloads, the strategies using goroutines saw a detrimental impact to timing performance and memory usage when compared to the sequential implementations in both types of layers.

The larger layers of the Conv2d did not see a large improvement in timing performance, either; the limitation likely being the size of the kernel. This means conv2d layers using larger kernels will be able to better take advantage of the concurrency. The fully connected linear layer was really able to utilize the extra compute for the larger vectors.
