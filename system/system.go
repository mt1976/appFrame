package system

import (
	"fmt"
	"net"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jaypipes/ghw"
	gInfo "github.com/matishsiao/goInfo"
	logs "github.com/mt1976/AppFrame/logs"
	txt "github.com/mt1976/AppFrame/translate"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
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

func init() {

	SYSINFO, _ = gInfo.GetInfo()

}

// Gets the current system information from various sources
func SystemInfoGet() SystemInfo {
	var thisSystem SystemInfo
	thisSystem.Hostname, _ = os.Hostname()
	thisSystem.Uptime = getUptime()
	thisSystem.CPU = getCPUInfo()
	thisSystem.OS = getOSInfo()
	thisSystem.Arch = getArchInfo()
	thisSystem.Kernel = getKernelInfo()
	thisSystem.Memory = getMemoryInfo()
	thisSystem.Network = getNetworkInfo()
	thisSystem.Docker = getDockerInfo()
	thisSystem.UserName = getUserName()
	thisSystem.UserHome = getUserHome()
	thisSystem.User = getUser()
	logs.WithFields(logs.Fields{"User": thisSystem.UserName, "Home": thisSystem.UserHome, "Name": thisSystem.User}).Info(txt.Get("user"))
	thisSystem.GoVersion = runtime.Version()
	thisSystem.CpuInfo, _ = ghw.CPU()
	thisSystem.MemoryInfo, _ = ghw.Memory()
	thisSystem.BlockInfo, _ = ghw.Block()
	thisSystem.NetworkInfo, _ = ghw.Network()
	thisSystem.TopologyInfo, _ = ghw.Topology()
	thisSystem.PCIInfo, _ = ghw.PCI()
	thisSystem.BIOSInfo, _ = ghw.BIOS()
	thisSystem.ChassisInfo, _ = ghw.Chassis()
	thisSystem.NetworkInfo, _ = ghw.Network()
	thisSystem.GPUInfo, _ = ghw.GPU()
	thisSystem.BIOSInfo, _ = ghw.BIOS()

	logs.WithFields(logs.Fields{"Hostname": thisSystem.Hostname, "Uptime": thisSystem.Uptime, "CPU": thisSystem.CPU, "OS": thisSystem.OS, "Arch": thisSystem.Arch, "Kernel": thisSystem.Kernel, "Memory": thisSystem.Memory, "Network": thisSystem.Network, "Docker": thisSystem.Docker, "User": thisSystem.User, "GoVersion": thisSystem.GoVersion}).Info(txt.Get("system"))

	//fmt.Printf("os.Environ(): %v\n", os.Environ())

	//spew.Dump(thisSystem)
	return thisSystem
}

func getUserHome() string {
	usr, err := user.Current()
	if err != nil {
		logs.Fatal(err)
	}
	return usr.HomeDir
}

func getUser() string {
	usr, err := user.Current()
	if err != nil {
		logs.Fatal(err)
	}
	return usr.Name
}

func getUserName() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.Username
}

func getDockerInfo() string {
	if isRunningInDockerContainer() {
		return "Yes"
	}
	return "No"
}

func getKernelInfo() string {
	return SYSINFO.Kernel + " " + SYSINFO.Core
}

func getOSInfo() string {

	os := runtime.GOOS
	switch os {
	case "windows":
		return "Windows"
	case "darwin":
		return "MACOS"
	case "linux":
		return "Linux"
	default:
	}

	return fmt.Sprintf("%s.\n", os)
}

func getArchInfo() string {
	return runtime.GOARCH
}

func getDiskInfo() DiskInfo {
	var thisDisk DiskInfo
	thisDisk.Total, thisDisk.Free, thisDisk.Used = getDiskUsage()
	logs.WithFields(logs.Fields{"Total": thisDisk.Total, "Free": thisDisk.Free, "Used": thisDisk.Used}).Info(txt.Get("disk"))
	return thisDisk
}

func getDiskUsage() (uint64, uint64, uint64) {

	return 0, 0, 0
}

func getNetworkInfo() NetworkInfo {
	var thisNetwork NetworkInfo
	thisNetwork.IP = getIP()
	thisNetwork.MAC = getMAC()
	thisNetwork.Name = getNetworkName()
	logs.WithFields(logs.Fields{"IP": thisNetwork.IP, "MAC": thisNetwork.MAC, "Name": thisNetwork.Name}).Info(txt.Get("network"))
	return thisNetwork
}

func getNetworkName() string {
	thisNetworkName := "TBD"
	return thisNetworkName
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logs.Fatal(err)
	}
	var ip string
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}

func getMAC() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		logs.Fatal(err)
	}
	var mac string
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			logs.Fatal(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			if ip.String() == getIP() {
				mac = i.HardwareAddr.String()
			}
		}
	}
	return mac
}

func getMemoryInfo() MemoryInfo {
	var thisMemory MemoryInfo
	thisMemory.Total = getMemoryTotal()
	thisMemory.Free = getMemoryFree()
	thisMemory.Used = thisMemory.Total - thisMemory.Free
	thisMemory.UsedPercent = (int64(thisMemory.Used) / int64(thisMemory.Total)) * 100
	thisMemory.HumanTotal = humanize.Bytes(thisMemory.Total)
	thisMemory.HumanFree = humanize.Bytes(thisMemory.Free)
	thisMemory.HumanUsed = humanize.Bytes(thisMemory.Used)
	thisMemory.HumanUsedPercent = fmt.Sprint(thisMemory.UsedPercent) + "%"
	logs.WithFields(logs.Fields{"Total": thisMemory.HumanTotal, "Free": thisMemory.HumanFree, "Used": thisMemory.HumanUsed, "UsedPercent": thisMemory.HumanUsedPercent}).Info(txt.Get("memory"))
	return thisMemory
}

func getMemoryTotal() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Total
}

func getMemoryFree() uint64 {
	v, _ := mem.VirtualMemory()
	logs.WithFields(logs.Fields{"Free": v.Free, "Total": v.Total, "Other": v.String()}).Info(txt.Get("memory"))
	return v.Free
}

func getCPUInfo() CPUInfo {
	var thisCPU CPUInfo
	thisCPU.NoCPUs = runtime.NumCPU()
	thisCPU.CPUs = strconv.Itoa(thisCPU.NoCPUs)
	logs.WithFields(logs.Fields{"CPUs": thisCPU.NoCPUs, "Info": thisCPU.CPUs}).Info(txt.Get("cpu"))
	return thisCPU
}

func getUptime() string {
	uptime, err := host.Uptime()
	if err != nil {
		logs.Error("error getting uptime: ", err)
	}
	upString, _ := time.ParseDuration(strconv.Itoa(int(uptime)) + "s")
	logs.WithField("uptime", upString).Debug("uptime")
	return upString.String()
}

func isRunningInDockerContainer() bool {
	// docker creates a .dockerenv file at the root
	// of the directory tree inside the container.
	// if this file exists then the viewer is running
	// from inside a container so return true

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}

func IsRunningInDockerContainer() bool {
	return isRunningInDockerContainer()
}
