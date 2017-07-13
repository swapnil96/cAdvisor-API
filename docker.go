package main

import (
	"flag"
	"fmt"
	"os"
	// "io"
	"encoding/json"
	// "github.com/golang/glog"
	"github.com/google/cadvisor/client"
	info "github.com/google/cadvisor/info/v1"
	"strings"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

// type ContainerInfo struct {
// 	ContainerReference

// 	// The direct subcontainers of the current container.
// 	Subcontainers []ContainerReference `json:"subcontainers,omitempty"`

// 	// The isolation used in the container.
// 	Spec ContainerSpec `json:"spec,omitempty"`

// 	// Historical statistics gathered from the container.
// 	Stats []*ContainerStats `json:"stats,omitempty"`
// }

// func docker(name string, url string){
// 	stat, err := client.NewClient(url)
// 	if err != nil{
// 		fmt.Println("Error\n")
// 	}

// }

// type ContainerReference struct {
// 	// The container id
// 	Id string `json:"id,omitempty"`

// 	// The absolute name of the container. This is unique on the machine.
// 	Name string `json:"name"`

// 	// Other names by which the container is known within a certain namespace.
// 	// This is unique within that namespace.
// 	Aliases []string `json:"aliases,omitempty"`

// 	// Namespace under which the aliases of a container are unique.
// 	// An example of a namespace is "docker" for Docker containers.
// 	Namespace string `json:"namespace,omitempty"`

// 	Labels map[string]string `json:"labels,omitempty"`
// }


// link - http://192.168.99.14:8080/api/v1.3/docker/
func docker_all(url string, num int){
	root, err := client.NewClient(url)
	check(err)
	
	// Returns 60 numstats value by default
	// query := info.DefaultContainerInfoRequest()
	query := info.ContainerInfoRequest{
		NumStats: num,
	}
	containers, err := root.AllDockerContainers(&query)
	check(err)

	file, err := os.Create("stats/docker.txt")
	check(err)
	defer file.Close()
	cpu := ""
	mem := ""
	for _, container := range containers {
		cpu += "Name - " + container.Aliases[0] + ", Image - " + container.Spec.Image +  "\n\tCPU -------------\n"
		mem += "\tMEMORY --------------\n"
		res, err := json.Marshal(container.Spec)
		check(err)

		decode := json.NewDecoder(strings.NewReader(string(res)))
		var spec info.ContainerSpec
		_ = decode.Decode(&spec)
			
			res, _ = json.MarshalIndent(container.Spec.Cpu, "", "\t\t")
			cpu += string(res) + "\n"
			res, _ = json.MarshalIndent(container.Spec.Memory, "", "\t\t")
			mem += string(res) + "\n"


		for _, history := range container.Stats {

			res, _ = json.Marshal(history)

			decode := json.NewDecoder(strings.NewReader(string(res)))
			var stat info.ContainerStats
			_ = decode.Decode(&stat)
			
			res, _ = json.MarshalIndent(history.Cpu, "", "\t\t")
			cpu += history.Timestamp.String() + "\n" + string(res) + "\n"
			res, _ = json.MarshalIndent(history.Memory, "", "\t\t")
			mem += history.Timestamp.String() + "\n" + string(res) + "\n"
		}
		cpu += "\n"
		fmt.Println("dsf")
	}
	_, _ = file.WriteString(cpu)
	_, _ = file.WriteString(mem)
		
}


// demonstrates how to use event clients
func main() {
	flag.Parse()
	docker_all("http://192.168.99.14:8080/", 10)

}
