package uuid

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"strconv"
	"strings"
	"time"
)

type UUID []byte

func Func(t string, n int) func() (UUID, error) {
	return func() (UUID, error) {
		return New(t, n)
	}
}

func New(t string, n int) (UUID, error) {
	mid, err := HostID()
	if err != nil {
		return nil, fmt.Errorf("uid: host: %w", err)
	}
	for len(t) < 4 {
		t = "_" + t
	}
	b := bytes.NewBuffer(nil)
	if _, err := io.CopyN(b, bytes.NewReader([]byte(t)), 4); err != nil {
		return nil, fmt.Errorf("uid: type: %w", err)
	}
	if err := binary.Write(b, binary.BigEndian, crc32.ChecksumIEEE([]byte(mid))); err != nil {
		return nil, fmt.Errorf("uid: host: %w", err)
	}
	if err := binary.Write(b, binary.BigEndian, uint32(time.Now().In(time.UTC).Unix())); err != nil {
		return nil, fmt.Errorf("uid: time: %w", err)
	}
	if _, err := io.CopyN(b, rand.Reader, int64(n)); err != nil {
		return nil, fmt.Errorf("uid: data: %w", err)
	}
	if err := binary.Write(b, binary.BigEndian, crc32.ChecksumIEEE(b.Bytes())); err != nil {
		return nil, fmt.Errorf("uid: host: %w", err)
	}
	if b.Len() != 16+n {
		return nil, fmt.Errorf("uid: length: %w", io.ErrShortWrite)
	}
	return b.Bytes(), nil
}

func Parse(s string) (UUID, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 5 {
		return nil, fmt.Errorf("uid: parts: expected 5, got %d", len(parts))
	}
	t := parts[0]
	if len(t) != 4 {
		return nil, fmt.Errorf("uid: type: expected 4, got %d", len(t))
	}
	host, err := hex.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("uid: host: %w", err)
	}
	secs, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("uid: time: %w", err)
	}
	unix := make([]byte, 4)
	binary.BigEndian.PutUint32(unix, uint32(secs))

	data, err := hex.DecodeString(parts[3])
	if err != nil {
		return nil, fmt.Errorf("uid: data: %w", err)
	}
	crc, err := hex.DecodeString(parts[4])
	if err != nil {
		return nil, fmt.Errorf("uid: crc: %w", err)
	}
	return bytes.Join([][]byte{
		[]byte(t),
		host,
		unix,
		data,
		crc,
	}, nil), nil
}

func Decode(s string) (UUID, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("uid: decode: %w", err)
	}
	return b, nil
}

// Type returns the type of the UUID.
func (uid UUID) Type() string {
	return string(uid[0:4])
}

// Host returns the host identifier part of the UUID.
func (uid UUID) Host() []byte {
	return uid[4:8]
}

// Time returns the time part of the UUID.
func (uid UUID) Time() time.Time {
	return time.Unix(int64(binary.BigEndian.Uint32(uid[8:12])), 0).In(time.UTC)
}

// Data returns the random data part of the UUID.
func (uid UUID) Data() []byte {
	return uid[12 : len(uid)-4]
}

func (uid UUID) CRC() []byte {
	return uid[len(uid)-4 : len(uid)]
}

// Marshal returns a byte slice representation of the UUID.
func (uid UUID) Marshal() []byte {
	return uid[:]
}

// Encode returns a base64 encoded string representation of the UUID.
func (uid UUID) Encode() string {
	return base64.RawURLEncoding.EncodeToString(uid[:])
}

// String returns a string representation of the UUID in the following format:
// "AAAA-BBBBBBBBBBB-CCCCCC-DDDDDDDDDD-EEEE"
// where:
// - A is the type of the UUID
// - B is the host identifier
// - C are unix seconds
// - D is the random data
// - E is a crc32 checksum
func (uid UUID) String() string {
	return fmt.Sprintf("%.4s-%X-%d-%X-%X", uid.Type(), uid.Host(), uint32(uid.Time().Unix()), uid.Data(), uid.CRC())
}
