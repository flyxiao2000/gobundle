package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

var (
	count                int
	startindex, getindex int
	set, get             bool
)

func init() {
	flag.IntVar(&count, "count", 0, "total record")
	flag.IntVar(&startindex, "startindex", 0, "start index")
	flag.IntVar(&getindex, "getindex", 0, "get index")

	flag.BoolVar(&set, "set", false, "add data to hbase")
	flag.BoolVar(&get, "get", false, "get data from hbase")
}

func gethbase() {
	if !get {
		return
	}
	client := gohbase.NewClient("192.168.151.246,192.168.151.247,192.168.151.248")

	str := fmt.Sprintf("%d", getindex)
	a := md5.Sum([]byte(str))

	getRequest, err := hrpc.NewGetStr(context.Background(), "largetable", hex.EncodeToString(a[:]))
	if err == nil {
		if getRsp, err := client.Get(getRequest); err == nil {
			fmt.Println("----------------------------------------------")
			fmt.Println("rowkey--------------Family----------Qualifier----------Value")
			for _, v := range getRsp.Cells {
				fmt.Println(string(v.Row), "--", string(v.Family), "--", string(v.Qualifier), "--", string(v.Value))
			}
		} else {
			fmt.Println("getRsp ", err)
		}
	}
}

func sethbase() {
	if !set {
		return
	}
	client := gohbase.NewClient("192.168.151.246,192.168.151.247,192.168.151.248")
	total := startindex + count

	fmt.Println("startindex ", startindex, " total ", total)

	for i := startindex; i < total; i++ {
		str := fmt.Sprintf("%d", i)
		a := md5.Sum([]byte(str))

		values := map[string]map[string][]byte{"cf1": make(map[string][]byte, 0)}

		for j := 0; j < 10; j++ {
			values["cf1"][fmt.Sprintf("column%d", j)] = []byte("abcdefghlm0123456789")
		}
		putRequest, err := hrpc.NewPutStr(context.Background(), "largetable", hex.EncodeToString(a[:]), values)
		if err == nil {
			if _, err := client.Put(putRequest); err != nil {
				fmt.Println("error pub ", err)
			}
		} else {
			fmt.Println("error putRequest ", err)
			return
		}

		if i%250 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	flag.Parse()

	sethbase()
	gethbase()
}

func ref() {
	client := gohbase.NewClient("192.168.151.246,192.168.151.247,192.168.151.248")

	values := map[string]map[string][]byte{"cf1": map[string][]byte{"b": []byte("nnnnnnn")}}
	putRequest, err := hrpc.NewPutStr(context.Background(), "largetable", "aaakey", values)
	if err == nil {
		if _, err := client.Put(putRequest); err != nil {
			fmt.Println("pub ", err)
		}
	} else {
		fmt.Println("putRequest ", err)
	}
	getRequest, err := hrpc.NewGetStr(context.Background(), "largetable", "aaakey")
	if err == nil {
		if getRsp, err := client.Get(getRequest); err == nil {
			fmt.Println("----------------------------------------------")
			for _, v := range getRsp.Cells {
				fmt.Println(string(v.Family))
				fmt.Println(string(v.Row))
				fmt.Println(string(v.Value))
			}
		} else {
			fmt.Println("getRsp ", err)
		}
	}
	/*
		family := map[string][]string{"aaakey": []string{"a"}}
			getRequest, err = hrpc.NewGetStr(context.Background(), "aaatable", "aaakey", hrpc.Families(family))
			if err == nil {
				if getRsp, err := client.Get(getRequest); err == nil {
					fmt.Println("----------------------------------------------")
					for _, v := range getRsp.Cells {
						fmt.Println(v)
						fmt.Println(string(v.Value))
					}
				} else {
					fmt.Println("getRsp ", err)
				}
			}
	*/

}
