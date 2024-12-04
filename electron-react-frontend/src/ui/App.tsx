import { useEffect, useMemo, useState } from 'react'
import './App.css'
import { Chart } from './Chart';
import { statistics } from './statistics';

function App() {
  const staticData = getStaticData();
  const stats = statistics(10);
  const [activeView, setActiveView] = useState<View>('CPU');
  const cpuUsageArray = useMemo(
    () => stats.map((stat) => stat.cpuUsage * 100),
    [stats]
  );
  const ramUsageArray = useMemo(
    () => stats.map((stat) => stat.ramUsage * 100),
    [stats]
  );
  const storageUsageArray = useMemo(
    () => stats.map((stat) => stat.storageUsage * 100),
    [stats]
  );
  const activeViewUsageArray = useMemo(() => {
    switch (activeView) {
      case 'CPU':
        return cpuUsageArray;
      case 'RAM':
        return ramUsageArray;
      case 'Storage':
        return storageUsageArray;
    }
  }, [activeView, cpuUsageArray, ramUsageArray, storageUsageArray]);

  useEffect(() => {
    return window.electron.subscribeChangeView((view) => setActiveView(view));
  }, []);

  return (
    <>
      <div className="main">
        <div>
          <SelectOption
            onClick={() => setActiveView('CPU')}
            title="CPU"
            view="CPU"
            subTitle={staticData?.cpuModel ?? ''}
            data={cpuUsageArray}
          />
          <SelectOption
            onClick={() => setActiveView('RAM')}
            title="RAM"
            view="RAM"
            subTitle={(staticData?.totalMemoryGB.toString() ?? '') + ' GB'}
            data={ramUsageArray}
          />
          <SelectOption
            onClick={() => setActiveView('Storage')}
            title="Storage"
            view="Storage"
            subTitle={(staticData?.totalStorage.toString() ?? '') + ' GB'}
            data={storageUsageArray}
          />
        </div>
        <div className="mainGrid">
          <Chart
            selectedView={activeView}
            data={activeViewUsageArray}
            maxDataPoints={10}
          />
        </div>
      </div>
    </>
  );
}

function SelectOption(props: {
  title: string;
  view: View;
  subTitle: string;
  data: number[];
  onClick: () => void;
}) {
  return (
    <button className="selectOption" onClick={props.onClick}>
      <div className="selectOptionTitle">
        <div>{props.title}</div>
        <div>{props.subTitle}</div>
      </div>
      <div className="selectOptionChart">
        <Chart selectedView={props.view} data={props.data} maxDataPoints={10} />
      </div>
    </button>
  );
}

function getStaticData() {
  const [staticData, setStaticData] = useState<StaticData | null>(null);

  useEffect(() => {
    (async () => {
      setStaticData(await window.electron.getStaticData());
    })();
  }, []);

  return staticData;
}

export default App
