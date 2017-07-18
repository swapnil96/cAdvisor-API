package main

import (
	"github.com/swapnil96/cAdvisor-API/utils"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func check(err error){
        if err != nil {
                log.Fatal(err)
        }
}

func run() {
	cmd := exec.Command("sh", "-c", "docker node ls")
	stdoutStderr, err := cmd.CombinedOutput()
	check(err)
	count := 0
	flag := 0
	var ids [10]string
	temp := ""
	idx := 0
	for _, b := range stdoutStderr{
		if b == ' ' || b == '\t'{
				if flag == 1{
								ids[idx] = temp
								idx++
				}
				flag = 0
				temp = ""
		}
		if count != 0 && flag == 1{
				temp += string(b)
		}
		if b == '\n'{
				count++
				flag = 1
		}
	}
	for _, node := range ids{
		cmd = exec.Command("sh", "-c", "docker inspect --format='{{.Status.Addr}}' " + node)
		stdoutStderr, err := cmd.CombinedOutput()
		check(err)
		utils.Host_cpu("http://" + string(stdoutStderr) + "/:8080/", "", 11)
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
	// run()
}
