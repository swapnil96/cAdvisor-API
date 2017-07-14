# cAdvisor-API
Metrics of docker swarm by using google cAdvisor API

## Structure of the Repo
**UTILS** folder has the scripts for polling metrics 


**host.go** gives the metrics of the host.
1) **Host_spec** function gives the specifications of the host.
2) **Host_stat** function gives the statistics of the host in a human readable form.
3) **Host_cpu** function gives cpu usage statistics in machine readable format.
4) **Host_memory** function gives memory usage statistics in machine readable format.

**docker.go** gives the metrics of all docker containers
1) **Docker_stat** function gives the statistics of all docker containers in a human readable form.
2) **Docker_cpu** function gives cpu usage of all docker containers in machine readable format.
3) **Dost_memory** function gives memory usage of all docker containers in machine readable format.

## Usage
**main.go** has sample way how to use the scripts in the utils folder

All the metrics will be stored inside the stats folder.
