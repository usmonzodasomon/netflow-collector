package models

type FlowSetHeader struct {
	ID     uint16
	Length uint16
}

type DataFlowSet struct {
	FlowSetHeader

	Records []DataRecord
}

type DataRecord struct {
	Values []DataField
}

type DataField struct {
	Type  uint16
	Value interface{}
}
