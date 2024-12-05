import { RefObject, useEffect, useRef } from 'react';
import './InputGrid.css';


export function InputGrid() {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  const inputGrid = Array<boolean[]>(28).fill(Array<boolean>(28));

  useEffect(() => {
    clearInputGrid(canvasRef);
  }, []);

  return (
    <div>
      <canvas ref={canvasRef} className='inputGridCanvas' />
      <div className="inputGridButtons">
        <button id='processGridInput'>Process</button>
        <button id='clearGridInput' onClick={() => clearInputGrid(canvasRef)}>Clear</button>
      </div>
    </div>
  );
}

function clearInputGrid(canvasRef: RefObject<HTMLCanvasElement>) {
  const canvas = canvasRef.current;
  if (canvas === null) return;
  const context = canvas.getContext('2d');
  if (context === null) return;

  context.fillStyle = 'green';
  context.fillRect(0, 0, 1600, 1600);
};

function beginRecording(drawing: Boolean) {
  drawing = true;
}

function endRecording(drawing: Boolean) {
  drawing = false;
}
