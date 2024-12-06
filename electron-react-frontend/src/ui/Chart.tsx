import './Chart.css'
import { useMemo } from 'react';
import { BaseChart } from './BaseChart';


export type View = 'CPU' | 'RAM';

interface ChartProps {
  view: View;
  subTitle: string;
  data: number[];
  maxDataPoints: number;
}

const COLOR_MAP = {
  CPU: {
    stroke: '#5DD4EE',
    fill: '#0A4D5C',
  },
  RAM: {
    stroke: '#E99311',
    fill: '#5F3C07',
  },
  Storage: {
    stroke: '#1ACF4D',
    fill: '#0B5B22',
  },
};

export function Chart(props: ChartProps) {
  const color = COLOR_MAP[props.view];

  const preparedData = useMemo(() => {
    const points = props.data.map((point) => ({ value: point }));
    return [
      ...points,
      ...Array.from({ length: props.maxDataPoints - points.length }).map(
        () => ({ value: undefined })
      ),
    ];
  }, [props.data, props.maxDataPoints]);

  return (
    <div>
      <div className='chartTitle'>
        <div>{props.view}</div>
        <div>{props.subTitle}</div>
      </div>
      <div className='chart'>
        <BaseChart data={preparedData} fill={color.fill} stroke={color.stroke} />
      </div>
    </div>
  );
}
