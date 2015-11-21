package gen3_test

import (
	"bytes"
	"errors"
	"github.com/anaminus/pkm"
	"github.com/anaminus/pkm/gen3"
	"io"
	"testing"
)

type dummyIO struct {
	buf          []byte
	err          error
	now          bool
	readBytes    int
	writtenBytes int
	writeCalled  bool
}

func (d *dummyIO) Read(p []byte) (n int, err error) {
	if d.err != nil && !d.now {
		return 0, d.err
	}
	if d.readBytes >= len(d.buf) {
		return 0, io.EOF
	}
	n = copy(p, d.buf[d.readBytes:])
	d.readBytes += n
	if d.now && d.readBytes >= len(d.buf) {
		err = io.EOF
	}
	if d.now && d.err != nil {
		err = d.err
	}
	return
}

func (d *dummyIO) Write(p []byte) (n int, err error) {
	d.writeCalled = true
	n = len(p)
	d.writtenBytes += n
	if d.err != nil {
		err = d.err
	}
	return
}

func (d *dummyIO) Copy() *dummyIO {
	c := *d
	return &c
}

func testIO(t *testing.T, codec pkm.Codec, testName string, readerTmpl *dummyIO) {
	var reader *dummyIO
	var writer *dummyIO
	var n int
	var err error

	// Test codec.Decode.
	writer = &dummyIO{}
	reader = readerTmpl.Copy()
	n, err = codec.Decode(writer, reader)
	// Check if the correct number of bytes were returned.
	if writer.writeCalled && n != writer.writtenBytes {
		t.Errorf("test writer decode: '%s' failed: expected %d bytes written, got %d", testName, writer.writtenBytes, n)
	} else if !writer.writeCalled && n != 0 {
		t.Errorf("test writer decode: '%s' failed: expected %d bytes written, got %d", testName, 0, n)
	}
	// Check if an error expected from the reader is returned.
	if reader.err != nil && err == nil {
		t.Errorf("test reader decode: '%s' failed: expected error", testName)
	} else if reader.err == nil && err != nil {
		t.Errorf("test reader decode: '%s' failed: expected no error, got '%s'", testName, err)
	}

	// Test codec.Encode.
	writer = &dummyIO{}
	reader = readerTmpl.Copy()
	n, err = codec.Encode(writer, reader)
	// Check if the correct number of bytes were returned.
	if writer.writeCalled && n != writer.writtenBytes {
		t.Errorf("test writer encode: '%s' failed: expected %d bytes written, got %d", testName, writer.writtenBytes, n)
	} else if !writer.writeCalled && n != 0 {
		t.Errorf("test writer encode: '%s' failed: expected %d bytes written, got %d", testName, 0, n)
	}
	// Check if an error expected from the reader is returned.
	if reader.err != nil && err == nil {
		t.Errorf("test reader decode: '%s' failed: expected error", testName)
	} else if reader.err == nil && err != nil {
		t.Errorf("test reader decode: '%s' failed: expected no error, got '%s'", testName, err)
	}

	// Test codec.Decode with an erroring writer.
	writer = &dummyIO{err: errors.New("writer error")}
	reader = readerTmpl.Copy()
	n, err = codec.Decode(writer, reader)
	// Check if the correct number of bytes were returned.
	if writer.writeCalled && n != writer.writtenBytes {
		t.Errorf("test error writer decode: '%s' failed: expected %d bytes written, got %d", testName, writer.writtenBytes, n)
	} else if !writer.writeCalled && n != 0 {
		t.Errorf("test error writer decode: '%s' failed: expected %d bytes written, got %d", testName, 0, n)
	}
	// Check if an error expected from the writer is returned.
	if reader.err == nil && writer.writeCalled && err == nil {
		t.Errorf("test error writer decode: '%s' failed: expected error", testName)
	} else if reader.err != nil && writer.writeCalled && err == nil {
		t.Errorf("test error writer decode: '%s' failed: expected reader error", testName)
	}

	// Test codec.Encode with an erroring writer.
	writer = &dummyIO{err: errors.New("writer error")}
	reader = readerTmpl.Copy()
	n, err = codec.Encode(writer, reader)
	// Check if the correct number of bytes were returned.
	if writer.writeCalled && n != writer.writtenBytes {
		t.Errorf("test error writer encode: '%s' failed: expected %d bytes written, got %d", testName, writer.writtenBytes, n)
	} else if !writer.writeCalled && n != 0 {
		t.Errorf("test error writer encode: '%s' failed: expected %d bytes written, got %d", testName, 0, n)
	}
	// Check if an error expected from the writer is returned.
	if reader.err == nil && writer.writeCalled && err == nil {
		t.Errorf("test error writer encode: '%s' failed: expected error", testName)
	} else if reader.err != nil && writer.writeCalled && err == nil {
		t.Errorf("test error writer encode: '%s' failed: expected reader error", testName)
	}
}

func testCodec(t *testing.T, codec pkm.Codec) {
	if len(codec.Name()) == 0 {
		t.Errorf("Name: expected non-empty name")
	}

	size := 1024
	err := errors.New("reader error")

	testIO(t, codec, "real size", &dummyIO{buf: make([]byte, 12), now: false, err: nil})
	testIO(t, codec, "real size; now", &dummyIO{buf: make([]byte, 12), now: true, err: nil})
	testIO(t, codec, "real size; error", &dummyIO{buf: make([]byte, 12), now: false, err: err})
	testIO(t, codec, "real size; error now", &dummyIO{buf: make([]byte, 12), now: true, err: err})

	testIO(t, codec, "half size", &dummyIO{buf: make([]byte, size/2), now: false, err: nil})
	testIO(t, codec, "half size; now", &dummyIO{buf: make([]byte, size/2), now: true, err: nil})
	testIO(t, codec, "half size; error", &dummyIO{buf: make([]byte, size/2), now: false, err: err})
	testIO(t, codec, "half size; error now", &dummyIO{buf: make([]byte, size/2), now: true, err: err})

	testIO(t, codec, "full size", &dummyIO{buf: make([]byte, size), now: false, err: nil})
	testIO(t, codec, "full size; now", &dummyIO{buf: make([]byte, size), now: true, err: nil})
	testIO(t, codec, "full size; error", &dummyIO{buf: make([]byte, size), now: false, err: err})
	testIO(t, codec, "full size; error now", &dummyIO{buf: make([]byte, size), now: true, err: err})

	testIO(t, codec, "full+half size", &dummyIO{buf: make([]byte, size+size/2), now: false, err: nil})
	testIO(t, codec, "full+half size; now", &dummyIO{buf: make([]byte, size+size/2), now: true, err: nil})
	testIO(t, codec, "full+half size; error", &dummyIO{buf: make([]byte, size+size/2), now: false, err: err})
	testIO(t, codec, "full+half size; error now", &dummyIO{buf: make([]byte, size+size/2), now: true, err: err})

	testIO(t, codec, "double size", &dummyIO{buf: make([]byte, size*2), now: false, err: nil})
	testIO(t, codec, "double size; now", &dummyIO{buf: make([]byte, size*2), now: true, err: nil})
	testIO(t, codec, "double size; error", &dummyIO{buf: make([]byte, size*2), now: false, err: err})
	testIO(t, codec, "double size; error now", &dummyIO{buf: make([]byte, size*2), now: true, err: err})

	content := []byte("Test Data 1234 \uF000\uF080\uF0FF")
	testIO(t, codec, "real content", &dummyIO{buf: content, now: false, err: nil})
	testIO(t, codec, "real content; now", &dummyIO{buf: content, now: true, err: nil})
	testIO(t, codec, "real content; error", &dummyIO{buf: content, now: false, err: err})
	testIO(t, codec, "real content; error now", &dummyIO{buf: content, now: true, err: err})
}

func testReencode(t *testing.T, codec pkm.Codec) {
	content := make([]byte, 256)
	for i := range content {
		content[i] = byte(i)
	}
	var decodedbuf bytes.Buffer
	var resultbuf bytes.Buffer

	n, err := codec.Decode(&decodedbuf, bytes.NewReader(content))
	if err != nil {
		t.Errorf("Re-encode test: Decode: unexpected error '%s'", err)
		return
	}
	decoded := decodedbuf.Bytes()[:n]
	n, err = codec.Encode(&resultbuf, bytes.NewReader(decoded))
	if err != nil {
		t.Errorf("Re-encode test: Encode: unexpected error '%s'", err)
		return
	}
	result := resultbuf.Bytes()[:n]
	if !bytes.Equal(content, result) {
		var n int
		if len(content) < len(result) {
			n = len(content)
		} else {
			n = len(result)
		}
		for i := 0; i < n; i++ {
			if result[i] != content[i] {
				t.Errorf("Re-encode test: result does not match original (@%d: 0x%02X != 0x%02X)", i, result[i], content[i])
				return
			}
		}
		t.Errorf("Re-encode test: result length does not match original length (%d != %d)", len(result), len(content))
	}
}

func TestCodecASCII(t *testing.T) {
	testCodec(t, gen3.CodecASCII)
}
func TestCodecString(t *testing.T) {
	testCodec(t, gen3.CodecString)
}
func TestCodecUTF8(t *testing.T) {
	testCodec(t, gen3.CodecUTF8)
	testReencode(t, gen3.CodecUTF8)
}
func TestCodecPUA(t *testing.T) {
	testCodec(t, gen3.CodecPUA)
	testReencode(t, gen3.CodecPUA)
}
