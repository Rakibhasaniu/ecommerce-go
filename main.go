package main

import (
	"fmt"

	"ecommerec/cmd"
	"ecommerec/config"
)

func main() {
	cnf := config.GetConfig()
	fmt.Println(cnf.HttpPort)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.Version)
	cmd.Serve()
}
