
const ipc = ipcRenderer
////// CLOSE APP
closeBtn.addEventListener('click', ()=>{
    ipc.send('closeApp')
})

////// MINIMIZE APP
minimizeBtn.addEventListener('click', ()=>{
    ipc.send('minimizeApp')
})

////// MAXIMIZE APP
maxResBtn.addEventListener('click', ()=>{
    ipc.send('maximizeApp')
})