package encode

import (
	"strings"
)

const base62Len = 62

type Base62 struct {
	table    string
	indexMap map[byte]int
}

func NewBase62(table string, blacklist []string) *Base62 {
	if len(table) != base62Len {
		return nil
	}
	idxMap := make(map[byte]int, base62Len)
	for i := range table {
		idxMap[table[i]] = i
	}

	return &Base62{
		table:    table,
		indexMap: idxMap,
	}
}

func (b *Base62) Encode(seq uint64) string {
	if seq == 0 {
		return "0"
	}
	buf := new(strings.Builder)
	for seq > 0 {
		r := seq % base62Len
		seq /= base62Len
		buf.WriteByte(b.table[r])
	}
	return buf.String()
}

func (b *Base62) Decode(cipher string) uint64 {
	seq := uint64(0)
	for i := range cipher {
		seq += uint64(b.indexMap[cipher[i]])
		seq *= base62Len
	}
	return seq
}
