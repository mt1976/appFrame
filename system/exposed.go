package system

import (
	"github.com/jaypipes/ghw"
	gInfo "github.com/matishsiao/goInfo"
)

// Gontains System Information
var SYSINFO gInfo.GoInfoObject

// Contains System Information
// The SystemInfo struct contains information about the system's hostname, operating system,
// architecture, kernel, uptime, CPU, memory, network, user, Docker, Go version, and various hardware
// information.
// @property {string} Hostname - The Hostname property represents the name of the system or machine. It
// is typically a unique identifier for the system within a network.
// @property {string} OS - The "OS" property in the SystemInfo struct represents the operating system
// of the system. It could be a string value indicating the name or version of the operating system,
// such as "Windows 10" or "Ubuntu 20.04".
// @property {string} Arch - The "Arch" property in the SystemInfo struct represents the architecture
// of the system. It typically indicates whether the system is running on a 32-bit or 64-bit
// architecture.
// @property {string} Kernel - The "Kernel" property refers to the operating system kernel that the
// system is running on. The kernel is the core component of the operating system that manages system
// resources and provides an interface for applications to interact with the hardware. Examples of
// popular kernels include Linux, Windows NT kernel, and macOS kernel.
// @property {string} Uptime - The "Uptime" property represents the amount of time that the system has
// been running since it was last booted. It is typically represented in a human-readable format, such
// as "1 day 2 hours 30 minutes".
// @property {CPUInfo} CPU - The CPU property represents information about the central processing unit
// of the system. It may include details such as the number of cores, clock speed, cache size, and
// other relevant information.
// @property {MemoryInfo} Memory - The `Memory` property in the `SystemInfo` struct represents
// information about the system's memory. It is of type `MemoryInfo`.
// @property {NetworkInfo} Network - The `Network` property in the `SystemInfo` struct represents
// information about the network configuration of the system. It may include details such as the
// network interfaces, IP addresses, MAC addresses, and network statistics.
// @property {string} UserName - The `UserName` property represents the username of the current user
// logged into the system.
// @property {string} Docker - The "Docker" property in the SystemInfo struct represents information
// about the Docker installation on the system. It could include details such as the Docker version,
// configuration, and any relevant information about Docker containers and images.
// @property {string} GoVersion - The GoVersion property represents the version of the Go programming
// language that is installed on the system. It can be used to determine the compatibility of the code
// with the installed Go version.
// @property {string} User - The "User" property in the SystemInfo struct represents the username of
// the current user logged into the system.
// @property {string} UserHome - The UserHome property represents the home directory of the current
// user. It is a string that stores the path to the user's home directory.
// @property CpuInfo - CpuInfo is a pointer to a struct that contains information about the CPU, such
// as the number of cores, clock speed, and cache size. It is part of the ghw package, which is a Go
// library for retrieving system hardware information.
// @property MemoryInfo - The MemoryInfo property is a struct that contains information about the
// system's memory. It may include details such as the total amount of memory, the amount of memory
// currently in use, and the memory speed.
// @property BlockInfo - The `BlockInfo` property is a pointer to a `ghw.BlockInfo` struct. This struct
// contains information about the block devices (such as hard drives and SSDs) on the system, including
// their names, sizes, and partitions.
// @property NetworkInfo - The NetworkInfo property is a struct that contains information about the
// network interfaces on the system. It may include details such as the interface name, MAC address, IP
// address, subnet mask, and other network-related information.
// @property TopologyInfo - The TopologyInfo property in the SystemInfo struct represents information
// about the system's hardware topology. This includes details about the system's processors, caches,
// and NUMA nodes. It provides information such as the number of processors, their IDs, the number of
// cores and threads per processor, and the cache
// @property PCIInfo - PCIInfo is a property of the SystemInfo struct that represents information about
// the PCI devices on the system. It provides details such as the PCI device ID, vendor ID, subsystem
// ID, and class information. This information can be useful for identifying and managing PCI devices
// on the system.
// @property BIOSInfo - The BIOSInfo property contains information about the BIOS (Basic Input/Output
// System) of the system. This includes details such as the vendor, version, release date, and features
// supported by the BIOS.
// @property ChassisInfo - ChassisInfo is a property of the SystemInfo struct that represents
// information about the chassis or enclosure of the system. This includes details such as the
// manufacturer, version, and serial number of the chassis.
// @property GPUInfo - The `GPUInfo` property in the `SystemInfo` struct represents information about
// the GPU (Graphics Processing Unit) of the system. It may include details such as the GPU model,
// vendor, memory size, and other relevant information about the graphics hardware.
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
// The function "Get" returns a SystemInfo object.
func Get() SystemInfo {
	return get()
}

func IsRunningInDockerContainer() bool {
	return isRunningInDockerContainer()
}
