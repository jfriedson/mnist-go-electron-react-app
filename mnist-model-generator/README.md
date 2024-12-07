## MNIST model generator
Need a simple MNIST model so why not create one. Code and NN arch adapted from https://github.com/tuomaso/train_mnist_fast/blob/master/8_Final_00s76.ipynb

### Setup 
conda create -n mnist
conda activate mnist
conda install pytorch torchvision pytorch-cuda=12.4 -c pytorch -c nvidia
conda install --file requirements.txt
