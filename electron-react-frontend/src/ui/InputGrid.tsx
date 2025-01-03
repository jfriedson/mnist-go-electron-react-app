import { useEffect, useRef, useState } from 'react';
import './InputGrid.css';
import { clearInputGrid, initInputGrid } from './inputGridCanvas';

const canvasDim = 28;

export function InputGrid() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [output, setOuput] = useState<string>(' ');

  const resetInputs = () => {
    clearInputGrid();
    setOuput(' ');
  };

  const sendInferReq = async () => {
    const canvasCtx = canvasRef.current?.getContext("2d");
    if (canvasCtx === undefined || canvasCtx === null)
      return;

    const imageData = canvasCtx.getImageData(0, 0, canvasDim, canvasDim).data;
    
    // convert to grayscale by extracting a single color channel
    const imagePixelCount = canvasDim * canvasDim;
    const inputData = new Uint8ClampedArray(imagePixelCount);
    for (let pixel = 0; pixel < imagePixelCount; pixel++)
      inputData[pixel] = imageData[pixel * 4]
    
    fetch('http://localhost:5122', {
      method: 'POST',
      body: inputData,
      headers: new Headers({
        'Content-Type': 'text/plain',
        'Accept': 'text/plain',
      })
    })
      .then(response => response.text())
      .then(text => setOuput(text))
      .catch(error => console.error(error))
  }

  useEffect(() => {
    return initInputGrid();
  }, []);

  return (
    <>
      <canvas
        ref={canvasRef}
        id="inputGridCanvas"
        className="inputGridCanvas"
        width="28"
        height="28"
      />
      <div className="inputPanel">
        <span id="output">{output}</span>
        <button id="inferGridInput" onClick={sendInferReq}>Infer</button>
        <button id="clearGridInput" onClick={resetInputs}>
          Clear
        </button>
      </div>
    </>
  );
}
