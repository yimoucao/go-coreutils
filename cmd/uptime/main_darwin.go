package main

import (
	"fmt"
	"time"

	"github.com/yimoucao/go-coreutils"
)

/*
https://stackoverflow.com/questions/12523704/mac-os-x-equivalent-header-file-for-sysinfo-h-in-linux

The closest equivalent to sysinfo in Mac OS X is sysctl / MIB. It doesn't return a sysinfo struct directly, but most of the values in that structure are available as sysctl keys. For instance:

uptime is approximated by kern.boottime (although that reflects the actual boot time, not the running time)
loads is available as vm.loadavg
totalram = hw.memsize (in bytes)
freeram, sharedram, and bufferram are complicated, as the XNU memory manager works differently from Linux's. I'm not sure if the closest equivalent values ("active" and "inactive" memory) are exposed.
totalswap and freeswap are reflected in vm.swapusage. (But note that OS X allocates swap space dynamically.)
procs doesn't appear to have any equivalent.
totalhigh and freehigh are specific to i386 Linux

*/

// https://ss64.com/osx/sysctl.html

// 21:47  up 10 days,  6:47, 6 users, load averages: 1.99 2.17 2.09
// time now, time since boot, ???, # of users, load average

func main() {
	boottime, err := coreutils.BootTime()
	if err != nil {
		panic(err)
	}
	now := time.Now()

	fmt.Println(boottime)
	delta := now.Sub(boottime)
	fmt.Println(delta)

	load, err := coreutils.LoadAvg()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.2f %.2f %.2f\n", load[0], load[1], load[2])

	totalram := coreutils.Totalram()
	fmt.Println("total ram: ", totalram)
}
