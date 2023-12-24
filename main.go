package main

import (
	"root/hash"
)

func main() {
	hash := hash.New(1 << 14)
	text := []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
	buffer := make([]byte, 2000)
	Search(buffer, text, hash)
}

func Search(buffer, data []byte, hash *hash.Hash) (d int) {
	pointer := 0

	for {
		skip := 32
		var position int
		for {
			//quick search (first 32 byte check every, after 32 check
			//every second, after that every third...)
			jump := skip >> 5
			pointer += jump
			skip += jump

			if pointer > len(data)-17 {
				return
			}

			block_byte := load32(data, pointer) //4 byte per block

			//check if our value exist
			if p, find := hash.GetValue(block_byte); find { //p is candidate
				if load32(data, p) == load32(data, pointer) {
					position = p
					break
				}
			}

			//write position in hash map
			hash.Append(block_byte, pointer)
		}

		d += emitLiteral(buffer, data[:position])

		base := pointer

		pointer, position = pointer+4, position+4

		for pointer < len(data) {
			if data[position] == data[pointer] {
				position, pointer = position+1, pointer+1
			} else {
				break
			}
		}

		d += emitCopy(buffer[d:], base-position, pointer-base)
	}
}

func load32(b []byte, i int) uint32 {
	b = b[i : i+4 : len(b)]
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func emitLiteral(dst, lit []byte) int {
	i, n := 0, uint(len(lit)-1)
	switch {
	case n < 60:
		dst[0] = uint8(n)<<2 | 0x00
		i = 1
	case n < 1<<8:
		dst[0] = 60<<2 | 0x00
		dst[1] = uint8(n)
		i = 2
	default:
		dst[0] = 61<<2 | 0x00
		dst[1] = uint8(n)
		dst[2] = uint8(n >> 8)
		i = 3
	}
	return i + copy(dst[i:], lit)
}

func emitCopy(dst []byte, offset, length int) int {
	i := 0

	for length >= 68 {
		dst[i+0] = 63<<2 | 0x11
		dst[i+1] = uint8(offset)
		dst[i+2] = uint8(offset >> 8)
		i += 3
		length -= 64
	}
	if length > 64 {
		dst[i+0] = 59<<2 | 0x11
		dst[i+1] = uint8(offset)
		dst[i+2] = uint8(offset >> 8)
		i += 3
		length -= 60
	}
	if length >= 12 || offset >= 2048 {
		dst[i+0] = uint8(length-1)<<2 | 0x11
		dst[i+1] = uint8(offset)
		dst[i+2] = uint8(offset >> 8)
		return i + 3
	}

	dst[i+0] = uint8(offset>>8)<<5 | uint8(length-4)<<2 | 0x01
	dst[i+1] = uint8(offset)
	return i + 2
}
