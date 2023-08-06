package saver

import (
	"encoding/binary"
	"io"
	"os"

	"udp/config"
	"udp/models"
)

var rec = []uint16{8, 12, 2, 1, 7, 11, 6, 4}

// Открывает файл, обрабатывает данные и сохраняет их в файл
func Save(dataRecord []models.DataRecord, headerUnixSeconds uint32) error {
	file, err := os.OpenFile(config.AppSettings.File.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	return ProcessData(file, dataRecord, headerUnixSeconds)
}

// Обрабатывает поля в том порядке, в котором указано в слайсе rec и сохраняет в файл
func ProcessData(file io.Writer, dataRecord []models.DataRecord, headerUnixSeconds uint32) error {
	for _, d := range dataRecord {
		for _, r := range rec {
			for _, v := range d.Values {
				if v.Type == r {
					if err := Write(file, v.Value); err != nil {
						return err
					}
				}
			}
		}
		if err := Write(file, headerUnixSeconds); err != nil {
			return err
		}
	}
	return nil
}

// Записывает данные в файл, используя блокировку
func Write(file io.Writer, data interface{}) error {
	models.FileMu.Lock()
	defer models.FileMu.Unlock()
	return binary.Write(file, binary.BigEndian, data)
}
