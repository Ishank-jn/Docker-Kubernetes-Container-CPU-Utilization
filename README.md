# Docker-Kubernetes-Container-CPU-Utilization

Implementing CPU Load goroutine requires the user to call the goroutine from the _main_ file.
>> go CPULoadCalc()

For multi-cpu env divide cpuLoad by cpu count.

Same code can be modified to calculate _memory_ for docker/kubernetes container.
Path: /sys/fs/cgroup/memory/memory.usage_in_bytes
refer: https://www.kernel.org/doc/Documentation/cgroup-v1/memory.txt


