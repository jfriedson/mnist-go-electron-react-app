import { app, BrowserWindow } from 'electron';
import { ipcMainHandle, isDev } from './util.js';
import { getPreloadPath, getUIPath } from './pathResolver.js';
import { createMenu } from './menu.js';
import { getSystemInfo, getSystemResourceUsage } from './systemResources.js';


const POLLING_INTERVAL = 500;

app.on('ready', () => {
	const mainWindow = new BrowserWindow({
		webPreferences: {
			preload: getPreloadPath(),
		},
	});
	
	if (isDev()) {
		mainWindow.loadURL('http://localhost:5123');
	} else {
		mainWindow.loadFile(getUIPath());
	}

	setInterval(async () => {
		getSystemResourceUsage(mainWindow);
	}, POLLING_INTERVAL);

	ipcMainHandle('getSystemInfo', () => {
		return getSystemInfo();
	});

	createMenu(mainWindow);
});
