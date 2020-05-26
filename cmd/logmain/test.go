package main

import "github.com/flyxiao2000/gobundle/pkg/glog"

func printlog() {
	glog.V(3).Info("v info3 ")
	glog.V(2).Info("v info2 ")
	glog.V(1).Info("v info1 ")

}
