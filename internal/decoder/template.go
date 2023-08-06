package decoder

import (
	"bytes"
	"udp/models"
)

// Декодирует Template
//
// Принимает на вход буфер, Header Count и указатель на TemplateFlowSet
func DecodeTemplates(buf *bytes.Buffer, count uint16, template *models.TemplateFlowSet) error {
	for i := uint16(0); i < count; i++ {
		if err := BinaryDecoder(buf, &template.Header); err != nil {
			return err
		}

		for i := uint16(0); i < template.Header.FieldCount && buf.Len() > 0; i++ {
			var templateField models.TemplateField
			if err := BinaryDecoder(buf, &templateField); err != nil {
				return err
			}
			template.Fields = append(template.Fields, templateField)
		}

		// log.Printf("buffer len end: %v\n", buf.Len())
		models.Templates.AddTemplate(template.Header.TemplateID, template.Fields)
	}
	return nil
}
