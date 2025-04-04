package cmd

import (
	"reflect"
	"testing"
)

func Test_start_kernel(t *testing.T) {
	if objs, err := Start_kernel(); err != nil {
		t.Error("Failed in loading madule in kernel")
	} else {
		defer objs.Close()
	}

}

func Test_dummy_packet(t *testing.T) {
	if objs, err := Start_kernel(); err != nil {
		t.Error("Failed in loading madule in kernel")
	} else {
		defer objs.Close()
		//constructing packet from eth header
		//this packet is a dummy packet with size of 14 bytes
		var input []byte = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4}
		if ret, out, err := objs.TestPktAccess.Test(input); err != nil {
			t.Error("Error in testing packet")
		} else {
			if ret != 0 {
				t.Error("input and output are not the same")
			}
			if !reflect.DeepEqual(out, input) {
				t.Error("input and output are not the same")
			}
		}
	}
}

func Test_skb_ipv4_validation(t *testing.T) {
	var skb __sk_buff = __sk_buff{}
	if skb.protocol == ETH_P_IP {
		

	}
}
