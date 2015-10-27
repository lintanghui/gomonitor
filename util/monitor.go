package util

import (
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"
)

func (w *GoMonitor) Monitor() {
    c := time.Tick(time.Second * time.Duration(w.Interval))
    for {
        select {
        case <-w.change: // file has been change
            w.BuildAndRun()
        case <-c: // walk every 5 seconds
            w.WalkFile()
        }
    }
}

func (w *GoMonitor) BuildAndRun() {
    // if the process is running ,kill it before build
    if w.cmd != nil && w.cmd.Process != nil {
        log.Printf("the process :%d\n", w.cmd.Process.Pid)
        if err := w.cmd.Process.Kill(); err != nil {
            log.Printf("%s %s\n", w.cmd.ProcessState.String(), err)
        } else {
            log.Printf("%s\n", w.cmd.ProcessState.String())
        }
    }
    if err := w.Build(); err != nil {
        log.Printf("build fail %s\n", err)
    } else {
        log.Printf("start process :%s\n", w.RunCmd)
        go w.Run()
    }
}

func (w *GoMonitor) Run() {
    args := strings.Split(w.RunCmd, " ")
    w.cmd = exec.Command(args[0], args[1:]...)
    w.cmd.Stdin = os.Stdin
    w.cmd.Stdout = os.Stdout
    w.cmd.Stderr = os.Stderr
    err := w.cmd.Run()
    if err != nil {
        log.Printf("%s\n", err)
    }
    w.cmd = nil
}

func (w *GoMonitor) Build() (err error) {
    args := strings.Split(w.BuildCmd, " ")
    out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
    log.Printf("[build cmd] %s", w.BuildCmd)

    if err != nil {
        log.Printf("%s\n", string(out))
        return
    }
    log.Printf("build sucess\n")

    return
}
func (w *GoMonitor) WalkFile() {
    var change bool
    for file, modtime := range w.FileStatus {
        info, err := os.Stat(file)
        if err != nil {
            log.Printf("file :%s does not exit or had been delete", err.Error())
            // file had been moved
            delete(w.FileStatus, file)
            continue
        }
        newModTime := info.ModTime().Unix()
        // update modtime
        if modtime != newModTime {
            log.Printf("File :%s has been changed", file)
            w.FileStatus[file] = newModTime
            change = true
            if info.IsDir() {
                filepath.Walk(file, w.updatedir)
            }
        }
    }

    if change {
        w.change <- true
    }

}

func (w *GoMonitor) updatedir(path string, f os.FileInfo, err error) error {
    if err != nil {
        log.Printf("%s\n", err)
        return err
    }

    w.FileStatus[path] = f.ModTime().Unix()
    return nil
}
