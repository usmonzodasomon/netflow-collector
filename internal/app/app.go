package app

import (
	"bytes"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"udp/config"
	"udp/internal/decoder"
	"udp/internal/parser"
)

// Точка старта программы
func Run() {
	conn, err := parser.GetConnUDP()
	if err != nil {
		log.Fatal("error while connecting upd: ", err.Error())
	}
	defer conn.Close()

	ch := make(chan *bytes.Buffer, config.AppSettings.Multithreading.BufferSize)
	shutdownCh := make(chan struct{})
	done := make(chan struct{})
	go parser.GetPacket(conn, ch, shutdownCh)
	go ProcessPackets(ch, done)

	log.Println("App started")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	shutdownCh <- struct{}{}
	log.Println("exiting...")
	<-done
	log.Println("App exited")
}

// Обрабатывает пакеты через горутины
//
// Получает канал для считывания данных и отправки сигнала о завершении работы
func ProcessPackets(ch chan *bytes.Buffer, done chan struct{}) {
	NumGoroutines := config.AppSettings.Multithreading.NumGoroutines
	var wg sync.WaitGroup
	wg.Add(NumGoroutines)
	for i := 0; i < NumGoroutines; i++ {
		go func() {
			defer wg.Done()
			for buf := range ch {
				if err := decoder.Decode(buf); err != nil {
					log.Println(err)
				}
			}
		}()
	}
	done <- struct{}{}
}
