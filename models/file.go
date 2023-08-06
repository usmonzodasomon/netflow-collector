package models

import "sync"

type FileRecord struct {
	Source      uint32
	Destination uint32
	Packets     uint32
	Bytes       uint32
	Sport       uint16
	Dport       uint16
	TcpFlags    uint8
	Proto       uint8
}

var FileMu sync.Mutex
