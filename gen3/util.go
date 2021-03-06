package gen3

import (
	"encoding/binary"
	"github.com/anaminus/pkm"
	"io"
)

const addrROM = 0x08000000
const addrGameCode ptr = 0x080000AC
const strTerm = 0xFF

var defaultCodec = CodecUTF8

var structPtr = makeStruct(
	4, // Pointer
)

// Decode a string from a Reader, delimited by string terminator, using
// default encoding.
func readTextString(r io.Reader) string {
	b := make([]byte, 0, 32)
	q := make([]byte, 1)
	for {
		if _, err := r.Read(q); err == io.EOF || q[0] == strTerm {
			break
		}
		b = append(b, q[0])
	}
	s, _ := pkm.DecodeText(defaultCodec, b)
	return s
}

// Truncate text to string terminator.
func truncateText(b []byte) []byte {
	for i, c := range b {
		if c == strTerm {
			return b[:i]
		}
	}
	return b
}

// Truncate and decode a slice into a string, using default encoding.
func decodeTextString(b []byte) string {
	s, _ := pkm.DecodeText(defaultCodec, truncateText(b))
	return s
}

// Encode a string using default encoding.
func encodeText(s string) []byte {
	b, _ := pkm.EncodeText(defaultCodec, s)
	return b
}

func decUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func decUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func decUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func readLZ77(r io.Reader) ([]byte, bool) {
	q := make([]byte, 4)
	r.Read(q[:1])
	if q[0] != 0x10 {
		return nil, false
	}
	r.Read(q[:3])
	q[3] = 0
	size := int(binary.LittleEndian.Uint32(q))
	output := make([]byte, size)
	p := 0
	for p < size {
		r.Read(q[:1])
		flags := q[0]
		for i := 0; i < 8; i++ {
			if flags&(0x80>>uint(i)) == 0 {
				r.Read(q[:1])
				output[p] = q[0]
				p++
			} else {
				r.Read(q[:2])
				b := binary.LittleEndian.Uint16(q[:2])
				c := int(((b >> 4) & 0xF) + 3)
				n := p - int(((b&0xF)<<8)|((b>>8)&0xFF)) - 1
				for i := 0; i < c; i++ {
					output[p+i] = output[n+i]
				}
				p += c
			}
			if p >= size {
				break
			}
		}
	}
	return output, true
}

////////////////////////////////////////////////////////////////

type ptr uint32

func (p ptr) ValidROM() bool {
	return addrROM <= p && p < addrROM+0x01000000
}

func (p ptr) ROM() int64 {
	if !p.ValidROM() {
		return 0
	}
	return int64(p - addrROM)
}

func decPtr(b []byte) ptr {
	p := binary.LittleEndian.Uint32(b)
	return ptr(p)
}

func readPtr(r io.Reader) ptr {
	b := make([]byte, 4)
	r.Read(b)
	return decPtr(b)
}

////////////////////////////////////////////////////////////////

type stct []int

func makeStruct(fields ...int) stct {
	s := make(stct, len(fields)+1)
	for i := 0; i < len(fields); i++ {
		s[i+1] = s[i] + fields[i]
	}
	return s
}

func (s stct) Len() int {
	return len(s) - 1
}

func (s stct) Size() int {
	return s[len(s)-1]
}

func (s stct) FieldSize(f int) int {
	return s[f+1] - s[f]
}

func (s stct) FieldOffset(f int) int {
	return s[f]
}

func readStruct(r io.ReadSeeker, addr ptr, index int, s stct, fields ...int) []byte {
	if len(fields) == 0 {
		fields = make([]int, s.Len())
		for i := range fields {
			fields[i] = i
		}
	}

	off, _ := r.Seek(addr.ROM()+int64(index*s.Size()), 0)

	n := 0
	for _, f := range fields {
		n += s.FieldSize(f)
	}

	b := make([]byte, n)
	n = 0
	for _, f := range fields {
		r.Seek(off+int64(s.FieldOffset(f)), 0)
		r.Read(b[n : n+s.FieldSize(f)])
		n += s.FieldSize(f)
	}
	return b
}
