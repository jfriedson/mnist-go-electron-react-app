import { useEffect, useMemo, useState } from 'react'
import './App.css'
import { Chart } from './Chart';
import { getSystemInfo, subscribeGetSystemResourceUsage } from './systemResourcesInfo';
import { InputGrid } from './InputGrid';


function App() {
  const systemInfo = getSystemInfo();
  const systemResourceUsageDataPointCount = 10;
  const systemResourceUsage = subscribeGetSystemResourceUsage();
  const [cpuUsageArray, setCpuUsageArray] = useState(Array<number>());
  const [ramUsageArray, setRamUsageArray] = useState(Array<number>());
  
  useMemo(() => {
    if (systemResourceUsage === undefined)
      return;

    var newCpuUsageArray = [...cpuUsageArray, systemResourceUsage.cpuUsage * 100];
    if (newCpuUsageArray.length > systemResourceUsageDataPointCount)
      newCpuUsageArray.shift();
    setCpuUsageArray(newCpuUsageArray);

    var newRamUsageArray = [...ramUsageArray, systemResourceUsage.cpuUsage  * 100];
    if (newRamUsageArray.length > systemResourceUsageDataPointCount)
      newRamUsageArray.shift();
    setRamUsageArray(newRamUsageArray);

  }, [systemResourceUsage]);

  return (
    <div id="app">
      <div id="appSidebar">
        <Chart
          view="CPU"
          title="CPU"
          subTitle={systemInfo?.cpuModel ?? 'CPU model not reported'}
          data={cpuUsageArray}
          maxDataPoints={systemResourceUsageDataPointCount}
        />
        <Chart
          view="RAM"
          title="RAM"
          subTitle={(systemInfo?.totalMemoryGB.toString() ?? 'Unreported') + ' GB'}
          data={ramUsageArray}
          maxDataPoints={systemResourceUsageDataPointCount}
        />
      </div>
      <div id="appMain">
        <InputGrid />
      </div>
    </div>
  );
}

export default App
