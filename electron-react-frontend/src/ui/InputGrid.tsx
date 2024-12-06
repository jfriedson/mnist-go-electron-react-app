import { MouseEvent, useEffect, useRef, useState } from 'react';
import './InputGrid.css';


const clearColor = '#E0E0E0';
const setColor = '#0F0F0F';
const pencilRadius = 1;

export function InputGrid() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [drawing, setDrawing] = useState<boolean>(false);

  //const inputGrid = boolean[boolean[]]>(28).fill(Array<boolean>(28));
  const [output, setOuput] = useState<string>(' ');

  const getCanvasCtx = () => {
    const canvas = canvasRef.current;
    var canvasCtx = null;
    if (canvas !== null) canvasCtx = canvas.getContext('2d');

    return { canvas, canvasCtx };
  }

  const clearInputGrid = () => {
    const { canvas, canvasCtx } = getCanvasCtx();
    if (canvas === null || canvasCtx === null)
      return;

    canvasCtx.fillStyle = clearColor;
    canvasCtx.fillRect(0, 0, canvas.width, canvas.height);
  }

  const resetInputs = () => {
    clearInputGrid();
    setOuput(' ');
  }

  const beginDrawing = (event: MouseEvent) => {
    setDrawing(true);
    moveDrawing(event, true);
  }

  const moveDrawing = (event: MouseEvent, forceDraw?: boolean) => {
    if (!drawing && !forceDraw)
      return;

    const { canvas, canvasCtx } = getCanvasCtx();
    if (canvas === null || canvasCtx === null)
      return;

    const mouseX = event.clientX - canvas.offsetLeft;
    const mouseY = event.clientY - canvas.offsetTop;
		const canvasX = mouseX * canvas.width / canvas.clientWidth;
    const canvasY = mouseY * canvas.height / canvas.clientHeight;
    
    fillCell(canvasCtx, canvasX, canvasY);
  }

  const fillCell = (canvasCtx: CanvasRenderingContext2D, cellX: number, cellY: number) => {
		const startX = cellX;
		const startY = cellY;

		canvasCtx.fillStyle = setColor;
		canvasCtx.fillRect(
			startX,
			startY,
			pencilRadius,
			pencilRadius
		);
	}

  useEffect(() => resetInputs, []);

  return (
    <>
      <canvas ref={canvasRef}
        id='inputGridCanvas'
        className='inputGridCanvas'
        width='28' height='28'
        onMouseDown={beginDrawing}
        onMouseMove={moveDrawing}
        onMouseUp={() => setDrawing(false) }
        onMouseLeave={() => setDrawing(false) }
      />
      <div className='inputPanel'>
        <span id='output'>{output}</span>
        <button id='inferGridInput'>Infer</button>
        <button id='clearGridInput' onClick={resetInputs}>Clear</button>
      </div>
    </>
  );
}
