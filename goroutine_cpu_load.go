// Go routine to capture CPU utilization for docker/kubernetes container

// cat imitates Unix cat
func cat(stream io.Reader) ([]byte, error) {
  data, err := ioutil.ReadAll(stream)

  if err != nil {
    panic(err)
    return nil, err
  }

  return data, nil
}

// int64frombytes converts bytes[] to int64
func int64frombytes(bytes []byte) int64 {
  s:= string(bytes[:len(bytes)-1])  // len-1 is needed to remove /n
  num, err := strconv.Atoi(s)
  if err != nil {
    panic(err)
  }
  return int64(num)
}

//readcpudata reads container cpu stats from container files
func readCPUStats() int64 {
  // filepath
  filepath := "/sys/fs/cgroup/cpu/cpuacct.usage"

  // open and read stats from file
  data, err := os.Open(filepath)
  if err != nil {
    panic(err)
  }

  output, _ := cat(data)    // calls cat func to immitate Unix cat
  CPUVal := int64frombytes(output)

  defer func() {
    err := data.Close()
    if err != nil {
      panic(err)
    }
  }()

  return CPUVal
}

// Goroutine CPULoadCalc computes CPU utilization over 5 seconds
func CPULoadCalc() {
  // infinte loop for goroutine
  for {
    previous := readCPUStats()
    start := time.Now().UnixNano()

    time.Sleep(5 * time.Second)

    after := readCPUStats()
    stop := time.Now().UnixNano()

    cpuLoad := float64(after-prev) / float64(stop-start)
    fmt.Println("CPU load percentage", cpuLoad * 100)
  }
}
