package main

import (
    "flag"
    "fmt"
    "github.com/astaxie/beego/config"
    "github.com/gomonitor/util"
    "log"
)

const (
    ROOTDIR  = "Monitor::rootdir"
    WORKDIR  = "Monitor::workdir"
    INTERVAL = "Monitor::interval"
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
    // todo multi rootdir
    monitor.Interval, _ = iniconf.Int(INTERVAL)
    monitor.WorkDir = iniconf.String(WORKDIR)

    err = monitor.AddRootDir(rootdir)
    if err != nil {
        fmt.Printf("%s", err)
    }
}
