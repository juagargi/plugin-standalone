package main

import (
	"fmt"

	"github.com/juagargi/plugin-test/pkg/snet"
)

// Declare declares that this is a path plugin.
func Declare() string {
	return "path"
}

type Main struct{}

func (Main) Doc() string {
	return `
	This is the documentation of the standalone path handler plugin.
	`
}

func (Main) PathType() string {
	return "epic"
}

func (Main) Handler() snet.PacketHandlerFunc {
	return Handle
}

func Handle(packet *snet.Packet) error {
	fmt.Printf("colibri handling %d sized packet\n", packet.SizeOfPacket)
	packet.CurrentHop++
	return nil
}
