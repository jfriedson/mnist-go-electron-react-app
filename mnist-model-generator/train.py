import json
from json import JSONEncoder

import torch
import torchvision
from torch.utils.data import Dataset
import torch.nn.functional as F
import torch.optim as optim
from torch.optim.lr_scheduler import OneCycleLR
from torch.fx import symbolic_trace

from tqdm import tqdm

from model import Model


batch_size = 256
num_workers = 8
persistent_dataloaders = True if num_workers else False
lr=1e-2
epochs=5

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")


train_dataset = torchvision.datasets.MNIST('dataset/', 
                                            train=True, 
                                            download=True,
                                            transform=torchvision.transforms.ToTensor())
train_loader = torch.utils.data.DataLoader(train_dataset,
                                           batch_size=batch_size, 
                                           num_workers=num_workers,
                                           persistent_workers=persistent_dataloaders,
                                           shuffle=True)

test_dataset = torchvision.datasets.MNIST('dataset/', 
                                          train=False, 
                                          download=True,
                                          transform=torchvision.transforms.ToTensor())
test_loader = torch.utils.data.DataLoader(test_dataset,
                                          batch_size=batch_size,
                                          num_workers=num_workers,
                                          persistent_workers=persistent_dataloaders,
                                          shuffle=True)

# Validate data
# import matplotlib.pyplot as plt
# example_datas, labels = next(iter(test_loader))
# sample = example_datas[0][0]
# plt.imshow(sample, cmap='gray', interpolation='none')
# print(example_datas[0][0])
# print("Label: "+ str(labels[0]))
# plt.show()


def train(model, device, optimizer, scheduler):
    model.train()
    total_loss = 0
    correct = 0
    for input, target in tqdm(train_loader):
        input, target = input.to(device), target.to(device)

        optimizer.zero_grad()
        output = model(input)

        loss = F.nll_loss(output, target)
        total_loss += loss.item()
        pred = output.argmax(dim=1, keepdim=True)
        correct += pred.eq(target.view_as(pred)).sum().item()

        loss.backward()
        optimizer.step()
        scheduler.step()

    total_loss /= len(train_dataset)
    print('loss: {:.4f}, acc: ({:.2f}%)\n'.format(
       total_loss, 100. * correct / len(train_dataset)))


def test(model, device):
    model.eval()
    total_loss = 0
    correct = 0
    with torch.no_grad():
        for input, target in tqdm(test_loader):
            input, target = input.to(device), target.to(device)

            output = model(input)

            total_loss += F.nll_loss(output, target, reduction='sum').item()
            pred = output.argmax(dim=1, keepdim=True)
            correct += pred.eq(target.view_as(pred)).sum().item()

    total_loss /= len(test_dataset)
    print('loss: {:.4f}, acc: ({:.2f}%)\n'.format(
       total_loss, 100. * correct / len(test_dataset)))


# https://stackoverflow.com/a/73251115
class EncodeTensor(JSONEncoder, Dataset):
    def default(self, obj):
        if isinstance(obj, torch.Tensor):
            return obj.cpu().detach().numpy().tolist()
        return super(EncodeTensor, self).default(obj)


def main():
    model = Model().to(device)
    optimizer = optim.Adam(model.parameters(), lr=lr, betas=(0.7, 0.9))

    scheduler = OneCycleLR(optimizer, max_lr=lr,
                    total_steps=int(((len(train_dataset)-1)//batch_size + 1)*epochs), 
                    cycle_momentum=False)
    
    for epoch in range(epochs):
        print('epoch {}'.format(epoch + 1))
        train(model, device, optimizer, scheduler)
        test(model, device)

    model.eval()
    model.to("cpu")
    # TODO: dump model arch to file
    # traced = symbolic_trace(model)
    # print(traced.graph)
    # for node in traced.graph.nodes:
    #     for arg in node.args:
    #         print(arg)
    with open('models/mnist.json', 'w') as json_file:
        json.dump(model.state_dict(), json_file, cls=EncodeTensor)

if __name__ == '__main__':
    main()
