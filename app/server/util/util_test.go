package util

import (
	"testing"
	"log"
	"github.com/shirou/gopsutil/mem"
	"encoding/json"
)

func TestHostPublicIP(t *testing.T) {
	ip := HostPublicIP()
	log.Println(ip)
}

func TestHostStatus(t *testing.T) {
	swapStat, _ := mem.SwapMemory()
	bs, _ := json.Marshal(swapStat)
	t.Log("swap", string(bs))
	virtual, _ := mem.VirtualMemory()
	bs, _ = json.Marshal(virtual)
	t.Log("virtual",string(bs))
}
