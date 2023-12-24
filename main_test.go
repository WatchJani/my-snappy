package main

import (
	"root/hash"
	"testing"
)

func Benchmark(b *testing.B) {
	hash := hash.New(1 << 14)
	text := []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
	buffer := make([]byte, 1000)

	for i := 0; i < b.N; i++ {
		Search(buffer, text, hash)
	}
}

func BenchmarkEqual(b *testing.B) {
	data := []byte("danas je lijep dana da je to super")
	position, pointer := 0, 0

	for i := 0; i < b.N; i++ {
		for pointer < len(data) {
			if data[position] == data[pointer] {
				position, pointer = position+1, pointer+1
			} else {
				break
			}
		}
	}
}

func BenchmarkEqual2(b *testing.B) {
	data := []byte("danas je lijep dana da je to super")
	position, pointer := 0, 0

	for i := 0; i < b.N; i++ {
		for b := position + 4; pointer < len(data) && data[b] == data[pointer]; b, pointer = b+1, pointer+1 {
		}
	}
}

func BenchmarkWriteNumber2(b *testing.B) {
	data := []byte("danas je lijep dana da je to super")

	buff := make([]byte, 8)

	for i := 0; i < b.N; i++ {
		emitLiteral(buff, data)
	}
}
