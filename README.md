# cAdvisor-API
Metrics of docker swarm by using google cAdvisor API

## Structure of the Repo
**host.go** gives the metrics of the host
	**host_spec** function gives the specifications of the host.
	**host_stat** function gives the statistics of the host in a human readable form.
	**host_cpu** function gives cpu usage statistics in machine readable format.
	**host_memory** function gives memory usage statistics in machine readable format.

**docker.go** gives the metrics of all docker containers
        **docker_stat** function gives the statistics of all docker containers in a human readable form.
        **docker_cpu** function gives cpu usage of all docker containers in machine readable format.
        **host_memory** function gives memory usage of all docker containers in machine readable format.

