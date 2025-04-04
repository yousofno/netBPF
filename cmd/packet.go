package cmd

import (
	"C"
	"errors"
	"fmt"
	"net"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type TCPPacket struct {
	SrcPort                 uint64
	DstPort                 uint64
	FIN, SYN, RST, PSH, ACK bool
}

func (packet *TCPPacket)createTCPPacket(ipLayer *[]byte) (*[]byte, error) {
	if ipLayer == nil {
		return nil, errors.New("ip layer is empty")
	}
	var tcpPacket layers.TCP = layers.TCP{SrcPort: layers.TCPPort(packet.SrcPort), DstPort: layers.TCPPort(packet.DstPort), FIN: packet.FIN,
		SYN: packet.SYN, RST: packet.RST, PSH: packet.PSH, ACK: packet.ACK}
	var ipPacketLayer layers.IPv4 = layers.IPv4{}
	if err := ipPacketLayer.DecodeFromBytes(*ipLayer, gopacket.NilDecodeFeedback); err != nil {
		fmt.Println("Error decoding:", err)
		return nil, errors.New("can't decode ip layer")
	}
	tcpPacket.SetNetworkLayerForChecksum(&ipPacketLayer)
	var buf gopacket.SerializeBuffer = gopacket.NewSerializeBuffer()
	if err := tcpPacket.SerializeTo(buf, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true}); err != nil {
		return nil, err
	}
	ans := buf.Bytes()
	return &ans, nil
}

type IPPacket struct {
	Version  uint8
	SrcIP    net.IP
	DstIP    net.IP
	Protocol layers.IPProtocol
}

func (packet *IPPacket) createIPPacket() (*[]byte, error) {
	var ipPacket layers.IPv4 = layers.IPv4{Version: packet.Version, SrcIP: packet.SrcIP, DstIP: packet.DstIP, Protocol: packet.Protocol}
	var buf gopacket.SerializeBuffer = gopacket.NewSerializeBuffer()
	if err := ipPacket.SerializeTo(buf, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true}); err != nil {
		return nil, err
	}
	ans := buf.Bytes()
	return &ans, nil
}


type EthPacket struct {
	SrcMAC, DstMAC net.HardwareAddr
    EthernetType   layers.EthernetType
}

func (packet *EthPacket)createEtherPacket() (*[]byte , error){
	var ethPacket layers.Ethernet = layers.Ethernet{SrcMAC:packet.SrcMAC , DstMAC: packet.DstMAC , EthernetType: packet.EthernetType }
	buf := gopacket.NewSerializeBuffer()
	if err := ethPacket.SerializeTo(buf , gopacket.SerializeOptions{FixLengths: true ,ComputeChecksums: true});err != nil {
		return nil , err
	}
	ans := buf.Bytes()
	return &ans , nil
}