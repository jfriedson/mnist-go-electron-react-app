import { useEffect, useRef, useState } from 'react';
import './InputGrid.css';
import InputGridCanvas from './inputGridCanvas';


export function InputGrid() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [output, setOuput] = useState<string>(' ');

  const resetInputs = () => {
    InputGridCanvas.clear();
    setOuput(' ');
  }

  useEffect(() => { return InputGridCanvas.init(); }, []);

  return (
    <>
      <canvas ref={canvasRef}
        id='inputGridCanvas'
        className='inputGridCanvas'
        width='28' height='28'
      />
      <div className='inputPanel'>
        <span id='output'>{output}</span>
        <button id='inferGridInput'>Infer</button>
        <button id='clearGridInput' onClick={resetInputs}>Clear</button>
      </div>
    </>
  );
}
