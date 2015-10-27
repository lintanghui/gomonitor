package main

import (
    "flag"
    "fmt"
    "github.com/astaxie/beego/config"
    "github.com/gomonitor/util"
    "log"
    "strings"
)

const (
    ROOTDIR  = "Monitor::rootdir"
    WORKDIR  = "Monitor::workdir"
    INTERVAL = "Monitor::interval"
    BUILDCMD = "Monitor::buildcmd"
    RUNCMD   = "Monitor::runcmd"
)

var monitor = util.DefMonitor

func main() {
    var c string
    flag.StringVar(&c, "c", "monitor.conf", "Usage: gomonitor -c=./monitor.conf")
    flag.Parse()

    parse(c)
    monitor.Monitor()

}

func parse(conf string) {
    iniconf, err := config.NewConfig("ini", conf)
    if err != nil {
        log.Panic(err)
    }
    rootdir := iniconf.String(ROOTDIR)
    monitorDir := strings.Split(rootdir, ";")
    // todo multi rootdir
    monitor.Interval, _ = iniconf.Int(INTERVAL)
    monitor.WorkDir = iniconf.String(WORKDIR)
    monitor.RunCmd = iniconf.String(RUNCMD)
    monitor.BuildCmd = iniconf.String(BUILDCMD)
    log.Printf("BuildCmd :%s\n", monitor.BuildCmd)
    log.Printf("RunCmd :%s\n", monitor.RunCmd)

    for _, dir := range monitorDir {
        err = monitor.AddRootDir(dir)
        if err != nil {
            fmt.Printf("%s", err)
        }
    }

}
