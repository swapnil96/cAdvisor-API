package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"encoding/json"
	// "github.com/golang/glog"
	"github.com/google/cadvisor/client"
	info "github.com/google/cadvisor/info/v1"
	"reflect"
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
func docker_all(url string){
	stat, err := client.NewClient(url)
	check(err)
	
	// Returns 60 numstats value by default
	query := info.DefaultContainerInfoRequest()
	containers, err := stat.AllDockerContainers(&query)
	check(err)

	file, err := os.Create("stats/docker.txt")
	check(err)
	defer file.Close()

	for _, container := range containers {
		// final := "Name - " + container.Aliases
		res, _ := json.MarshalIndent(container.Spec, "", "")
		// for _, data := range res{
		// 	fmt.Println(string(data))
		// }
		dec := json.NewDecoder(strings.NewReader(string(res)))
		
		var m info.ContainerSpec
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			check(err)
		}
		fmt.Printf("%t: %t\n", m.HasCpu, m.HasMemory)
		fmt.Println(m.Cpu.Limit)
	
		// fmt.Println(dec)


		fmt.Println(reflect.TypeOf(res))

		fmt.Println(string(res))

		_, err := file.WriteString("Name - " + container.Aliases[0])
		check(err)
		
		// fmt.Println(container.Aliases[0], container.Stats)

		
	}


	fmt.Println(len(containers));
	// fmt.Println(data[0].Name)	
	// res, _ := json.Marshal(data[0]) 
	// fmt.Println(string(res))
	// json.Unmarshal([]byte(res), &data)
	// fmt.Println(data)
	
	// fmt.Println(data)
}


// demonstrates how to use event clients
func main() {
	flag.Parse()
	docker_all("http://192.168.99.14:8080/")

}
