package stl

import (
	"bufio"
	"fmt"
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
	sl := strings.Split(input, "\n")

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
	sl := strings.Split(strings.TrimSpace(s), " ")
	if len(sl) != 4 {
		return Coordinate{}, fmt.Errorf("invalid input for coordinate: %s", strings.TrimSpace(s))
	}

	x, err := strconv.ParseFloat(sl[1], 32)
	if err != nil {
		return Coordinate{}, fmt.Errorf("invalid input for coordinate x: %v", err)
	}
	y, err := strconv.ParseFloat(sl[2], 32)
	if err != nil {
		return Coordinate{}, fmt.Errorf("invalid input for coordinate y: %v", err)
	}
	z, err := strconv.ParseFloat(sl[3], 32)
	if err != nil {
		return Coordinate{}, fmt.Errorf("invalid input for coordinate z: %v", err)
	}

	return Coordinate{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}, nil
}

func extractUnitVector(s string) (UnitVector, error) {
	sl := strings.Split(strings.TrimSpace(s), " ")
	if len(sl) != 5 {
		return UnitVector{}, fmt.Errorf("invalid input for unit vector: %s", strings.TrimSpace(s))
	}

	i, err := strconv.ParseFloat(sl[2], 32)
	if err != nil {
		return UnitVector{}, fmt.Errorf("invalid input for unit vector i: %v", err)
	}
	j, err := strconv.ParseFloat(sl[3], 32)
	if err != nil {
		return UnitVector{}, fmt.Errorf("invalid input for unit vector j: %v", err)
	}
	k, err := strconv.ParseFloat(sl[4], 32)
	if err != nil {
		return UnitVector{}, fmt.Errorf("invalid input for unit vector k: %v", err)
	}

	return UnitVector{
		Ni: float32(i),
		Nj: float32(j),
		Nk: float32(k),
	}, nil
}
