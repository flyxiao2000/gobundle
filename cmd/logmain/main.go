package main

import (
	"flag"
	"time"

	"github.com/flyxiao2000/gobundle/pkg/glog"
)

func main() {
	flag.Parse()
	glog.Warning("hello warning")
	printlog()

	defer glog.Flush()
	for {
		time.Sleep(2 * time.Second)
		glog.Warning("hello warning")
		glog.Info("hello info")
	}

}
