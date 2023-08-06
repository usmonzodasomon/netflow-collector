package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"udp/models"
)

// Декодирует Data
//
// Принимает буфер, слайс TemplateField и слайс DataRecord
func DecodeData(buf *bytes.Buffer, fields []models.TemplateField) ([]models.DataRecord, error) {
	fieldsSize := GetTemplateSize(fields)
	records := make([]models.DataRecord, 0)

	for buf.Len() >= fieldsSize {
		bufferFields := bytes.NewBuffer(buf.Next(fieldsSize))
		values, err := DecodeDataSetUsingFields(bufferFields, fields)
		if err != nil {
			return nil, err
		}

		record := models.DataRecord{
			Values: values,
		}

		records = append(records, record)
	}
	return records, nil
}

// Возвращает размер полей шаблона
func GetTemplateSize(template []models.TemplateField) int {
	size := 0
	for _, templateField := range template {
		size += int(templateField.FieldLength)
	}
	return size
}

// Декодирует поля FlowSetData
//
// Получает буфер и слайс полей, возвращает слайс декодированных полей
func DecodeDataSetUsingFields(buf *bytes.Buffer, fields []models.TemplateField) ([]models.DataField, error) {
	if buf.Len() >= GetTemplateSize(fields) {

		dataFields := make([]models.DataField, len(fields))

		for i, templateField := range fields {
			value, err := DecodeFieldValue(buf, templateField.FieldLength)
			if err != nil {
				return []models.DataField{}, err
			}
			some := models.DataField{
				Type:  templateField.FieldType,
				Value: value,
			}
			dataFields[i] = some
		}
		return dataFields, nil
	}
	return []models.DataField{}, errors.New("buffer is incorrect")
}

// Декодирует поля FlowSet Data
func DecodeFieldValue(buf *bytes.Buffer, fieldType uint16) (interface{}, error) {
	switch fieldType {
	case 1:
		var value uint8
		return value, binary.Read(buf, binary.BigEndian, &value)
	case 2:
		var value uint16
		return value, binary.Read(buf, binary.BigEndian, &value)
	case 4:
		var value uint32
		return value, binary.Read(buf, binary.BigEndian, &value)
	default:
		return -1, errors.New("error fieldtype")
	}
}
