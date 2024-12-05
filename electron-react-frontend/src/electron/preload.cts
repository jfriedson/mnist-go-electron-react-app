const electron = require('electron');


electron.contextBridge.exposeInMainWorld("electron", {
	getSystemInfo: () => ipcInvoke("getSystemInfo"),
	subscribeGetSystemResourceUsage: (callback) =>
		ipcOn("getSystemResourceUsage", (stats) => {
			callback(stats);
		}),
} satisfies Window["electron"]);

function ipcInvoke<Key extends keyof EventPayloadMapping>(
	key: Key, 
): Promise<EventPayloadMapping[Key]> {
	return electron.ipcRenderer.invoke(key);
}

function ipcOn<Key extends keyof EventPayloadMapping>(
	key: Key,
	callback: (payload: EventPayloadMapping[Key]) => void
) {
    const cb = (_: Electron.IpcRendererEvent, payload: any) => callback(payload);
    electron.ipcRenderer.on(key, (_, payload) => callback(payload));
    return () => electron.ipcRenderer.off(key, cb);
}
