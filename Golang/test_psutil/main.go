package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"time"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	fmt.Println("====== Test go-psutil ======")

	fmt.Println("CPU information:")

	printMemStat()
}

var cpuStat []cpu.TimesStat
var cpuPrevStat []cpu.TimesStat

type CPUPercentages struct {
	CPU       string
	User      float64
	System    float64
	Idle      float64
	Nice      float64
	IOWait    float64
	IRQ       float64
	SoftIRQ   float64
	Steal     float64
	Guest     float64
	GuestNice float64
	Stolen    float64
	Total     float64
}

func printCpuStat() {
	cpuInfo, err := cpu.Info()

	if err != nil {
		panic(fmt.Sprintf("cpu.Info() failed! error:%v", err.Error()))
	}

	for _, info := range cpuInfo {
		fmt.Printf("[CPU-%v] core id:%v, cores:%v\r\n", info.CPU, info.CoreID, info.Cores)
	}

	fmt.Println("CPU stat:")

	cpuPrevStat = cpuStat
	cpuStat, err = cpu.Times(false)

	if err != nil {
		panic(fmt.Sprintf("cpu.Times() failed! error:%v", err.Error()))
	}

	if len(cpuPrevStat) == 0 {
		fmt.Println("first time check cpu stat, don't deal it")
		return
	}

	for index, stat := range cpuStat {
		//fmt.Printf("[CPU-%v] core id:%v, cores:%v\r\n", info.CPU, info.CoreID, info.Cores)
		fmt.Println(stat)

		prevStat := cpuPrevStat[index]

		total := stat.User + stat.System + stat.Idle + stat.Nice + stat.Iowait + stat.Irq + stat.Softirq + stat.Steal + stat.Guest + stat.GuestNice + stat.Stolen
		prevTotal := prevStat.User + prevStat.System + prevStat.Idle + prevStat.Nice + prevStat.Iowait + prevStat.Irq + prevStat.Softirq + prevStat.Steal + prevStat.Guest + prevStat.GuestNice + prevStat.Stolen

		diff := total - prevTotal

		s := &CPUPercentages{}
		s.CPU = stat.CPU
		s.User = (stat.User - prevStat.User) / diff * 100
		s.System = (stat.System - prevStat.System) / diff * 100
		s.Idle = (stat.Idle - prevStat.Idle) / diff * 100
		s.Nice = (stat.Nice - prevStat.Nice) / diff * 100
		s.IOWait = (stat.Iowait - prevStat.Iowait) / diff * 100
		s.IRQ = (stat.Irq - prevStat.Irq) / diff * 100
		s.SoftIRQ = (stat.Softirq - prevStat.Softirq) / diff * 100
		s.Steal = (stat.Steal - prevStat.Steal) / diff * 100
		s.Guest = (stat.Guest - prevStat.Guest) / diff * 100
		s.GuestNice = (stat.GuestNice - prevStat.GuestNice) / diff * 100
		s.Stolen = (stat.Stolen - prevStat.Stolen) / diff * 100
		s.Total = 100 * (diff - (stat.Idle - prevStat.Idle)) / diff

		fmt.Println(s)
	}

	time.Sleep(10 * time.Second)
}

func printMemStat() {
	memStat, err := mem.VirtualMemory()

	if err != nil {
		panic(fmt.Sprintf("mem.VirtualMemory() failed! error:%v", err.Error()))
	}

	fmt.Println(memStat)
}