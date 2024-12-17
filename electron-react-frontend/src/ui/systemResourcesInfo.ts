import { useEffect, useState } from 'react';

export function GetSystemInfo(): SystemInfo | null {
  const [systemInfo, setSystemInfo] = useState<SystemInfo | null>(null);

  useEffect(() => {
    if (!window.electron) return;

    (async () => {
      setSystemInfo(await window.electron.getSystemInfo());
    })();
  }, []);

  return systemInfo;
}

export function SubscribeGetSystemResourceUsage():
  | SystemResourceUsage
  | null {
  const [systemResourceUsage, setSystemResourceUsage] =
    useState<SystemResourceUsage | null>(null);

  useEffect(() => {
    if (!window.electron) return;

    const unsub = window.electron.subscribeGetSystemResourceUsage((stats) =>
      setSystemResourceUsage(stats)
    );
    return unsub;
  }, []);

  return systemResourceUsage;
}
