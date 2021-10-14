package main

import (
	"flag"
	"fmt"
)

var (
	verbose bool
	port    int
)

func main() {

	flag.BoolVar(&verbose, "verbose", false, "是否输出日志")
	flag.IntVar(&port, "port", 7878, "运行端口号")
	flag.Parse()

	if err := initDB(verbose); err != nil {
		panic(err)
	}

	server := createServer(verbose)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf(`
    =======================
    Listening on port %s
    =======================

`, addr)
	panic(server.Listen(addr))
}
