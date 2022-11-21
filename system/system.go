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
	xlogs "github.com/mt1976/appFrame/logs"
	xtl "github.com/mt1976/appFrame/translate"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func init() {

	SYSINFO, _ = gInfo.GetInfo()

}

func get() SystemInfo {
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

	xlogs.WithFields(xlogs.Fields{"User": thisSystem.UserName, "Home": thisSystem.UserHome, "Name": thisSystem.User}).Info(xtl.Get("user"))
	xlogs.WithFields(xlogs.Fields{"Hostname": thisSystem.Hostname, "Uptime": thisSystem.Uptime, "CPU": thisSystem.CPU, "OS": thisSystem.OS, "Arch": thisSystem.Arch, "Kernel": thisSystem.Kernel, "Memory": thisSystem.Memory, "Network": thisSystem.Network, "Docker": thisSystem.Docker, "User": thisSystem.User, "GoVersion": thisSystem.GoVersion}).Info(xtl.Get("system"))

	//fmt.Printf("os.Environ(): %v\n", os.Environ())

	//spew.Dump(thisSystem)
	return thisSystem
}

func getUserHome() string {
	usr, err := user.Current()
	if err != nil {
		xlogs.Fatal(err)
	}
	return usr.HomeDir
}

func getUser() string {
	usr, err := user.Current()
	if err != nil {
		xlogs.Fatal(err)
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
	xlogs.WithFields(xlogs.Fields{"Total": thisDisk.Total, "Free": thisDisk.Free, "Used": thisDisk.Used}).Info(xtl.Get("disk"))
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
	xlogs.WithFields(xlogs.Fields{"IP": thisNetwork.IP, "MAC": thisNetwork.MAC, "Name": thisNetwork.Name}).Info(xtl.Get("network"))
	return thisNetwork
}

func getNetworkName() string {
	thisNetworkName := "TBD"
	return thisNetworkName
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		xlogs.Fatal(err)
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
		xlogs.Fatal(err)
	}
	var mac string
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			xlogs.Fatal(err)
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
	xlogs.WithFields(xlogs.Fields{"Total": thisMemory.HumanTotal, "Free": thisMemory.HumanFree, "Used": thisMemory.HumanUsed, "UsedPercent": thisMemory.HumanUsedPercent}).Info(xtl.Get("memory"))
	return thisMemory
}

func getMemoryTotal() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Total
}

func getMemoryFree() uint64 {
	v, _ := mem.VirtualMemory()
	xlogs.WithFields(xlogs.Fields{"Free": humanize.Bytes(v.Free), "Total": humanize.Bytes(v.Total)}).Info(xtl.Get("memory"))
	return v.Free
}

func getCPUInfo() CPUInfo {
	var thisCPU CPUInfo
	thisCPU.NoCPUs = runtime.NumCPU()
	thisCPU.CPUs = strconv.Itoa(thisCPU.NoCPUs)
	xlogs.WithFields(xlogs.Fields{"CPUs": thisCPU.NoCPUs, "Info": thisCPU.CPUs}).Info(xtl.Get("cpu"))
	return thisCPU
}

func getUptime() string {
	uptime, err := host.Uptime()
	if err != nil {
		xlogs.Error("error getting uptime: ", err)
	}
	upString, _ := time.ParseDuration(strconv.Itoa(int(uptime)) + "s")
	xlogs.WithField("uptime", upString).Debug("uptime")
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
