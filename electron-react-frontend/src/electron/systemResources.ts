import osUtils from 'os-utils';
import os from 'os';
import { BrowserWindow } from 'electron';
import { ipcWebContentsSend } from './util.js';

export async function getSystemResourceUsage(mainWindow: BrowserWindow) {
  const cpuUsage = await getCpuUsage();
  const ramUsage = getRamUsage();

  ipcWebContentsSend('getSystemResourceUsage', mainWindow.webContents, {
    cpuUsage,
    ramUsage,
  });
}

export function getSystemInfo() {
  const cpuModel = os.cpus()[0].model;
  const totalMemoryGB = +(osUtils.totalmem() / 1_024).toFixed(1);

  return {
    cpuModel,
    totalMemoryGB,
  };
}

function getCpuUsage(): Promise<number> {
  return new Promise((resolve) => {
    osUtils.cpuUsage(resolve);
  });
}

function getRamUsage() {
  return 1 - osUtils.freememPercentage();
}
