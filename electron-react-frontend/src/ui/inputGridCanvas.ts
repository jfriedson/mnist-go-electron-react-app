let canvas: HTMLCanvasElement;
let drawingContext: CanvasRenderingContext2D;

const clearColor = '#000000';
const setColor = '#FFFFFF';
const brushRadius = 0.9;

let painting: boolean = false;

export function initInputGrid(): UnsubscribeFunction {
  canvas = document.getElementById('inputGridCanvas')! as HTMLCanvasElement;
  drawingContext = canvas.getContext('2d')!;

  clearInputGrid();

  canvas.addEventListener('mousedown', placeBrush);
  canvas.addEventListener('mousemove', moveBrush);
  canvas.addEventListener('mouseup', liftBrush);
  canvas.addEventListener('mouseleave', liftBrush);

  return () => {
    canvas.removeEventListener('mousedown', placeBrush);
    canvas.removeEventListener('mousemove', moveBrush);
    canvas.removeEventListener('mouseup', liftBrush);
    canvas.removeEventListener('mouseleave', liftBrush);
  };
}

export function clearInputGrid() {
  drawingContext.fillStyle = clearColor;
  drawingContext.fillRect(0, 0, canvas.width, canvas.height);
}

const placeBrush = (ev: MouseEvent) => {
  if (ev.button !== 0) return;

  painting = true;

  const mouseX = ev.clientX - canvas.offsetLeft;
  const mouseY = ev.clientY - canvas.offsetTop;
  const canvasX = (mouseX * canvas.width) / canvas.clientWidth;
  const canvasY = (mouseY * canvas.height) / canvas.clientHeight;

  fillCell(Math.round(canvasX), Math.round(canvasY));
};

const moveBrush = (ev: MouseEvent) => {
  if (ev.button !== 0 || painting === false) return;

  const mouseX = ev.clientX - canvas.offsetLeft;
  const mouseY = ev.clientY - canvas.offsetTop;
  const canvasX = (mouseX * canvas.width) / canvas.clientWidth;
  const canvasY = (mouseY * canvas.height) / canvas.clientHeight;

  fillCell(Math.round(canvasX), Math.round(canvasY));
};

const liftBrush = (ev: MouseEvent) => {
  if (ev.button === 0) painting = false;
};

function fillCell(cellX: number, cellY: number) {
  drawingContext.fillStyle = setColor;
  drawingContext.fillRect(
    cellX - brushRadius,
    cellY - brushRadius,
    brushRadius * 2,
    brushRadius * 2
  );
}
