package stl

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

func fromBinary(br *bufio.Reader) (Solid, error) {
	header, err := extractBinaryHeader(br)
	if err != nil {
		return Solid{}, err
	}

	triCount, err := extractBinaryTriangleCount(br)
	if err != nil {
		return Solid{}, err
	}
	tris, err := extractBinaryTriangles(triCount, br)
	if err != nil {
		return Solid{}, err
	}

	return Solid{
		Header:        header,
		TriangleCount: triCount,
		Triangles:     tris,
	}, nil
}

func extractBinaryHeader(br *bufio.Reader) (string, error) {
	hBytes := make([]byte, 80)
	_, err := br.Read(hBytes)
	if err != nil {
		return "", fmt.Errorf("could not read header: %v", err)
	}

	return strings.TrimSpace(string(hBytes)), nil
}

func extractBinaryTriangleCount(br *bufio.Reader) (uint32, error) {
	cntBytes := make([]byte, 4)
	_, err := br.Read(cntBytes)
	if err != nil {
		return 0, fmt.Errorf("could not read triangle count: %v", err)
	}

	return binary.LittleEndian.Uint32(cntBytes), nil
}

// Each triangle is 50 bytes.
func extractBinaryTriangles(triCount uint32, br *bufio.Reader) (tris []Triangle, err error) {
	// Create Scanner with split func for binary triangles
	scanner := bufio.NewScanner(br)
	scanner.Split(splitTrianglesBinary)

	// Parse the triangles
	var t []Triangle
	for scanner.Scan() {
		t, err = parseChunksOfBinary(scanner.Bytes())
		if err != nil {
			return
		}
		for _, j := range t {
			tris = append(tris, j)
		}
	}
	return
}

func parseChunksOfBinary(raw []byte) (triParsed []Triangle, err error) {
	t := make([]Triangle, 0, len(raw)/50)
	for i := 0; i < len(raw); i += 50 {
		t = append(t, triangleFromBinary(raw[i:i+50]))
	}

	return t, err
}

func triangleFromBinary(bin []byte) Triangle {
	return Triangle{
		Normal: unitVectorFromBinary(bin[0:12]),
		Vertices: [3]Coordinate{
			coordinateFromBinary(bin[12:24]),
			coordinateFromBinary(bin[24:36]),
			coordinateFromBinary(bin[36:48]),
		},
		AttrByteCnt: uint16(bin[48])<<8 | uint16(bin[49]),
	}
}

func coordinateFromBinary(bin []byte) Coordinate {
	return Coordinate{
		X: math.Float32frombits(binary.LittleEndian.Uint32(bin[0:4])),
		Y: math.Float32frombits(binary.LittleEndian.Uint32(bin[4:8])),
		Z: math.Float32frombits(binary.LittleEndian.Uint32(bin[8:12])),
	}
}

func unitVectorFromBinary(bin []byte) UnitVector {
	return UnitVector{
		Ni: math.Float32frombits(binary.LittleEndian.Uint32(bin[0:4])),
		Nj: math.Float32frombits(binary.LittleEndian.Uint32(bin[4:8])),
		Nk: math.Float32frombits(binary.LittleEndian.Uint32(bin[8:12])),
	}
}
