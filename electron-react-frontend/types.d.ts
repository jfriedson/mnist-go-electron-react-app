type SystemResourceUsage = {
  cpuUsage: number;
  ramUsage: number;
};

type SystemInfo = {
  cpuModel: string;
  totalMemoryGB: number;
};

type EventPayloadMapping = {
  getSystemResourceUsage: SystemResourceUsage;
  getSystemInfo: SystemInfo;
};

type UnsubscribeFunction = () => void;

interface Window {
  electron: {
    getSystemInfo: () => Promise<SystemInfo>;
    subscribeGetSystemResourceUsage: (
      callback: (systemResourceUsage: SystemResourceUsage) => void,
    ) => UnsubscribeFunction;
  };
}
