// Modules to control application life and create native browser window
const {app, BrowserWindow, ipcMain} = require('electron')
const os = require('os-utils')
const path = require('path')
const ipc = ipcMain

function createWindow () {
  // Create the browser window.
  const mainWindow = new BrowserWindow({
    width: 1200,
    height: 680,
    minWidth: 940,
    minHeight: 560,
    frame: false,
    icon: __dirname + '/images/icon.ico',
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false,
      devTools: true,
      preload: path.join(__dirname, 'preload.js')
    }
  })
  // close dev bar
 // mainWindow.setMenuBarVisibility(false)
  // and load the index.html of the app.
  mainWindow.loadFile('index.html')
  mainWindow.setBackgroundColor('#343B48')

  // Open the DevTools.
  mainWindow.webContents.openDevTools()

  // PC resource uses
  setInterval(() => {
    os.cpuUsage(function(v){
      mainWindow.webContents.send('cpu',v*100);
      mainWindow.webContents.send('mem',os.freememPercentage()*100);
      mainWindow.webContents.send('total-mem',os.totalmem()/1024);
    });
  },1000);

  /////CLOSE APP
  ipc.on('closeApp', ()=> {
    mainWindow.close()
  })

  /////MINIMIZE APP
  ipc.on('minimizeApp', ()=> {
    mainWindow.minimize()
  })


   /////MAXIMIZE APP
   ipc.on('maximizeApp', ()=> {
    mainWindow.maximize()
  })
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(() => {
  createWindow()

  app.on('activate', function () {
    // On macOS it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
  })
})

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', function () {
  if (process.platform !== 'darwin') app.quit()
})

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.
