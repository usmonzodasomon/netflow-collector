package models

type Header struct {
	Version     uint16
	Count       uint16
	SySUPTime   uint32
	UnixSeconds uint32
	SequenceNum uint32
	SourceID    uint32
}
