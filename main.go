package main

import (
	"github.com/swapnil96/cAdvisor-API/utils"
	"flag"
	"fmt"
	"log"
)

func check(err error){
        if err != nil {
                log.Fatal(err)
        }
}

// demonstrates how to use event clients
func main() {
	flag.Parse()
	fmt.Println()
	utils.Host_spec("http://192.168.99.14:8080/")
	utils.Host_stat("http://192.168.99.14:8080/", "", 10)
	utils.Host_cpu("http://192.168.99.14:8080/", "", 11)
	utils.Host_memory("http://192.168.99.14:8080/", "", 11)
	utils.Swarm()
}
