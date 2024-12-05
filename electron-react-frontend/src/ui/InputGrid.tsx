import './InputGrid.css';

export function InputGrid() {
    const inputGrid = Array<boolean[]>(28).fill(Array<boolean>(28));

    return (
        <div id='inputGridCanvas'>
            {inputGrid}
        </div>
    );
}
