import { useEffect, useState } from "react";

export function getSystemInfo() {
	const [systemInfo, setSystemInfo] = useState<SystemInfo | null>(null);

	useEffect(() => {
		(async () => {
			setSystemInfo(await window.electron.getSystemInfo());
		})();
	}, []);

	return systemInfo;
}

export function subscribeGetSystemResourceUsage(dataPointCountLimit: number): SystemResourceUsage[] {
    const [systemResourceUsageArray, setSystemResourceUsageArray] = useState<
		SystemResourceUsage[]
	>([]);

    useEffect(() => {
        const unsub = window.electron.subscribeGetSystemResourceUsage((stats) =>
            // append new system resource usage data points to the end of the array
            // remove oldest data point if data point count exceeds limit
			setSystemResourceUsageArray((prev) => {
				const newArray = [...prev, stats];

				if (newArray.length > dataPointCountLimit) {
					newArray.shift();
				}

				return newArray;
			})
		);
		return unsub;
    }, []);
    
    return systemResourceUsageArray;
}
