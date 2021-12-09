# Docker-Kubernetes-Container-CPU-Utilization

Implementing CPU Load goroutine requires the user to call the goroutine from the _main_ file.
>> go CPULoadCalc()
>> for multi-cpu env divide cpuLoad by cpu count.

Giving credit where it is due:
https://unix.stackexchange.com/questions/450748/calculating-cpu-usage-of-a-cgroup-over-a-period-of-time
I have written both Python and Go implementation of the unix shell script shared in the above link
