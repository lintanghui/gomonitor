package main

import (
    "flag"
    "fmt"
    "github.com/astaxie/beego/config"
    "github.com/gomonitor/util"
    "log"
)

const (
    ROOTDIR = "Monitor::rootdir"
)

var monitor = util.DefMonitor

func main() {
    var c string
    flag.StringVar(&c, "c", "monitor.conf", "Usage: gomonitor -c=./monitor.conf")
    flag.Parse()

    parse(c)

    // monitor.PrintFile()
}

func parse(conf string) {
    iniconf, err := config.NewConfig("ini", conf)
    if err != nil {
        log.Panic(err)
    }
    rootdir := iniconf.String(ROOTDIR)
    err = monitor.AddRootDir(rootdir)
    if err != nil {
        fmt.Printf("%s", err)
    }
}
