import { useMemo } from 'react'
import './App.css'
import { Chart } from './Chart';
import { getSystemInfo, subscribeGetSystemResourceUsage } from './systemResourcesInfo';
import { InputGrid } from './inputGrid';


function App() {
  const staticData = getSystemInfo();
  const systemResourceUsageDataPointCount = 10;
  const stats = subscribeGetSystemResourceUsage(systemResourceUsageDataPointCount);

  const cpuUsageArray = useMemo(
    () => stats.map((stat) => stat.cpuUsage * 100),
    [stats]
  );
  const ramUsageArray = useMemo(
    () => stats.map((stat) => stat.ramUsage * 100),
    [stats]
  );

  return (
    <>
      <div className="main">
        <div>
          <Chart
            view="CPU"
            title="CPU"
            subTitle={staticData?.cpuModel ?? 'CPU model not reported'}
            data={cpuUsageArray}
            maxDataPoints={systemResourceUsageDataPointCount}
          />
          <Chart
            view="RAM"
            title="RAM"
            subTitle={(staticData?.totalMemoryGB.toString() ?? '') + ' GB'}
            data={ramUsageArray}
            maxDataPoints={systemResourceUsageDataPointCount}
          />
        </div>
        <div className="mainGrid">
          <InputGrid />
        </div>
      </div>
    </>
  );
}

export default App
