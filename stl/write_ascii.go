package stl

import (
	"bufio"
	"errors"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// ToASCII writes the Solid out in ASCII form
func (s *Solid) ToASCII(w io.Writer) error {
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	_, err := bw.WriteString("solid " + s.Header + "\n")
	if err != nil {
		return errors.New("did not write header: " + err.Error())
	}

	for _, t := range s.Triangles {
		if _, err := bw.WriteString(triangleASCII(t)); err != nil {
			return errors.New("did not write triangle: " + err.Error())
		}
	}

	_, err = bw.WriteString("endsolid " + s.Header + "\n")
	if err != nil {
		return errors.New("did not write footer: " + err.Error())
	}

	return nil
}

// ToASCIIFile writes the Solid to a file in ASCII format
// See stl.ToASCII for more info
func (s *Solid) ToASCIIFile(filename string) error {
	file, err := os.OpenFile(strings.TrimSpace(filename), os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		return err
	}
	defer file.Close()

	return s.ToASCII(file)
}
func triangleASCII(t Triangle) string {
	s := " facet normal " + shortFloat(t.Normal.Ni) + " " + shortFloat(t.Normal.Nj) + " " + shortFloat(t.Normal.Nk) + "\n"
	s += "  outer loop\n"
	s += "   vertex " + shortFloat(t.Vertices[0].X) + " " + shortFloat(t.Vertices[0].Y) + " " + shortFloat(t.Vertices[0].Z) + "\n"
	s += "   vertex " + shortFloat(t.Vertices[1].X) + " " + shortFloat(t.Vertices[1].Y) + " " + shortFloat(t.Vertices[1].Z) + "\n"
	s += "   vertex " + shortFloat(t.Vertices[2].X) + " " + shortFloat(t.Vertices[2].Y) + " " + shortFloat(t.Vertices[2].Z) + "\n"
	s += "  endloop\n"
	s += " endfacet\n"
	return s
}
func shortFloat(f float32) string {
	// Scientific notation
	sn := strconv.FormatFloat(float64(f), 'g', -1, 32)

	// If f is an integer, and its shorter than scientific notation form, return an integer
	if float64(f) == math.Floor(float64(f)) {
		in := strconv.FormatFloat(float64(f), 'f', 0, 64)
		if len(sn) > len(in) {
			return in
		}
	}

	return sn
}
