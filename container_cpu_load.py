# readcpudata reads container cpu stats from container files
def readCPUStats():
  # filepath
  filepath = "/sys/fs/cgroup/cpu/cpuacct.usage"

  if os.path.exists(filepath):
   with open(filepath, 'rb') as tobj:
     CPUVal = tobj.read()
     return CPUVal
  # if file not exists
  return 0


# CPULoadCalc computes CPU utilization over 5 seconds
def CPULoadCalc():
  // infinte loop
  interval = 5
  while True:
    previous = readCPUStats()
    start = time.time_ns()

    time.sleep(interval)

    after = readCPUStats()
    stop = time.time_ns()

    cpuLoad = float64((after-prev) / (stop-start))
    print("CPU load percentage", cpuLoad * 100) 
