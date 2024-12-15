let canvas: HTMLCanvasElement;
let drawingContext: CanvasRenderingContext2D;

const clearColor = '#000000';
const setColor = '#FFFFFF';
const brushRadius = 0.9;

export function initInputGrid(): UnsubscribeFunction {
  canvas = document.getElementById('inputGridCanvas')! as HTMLCanvasElement;
  drawingContext = canvas.getContext('2d')!;

  clearInputGrid();

  canvas.addEventListener('mousedown', placeBrush);
  canvas.addEventListener('mousemove', moveBrush);

  return () => {
    canvas.removeEventListener('mousedown', placeBrush);
    canvas.removeEventListener('mousemove', moveBrush);
  };
}

export function clearInputGrid() {
  drawingContext.fillStyle = clearColor;
  drawingContext.fillRect(0, 0, canvas.width, canvas.height);
}

const placeBrush = (ev: MouseEvent) => {
  let color: string;
  let radius: number;
  if (ev.button === 0) {
    color = setColor;
    radius = brushRadius;
  }
  else if (ev.button === 2) {
    color = clearColor;
    radius = 1;
  }
  else return;

  const mouseX = ev.clientX - canvas.offsetLeft;
  const mouseY = ev.clientY - canvas.offsetTop;
  const canvasX = (mouseX * canvas.width) / canvas.clientWidth;
  const canvasY = (mouseY * canvas.height) / canvas.clientHeight;

  fillCell(Math.round(canvasX), Math.round(canvasY), color, radius);
};

const moveBrush = (ev: MouseEvent) => {
  let color: string;
  let radius: number;
  if (ev.buttons === 1) {
    color = setColor;
    radius = brushRadius;
  } else if (ev.buttons === 2) {
    color = clearColor;
    radius = 1;
  } else return;

  const mouseX = ev.clientX - canvas.offsetLeft;
  const mouseY = ev.clientY - canvas.offsetTop;
  const canvasX = (mouseX * canvas.width) / canvas.clientWidth;
  const canvasY = (mouseY * canvas.height) / canvas.clientHeight;

  fillCell(Math.round(canvasX), Math.round(canvasY), color, radius);
};

function fillCell(cellX: number, cellY: number, color: string, radius: number) {
  drawingContext.fillStyle = color;
  drawingContext.fillRect(
    cellX - radius,
    cellY - radius,
    radius * 2,
    radius * 2
  );
}
