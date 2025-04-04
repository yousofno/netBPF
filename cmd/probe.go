package cmd

// #include "pkt.h"
import "C"
import (
	"log"

	"github.com/cilium/ebpf/rlimit"
)

type __sk_buff C.struct___sk_buff

const BPF_OK = C.BPF_OK
const ETH_P_IP = C.ETH_P_IP

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang pkt_counter pkt.c -- -O3 -Wall -Werror -Wno-address-of-packed-member
func Start_kernel() (*pkt_counterObjects, error) {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
		return nil, err
	}
	// Load the compiled eBPF ELF and load it into the kernel.
	var objs pkt_counterObjects
	if err := loadPkt_counterObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
		return nil, err
	}
	return &objs, nil
}
func pkt_access(sk *__sk_buff) int32 {
	if sk == nil {
		return -1
	}
	var skb C.struct___sk_buff = C.struct___sk_buff(*sk)
	return int32(C.test_pkt_access(&skb))
}
