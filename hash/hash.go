package hash

func hash(u, shift uint32) uint32 {
	return (u * 0x1e35a7bd) >> shift
}

type Hash struct {
	table []int
}

func New(capacity int) *Hash {
	return &Hash{
		table: make([]int, capacity),
	}
}

func (h *Hash) Append(key uint32, value int) {
	h.table[hash(uint32(key), 18)] = value
}

func (h *Hash) GetValue(key uint32) (int, bool) {
	data := h.table[hash(uint32(key), 18)]

	return data, data != 0
}
