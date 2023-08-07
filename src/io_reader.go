package src

import (
	"errors"
	"io"
	"strings"
	"unicode"
)

/*
Задание:
- Изменить структуру так, чтобы она удовлетворяла интерфейсу Reader
- Метод BytesRead  возвращает кол-во байтов считанное ридером на текущий момент /////
- Метод Read : принимает []byte и пишет туда часть байт из своего CountingReaderImplementation.Reader или все байты, если длинна []byte позволяет
- Метод Read так же приводит заглавные латинские символы в строчным GaJA → gaja. Для упрощения будем использовать только латиницу
- Метод ReadAll принимает размер буфера и реализует необходимые действия чтобы вычитать все данные из CountingReaderImplementation.Reader . Можно в цикле вызывать Read . Возвращает финальную строку из всех данных
*/
type Reader interface {
	Read(p []byte) (int, error)
	ReadAll(bufSize int) (string, error)
	BytesRead() int64
}

type CountingToLowerReaderImpl struct {
	Reader         io.Reader
	TotalBytesRead int64
}

func NewCountingReader(r io.Reader) *CountingToLowerReaderImpl {
	return &CountingToLowerReaderImpl{
		Reader: r,
	}
}

func (cr *CountingToLowerReaderImpl) Read(p []byte) (int, error) {
	n, err := cr.Reader.Read(p)

	cr.TotalBytesRead += int64(n)

	for i := 0; i < n; i++ {
		if unicode.IsUpper(rune(p[i])) {
			p[i] += 32
		}
	}
	return n, err
}

func (cr *CountingToLowerReaderImpl) ReadAll(bufSize int) (string, error) {
	buff := make([]byte, bufSize)
	str := strings.Builder{}

	for {
		n, err := cr.Read(buff)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return "", err
			}
			break
		}

		str.Write(buff[:n])
	}
	return str.String(), nil
}

func (cr *CountingToLowerReaderImpl) BytesRead() int64 {
	return cr.TotalBytesRead
}
