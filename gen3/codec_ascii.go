package gen3

import (
	"io"
)

// CodecASCII is a pkm.Codec that decodes as single bytes. It has the least
// detail, but each character fits in one byte, and will not be interpreted as
// invalid unicode.
//
// Decode decodes text data from src into bytes, which are written to dst.
// Unrepresentable characters are replaced by "#", or a reasonable equivalent.
//
// Encode encodes bytes in src into text data, which are written to dst. Most
// printable characters are converted as expected. All other characters,
// including "#", are encoded as 0xFF.
var CodecASCII codecASCII

type codecASCII struct{}

func (codecASCII) Name() string {
	return "ASCII"
}

func (codecASCII) Decode(dst io.Writer, src io.Reader) (written int, err error) {
	buf := make([]byte, 1024)
	for {
		n, e := src.Read(buf)
		if e != nil && e != io.EOF {
			err = e
			return
		}
		if n > 0 {
			for i := 0; i < n; i++ {
				buf[i] = codecASCIILookup[int(buf[i])]
			}
			n, e := dst.Write(buf[:n])
			written += n
			if e != nil {
				err = e
				return
			}
		}
		if e != nil { // EOF
			return
		}
	}
}

func (codecASCII) Encode(dst io.Writer, src io.Reader) (written int, err error) {
	buf := make([]byte, 1024)
	for {
		n, e := src.Read(buf)
		if e != nil && e != io.EOF {
			err = e
			return
		}
		if n > 0 {
			for i := 0; i < n; i++ {
				buf[i] = codecASCIILookup[int(buf[i])+256]
			}
			n, e := dst.Write(buf[:n])
			written += n
			if e != nil {
				err = e
				return
			}
		}
		if e != nil { // EOF
			return
		}
	}
}

var codecASCIILookup = [512]byte{
	// Decode
	0x00: ' ',
	0x01: 'A',
	0x02: 'A',
	0x03: 'A',
	0x04: 'C',
	0x05: 'E',
	0x06: 'E',
	0x07: 'E',
	0x08: 'E',
	0x09: 'I',
	0x0A: '#',
	0x0B: 'I',
	0x0C: 'I',
	0x0D: 'O',
	0x0E: 'O',
	0x0F: 'O',
	0x10: '#',
	0x11: 'U',
	0x12: 'U',
	0x13: 'U',
	0x14: 'N',
	0x15: 'B',
	0x16: 'a',
	0x17: 'a',
	0x18: '#',
	0x19: 'c',
	0x1A: 'e',
	0x1B: 'e',
	0x1C: 'e',
	0x1D: 'e',
	0x1E: 'i',
	0x1F: '#',
	0x20: 'i',
	0x21: 'i',
	0x22: 'o',
	0x23: 'o',
	0x24: 'o',
	0x25: '#',
	0x26: 'u',
	0x27: 'u',
	0x28: 'u',
	0x29: 'n',
	0x2A: 'o',
	0x2B: 'a',
	0x2C: '#',
	0x2D: '&',
	0x2E: '+',
	0x2F: '#',
	0x30: '#',
	0x31: '#',
	0x32: '#',
	0x33: '#',
	0x34: '#',
	0x35: '=',
	0x36: '#',
	0x37: '#',
	0x38: '#',
	0x39: '#',
	0x3A: '#',
	0x3B: '#',
	0x3C: '#',
	0x3D: '#',
	0x3E: '#',
	0x3F: '#',
	0x40: '#',
	0x41: '#',
	0x42: '#',
	0x43: '#',
	0x44: '#',
	0x45: '#',
	0x46: '#',
	0x47: '#',
	0x48: '#',
	0x49: '#',
	0x4A: '#',
	0x4B: '#',
	0x4C: '#',
	0x4D: '#',
	0x4E: '#',
	0x4F: '#',
	0x50: '#',
	0x51: '?',
	0x52: '!',
	0x53: '#',
	0x54: '#',
	0x55: '#',
	0x56: '#',
	0x57: '#',
	0x58: '#',
	0x59: '#',
	0x5A: 'I',
	0x5B: '%',
	0x5C: '(',
	0x5D: ')',
	0x5E: '#',
	0x5F: '#',
	0x60: '#',
	0x61: '#',
	0x62: '#',
	0x63: '#',
	0x64: '#',
	0x65: '#',
	0x66: '#',
	0x67: '#',
	0x68: 'a',
	0x69: '#',
	0x6A: '#',
	0x6B: '#',
	0x6C: '#',
	0x6D: '#',
	0x6E: '#',
	0x6F: '#',
	0x70: '#',
	0x71: '#',
	0x72: '#',
	0x73: '#',
	0x74: '#',
	0x75: '#',
	0x76: '#',
	0x77: '#',
	0x78: '#',
	0x79: '^',
	0x7A: 'v',
	0x7B: '<',
	0x7C: '>',
	0x7D: '#',
	0x7E: '#',
	0x7F: '#',
	0x80: '#',
	0x81: '#',
	0x82: '#',
	0x83: '#',
	0x84: '#',
	0x85: '#',
	0x86: '#',
	0x87: '#',
	0x88: '#',
	0x89: '#',
	0x8A: '#',
	0x8B: '#',
	0x8C: '#',
	0x8D: '#',
	0x8E: '#',
	0x8F: '#',
	0x90: '#',
	0x91: '#',
	0x92: '#',
	0x93: '#',
	0x94: '#',
	0x95: '#',
	0x96: '#',
	0x97: '#',
	0x98: '#',
	0x99: '#',
	0x9A: '#',
	0x9B: '#',
	0x9C: '#',
	0x9D: '#',
	0x9E: '#',
	0x9F: '#',
	0xA0: '#',
	0xA1: '0',
	0xA2: '1',
	0xA3: '2',
	0xA4: '3',
	0xA5: '4',
	0xA6: '5',
	0xA7: '6',
	0xA8: '7',
	0xA9: '8',
	0xAA: '9',
	0xAB: '!',
	0xAC: '?',
	0xAD: '.',
	0xAE: '-',
	0xAF: '*',
	0xB0: '_',
	0xB1: '"',
	0xB2: '"',
	0xB3: '\'',
	0xB4: '\'',
	0xB5: '#',
	0xB6: '#',
	0xB7: '$',
	0xB8: ',',
	0xB9: 'x',
	0xBA: '/',
	0xBB: 'A',
	0xBC: 'B',
	0xBD: 'C',
	0xBE: 'D',
	0xBF: 'E',
	0xC0: 'F',
	0xC1: 'G',
	0xC2: 'H',
	0xC3: 'I',
	0xC4: 'J',
	0xC5: 'K',
	0xC6: 'L',
	0xC7: 'M',
	0xC8: 'N',
	0xC9: 'O',
	0xCA: 'P',
	0xCB: 'Q',
	0xCC: 'R',
	0xCD: 'S',
	0xCE: 'T',
	0xCF: 'U',
	0xD0: 'V',
	0xD1: 'W',
	0xD2: 'X',
	0xD3: 'Y',
	0xD4: 'Z',
	0xD5: 'a',
	0xD6: 'b',
	0xD7: 'c',
	0xD8: 'd',
	0xD9: 'e',
	0xDA: 'f',
	0xDB: 'g',
	0xDC: 'h',
	0xDD: 'i',
	0xDE: 'j',
	0xDF: 'k',
	0xE0: 'l',
	0xE1: 'm',
	0xE2: 'n',
	0xE3: 'o',
	0xE4: 'p',
	0xE5: 'q',
	0xE6: 'r',
	0xE7: 's',
	0xE8: 't',
	0xE9: 'u',
	0xEA: 'v',
	0xEB: 'w',
	0xEC: 'x',
	0xED: 'y',
	0xEE: 'z',
	0xEF: '>',
	0xF0: ':',
	0xF1: 'A',
	0xF2: 'O',
	0xF3: 'U',
	0xF4: 'a',
	0xF5: 'o',
	0xF6: 'u',
	0xF7: '^',
	0xF8: 'v',
	0xF9: '<',
	0xFA: '#',
	0xFB: '#',
	0xFC: '#',
	0xFD: '#',
	0xFE: '\n',
	0xFF: '#',
	// Encode
	0x100:        0xFF,
	0x101:        0xFF,
	0x102:        0xFF,
	0x103:        0xFF,
	0x104:        0xFF,
	0x105:        0xFF,
	0x106:        0xFF,
	0x107:        0xFF,
	0x108:        0xFF,
	0x109:        0xFF,
	'\n' + 0x100: 0xFE,
	0x10B:        0xFF,
	0x10C:        0xFF,
	0x10D:        0xFF,
	0x10E:        0xFF,
	0x10F:        0xFF,
	0x110:        0xFF,
	0x111:        0xFF,
	0x112:        0xFF,
	0x113:        0xFF,
	0x114:        0xFF,
	0x115:        0xFF,
	0x116:        0xFF,
	0x117:        0xFF,
	0x118:        0xFF,
	0x119:        0xFF,
	0x11A:        0xFF,
	0x11B:        0xFF,
	0x11C:        0xFF,
	0x11D:        0xFF,
	0x11E:        0xFF,
	0x11F:        0xFF,
	' ' + 0x100:  0x00,
	'!' + 0x100:  0xAB,
	'"' + 0x100:  0xB1,
	'#' + 0x100:  0xFF,
	'$' + 0x100:  0xB7,
	'%' + 0x100:  0x5B,
	'&' + 0x100:  0x2D,
	'\'' + 0x100: 0xB3,
	'(' + 0x100:  0x5C,
	')' + 0x100:  0x5D,
	'*' + 0x100:  0xAF,
	'+' + 0x100:  0x2E,
	',' + 0x100:  0xB8,
	'-' + 0x100:  0xAE,
	'.' + 0x100:  0xAD,
	'/' + 0x100:  0xBA,
	'0' + 0x100:  0xA1,
	'1' + 0x100:  0xA2,
	'2' + 0x100:  0xA3,
	'3' + 0x100:  0xA4,
	'4' + 0x100:  0xA5,
	'5' + 0x100:  0xA6,
	'6' + 0x100:  0xA7,
	'7' + 0x100:  0xA8,
	'8' + 0x100:  0xA9,
	'9' + 0x100:  0xAA,
	':' + 0x100:  0xF0,
	0x13B:        0xFF,
	'<' + 0x100:  0x7B,
	'=' + 0x100:  0x35,
	'>' + 0x100:  0x7C,
	'?' + 0x100:  0xAC,
	0x140:        0xFF,
	'A' + 0x100:  0xBB,
	'B' + 0x100:  0xBC,
	'C' + 0x100:  0xBD,
	'D' + 0x100:  0xBE,
	'E' + 0x100:  0xBF,
	'F' + 0x100:  0xC0,
	'G' + 0x100:  0xC1,
	'H' + 0x100:  0xC2,
	'I' + 0x100:  0xC3,
	'J' + 0x100:  0xC4,
	'K' + 0x100:  0xC5,
	'L' + 0x100:  0xC6,
	'M' + 0x100:  0xC7,
	'N' + 0x100:  0xC8,
	'O' + 0x100:  0xC9,
	'P' + 0x100:  0xCA,
	'Q' + 0x100:  0xCB,
	'R' + 0x100:  0xCC,
	'S' + 0x100:  0xCD,
	'T' + 0x100:  0xCE,
	'U' + 0x100:  0xCF,
	'V' + 0x100:  0xD0,
	'W' + 0x100:  0xD1,
	'X' + 0x100:  0xD2,
	'Y' + 0x100:  0xD3,
	'Z' + 0x100:  0xD4,
	0x15B:        0xFF,
	0x15C:        0xFF,
	0x15D:        0xFF,
	'^' + 0x100:  0x79,
	'_' + 0x100:  0xB0,
	0x160:        0xFF,
	'a' + 0x100:  0xD5,
	'b' + 0x100:  0xD6,
	'c' + 0x100:  0xD7,
	'd' + 0x100:  0xD8,
	'e' + 0x100:  0xD9,
	'f' + 0x100:  0xDA,
	'g' + 0x100:  0xDB,
	'h' + 0x100:  0xDC,
	'i' + 0x100:  0xDD,
	'j' + 0x100:  0xDE,
	'k' + 0x100:  0xDF,
	'l' + 0x100:  0xE0,
	'm' + 0x100:  0xE1,
	'n' + 0x100:  0xE2,
	'o' + 0x100:  0xE3,
	'p' + 0x100:  0xE4,
	'q' + 0x100:  0xE5,
	'r' + 0x100:  0xE6,
	's' + 0x100:  0xE7,
	't' + 0x100:  0xE8,
	'u' + 0x100:  0xE9,
	'v' + 0x100:  0xEA,
	'w' + 0x100:  0xEB,
	'x' + 0x100:  0xEC,
	'y' + 0x100:  0xED,
	'z' + 0x100:  0xEE,
	0x17B:        0xFF,
	0x17C:        0xFF,
	0x17D:        0xFF,
	0x17E:        0xFF,
	0x17F:        0xFF,
	0x180:        0xFF,
	0x181:        0xFF,
	0x182:        0xFF,
	0x183:        0xFF,
	0x184:        0xFF,
	0x185:        0xFF,
	0x186:        0xFF,
	0x187:        0xFF,
	0x188:        0xFF,
	0x189:        0xFF,
	0x18A:        0xFF,
	0x18B:        0xFF,
	0x18C:        0xFF,
	0x18D:        0xFF,
	0x18E:        0xFF,
	0x18F:        0xFF,
	0x190:        0xFF,
	0x191:        0xFF,
	0x192:        0xFF,
	0x193:        0xFF,
	0x194:        0xFF,
	0x195:        0xFF,
	0x196:        0xFF,
	0x197:        0xFF,
	0x198:        0xFF,
	0x199:        0xFF,
	0x19A:        0xFF,
	0x19B:        0xFF,
	0x19C:        0xFF,
	0x19D:        0xFF,
	0x19E:        0xFF,
	0x19F:        0xFF,
	0x1A0:        0xFF,
	0x1A1:        0xFF,
	0x1A2:        0xFF,
	0x1A3:        0xFF,
	0x1A4:        0xFF,
	0x1A5:        0xFF,
	0x1A6:        0xFF,
	0x1A7:        0xFF,
	0x1A8:        0xFF,
	0x1A9:        0xFF,
	0x1AA:        0xFF,
	0x1AB:        0xFF,
	0x1AC:        0xFF,
	0x1AD:        0xFF,
	0x1AE:        0xFF,
	0x1AF:        0xFF,
	0x1B0:        0xFF,
	0x1B1:        0xFF,
	0x1B2:        0xFF,
	0x1B3:        0xFF,
	0x1B4:        0xFF,
	0x1B5:        0xFF,
	0x1B6:        0xFF,
	0x1B7:        0xFF,
	0x1B8:        0xFF,
	0x1B9:        0xFF,
	0x1BA:        0xFF,
	0x1BB:        0xFF,
	0x1BC:        0xFF,
	0x1BD:        0xFF,
	0x1BE:        0xFF,
	0x1BF:        0xFF,
	0x1C0:        0xFF,
	0x1C1:        0xFF,
	0x1C2:        0xFF,
	0x1C3:        0xFF,
	0x1C4:        0xFF,
	0x1C5:        0xFF,
	0x1C6:        0xFF,
	0x1C7:        0xFF,
	0x1C8:        0xFF,
	0x1C9:        0xFF,
	0x1CA:        0xFF,
	0x1CB:        0xFF,
	0x1CC:        0xFF,
	0x1CD:        0xFF,
	0x1CE:        0xFF,
	0x1CF:        0xFF,
	0x1D0:        0xFF,
	0x1D1:        0xFF,
	0x1D2:        0xFF,
	0x1D3:        0xFF,
	0x1D4:        0xFF,
	0x1D5:        0xFF,
	0x1D6:        0xFF,
	0x1D7:        0xFF,
	0x1D8:        0xFF,
	0x1D9:        0xFF,
	0x1DA:        0xFF,
	0x1DB:        0xFF,
	0x1DC:        0xFF,
	0x1DD:        0xFF,
	0x1DE:        0xFF,
	0x1DF:        0xFF,
	0x1E0:        0xFF,
	0x1E1:        0xFF,
	0x1E2:        0xFF,
	0x1E3:        0xFF,
	0x1E4:        0xFF,
	0x1E5:        0xFF,
	0x1E6:        0xFF,
	0x1E7:        0xFF,
	0x1E8:        0xFF,
	0x1E9:        0xFF,
	0x1EA:        0xFF,
	0x1EB:        0xFF,
	0x1EC:        0xFF,
	0x1ED:        0xFF,
	0x1EE:        0xFF,
	0x1EF:        0xFF,
	0x1F0:        0xFF,
	0x1F1:        0xFF,
	0x1F2:        0xFF,
	0x1F3:        0xFF,
	0x1F4:        0xFF,
	0x1F5:        0xFF,
	0x1F6:        0xFF,
	0x1F7:        0xFF,
	0x1F8:        0xFF,
	0x1F9:        0xFF,
	0x1FA:        0xFF,
	0x1FB:        0xFF,
	0x1FC:        0xFF,
	0x1FD:        0xFF,
	0x1FE:        0xFF,
	0x1FF:        0xFF,
}
