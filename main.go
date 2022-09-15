package main

import (
	"io"
)

// Написать очередь сообщений который кладет сообщения в очередь
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
	/*in  chan<- []byte
	out <-chan []byte*/

	dataCh chan []byte
	errCh  chan error
}

func (b *SomeBroker) Write(data []byte) {
	//b.in <- data
	b.dataCh <- data
}

func (b *SomeBroker) Read() <-chan []byte {
	//return b.out
	return b.dataCh
}

func (b *SomeBroker) Wait() {
	// ???
}

func (b *SomeBroker) Close() error {
	defer close(b.dataCh)

	err, ok := <-b.errCh
	if !ok {
		return err
	}

	return nil
}
