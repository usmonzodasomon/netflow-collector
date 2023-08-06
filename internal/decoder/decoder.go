package decoder

import (
	"bytes"
	"errors"

	"udp/internal/saver"
	"udp/models"
)

// Декодирует данные получённые из буфера
func Decode(buf *bytes.Buffer) error {
	var header models.Header
	if err := DecodeHeader(buf, &header); err != nil {
		return err
	}

	if header.Version != 9 {
		return errors.New("получен пакет, который не является netflow-пакетом версии 9")
	}

	var flowSetHeader models.FlowSetHeader
	if err := BinaryDecoder(buf, &flowSetHeader); err != nil {
		return err
	}

	if flowSetHeader.ID == 0 {
		var template models.TemplateFlowSet
		template.FlowSetHeader = flowSetHeader
		return DecodeTemplates(buf, header.Count, &template)
	} else {
		fields := models.Templates.GetTemplates(flowSetHeader.ID)
		if len(fields) == 0 {
			return nil
		}
		var data models.DataFlowSet
		data.FlowSetHeader = flowSetHeader

		var err error
		data.Records, err = DecodeData(buf, fields)
		if err != nil {
			return err
		}

		if err := saver.Save(data.Records, header.UnixSeconds); err != nil {
			return err
		}
		return nil
	}
}
