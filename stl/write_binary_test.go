package stl

import (
	"bytes"
	"strconv"
	"testing"
)

func Test_headerBinary(t *testing.T) {
	for _, tst := range []struct {
		h        string
		expected []byte
	}{
		{
			h:        "",
			expected: make([]byte, 80),
		},
		{
			h:        "This is a header",
			expected: []byte{0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			h:        "This header is too long, so it will be trimmed.  This header is too long, so it will be trimmed.",
			expected: []byte{0x54, 0x68, 0x69, 0x73, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x74, 0x6f, 0x6f, 0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x2c, 0x20, 0x73, 0x6f, 0x20, 0x69, 0x74, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x74, 0x72, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x2e, 0x20, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x74, 0x6f, 0x6f, 0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x2c, 0x20, 0x73, 0x6f, 0x20, 0x69, 0x74, 0x20},
		},
	} {
		tst := tst
		t.Run(tst.h, func(t *testing.T) {
			t.Parallel()
			got := headerBinary(tst.h)
			if !bytes.Equal(got, tst.expected) {
				t.Errorf("Expecting %x, got %x", tst.expected, got)
			}
		})
	}
}
func Test_triCountBinary(t *testing.T) {
	for _, tst := range []struct {
		c        uint32
		expected []byte
	}{
		{
			c:        0,
			expected: []byte{0x00, 0x00, 0x00, 0x00},
		},
		{
			c:        500,
			expected: []byte{0xf4, 0x01, 0x00, 0x00},
		},
		{
			c:        1000222,
			expected: []byte{0x1e, 0x43, 0x0f, 0x00},
		},
	} {
		tst := tst
		t.Run(strconv.FormatFloat(float64(tst.c), 'f', 0, 64), func(t *testing.T) {
			t.Parallel()
			got := triCountBinary(tst.c)
			if !bytes.Equal(got, tst.expected) {
				t.Errorf("Expecting %x, got %x", tst.expected, got)
			}
		})
	}
}
func Test_triangleBinary(t *testing.T) {
	for _, tst := range []struct {
		t        Triangle
		expected []byte
	}{
		{
			t: Triangle{
				Normal: UnitVector{
					Ni: 0,
					Nj: 0,
					Nk: 0,
				},
				Vertices: [3]Coordinate{
					{
						X: 0,
						Y: 0,
						Z: 0,
					},
					{
						X: 0,
						Y: 0,
						Z: 0,
					},
					{
						X: 0,
						Y: 0,
						Z: 0,
					},
				},
				AttrByteCnt: 0,
			},
			expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			t: Triangle{
				Normal: UnitVector{
					Ni: 5,
					Nj: 0,
					Nk: 0,
				},
				Vertices: [3]Coordinate{
					{
						X: 0,
						Y: 2,
						Z: 0,
					},
					{
						X: 0,
						Y: 0,
						Z: 1,
					},
					{
						X: 123,
						Y: 0,
						Z: 0,
					},
				},
				AttrByteCnt: 0,
			},
			expected: []byte{0x00, 0x00, 0xa0, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x3f, 0x00, 0x00, 0xf6, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			t: Triangle{
				Normal: UnitVector{
					Ni: 0,
					Nj: 0,
					Nk: 0,
				},
				Vertices: [3]Coordinate{
					{
						X: 0,
						Y: 0,
						Z: 0,
					},
					{
						X: 0,
						Y: 0,
						Z: 0,
					},
					{
						X: 0,
						Y: 0,
						Z: 0,
					},
				},
				AttrByteCnt: 5,
			},
			expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00},
		},
	} {
		tst := tst
		t.Run("triangleBinary", func(t *testing.T) {
			t.Parallel()
			got := triangleBinary(tst.t)
			if !bytes.Equal(got, tst.expected) {
				t.Errorf("Expecting %x, got %x", tst.expected, got)
			}
		})
	}
}
