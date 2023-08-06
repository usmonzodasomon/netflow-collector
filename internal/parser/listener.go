package parser

import (
	"bytes"
	"log"
	"net"
	"udp/config"
)

// Возвращает UDP соединение
func GetConnUDP() (net.PacketConn, error) {
	return net.ListenPacket("udp", config.AppSettings.Udp.Address)
}

// Получает UDP пакеты из соединения conn и отправляет их в канал ch
func GetPacket(conn net.PacketConn, ch chan *bytes.Buffer, shutdownCh chan struct{}) {
	for {
		buffer := make([]byte, 1500)
		select {
		case <-shutdownCh:
			close(ch)
			return
		default:
			n, _, err := conn.ReadFrom(buffer)
			if err != nil {
				log.Println("Ошибка чтения пакета: ", err)
				continue
			}
			ch <- bytes.NewBuffer(buffer[:n])
		}
	}
}
