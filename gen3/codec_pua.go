package gen3

import (
	"bufio"
	"io"
)

// CodecPUA is a pkm.Codec that decodes into UTF-8. Characters are located in
// the Private Use Area in the range F000-F0FF.
//
// Decode decodes text data from src into a UTF-8 string, which is written to
// dst.
//
// Encode encodes a UTF-8 string in src into text data, which are written to
// dst. Characters outside the range F000-F0FF are ignored.
var CodecPUA codecPUA

type codecPUA struct{}

func (codecPUA) Name() string {
	return "UTF-8 PUA"
}

func (codecPUA) Decode(dst io.Writer, src io.Reader) (written int, err error) {
	bufin := make([]byte, 1024)
	bufout := make([]rune, 1024)
	for {
		n, e := src.Read(bufin)
		if e != nil && e != io.EOF {
			err = e
			return
		}
		if n > 0 {
			for i := 0; i < n; i++ {
				bufout[i] = rune(bufin[i]) + 0xF000
			}
			n, e := dst.Write([]byte(string(bufout[:n])))
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

func (codecPUA) Encode(dst io.Writer, src io.Reader) (written int, err error) {
	buf := bufio.NewReader(src)
	bufout := make([]byte, 1)
	for {
		r, n, e := buf.ReadRune()
		if e != nil && e != io.EOF {
			err = e
			return
		}
		if n > 0 && 0xF000 <= r && r <= 0xF0FF {
			bufout[0] = byte(r - 0xF000)
			n, e := dst.Write(bufout)
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
