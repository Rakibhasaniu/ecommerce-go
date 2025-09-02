package cmd

import (
	"ecommerec/config"
	"ecommerec/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)

}
