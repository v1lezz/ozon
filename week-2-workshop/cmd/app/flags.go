package main

import (
	"flag"
	"fmt"
	"gitlab.ozon.dev/14/week-2-workshop/internal/app"
	"os"
)

const (
	defaultAddr        = ":8082"
	defaultProductAddr = "http://route256.pavl.uk:8080"

	envToken = "TOKEN"
)

var opts = app.Options{}

func initOpts() {
	flag.StringVar(&opts.Addr, "addr", defaultAddr, fmt.Sprintf("server address, default: %q", defaultAddr))
	flag.StringVar(&opts.ProductAddr, "product_addr", defaultProductAddr, fmt.Sprintf("products-service address, default: %q", defaultProductAddr))
	flag.Parse()

	opts.ProductToken = os.Getenv(envToken)
}
