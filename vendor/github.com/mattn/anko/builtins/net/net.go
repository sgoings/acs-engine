// +build !appengine

// Package net implements net interface for anko script.
package net

import (
	pkg "net"

	"github.com/mattn/anko/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("net")
	m.Define("CIDRMask", pkg.CIDRMask)
	m.Define("Dial", pkg.Dial)
	m.Define("DialIP", pkg.DialIP)
	m.Define("DialTCP", pkg.DialTCP)
	m.Define("DialTimeout", pkg.DialTimeout)
	m.Define("DialUDP", pkg.DialUDP)
	m.Define("DialUnix", pkg.DialUnix)
	m.Define("ErrWriteToConnected", pkg.ErrWriteToConnected)
	m.Define("FileConn", pkg.FileConn)
	m.Define("FileListener", pkg.FileListener)
	m.Define("FilePacketConn", pkg.FilePacketConn)
	m.Define("FlagBroadcast", pkg.FlagBroadcast)
	m.Define("FlagLoopback", pkg.FlagLoopback)
	m.Define("FlagMulticast", pkg.FlagMulticast)
	m.Define("FlagPointToPoint", pkg.FlagPointToPoint)
	m.Define("FlagUp", pkg.FlagUp)
	m.Define("IPv4", pkg.IPv4)
	m.Define("IPv4Mask", pkg.IPv4Mask)
	m.Define("IPv4allrouter", pkg.IPv4allrouter)
	m.Define("IPv4allsys", pkg.IPv4allsys)
	m.Define("IPv4bcast", pkg.IPv4bcast)
	m.Define("IPv4len", pkg.IPv4len)
	m.Define("IPv4zero", pkg.IPv4zero)
	m.Define("IPv6interfacelocalallnodes", pkg.IPv6interfacelocalallnodes)
	m.Define("IPv6len", pkg.IPv6len)
	m.Define("IPv6linklocalallnodes", pkg.IPv6linklocalallnodes)
	m.Define("IPv6linklocalallrouters", pkg.IPv6linklocalallrouters)
	m.Define("IPv6loopback", pkg.IPv6loopback)
	m.Define("IPv6unspecified", pkg.IPv6unspecified)
	m.Define("IPv6zero", pkg.IPv6zero)
	m.Define("InterfaceAddrs", pkg.InterfaceAddrs)
	m.Define("InterfaceByIndex", pkg.InterfaceByIndex)
	m.Define("InterfaceByName", pkg.InterfaceByName)
	m.Define("Interfaces", pkg.Interfaces)
	m.Define("JoinHostPort", pkg.JoinHostPort)
	m.Define("Listen", pkg.Listen)
	m.Define("ListenIP", pkg.ListenIP)
	m.Define("ListenMulticastUDP", pkg.ListenMulticastUDP)
	m.Define("ListenPacket", pkg.ListenPacket)
	m.Define("ListenTCP", pkg.ListenTCP)
	m.Define("ListenUDP", pkg.ListenUDP)
	m.Define("ListenUnix", pkg.ListenUnix)
	m.Define("ListenUnixgram", pkg.ListenUnixgram)
	m.Define("LookupAddr", pkg.LookupAddr)
	m.Define("LookupCNAME", pkg.LookupCNAME)
	m.Define("LookupHost", pkg.LookupHost)
	m.Define("LookupIP", pkg.LookupIP)
	m.Define("LookupMX", pkg.LookupMX)
	m.Define("LookupNS", pkg.LookupNS)
	m.Define("LookupPort", pkg.LookupPort)
	m.Define("LookupSRV", pkg.LookupSRV)
	m.Define("LookupTXT", pkg.LookupTXT)
	m.Define("ParseCIDR", pkg.ParseCIDR)
	m.Define("ParseIP", pkg.ParseIP)
	m.Define("ParseMAC", pkg.ParseMAC)
	m.Define("Pipe", pkg.Pipe)
	m.Define("ResolveIPAddr", pkg.ResolveIPAddr)
	m.Define("ResolveTCPAddr", pkg.ResolveTCPAddr)
	m.Define("ResolveUDPAddr", pkg.ResolveUDPAddr)
	m.Define("ResolveUnixAddr", pkg.ResolveUnixAddr)
	m.Define("SplitHostPort", pkg.SplitHostPort)
	return m
}
