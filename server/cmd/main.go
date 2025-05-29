package main

import (
	"fmt"
	"server/pkg/packets"

	"google.golang.org/protobuf/proto"
)

func main() {
	// packet := &packets.Packet{
	// 	SenderId: 12,
	// 	Msg:      packets.NewChat("hello"),
	// }

	// data, err := proto.Marshal(packet)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(data)

	data := []byte{8, 12, 18, 7, 10, 5, 104, 101, 108, 108, 111}

	packet := DeseralizePacket(data)
	fmt.Println(packet)
}

func DeseralizePacket(data []byte) *packets.Packet {
	packet := &packets.Packet{}

	err := proto.Unmarshal(data, packet)

	if err != nil {
		panic(err)
	}
	return packet
}
