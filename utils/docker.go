package utils

import (
	"os"
	"encoding/json"
	// "github.com/golang/glog"
	"github.com/google/cadvisor/client"
	info "github.com/google/cadvisor/info/v1"
	"strconv"
)

func Docker_stat(url string, num int){
	root, err := client.NewClient(url)
	check_error(err)
	
	query := info.ContainerInfoRequest{
		NumStats: num,
	}

	check_dir()
	file, err := os.Create("stats/docker_stat.txt")
	check_error(err)
	
	defer file.Close()

	docker_info, _ := root.AllDockerContainers(&query)
	cpu := ""
	mem := ""
	for _, container := range docker_info{

		cpu = "\n\nName - " + container.Aliases[0] + ", Image - " + container.Spec.Image +  "\n-------------------------------------------CPU-------------------------------------------\n"
		mem = "\n\n-------------------------------------------MEMORY-------------------------------------------\n"

		res, _ := json.MarshalIndent(container.Spec.Cpu, "", "\t\t")
		cpu += string(res) + "\n"
		res, _ = json.MarshalIndent(container.Spec.Memory, "", "\t\t")
		mem += string(res) + "\n"

		for _, history := range container.Stats {
			
			res, _ = json.MarshalIndent(history.Cpu, "", "\t\t")
			cpu += history.Timestamp.String() + "\n" + string(res) + "\n"
			res, _ = json.MarshalIndent(history.Memory, "", "\t\t")
			mem += history.Timestamp.String() + "\n" + string(res) + "\n"
		}

		_, _ = file.WriteString(cpu)
		_, _ = file.WriteString(mem)

	}
}

func Docker_cpu(url string, num int){
	root, err := client.NewClient(url)
	check_error(err)
	
	query := info.ContainerInfoRequest{
		NumStats: num,
	}

	check_dir()
	file, err := os.Create("stats/docker_cpu.dat")
	check_error(err)
	defer file.Close()

	docker_info, _ := root.AllDockerContainers(&query)
	res :=  ""
	
	for _, container := range docker_info{
		
		res = container.Aliases[0] + "\n\n\n"
		initial_time := container.Stats[0].Timestamp
		initial_usage_total := container.Stats[0].Cpu.Usage.Total
		initial_usage_core0 := container.Stats[0].Cpu.Usage.PerCpu[0]
		initial_usage_core1 := container.Stats[0].Cpu.Usage.PerCpu[1]
		initial_usage_core2 := container.Stats[0].Cpu.Usage.PerCpu[2]
		initial_usage_core3 := container.Stats[0].Cpu.Usage.PerCpu[3]

		for i, history := range container.Stats {
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
		
		res += "-----------------------------------------------------------------------------------------------------\n"
		_, _ = file.WriteString(res)
	}
}	

func Docker_memory(url string, num int){
	root, err := client.NewClient(url)
	check_error(err)
	
	query := info.ContainerInfoRequest{
		NumStats: num,
	}

	check_dir()
	file, err := os.Create("stats/docker_memory.dat")
	check_error(err)
	defer file.Close()

	docker_info, _ := root.AllDockerContainers(&query)
	res :=  ""
	
	for _, container := range docker_info{
		res = container.Aliases[0] + "\n\n"
		for _, history := range container.Stats {

			res += strconv.FormatUint(history.Memory.Usage, 10) + " " + strconv.FormatUint(history.Memory.Cache, 10) + "\n"
		}
		
		res += "-----------------------------------------------------------------------------------------------------\n\n"
		_, _ = file.WriteString(res)
	}

}	
