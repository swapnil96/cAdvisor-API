# cAdvisor-API
Metrics of docker swarm by using google cAdvisor API

To find CPU usage % the following formula is used 

    Usage % = (Used CPU Time (in nanoseconds) for the interval) /(interval (in nano secs) * num cores)

## Install
System should have **go** installed.

To use as a go package, install this package by - 
    
    go get github.com/swapnil96/cAdvisor-API

This will make the package available for use.Then import this package by - 

    import "github.com/swapnil96/cAdvisor-API/utils"

Then you are ready to go !!

## Structure of the Repo
**UTILS** folder has the scripts for polling metrics 


**host.go** gives the metrics of the host.
1) **Host_spec** function gives the specifications of the host.
2) **Host_stat** function gives the statistics of the host in a human readable form.
3) **Host_cpu** function gives cpu usage statistics in machine readable format.
4) **Host_memory** function gives memory usage statistics in machine readable format.

**docker.go** gives the metrics of all docker containers.
1) **Docker_stat** function gives the statistics of all docker containers in a human readable form.
2) **Docker_cpu** function gives cpu usage of all docker containers in machine readable format.
3) **Dost_memory** function gives memory usage of all docker containers in machine readable format.

**swarm.go** gives the metrics for all the nodes in the swarm.
1) **Swarm** function gives the statistics of all the nodes that are in the swarm.

## Usage
**main.go** has sample way how to use the scripts in the utils folder

All the metrics will be stored inside the stats folder.
