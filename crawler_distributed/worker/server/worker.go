package main

import (
	"flag"
	"fmt"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
