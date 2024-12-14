## Go service
Listens for an HTTP Post request at root "/" containing a flattened 28x28 single channel image and responds with the result of running the input through a neural network trained on MNIST for handwritten digit recognition from the [mnist-model generator module](../mnist-model-generator).

### Features
- Dynamic typing by means of reflection and type assertion in the neuralnet modules allows for easy to read and performant code that is simple to call using an interface method
- Modules accept pointers to their input so as to allow inplace operations to drastically improve memory use and minimize garbage to be collect
