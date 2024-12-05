## Objectives of this Project
- Showcase Go's concurrency to accelerate neural network calculations on the CPU and personally implement the matrix operations
- Electron app and React UI in typescript for the front end interface

### Architecture
- Communication between the React UI and the Go service is done by piping the neural network's input from React to Electron's Main process through IPC and making a gRPC call from the Electron Main process to the Go server. The neural network's output is returned through the pipeline in reverse.
