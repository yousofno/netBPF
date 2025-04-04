#include "pkt.h"
int test_pkt_access(struct __sk_buff *skb)
{
    bpf_printk("Im Yousof and i am in the kernel");
    if(skb->protocol != ETH_P_IP){
        bpf_printk("This is not an ip packet , prtcl number is : %d" , skb->protocol);
        return BPF_OK;
    }
    bpf_printk("This is an ip packet");
    return BPF_OK;
}
char __license[] SEC("license") = "Dual MIT/GPL";

