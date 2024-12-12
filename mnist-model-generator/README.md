## MNIST model generator
Need an MNIST model so why not create one. Code and neural network architecture adapted from https://github.com/tuomaso/train_mnist_fast/blob/master/8_Final_00s76.ipynb

### Setup 
conda create -n mnist
conda activate mnist
conda install pytorch torchvision pytorch-cuda=12.4 -c pytorch -c nvidia
conda install --file requirements.txt
