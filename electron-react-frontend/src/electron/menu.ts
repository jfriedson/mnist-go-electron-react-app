import { app, BrowserWindow, Menu } from 'electron';
import { ipcWebContentsSend } from './util.js';

export function createMenu(mainWindow: BrowserWindow) {
  Menu.setApplicationMenu(
    Menu.buildFromTemplate([
      {
        label: process.platform === 'darwin' ? undefined : 'App',
        type: 'submenu',
        submenu: [
          {
            label: 'DevTools',
            click: () => mainWindow.webContents.openDevTools(),
          },
          {
            label: 'Quit',
            click: app.quit,
          },
        ],
      },
    ])
  );
}
