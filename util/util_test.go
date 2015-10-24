package util

import (
    "testing"
)

func TestAddRootDir(t *testing.T) {
    err := DefMonitor.AddRootDir("./")
    if err != nil {
        t.Log(err)
    }

}
