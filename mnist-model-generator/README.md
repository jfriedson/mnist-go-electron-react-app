## MNIST model generator
Need an MNIST model so why not create one. Code and neural network architecture adapted from https://github.com/tuomaso/train_mnist_fast/blob/master/8_Final_00s76.ipynb

### Architecture
Input is first operated on by two 2d convolutional layers each followed by max pooling downsampling and ReLU activation. The result is then flattened and passed through two fully connected linear layers padded by a ReLU activation.

### Setup 
conda create -n mnist
conda activate mnist
conda install pytorch torchvision pytorch-cuda=12.4 -c pytorch -c nvidia
conda install --file requirements.txt
