import osUtils from 'node-os-utils';
import { BrowserWindow } from 'electron';
import { ipcWebContentsSend } from './util.js';

export async function getSystemResourceUsage(mainWindow: BrowserWindow) {
  const cpuUsage = await getCpuUsage();
  const ramUsage = await getRamUsage();

  ipcWebContentsSend('getSystemResourceUsage', mainWindow.webContents, {
    cpuUsage,
    ramUsage,
  });
}

export function getSystemInfo(): SystemInfo {
  const cpuModel = osUtils.cpu.model();
  const totalMemoryGB = +(osUtils.mem.totalMem() / (1024**3)).toFixed(1);

  return {
    cpuModel,
    totalMemoryGB,
  };
}

function getCpuUsage(): Promise<number> {
  return osUtils.cpu
    .usage()
    .then((val) => +((val * 100) / osUtils.cpu.count()).toFixed(1));
}

function getRamUsage(): Promise<number> {
  return osUtils.mem
    .used()
    .then((memUsedInfo) => +(memUsedInfo.usedMemMb / 1024).toFixed(1));
}
