## Objectives of this Project
- Showcase Go's concurrency by accelerating a neural network's inference on the CPU in tensor operations
- Utilize Electron desktop app and React UI, facilitating type safety by use of typescript, for the front end interface

![animated demo](showcase.webp?raw=true)

### Elements of this project
1. [MNIST model generator](/mnist-model-generator/) - train and export MNIST digit recognizition model using pytorch
1. [Go service](/go-service/) - accept MNIST digit input data by HTTP, process neural network inference showcasing a multitude of goroutine strategies, and respond with result
1. [Electron app with React UI](/electron-react-frontend/) - HTML canvas enabling the app user to draw MNIST digit input that is sent to the go service by Javascript Fetch API, and the neural network output is displayed using a React hook in a callback when the promise is fulfilled. System CPU and RAM info and live usage is periodically reported by the Electron main process to the render process over IPC which then updates the subscribed React UI by callback method.

### Architecture
Currently, communication between the React UI and the Go service is done by making an HTTP Post request containing the MNIST input from the React UI. The neural network inference result is returned in the HTTP response and displayed by React.

This will pipeline will be offered alongside the option to pipe the MNIST input data from the React UI to Electron's Main process through IPC and use gRPC to invoke the neural network inference on the Go server from Electron Main process. The neural network's output will be returned through the pipeline.

### Remaining features to implement
- pass neural network input and output using gRPC
- generate neural network model arch file programmatically
- add testing to electron-react front end
