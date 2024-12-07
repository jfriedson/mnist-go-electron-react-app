import torch
import torchvision
import torch.nn as nn
import torch.nn.functional as F
import torch.optim as optim
from torch.optim.lr_scheduler import OneCycleLR
from tqdm import tqdm


batch_size = 256
num_workers = 8
lr=1e-2
epochs=5

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")

train_dataset = torchvision.datasets.MNIST('dataset/', 
                                            train=True, 
                                            download=True,
                                            transform=torchvision.transforms.ToTensor())
train_loader = torch.utils.data.DataLoader(train_dataset,
                                           batch_size=batch_size, 
                                        #    num_workers=num_workers,
                                           shuffle=True)

test_dataset = torchvision.datasets.MNIST('dataset/', 
                                          train=False, 
                                          download=True,
                                          transform=torchvision.transforms.ToTensor())
test_loader = torch.utils.data.DataLoader(test_dataset,
                                          batch_size=batch_size,
                                        #   num_workers=num_workers,
                                          shuffle=True)

# Validate data
# import matplotlib.pyplot as plt
# example_datas, labels = next(iter(test_loader))
# sample = example_datas[0][0]
# plt.imshow(sample, cmap='gray', interpolation='none')
# print(example_datas[0][0])
# print("Label: "+ str(labels[0]))
# plt.show()


class Net(nn.Module):
    def __init__(self):
        super(Net, self).__init__()
        self.conv1 = nn.Conv2d(1, 24, 5, 1)
        self.conv2 = nn.Conv2d(24, 32, 3, 1)
        self.fc1 = nn.Linear(800, 256)
        self.fc2 = nn.Linear(256, 10)

    def forward(self, x):
        x = self.conv1(x)
        x = F.max_pool2d(x, 2)
        x = F.relu(x)
        
        x = self.conv2(x)
        x = F.max_pool2d(x, 2)
        x = F.relu(x)
        
        x = torch.flatten(x, 1)
        x = self.fc1(x)
        x = F.relu(x)
        x = self.fc2(x)
        output = F.log_softmax(x, dim=1)
        return output


def train(model, device, optimizer, scheduler):
    model.train()
    total_loss = 0
    for input, target in tqdm(train_loader):
        optimizer.zero_grad()
        output = model(input)
        loss = F.nll_loss(output, target)
        total_loss += loss.item()
        loss.backward()
        optimizer.step()
        scheduler.step()

    total_loss /= len(test_dataset)
    print('loss: {:.4f}\n'.format(total_loss))

def test(model, device):
    model.eval()
    total_loss = 0
    correct = 0
    with torch.no_grad():
        for input, target in tqdm(test_loader):
            output = model(input)
            total_loss += F.nll_loss(output, target, reduction='sum').item()
            pred = output.argmax(dim=1, keepdim=True)
            correct += pred.eq(target.view_as(pred)).sum().item()

    total_loss /= len(test_dataset)
    print('loss: {:.4f}, acc: ({:.2f}%)\n'.format(
       total_loss, 100. * correct / len(test_data)))


def main():
    model = Net().to(device)
    optimizer = optim.Adam(model.parameters(), lr=lr, betas=(0.7, 0.9))

    scheduler = OneCycleLR(optimizer, max_lr=lr,
                           total_steps=int(((len(train_dataset)-1)//batch_size + 1)*epochs), 
                           cycle_momentum=False)
    
    for _ in range(epochs):
        train(model, device, optimizer, scheduler)
        test(model, device)

main()
torch.save(model.state_dict(), "mnist_cnn.pt")