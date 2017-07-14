package utils

import (
	// "flag"
	// "fmt"
	"os"
	"encoding/json"
	// "github.com/golang/glog"
	"github.com/google/cadvisor/client"
	info "github.com/google/cadvisor/info/v1"
	"strconv"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

// -------------------------------------------------------------------------------------------------------------------------------------------
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

// -------------------------------------------------------------------------------------------------------------------------------------------
// type ContainerInfo struct {
// 	ContainerReference

// 	// The direct subcontainers of the current container.
// 	Subcontainers []ContainerReference `json:"subcontainers,omitempty"`

// 	// The isolation used in the container.
// 	Spec ContainerSpec `json:"spec,omitempty"`

// 	// Historical statistics gathered from the container.
// 	Stats []*ContainerStats `json:"stats,omitempty"`
// }

// -------------------------------------------------------------------------------------------------------------------------------------------
// type CpuSpec struct {
// 	Limit    uint64 `json:"limit"`
// 	MaxLimit uint64 `json:"max_limit"`
// 	Mask     string `json:"mask,omitempty"`
// 	Quota    uint64 `json:"quota,omitempty"`
// 	Period   uint64 `json:"period,omitempty"`
// }

// -------------------------------------------------------------------------------------------------------------------------------------------
// type MemorySpec struct {
// 	// The amount of memory requested. Default is unlimited (-1).
// 	// Units: bytes.
// 	Limit uint64 `json:"limit,omitempty"`

// 	// The amount of guaranteed memory.  Default is 0.
// 	// Units: bytes.
// 	Reservation uint64 `json:"reservation,omitempty"`

// 	// The amount of swap space requested. Default is unlimited (-1).
// 	// Units: bytes.
// 	SwapLimit uint64 `json:"swap_limit,omitempty"`
// }


// -------------------------------------------------------------------------------------------------------------------------------------------
// type ContainerStats struct {
// 	// The time of this stat point.
// 	Timestamp time.Time    `json:"timestamp"`
// 	Cpu       CpuStats     `json:"cpu,omitempty"`
// 	DiskIo    DiskIoStats  `json:"diskio,omitempty"`
// 	Memory    MemoryStats  `json:"memory,omitempty"`
// 	Network   NetworkStats `json:"network,omitempty"`

// 	// Filesystem statistics
// 	Filesystem []FsStats `json:"filesystem,omitempty"`

// 	// Task load stats
// 	TaskStats LoadStats `json:"task_stats,omitempty"`

// 	// Custom metrics from all collectors
// 	CustomMetrics map[string][]MetricVal `json:"custom_metrics,omitempty"`
// }

// -------------------------------------------------------------------------------------------------------------------------------------------
// type CpuStats struct {
// 	Usage CpuUsage `json:"usage"`
// 	CFS   CpuCFS   `json:"cfs"`
// 	// Smoothed average of number of runnable threads x 1000.
// 	// We multiply by thousand to avoid using floats, but preserving precision.
// 	// Load is smoothed over the last 10 seconds. Instantaneous value can be read
// 	// from LoadStats.NrRunning.
// 	LoadAverage int32 `json:"load_average"`
// }

// -------------------------------------------------------------------------------------------------------------------------------------------
// type MemoryStats struct {
// 	// Current memory usage, this includes all memory regardless of when it was
// 	// accessed.
// 	// Units: Bytes.
// 	Usage uint64 `json:"usage"`

// 	// Number of bytes of page cache memory.
// 	// Units: Bytes.
// 	Cache uint64 `json:"cache"`

// 	// The amount of anonymous and swap cache memory (includes transparent
// 	// hugepages).
// 	// Units: Bytes.
// 	RSS uint64 `json:"rss"`

// 	// The amount of swap currently used by the processes in this cgroup
// 	// Units: Bytes.
// 	Swap uint64 `json:"swap"`

// 	// The amount of working set memory, this includes recently accessed memory,
// 	// dirty memory, and kernel memory. Working set is <= "usage".
// 	// Units: Bytes.
// 	WorkingSet uint64 `json:"working_set"`

// 	Failcnt uint64 `json:"failcnt"`

// 	ContainerData    MemoryStatsMemoryData `json:"container_data,omitempty"`
// 	HierarchicalData MemoryStatsMemoryData `json:"hierarchical_data,omitempty"`
// }
// -------------------------------------------------------------------------------------------------------------------------------------------

// type CpuUsage struct {
// 	// Total CPU usage.
// 	// Unit: nanoseconds.
// 	Total uint64 `json:"total"`

// 	// Per CPU/core usage of the container.
// 	// Unit: nanoseconds.
// 	PerCpu []uint64 `json:"per_cpu_usage,omitempty"`

// 	// Time spent in user space.
// 	// Unit: nanoseconds.
// 	User uint64 `json:"user"`

// 	// Time spent in kernel space.
// 	// Unit: nanoseconds.
// 	System uint64 `json:"system"`
// }


func Host_spec(url string){
	root, err := client.NewClient(url)
	check(err)
	
	file, err := os.Create("stats/host_spec.txt")
	check(err)
	
	defer file.Close()

	mac_info, _ := root.MachineInfo()
	res, _ := json.MarshalIndent(mac_info, "", "\t") 

	_, _ = file.WriteString(string(res))
}

func Host_stat(url string, link string, num int){
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
	mem := "\n\n-------------------------------------------MEMORY-------------------------------------------\n"

		res, _ := json.MarshalIndent(mac_info.Spec.Cpu, "", "\t\t")
		cpu += string(res) + "\n"
		res, _ = json.MarshalIndent(mac_info.Spec.Memory, "", "\t\t")
		mem += string(res) + "\n"

		for _, history := range mac_info.Stats {
			res, _ = json.MarshalIndent(history.Cpu, "", "\t\t")
			cpu += history.Timestamp.String() + "\n" + string(res) + "\n"
			res, _ = json.MarshalIndent(history.Memory, "", "\t\t")
			mem += history.Timestamp.String() + "\n" + string(res) + "\n"
		}

	_, _ = file.WriteString(cpu)
	_, _ = file.WriteString(mem)
}

func Host_cpu(url string, link string, num int){
	root, err := client.NewClient(url)
	check(err)
	
	query := info.ContainerInfoRequest{
		NumStats: num,
	}

	file, err := os.Create("stats/host_cpu.dat")
	check(err)
	defer file.Close()

	mac_info, _ := root.ContainerInfo(link , &query)

	res :=  ""
	initial_time := mac_info.Stats[0].Timestamp
	initial_usage_total := mac_info.Stats[0].Cpu.Usage.Total
	initial_usage_core0 := mac_info.Stats[0].Cpu.Usage.PerCpu[0]
	initial_usage_core1 := mac_info.Stats[0].Cpu.Usage.PerCpu[1]
	initial_usage_core2 := mac_info.Stats[0].Cpu.Usage.PerCpu[2]
	initial_usage_core3 := mac_info.Stats[0].Cpu.Usage.PerCpu[3]

	for i, history := range mac_info.Stats {
		if i == 0{
			continue
		}
		temp := (float64(history.Timestamp.Sub(initial_time).Nanoseconds()) * float64(4))

		total := (float64(history.Cpu.Usage.Total - initial_usage_total) * float64(100))/ temp 
		core0 := (float64(history.Cpu.Usage.PerCpu[0] - initial_usage_core0) * float64(100))/ temp
		core1 := (float64(history.Cpu.Usage.PerCpu[1] - initial_usage_core1) * float64(100))/ temp
		core2 := (float64(history.Cpu.Usage.PerCpu[2] - initial_usage_core2) * float64(100))/ temp
		core3 := (float64(history.Cpu.Usage.PerCpu[3] - initial_usage_core3) * float64(100))/ temp

		res += strconv.FormatFloat(core0, 'f', -1, 64) + " " + strconv.FormatFloat(core1, 'f', -1, 64) + " " + strconv.FormatFloat(core2, 'f', -1, 64) + " " + strconv.FormatFloat(core3, 'f', -1, 64) + " " + strconv.FormatFloat(total, 'f', -1, 64) + "\n"
	}

	_, _ = file.WriteString(res)

}	

func Host_memory(url string, link string, num int){
	root, err := client.NewClient(url)
	check(err)
	
	query := info.ContainerInfoRequest{
		NumStats: num,
	}

	file, err := os.Create("stats/host_memory.dat")
	check(err)
	defer file.Close()

	mac_info, _ := root.ContainerInfo(link , &query)

	res :=  ""
	for _, history := range mac_info.Stats {		
		res += strconv.FormatUint(history.Memory.Usage, 10) + " " + strconv.FormatUint(history.Memory.Cache, 10) + "\n"
	}

	_, _ = file.WriteString(res)

}	
