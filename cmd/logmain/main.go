package main

import (
	"flag"
	"time"

	"github.com/flyxiao2000/gobundle/pkg/glog"
)

func main() {
	flag.Parse()
	glog.Warning("hello warning")
	glog.V(3).Info("v info ")
	for {
		time.Sleep(2 * time.Second)
		glog.Warning("hello warning")
		glog.Info("hello info")
	}
	defer glog.Flush()
}
