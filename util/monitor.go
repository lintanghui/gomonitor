package util

import (
    "log"
    "os"
    "os/exec"
    "time"
)

func (w *GoMonitor) Monitor() {
    c := time.Tick(time.Second * time.Duration(w.Interval))
    for {
        select {
        case <-w.change: // file has been change
            log.Printf("file had been change rebuild and run:%s", w.RootDir)
            w.BuildAndRun()
        case <-c: // walk every 5 seconds
            w.WalkFile()
        }
    }
}

func (w *GoMonitor) BuildAndRun() {
    w.cmd = exec.Command("go", "build", w.WorkDir)
    w.cmd.Stdin = os.Stdin
    w.cmd.Stdout = os.Stdout
    w.cmd.Stderr = os.Stderr
    err := w.cmd.Run()
    if err != nil {
        log.Printf("%s\n", err)
    }
    w.cmd = nil
}

func (w *GoMonitor) WalkFile() {
    var change bool
    for file, modtime := range w.FileStatus {
        info, err := os.Stat(file)
        if err != nil {
            log.Printf("filename err :%s", err.Error())
        }
        newModTime := info.ModTime().Unix()
        if modtime != newModTime {
            log.Printf("File :%s has been changed", file)
            w.FileStatus[file] = newModTime
            change = true
        }
    }

    if change {
        w.change <- true
    }

}
