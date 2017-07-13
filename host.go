package main

import (
	"flag"
	"fmt"
	"os"
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

// type MachineInfo struct {
// 	// The number of cores in this machine.
// 	NumCores int `json:"num_cores"`

// 	// Maximum clock speed for the cores, in KHz.
// 	CpuFrequency uint64 `json:"cpu_frequency_khz"`

// 	// The amount of memory (in bytes) in this machine
// 	MemoryCapacity uint64 `json:"memory_capacity"`

// 	// The machine id
// 	MachineID string `json:"machine_id"`

// 	// The system uuid
// 	SystemUUID string `json:"system_uuid"`

// 	// The boot id
// 	BootID string `json:"boot_id"`

// 	// Filesystems on this machine.
// 	Filesystems []FsInfo `json:"filesystems"`

// 	// Disk map
// 	DiskMap map[string]DiskInfo `json:"disk_map"`

// 	// Network devices
// 	NetworkDevices []NetInfo `json:"network_devices"`

// 	// Machine Topology
// 	// Describes cpu/memory layout and hierarchy.
// 	Topology []Node `json:"topology"`

// 	// Cloud provider the machine belongs to.
// 	CloudProvider CloudProvider `json:"cloud_provider"`

// 	// Type of cloud instance (e.g. GCE standard) the machine is.
// 	InstanceType InstanceType `json:"instance_type"`

// 	// ID of cloud instance (e.g. instance-1) given to it by the cloud provider.
// 	InstanceID InstanceID `json:"instance_id"`
// }

// type ContainerInfo struct {
// 	ContainerReference

// 	// The direct subcontainers of the current container.
// 	Subcontainers []ContainerReference `json:"subcontainers,omitempty"`

// 	// The isolation used in the container.
// 	Spec ContainerSpec `json:"spec,omitempty"`

// 	// Historical statistics gathered from the container.
// 	Stats []*ContainerStats `json:"stats,omitempty"`
// }


func host_spec(url string){
	root, err := client.NewClient(url)
	check(err)
	
	file, err := os.Create("stats/host_spec.txt")
	check(err)
	
	defer file.Close()

	mac_info, _ := root.MachineInfo()
	res, _ := json.MarshalIndent(mac_info, "", "\t") 

	_, _ = file.WriteString(string(res))
}

func host_stat(url string, link string, num int){
	root, err := client.NewClient(url)
	check(err)
	
	query := info.ContainerInfoRequest{
		NumStats: num,
	}

	file, err := os.Create("stats/host_stat.txt")
	check(err)
	
	defer file.Close()

	mac_info, _ := root.ContainerInfo(link , &query)

	cpu := "Name - " + mac_info.Name + ", Image - " + mac_info.Spec.Image +  "\n-------------------------------------------CPU-------------------------------------------\n"
	mem := "\n\n-------------------------------------------MEMORY-------------------------------------------"
	res, err := json.Marshal(mac_info.Spec)
	check(err)

	decode := json.NewDecoder(strings.NewReader(string(res)))
	var spec info.ContainerSpec
	_ = decode.Decode(&spec)
		
		res, _ = json.MarshalIndent(mac_info.Spec.Cpu, "", "\t\t")
		cpu += string(res) + "\n"
		res, _ = json.MarshalIndent(mac_info.Spec.Memory, "", "\t\t")
		mem += string(res) + "\n"

		for _, history := range mac_info.Stats {

			res, _ = json.Marshal(history)

			decode := json.NewDecoder(strings.NewReader(string(res)))
			var stat info.ContainerStats
			_ = decode.Decode(&stat)
			
			res, _ = json.MarshalIndent(history.Cpu, "", "\t\t")
			cpu += history.Timestamp.String() + "\n" + string(res) + "\n"
			res, _ = json.MarshalIndent(history.Memory, "", "\t\t")
			mem += history.Timestamp.String() + "\n" + string(res) + "\n"
		}

	_, _ = file.WriteString(cpu)
	_, _ = file.WriteString(mem)
}

// demonstrates how to use event clients
func main() {
	flag.Parse()
	fmt.Println()
	host_spec("http://192.168.99.14:8080/")
	host_stat("http://192.168.99.14:8080/", "", 10)
}
