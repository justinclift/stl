package stl

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

func fromASCII(br *bufio.Reader) (Solid, error) {
	header, err := extractASCIIHeader(br)
	if err != nil {
		return Solid{}, err
	}

	tris, err := extractASCIITriangles(br)
	if err != nil {
		return Solid{}, err
	}

	return Solid{
		Header:        header,
		TriangleCount: uint32(len(tris)),
		Triangles:     tris,
	}, nil
}

func extractASCIIHeader(br *bufio.Reader) (string, error) {
	s, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.TrimPrefix(string(s), "solid")), nil
}

func extractASCIITriangles(br *bufio.Reader) (t []Triangle, err error) {
	// Create Scanner with split func for ASCII triangles
	scanner := bufio.NewScanner(br)
	scanner.Split(splitTrianglesASCII)

	// Parse the triangles
	tris := make([]Triangle, 0, 1024)
	var triParsed Triangle
	for scanner.Scan() {
		triParsed, err = parseTriangles(scanner.Text())
		if err != nil {
			return
		}
		tris = append(tris, triParsed)
	}
	return tris, nil
}

func parseTriangles(input string) (triParsed Triangle, err error) {
	sl := ourSplit(strings.TrimSpace(input), '\n')

	// Get the normal for a triangle
	var norm UnitVector
	norm, err = extractUnitVector(sl[0])
	if err != nil {
		return
	}

	// Get coordinates
	var v [3]Coordinate
	for i := 0; i < 3; i++ {
		v[i], err = extractCoordinate(sl[i+2])
		if err != nil {
			return
		}
	}
	triParsed.Normal = norm
	triParsed.Vertices = v
	return
}

func extractCoordinate(s string) (Coordinate, error) {
	sl := ourSplit(strings.TrimSpace(s), ' ')
	if len(sl) != 4 {
		return Coordinate{}, errors.New("invalid input for coordinate: " + s)
	}

	x, err := strconv.ParseFloat(sl[1], 32)
	if err != nil {
		return Coordinate{}, err
	}
	y, err := strconv.ParseFloat(sl[2], 32)
	if err != nil {
		return Coordinate{}, err
	}
	z, err := strconv.ParseFloat(sl[3], 32)
	if err != nil {
		return Coordinate{}, err
	}

	return Coordinate{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}, nil
}

func extractUnitVector(s string) (UnitVector, error) {
	sl := ourSplit(strings.TrimSpace(s), ' ')
	if len(sl) != 5 {
		return UnitVector{}, errors.New("invalid input for unit vector: " + s)
	}

	i, err := strconv.ParseFloat(sl[2], 32)
	if err != nil {
		return UnitVector{}, err
	}
	j, err := strconv.ParseFloat(sl[3], 32)
	if err != nil {
		return UnitVector{}, err
	}
	k, err := strconv.ParseFloat(sl[4], 32)
	if err != nil {
		return UnitVector{}, err
	}

	return UnitVector{
		Ni: float32(i),
		Nj: float32(j),
		Nk: float32(k),
	}, nil
}

// Quick and dirty initial attempt at replacement split function (written at ~3:30am!), to avoid using strings.Split()
// which is currently triggering a TinyGo bug: https://github.com/tinygo-org/tinygo/issues/699
func ourSplit(input string, sep rune) (val []string) {
	var s string
	var endAdded bool
	for _, j := range input {
		if j != sep {
			// Not the separator character
			if j != ' ' || len(s) != 0 { // Only add a space character if it's not at the start
				s += string(j)
				endAdded = false
			}
		} else {
			// Separator character found.  If the current string has contents, add it to the string list
			if len(s) != 0 {
				val = append(val, s)
				s = ""
				endAdded = true
			}
		}
	}
	if !endAdded {
		val = append(val, s)
	}
	return
}
