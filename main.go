package main

import (
	"fmt"
	"io"
	"time"
)

// Написать брокер сообщений который кладет сообщения в очередь
// и из которого есть возможность читать эти сообщения.

// Имплементировать следующий интерфейс:

// Queue интерфейс описывающий ПОТОКОБЕЗОПАСНУЮ работу с очередью
type Queue interface {
	// Write Запись сообщений в брокер
	Write(bb []byte)

	// Read чтение сообщений из брокера
	Read() <-chan []byte

	// Wait ожидание завершения исполнения горутин
	Wait()

	// Closer Завершение брокера
	io.Closer
}

type SomeBroker struct {
	dataCh chan []byte

	closedBool bool
}

func NewSomeBroker() *SomeBroker {
	return &SomeBroker{
		dataCh: make(chan []byte),
	}
}

func (b *SomeBroker) Write(data []byte) {
	b.dataCh <- data
}

func (b *SomeBroker) Read() <-chan []byte {
	return b.dataCh
}

func (b *SomeBroker) Wait() {
	for !b.closedBool {
	}
}

func (b *SomeBroker) Close() error {
	defer close(b.dataCh)

	b.closedBool = true

	return nil
}

func main() {
	sb := NewSomeBroker()

	go func() {
		time.Sleep(5 * time.Second)
		sb.Close()
	}()

	go sb.Write([]byte("Str1"))
	go sb.Write([]byte("Str2"))
	go sb.Write([]byte("Str3"))
	go sb.Write([]byte("Str4"))

	go func() {
		for v := range sb.Read() {
			fmt.Println(string(v))
		}
	}()

	sb.Wait()
}
