import { useCallback, useEffect, useMemo, useState } from "react";

export function getSystemInfo() {
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
			setSystemResourceUsage(() => {
				return stats;
			})
		);
		return unsub;
	}, []);
    
    return systemResourceUsage;
}
