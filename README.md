## Objectives of this Project
- Showcase Go's concurrency by accelerating a neural network's inference on the CPU in tensor operations
- Electron desktop app and React UI in typescript for the front end interface

![animated demo](showcase.webp?raw=true)

### Architecture
Currently, communication between the React UI and the Go service is done by making an HTTP Post request containing the MNIST input from the React UI. The neural network inference result is returned in the HTTP response and displayed by React.

This will pipeline will be offered alongside the option to pipe the MNIST input data from the React UI to Electron's Main process through IPC and use gRPC to invoke the neural network inference on the Go server from Electron Main process. The neural network's output will be returned through the pipeline.

### Elements of this project
1. [MNIST model generator](/mnist-model-generator/) - train and generate

### Remaining features to implement
- accelerate neural network modules using goroutines
- pass neural network input and output using gRPC
- generate neural network model arch file programmatically
- add testing to electron-react front end
- improve testing of go-service
