import { useEffect, useState } from 'react'
import './App.css'
import { Chart } from './Chart';
import { getSystemInfo, subscribeGetSystemResourceUsage } from './systemResourcesInfo';
import { InputGrid } from './InputGrid';


const systemResourceUsageDataPointCount = 10;

function App() {
  const systemInfo = getSystemInfo();
  const systemResourceUsage = subscribeGetSystemResourceUsage();
  const [cpuUsageArray, setCpuUsageArray] = useState<number[]>([]);
  const [ramUsageArray, setRamUsageArray] = useState<number[]>([]);
  
  useEffect(() => {
    if (systemResourceUsage === undefined)
      return;

    const updateArray = (oldArray: number[], newValue: number) => {
      var newArray = [...oldArray, newValue];
      if (newArray.length > systemResourceUsageDataPointCount)
        newArray.shift();
      return newArray;
    }

    setCpuUsageArray(updateArray(cpuUsageArray, systemResourceUsage.cpuUsage * 100));
    setRamUsageArray(updateArray(ramUsageArray, systemResourceUsage.ramUsage * 100));
  }, [systemResourceUsage]);

  return (
    <div id='app'>
      <div id='appSidebar'>
        <Chart
          view='CPU'
          subTitle={cpuChartSubtitle(systemInfo?.cpuModel, systemResourceUsage?.cpuUsage)}
          data={cpuUsageArray}
          maxDataPoints={systemResourceUsageDataPointCount}
        />
        <Chart
          view='RAM'
          subTitle={ramChartSubtitle(systemResourceUsage?.ramUsage, systemInfo?.totalMemoryGB)}
          data={ramUsageArray}
          maxDataPoints={systemResourceUsageDataPointCount}
        />
      </div>
      <div id='appMain'>
        <InputGrid />
      </div>
    </div>
  );
}

function cpuChartSubtitle(cpuModel: string | undefined, cpuUsage: number | undefined) {
  if (cpuModel === undefined || cpuUsage === undefined) {
    if (cpuModel !== undefined)
      return cpuModel;

    if (cpuUsage !== undefined)
      return (cpuUsage * 100).toFixed(1) + '%';
    
    return 'Unreported';
  }

  return cpuModel + ' - ' + (cpuUsage * 100).toFixed(1) + '%';
}

function ramChartSubtitle(ramUsage: number | undefined, ramTotal: number | undefined) {
  if (ramUsage === undefined || ramTotal === undefined) {
    if (ramTotal !== undefined)
      return ramTotal.toString() + ' GB';

    if (ramUsage !== undefined)
      return (ramUsage * 100).toFixed(1) + '%';
    
    return 'Unreported';
  }

  return (ramUsage * ramTotal).toFixed(1) + ' / ' + ramTotal.toString() + ' GB';
}

export default App
