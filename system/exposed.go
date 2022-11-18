package system

import (
	"github.com/jaypipes/ghw"
	gInfo "github.com/matishsiao/goInfo"
)

// Gontains System Information
var SYSINFO gInfo.GoInfoObject

// Contains System Information
type SystemInfo struct {
	Hostname     string
	OS           string
	Arch         string
	Kernel       string
	Uptime       string
	CPU          CPUInfo
	Memory       MemoryInfo
	Network      NetworkInfo
	UserName     string
	Docker       string
	GoVersion    string
	User         string
	UserHome     string
	CpuInfo      *ghw.CPUInfo
	MemoryInfo   *ghw.MemoryInfo
	BlockInfo    *ghw.BlockInfo
	NetworkInfo  *ghw.NetworkInfo
	TopologyInfo *ghw.TopologyInfo
	PCIInfo      *ghw.PCIInfo
	BIOSInfo     *ghw.BIOSInfo
	ChassisInfo  *ghw.ChassisInfo
	GPUInfo      *ghw.GPUInfo
}

// Embodies CPU Information
type CPUInfo struct {
	NoCPUs int
	CPUs   string
}

// Contains Memory Information
type MemoryInfo struct {
	Total            uint64
	Free             uint64
	Used             uint64
	UsedPercent      int64
	HumanTotal       string
	HumanFree        string
	HumanUsed        string
	HumanUsedPercent string
}

// Contains Network Information
type NetworkInfo struct {
	IP   string
	MAC  string
	Name string
}

// Contains Storage Information
type DiskInfo struct {
	Total uint64
	Free  uint64
	Used  uint64
}

// Gets the current system information from various sources
func Get() SystemInfo {
	return get()
}

func IsRunningInDockerContainer() bool {
	return isRunningInDockerContainer()
}
