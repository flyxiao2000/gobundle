package main

import (
	"flag"
	"fmt"

	"github.com/koding/multiconfig"
)

type Server struct {
	Name    string `required:"true"`
	Port    int    `default:"6060"`
	Enabled bool
	Users   []string
}

func main() {
	var p string
	flag.StringVar(&p, "config", "config.toml", "set `prefix` path")
	flag.Parse()

	fmt.Println(p)
	m := multiconfig.NewWithPath(p) // supports TOML, JSON and YAML

	// Get an empty struct for your configuration
	serverConf := new(Server)

	// Populated the serverConf struct
	if err := m.Load(serverConf); err != nil { // Check for error
		fmt.Println(err)
	}
	m.MustLoad(serverConf) // Panic's if there is any error
	fmt.Println(serverConf)
	fmt.Println("vim-go")
}
