package main

import (
	"fmt"
	"root/hash"
)

func main() {
	hash := hash.New(1 << 14)
	text := []byte("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

	hash.Append(14, 1)
	hash.Append(15, 2)

	data, _ := hash.GetValue(14)
	fmt.Println(data)

	data, _ = hash.GetValue(15)
	fmt.Println(data)

	Search(text, hash)
}

func Search(data []byte, hash *hash.Hash) {
	pointer := 0

	for {
		skip := 32
		// var position uint32
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
			if _, find := hash.GetValue(block_byte); find {
				// fmt.Println(p, pointer)
				// position = p
				// fmt.Println(string(data[p:p+4]), string(data[pointer:pointer+4]))
				break
			}

			//write position in hash map
			hash.Append(block_byte, pointer)
		}
	}
}

func load32(b []byte, i int) uint32 {
	b = b[i : i+4 : len(b)]
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
