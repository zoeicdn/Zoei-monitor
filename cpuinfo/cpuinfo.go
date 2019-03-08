package cpuinfo

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// CPUInfo is the structure of the machine's cpu data.
type CPUInfo struct {
	Processor       int
	VendorID        string
	CpuFamily       int
	Model           int
	ModelName       string
	Stepping        int
	Microcode       string
	CpuMHz          float64
	CacheSize       int
	PhysicalId      int
	Siblings        int
	CoreID          int
	CpuCores        int
	Apicid          int
	InitialApicid   int
	Fpu             bool
	FpuException    bool
	CpuIdLevel      int
	WP              bool
	Flags           string
	Bugs            string
	Bogomips        float64
	ClFlushSize     int
	CacheAlignment  int
	AddressSizes    string
	PowerManagement string
}

// Get function returns the machine's CPU data.
func Get() (CPUInfo, error) {
	bytes, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return CPUInfo{}, err
	}

	cpuInfo, err := parse(bytes)

	return cpuInfo, err
}

func parse(bytes []byte) (cpuInfo CPUInfo, err error) {
	lines := strings.Split(string(bytes), "\n")

	for i, n := range lines {
		var data string
		if arr := strings.Split(n, ":"); len(arr) == 2 {
			data = strings.TrimLeft(arr[1], " ")
		}

		switch i {
		case 0:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.Processor = d
			break

		case 1:
			cpuInfo.VendorID = data
			break

		case 2:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.CpuFamily = d
			break

		case 3:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.Model = d
			break

		case 4:
			cpuInfo.ModelName = data
			break

		case 5:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.Stepping = d
			break

		case 6:
			cpuInfo.Microcode = data
			break

		case 7:
			d, err := strconv.ParseFloat(data, 64)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.CpuMHz = d
			break

		case 8:
			d, err := strconv.Atoi(strings.TrimRight(data, " KB"))
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.CacheSize = d
			break

		case 9:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.PhysicalId = d
			break

		case 10:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.Siblings = d
			break

		case 11:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.CoreID = d
			break

		case 12:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.CpuCores = d
			break

		case 13:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.Apicid = d
			break

		case 14:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.InitialApicid = d
			break

		case 15:
			if data == "yes" {
				cpuInfo.Fpu = true
			} else {
				cpuInfo.Fpu = false
			}
			break

		case 16:
			if data == "yes" {
				cpuInfo.FpuException = true
			} else {
				cpuInfo.FpuException = false
			}
			break

		case 17:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.CpuIdLevel = d
			break

		case 18:
			if data == "yes" {
				cpuInfo.WP = true
			} else {
				cpuInfo.WP = false
			}
			break

		case 19:
			cpuInfo.Flags = data
			break

		case 20:
			cpuInfo.Bugs = data
			break

		case 21:
			d, err := strconv.ParseFloat(data, 64)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.Bogomips = d
			break

		case 22:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.ClFlushSize = d
			break

		case 23:
			d, err := strconv.Atoi(data)
			if err != nil {
				return cpuInfo, err
			}

			cpuInfo.CacheAlignment = d
			break

		case 24:
			cpuInfo.AddressSizes = data
			break

		case 25:
			cpuInfo.PowerManagement = data
			break
		}
	}

	return
}
