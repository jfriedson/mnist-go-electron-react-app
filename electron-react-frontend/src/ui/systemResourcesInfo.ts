import { useEffect, useState } from 'react';

export function getSystemInfo(): SystemInfo | null {
	const [systemInfo, setSystemInfo] = useState<SystemInfo | null>(null);

	useEffect(() => {
		(async () => {
			setSystemInfo(await window.electron.getSystemInfo());
		})();
	}, []);

	return systemInfo;
}

export function subscribeGetSystemResourceUsage(): SystemResourceUsage | undefined {
    const [systemResourceUsage, setSystemResourceUsage] =
		useState<SystemResourceUsage>();

	useEffect(() => {
		const unsub = window.electron.subscribeGetSystemResourceUsage((stats) =>
			setSystemResourceUsage(stats)
		);
		return unsub;
	}, []);
    
    return systemResourceUsage;
}
