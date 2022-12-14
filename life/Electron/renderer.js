// This file is required by the index.html file and will
// be executed in the renderer process for that window.
// No Node.js APIs are available in this process because
// `nodeIntegration` is turned off. Use `preload.js` to
// selectively enable features needed in the rendering
// process.
const electron = require('electron');
const ipcRenderer = electron.ipcRenderer;
ipcRenderer.on('cpu',(event,data) => {
    document.getElementById('cpu').innerHTML = data.toFixed(2);
  });
  ipcRenderer.on('mem',(event,data) => {
    document.getElementById('mem').innerHTML = data.toFixed(2);
  });
  ipcRenderer.on('total-mem',(event,data) => {
    document.getElementById('total-mem').innerHTML = data.toFixed(2);
  });