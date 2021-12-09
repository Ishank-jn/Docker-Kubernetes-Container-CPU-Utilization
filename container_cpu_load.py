# readcpudata reads container cpu stats from container files
def readCPUStats():
  # filepath
  filepath = "/sys/fs/cgroup/cpu/cpuacct.usage"

  if os.path.exists(filepath):
   with open(filepath, 'rb') as tobj:
     CPUVal = tobj.read()

  return CPUVal


# CPULoadCalc computes CPU utilization over 5 seconds
func CPULoadCalc() {
  // infinte loop for goroutine
  interval = 5
  for {
    previous = readCPUStats()

    time.sleep(interval)

    after = readCPUStats()

    cpuLoad := float64(after-prev) / (5 * 1000000000)
    fmt.Println("CPU load percentage", cpuLoad * 100)
  }
}
