type Statistics = {
	cpuUsage: number;
	ramUsage: number;
	storageUsage: number;
};

type StaticData = {
	cpuModel: string;
	totalMemoryGB: number;
	totalStorage: number;
};

type View = 'CPU' | 'RAM' | 'Storage';

type EventPayloadMapping = {
	statistics: Statistics;
	getStaticData: StaticData;
	changeView: View;
};

type UnsubscribeFunction = () => void;

interface Window {
	electron: {
		subscribeStatistics: (
			callback: (statistics: Statistics) => void
		) => UnsubscribeFunction;
		getStaticData: () => Promise<StaticData>;
		subscribeChangeView: (
			callback: (view: View) => void
		) => UnsubscribeFunction;
	};
}
