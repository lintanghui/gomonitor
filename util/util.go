package util

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

type GoMonitor struct {
    RootDir    []string         // the dir to monitor that set at configuration file
    sourceDir  []string         // directory files in RootDir
    FileStatus map[string]int64 // file path and modtime
}

var DefMonitor = NewGoMonitor()

func NewGoMonitor() (goMonitor *GoMonitor) {
    m := make(map[string]int64, 20)
    goMonitor = &GoMonitor{FileStatus: m}

    return
}

// add config dir to RootDir
func (w *GoMonitor) AddRootDir(path string) (err error) {
    dirinfo, err := os.Stat(path)
    if err != nil {
        log.Printf("%s\n", err)
        return
    }
    if dirinfo.IsDir() {
        w.RootDir = append(w.RootDir, path)
    }

    // walk the file tree at path
    filepath.Walk(path, w.walkFn)
    return
}

func (w *GoMonitor) walkFn(path string, f os.FileInfo, err error) error {
    if err != nil {
        log.Printf("%s\n", err)
        return err
    }
    if f.IsDir() {
        w.sourceDir = append(w.sourceDir, path)
    }
    // add filepath and modtime to map
    w.FileStatus[path] = f.ModTime().Unix()
    return nil
}

func (w *GoMonitor) PrintFile() {
    for file, time := range w.FileStatus {
        fmt.Printf("%s   %v \n", file, time)
    }
}
