import torch
import torch.nn as nn
import torch.nn.functional as F


class Model(nn.Module):
    def __init__(self):
        super(Model, self).__init__()
        self.maxpool = nn.MaxPool2d(2)
        self.relu = nn.ReLU(False)
        self.conv1 = nn.Conv2d(1, 24, 5)
        self.conv2 = nn.Conv2d(24, 32, 3)
        self.flatten = nn.Flatten(1)
        self.fc1 = nn.Linear(800, 256)
        self.fc2 = nn.Linear(256, 10)
        self.logsoftmax = nn.LogSoftmax(1)

    def forward(self, x):
        x = self.conv1(x)
        x = self.maxpool(x)
        x = self.relu(x)
        
        x = self.conv2(x)
        x = self.maxpool(x)
        x = self.relu(x)
        
        x = self.flatten(x)
        x = self.fc1(x)
        x = self.relu(x)
        x = self.fc2(x)

        return self.logsoftmax(x)
