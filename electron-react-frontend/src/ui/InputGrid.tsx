import { useEffect, useRef, useState } from 'react';
import './InputGrid.css';
import { clearInputGrid, initInputGrid } from './inputGridCanvas';

export function InputGrid() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [output, setOuput] = useState<string>(' ');

  const resetInputs = () => {
    clearInputGrid();
    setOuput(' ');
  };

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
        <button id="inferGridInput">Infer</button>
        <button id="clearGridInput" onClick={resetInputs}>
          Clear
        </button>
      </div>
    </>
  );
}
