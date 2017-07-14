# cAdvisor-API
Metrics of docker swarm by using google cAdvisor API

## Structure of the Repo
**host.go** gives the metrics of the host.
1) **host_spec** function gives the specifications of the host.
2) **host_stat** function gives the statistics of the host in a human readable form.
3) **host_cpu** function gives cpu usage statistics in machine readable format.
4) **host_memory** function gives memory usage statistics in machine readable format.

**docker.go** gives the metrics of all docker containers
1) **docker_stat** function gives the statistics of all docker containers in a human readable form.
2) **docker_cpu** function gives cpu usage of all docker containers in machine readable format.
3) **host_memory** function gives memory usage of all docker containers in machine readable format.

