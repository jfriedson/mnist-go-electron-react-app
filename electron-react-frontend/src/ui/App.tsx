import { useEffect, useState } from 'react';
import './App.css';
import { Chart } from './Chart';
import { InputGrid } from './InputGrid';
import {
  GetSystemInfo,
  SubscribeGetSystemResourceUsage,
} from './SystemResourcesInfo';

const systemResourceUsageDataPointCount = 20;

function App() {
  const systemInfo = GetSystemInfo();
  const systemResourceUsage = SubscribeGetSystemResourceUsage();
  const [cpuUsageArray, setCpuUsageArray] = useState<number[]>([]);
  const [ramUsageArray, setRamUsageArray] = useState<number[]>([]);

  useEffect(() => {
    if (systemResourceUsage === null) return;

    const updateArray = (oldArray: number[], newValue: number) => {
      const newArray = [...oldArray, newValue];
      if (newArray.length > systemResourceUsageDataPointCount) newArray.shift();
      return newArray;
    };

    setCpuUsageArray(
      updateArray(cpuUsageArray, systemResourceUsage.cpuUsage)
    );
    setRamUsageArray(
      updateArray(ramUsageArray, 100 * systemResourceUsage.ramUsage / (systemInfo?.totalMemoryGB ?? 1))
    );

    // cpuUsageArray and ramUsageArray alterations should not trigger effect
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [systemResourceUsage]);

  return (
    <div id="app">
      <div id="appSidebar">
        <Chart
          view="CPU"
          subTitle={cpuChartSubtitle(
            systemInfo?.cpuModel,
            systemResourceUsage?.cpuUsage
          )}
          data={cpuUsageArray}
          maxDataPoints={systemResourceUsageDataPointCount}
        />
        <Chart
          view="RAM"
          subTitle={ramChartSubtitle(
            systemResourceUsage?.ramUsage,
            systemInfo?.totalMemoryGB
          )}
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

function cpuChartSubtitle(
  cpuModel: string | undefined,
  cpuUsage: number | undefined
) {
  if (cpuModel !== undefined && cpuUsage !== undefined)
    return cpuModel + ' - ' + cpuUsage + '%';

  if (cpuModel !== undefined)
    return cpuModel;

  if (cpuUsage !== undefined)
    return cpuUsage + '%';

  return 'Unreported';
}

function ramChartSubtitle(
  ramUsage: number | undefined,
  ramTotal: number | undefined
) {
  if (ramUsage !== undefined && ramTotal !== undefined)
    return ramUsage + ' / ' + ramTotal + ' GB';

  if (ramTotal !== undefined)
    return ramTotal + ' GB';

  if (ramUsage !== undefined)
    return ramUsage + '%';

  return 'Unreported';
}

export default App;
