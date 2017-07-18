package utils

import "os/exec"

func Swarm() {
	cmd := exec.Command("sh", "-c", "docker node ls")
	stdoutStderr, err := cmd.CombinedOutput()
	check_error(err)
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
		if node != "" {
			cmd = exec.Command("sh", "-c", "docker inspect --format='{{.Status.Addr}}' " + node)
			stdoutStderr, err := cmd.CombinedOutput()
			check_error(err)
			Host_cpu("http://" + string(stdoutStderr)[:len(stdoutStderr)-1] + ":8080/", "", 11)
		}
	}
}
