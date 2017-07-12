package main

import (
	"flag"
	"fmt"
	"encoding/json"
	// "github.com/golang/glog"
	"github.com/google/cadvisor/client"
	// info "github.com/google/cadvisor/info/v1"

)

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

func machine(url string){
	stat, err := client.NewClient(url)
	if err != nil{
		fmt.Println("Error\n")
	}
	MI, err := stat.MachineInfo()
	if err != nil{
		fmt.Println("Error\n")
	}
	res, _ := json.Marshal(MI) 
	// fmt.Println(stat.MachineInfo())
	fmt.Println(string(res))
}


// demonstrates how to use event clients
func main() {
	flag.Parse()
	machine("http://192.168.99.14:8080/")

}
