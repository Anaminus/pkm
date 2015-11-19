package emerald

import (
	"bufio"
	"io"
	"strings"
)

// CodecString is a pkm.Codec that decodes into UTF-8. Single text bytes can
// be decoded into multiple UTF-8 characters. It has the best detail, but
// decoded characters cannot be re-encoded 1-to-1.
//
// Decode decodes text data from src into a UTF-8 string, which is written to
// dst. Unrepresentable characters are replaced with the corresponding rune in
// CodecPUA.
//
// Encode encodes a UTF-8 string in src into text data, which are written to
// dst. Unsupported characters are ignored.
var CodecString codecString

type codecString struct{}

func (codecString) Name() string {
	return "String"
}

func (codecString) Decode(dst io.Writer, src io.Reader) (written int, err error) {
	bufin := make([]byte, 1024)
	bufout := make([]string, 1024)
	for {
		n, e := src.Read(bufin)
		if e != nil && e != io.EOF {
			err = e
			return
		}
		for i := 0; i < n; i++ {
			bufout[i] = codecStringDecode[bufin[i]]
		}
		if _, e := dst.Write([]byte(strings.Join(bufout, ""))); e != nil {
			err = e
			return
		}
		if e != nil { // EOF
			return
		}
	}
}

func (codecString) Encode(dst io.Writer, src io.Reader) (written int, err error) {
	buf := bufio.NewReader(src)
	bufout := make([]byte, 1)
	for {
		r, _, e := buf.ReadRune()
		if e != nil && e != io.EOF {
			err = e
			return
		}
		if b, ok := codecStringEncode[r]; ok {
			bufout[0] = b
			if _, e := dst.Write(bufout); e != nil {
				err = e
				return
			}
		}
		if e != nil { // EOF
			return
		}
	}
}

var (
	codecStringDecode = [256]string{
		0x00: " ",
		0x01: "À",
		0x02: "Á",
		0x03: "Â",
		0x04: "Ç",
		0x05: "È",
		0x06: "É",
		0x07: "Ê",
		0x08: "Ë",
		0x09: "Ì",
		0x0A: "こ",
		0x0B: "Î",
		0x0C: "Ï",
		0x0D: "Ò",
		0x0E: "Ó",
		0x0F: "Ô",
		0x10: "Œ",
		0x11: "Ù",
		0x12: "Ú",
		0x13: "Û",
		0x14: "Ñ",
		0x15: "ß",
		0x16: "à",
		0x17: "á",
		0x18: "ね",
		0x19: "ç",
		0x1A: "è",
		0x1B: "é",
		0x1C: "ê",
		0x1D: "ë",
		0x1E: "ì",
		0x1F: "ま",
		0x20: "î",
		0x21: "ï",
		0x22: "ò",
		0x23: "ó",
		0x24: "ô",
		0x25: "œ",
		0x26: "ù",
		0x27: "ú",
		0x28: "û",
		0x29: "ñ",
		0x2A: "º",
		0x2B: "ª",
		0x2C: "er",
		0x2D: "&",
		0x2E: "+",
		0x2F: "あ",
		0x30: "ぃ",
		0x31: "ぅ",
		0x32: "ぇ",
		0x33: "ぉ",
		0x34: "Lv",
		0x35: "=",
		0x36: "ょ",
		0x37: "が",
		0x38: "ぎ",
		0x39: "ぐ",
		0x3A: "げ",
		0x3B: "ご",
		0x3C: "ざ",
		0x3D: "じ",
		0x3E: "ず",
		0x3F: "ぜ",
		0x40: "ぞ",
		0x41: "だ",
		0x42: "ぢ",
		0x43: "づ",
		0x44: "で",
		0x45: "ど",
		0x46: "ば",
		0x47: "び",
		0x48: "ぶ",
		0x49: "べ",
		0x4A: "ぼ",
		0x4B: "ぱ",
		0x4C: "ぴ",
		0x4D: "ぷ",
		0x4E: "ぺ",
		0x4F: "ぽ",
		0x50: "っ",
		0x51: "¿",
		0x52: "¡",
		0x53: "PK",
		0x54: "MN",
		0x55: "PO",
		0x56: "Ké",
		0x57: "BL",
		0x58: "OC",
		0x59: "K",
		0x5A: "Í",
		0x5B: "%",
		0x5C: "(",
		0x5D: ")",
		0x5E: "セ",
		0x5F: "ソ",
		0x60: "タ",
		0x61: "チ",
		0x62: "ツ",
		0x63: "テ",
		0x64: "ト",
		0x65: "ナ",
		0x66: "ニ",
		0x67: "ヌ",
		0x68: "â",
		0x69: "ノ",
		0x6A: "ハ",
		0x6B: "ヒ",
		0x6C: "フ",
		0x6D: "ヘ",
		0x6E: "ホ",
		0x6F: "í",
		0x70: "ミ",
		0x71: "ム",
		0x72: "メ",
		0x73: "モ",
		0x74: "ヤ",
		0x75: "ユ",
		0x76: "ヨ",
		0x77: "ラ",
		0x78: "リ",
		0x79: "⬆",
		0x7A: "⬇",
		0x7B: "⬅",
		0x7C: "➡",
		0x7D: "ヲ",
		0x7E: "ン",
		0x7F: "ァ",
		0x80: "ィ",
		0x81: "ゥ",
		0x82: "ェ",
		0x83: "ォ",
		0x84: "ャ",
		0x85: "ュ",
		0x86: "ョ",
		0x87: "ガ",
		0x88: "ギ",
		0x89: "グ",
		0x8A: "ゲ",
		0x8B: "ゴ",
		0x8C: "ザ",
		0x8D: "ジ",
		0x8E: "ズ",
		0x8F: "ゼ",
		0x90: "ゾ",
		0x91: "ダ",
		0x92: "ヂ",
		0x93: "ヅ",
		0x94: "デ",
		0x95: "ド",
		0x96: "バ",
		0x97: "ビ",
		0x98: "ブ",
		0x99: "ベ",
		0x9A: "ボ",
		0x9B: "パ",
		0x9C: "ピ",
		0x9D: "プ",
		0x9E: "ペ",
		0x9F: "ポ",
		0xA0: "ッ",
		0xA1: "0",
		0xA2: "1",
		0xA3: "2",
		0xA4: "3",
		0xA5: "4",
		0xA6: "5",
		0xA7: "6",
		0xA8: "7",
		0xA9: "8",
		0xAA: "9",
		0xAB: "!",
		0xAC: "?",
		0xAD: ".",
		0xAE: "-",
		0xAF: "・",
		0xB0: "…",
		0xB1: "“",
		0xB2: "”",
		0xB3: "‘",
		0xB4: "’",
		0xB5: "♂",
		0xB6: "♀",
		0xB7: "₽",
		0xB8: ",",
		0xB9: "×",
		0xBA: "/",
		0xBB: "A",
		0xBC: "B",
		0xBD: "C",
		0xBE: "D",
		0xBF: "E",
		0xC0: "F",
		0xC1: "G",
		0xC2: "H",
		0xC3: "I",
		0xC4: "J",
		0xC5: "K",
		0xC6: "L",
		0xC7: "M",
		0xC8: "N",
		0xC9: "O",
		0xCA: "P",
		0xCB: "Q",
		0xCC: "R",
		0xCD: "S",
		0xCE: "T",
		0xCF: "U",
		0xD0: "V",
		0xD1: "W",
		0xD2: "X",
		0xD3: "Y",
		0xD4: "Z",
		0xD5: "a",
		0xD6: "b",
		0xD7: "c",
		0xD8: "d",
		0xD9: "e",
		0xDA: "f",
		0xDB: "g",
		0xDC: "h",
		0xDD: "i",
		0xDE: "j",
		0xDF: "k",
		0xE0: "l",
		0xE1: "m",
		0xE2: "n",
		0xE3: "o",
		0xE4: "p",
		0xE5: "q",
		0xE6: "r",
		0xE7: "s",
		0xE8: "t",
		0xE9: "u",
		0xEA: "v",
		0xEB: "w",
		0xEC: "x",
		0xED: "y",
		0xEE: "z",
		0xEF: "▶",
		0xF0: ":",
		0xF1: "Ä",
		0xF2: "Ö",
		0xF3: "Ü",
		0xF4: "ä",
		0xF5: "ö",
		0xF6: "ü",
		0xF7: "⬆",
		0xF8: "⬇",
		0xF9: "⬅",
		0xFA: "\uF0FA",
		0xFB: "\uF0FB",
		0xFC: "\uF0FC",
		0xFD: "\uF0FD",
		0xFE: "\uF0FE",
		0xFF: "\uF0FF",
	}
	codecStringEncode = map[rune]byte{
		' ': 0x00,
		'À': 0x01,
		'Á': 0x02,
		'Â': 0x03,
		'Ç': 0x04,
		'È': 0x05,
		'É': 0x06,
		'Ê': 0x07,
		'Ë': 0x08,
		'Ì': 0x09,
		'こ': 0x0A,
		'Î': 0x0B,
		'Ï': 0x0C,
		'Ò': 0x0D,
		'Ó': 0x0E,
		'Ô': 0x0F,
		'Œ': 0x10,
		'Ù': 0x11,
		'Ú': 0x12,
		'Û': 0x13,
		'Ñ': 0x14,
		'ß': 0x15,
		'à': 0x16,
		'á': 0x17,
		'ね': 0x18,
		'ç': 0x19,
		'è': 0x1A,
		'é': 0x1B,
		'ê': 0x1C,
		'ë': 0x1D,
		'ì': 0x1E,
		'ま': 0x1F,
		'î': 0x20,
		'ï': 0x21,
		'ò': 0x22,
		'ó': 0x23,
		'ô': 0x24,
		'œ': 0x25,
		'ù': 0x26,
		'ú': 0x27,
		'û': 0x28,
		'ñ': 0x29,
		'º': 0x2A,
		'ª': 0x2B,
		// 'er': 0x2C,
		'&': 0x2D,
		'+': 0x2E,
		'あ': 0x2F,
		'ぃ': 0x30,
		'ぅ': 0x31,
		'ぇ': 0x32,
		'ぉ': 0x33,
		// 'Lv': 0x34,
		'=': 0x35,
		'ょ': 0x36,
		'が': 0x37,
		'ぎ': 0x38,
		'ぐ': 0x39,
		'げ': 0x3A,
		'ご': 0x3B,
		'ざ': 0x3C,
		'じ': 0x3D,
		'ず': 0x3E,
		'ぜ': 0x3F,
		'ぞ': 0x40,
		'だ': 0x41,
		'ぢ': 0x42,
		'づ': 0x43,
		'で': 0x44,
		'ど': 0x45,
		'ば': 0x46,
		'び': 0x47,
		'ぶ': 0x48,
		'べ': 0x49,
		'ぼ': 0x4A,
		'ぱ': 0x4B,
		'ぴ': 0x4C,
		'ぷ': 0x4D,
		'ぺ': 0x4E,
		'ぽ': 0x4F,
		'っ': 0x50,
		'¿': 0x51,
		'¡': 0x52,
		// 'PK': 0x53,
		// 'MN': 0x54,
		// 'PO': 0x55,
		// 'Ké': 0x56,
		// 'BL': 0x57,
		// 'OC': 0x58,
		// 'K':  0x59,
		'Í':      0x5A,
		'%':      0x5B,
		'(':      0x5C,
		')':      0x5D,
		'セ':      0x5E,
		'ソ':      0x5F,
		'タ':      0x60,
		'チ':      0x61,
		'ツ':      0x62,
		'テ':      0x63,
		'ト':      0x64,
		'ナ':      0x65,
		'ニ':      0x66,
		'ヌ':      0x67,
		'â':      0x68,
		'ノ':      0x69,
		'ハ':      0x6A,
		'ヒ':      0x6B,
		'フ':      0x6C,
		'ヘ':      0x6D,
		'ホ':      0x6E,
		'í':      0x6F,
		'ミ':      0x70,
		'ム':      0x71,
		'メ':      0x72,
		'モ':      0x73,
		'ヤ':      0x74,
		'ユ':      0x75,
		'ヨ':      0x76,
		'ラ':      0x77,
		'リ':      0x78,
		'⬆':      0x79,
		'⬇':      0x7A,
		'⬅':      0x7B,
		'➡':      0x7C,
		'ヲ':      0x7D,
		'ン':      0x7E,
		'ァ':      0x7F,
		'ィ':      0x80,
		'ゥ':      0x81,
		'ェ':      0x82,
		'ォ':      0x83,
		'ャ':      0x84,
		'ュ':      0x85,
		'ョ':      0x86,
		'ガ':      0x87,
		'ギ':      0x88,
		'グ':      0x89,
		'ゲ':      0x8A,
		'ゴ':      0x8B,
		'ザ':      0x8C,
		'ジ':      0x8D,
		'ズ':      0x8E,
		'ゼ':      0x8F,
		'ゾ':      0x90,
		'ダ':      0x91,
		'ヂ':      0x92,
		'ヅ':      0x93,
		'デ':      0x94,
		'ド':      0x95,
		'バ':      0x96,
		'ビ':      0x97,
		'ブ':      0x98,
		'ベ':      0x99,
		'ボ':      0x9A,
		'パ':      0x9B,
		'ピ':      0x9C,
		'プ':      0x9D,
		'ペ':      0x9E,
		'ポ':      0x9F,
		'ッ':      0xA0,
		'0':      0xA1,
		'1':      0xA2,
		'2':      0xA3,
		'3':      0xA4,
		'4':      0xA5,
		'5':      0xA6,
		'6':      0xA7,
		'7':      0xA8,
		'8':      0xA9,
		'9':      0xAA,
		'!':      0xAB,
		'?':      0xAC,
		'.':      0xAD,
		'-':      0xAE,
		'・':      0xAF,
		'…':      0xB0,
		'“':      0xB1,
		'”':      0xB2,
		'‘':      0xB3,
		'’':      0xB4,
		'♂':      0xB5,
		'♀':      0xB6,
		'₽':      0xB7,
		',':      0xB8,
		'×':      0xB9,
		'/':      0xBA,
		'A':      0xBB,
		'B':      0xBC,
		'C':      0xBD,
		'D':      0xBE,
		'E':      0xBF,
		'F':      0xC0,
		'G':      0xC1,
		'H':      0xC2,
		'I':      0xC3,
		'J':      0xC4,
		'K':      0xC5,
		'L':      0xC6,
		'M':      0xC7,
		'N':      0xC8,
		'O':      0xC9,
		'P':      0xCA,
		'Q':      0xCB,
		'R':      0xCC,
		'S':      0xCD,
		'T':      0xCE,
		'U':      0xCF,
		'V':      0xD0,
		'W':      0xD1,
		'X':      0xD2,
		'Y':      0xD3,
		'Z':      0xD4,
		'a':      0xD5,
		'b':      0xD6,
		'c':      0xD7,
		'd':      0xD8,
		'e':      0xD9,
		'f':      0xDA,
		'g':      0xDB,
		'h':      0xDC,
		'i':      0xDD,
		'j':      0xDE,
		'k':      0xDF,
		'l':      0xE0,
		'm':      0xE1,
		'n':      0xE2,
		'o':      0xE3,
		'p':      0xE4,
		'q':      0xE5,
		'r':      0xE6,
		's':      0xE7,
		't':      0xE8,
		'u':      0xE9,
		'v':      0xEA,
		'w':      0xEB,
		'x':      0xEC,
		'y':      0xED,
		'z':      0xEE,
		'▶':      0xEF,
		':':      0xF0,
		'Ä':      0xF1,
		'Ö':      0xF2,
		'Ü':      0xF3,
		'ä':      0xF4,
		'ö':      0xF5,
		'ü':      0xF6,
		'↑':      0xF7,
		'↓':      0xF8,
		'←':      0xF9,
		'\uF0FA': 0xFA,
		'\uF0FB': 0xFB,
		'\uF0FC': 0xFC,
		'\uF0FD': 0xFD,
		'\uF0FE': 0xFE,
		'\uF0FF': 0xFF,
	}
)
