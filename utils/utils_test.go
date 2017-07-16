package utils

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/mailtruck/longma/generate"
	"github.com/stretchr/testify/assert"
)

func TestSliceToString(t *testing.T) {
	in := []string{
		"one",
		"two",
		"three",
	}

	out := SliceToString(in)

	assert.Equal(t, "one,two,three", out)
}

func TestStringMap(t *testing.T) {
	f := func(s string) string {
		return s + "bar"
	}

	input := []string{
		"one",
		"two",
		"three",
	}

	expected := []string{
		"onebar",
		"twobar",
		"threebar",
	}

	testVal := StringMap(input, f)

	assert.Equal(t, expected[0], testVal[0])
	assert.Equal(t, expected[1], testVal[1])
	assert.Equal(t, expected[2], testVal[2])

	input = []string{
		"ds/dfds/one.csv",
		"ds/dsds/two.csv",
		"sd/dsd/three.csv",
	}

	expected = []string{
		"one.csv",
		"two.csv",
		"three.csv",
	}

	testVal = StringMap(input, filepath.Base)

	assert.Equal(t, expected[0], testVal[0])
	assert.Equal(t, expected[1], testVal[1])
	assert.Equal(t, expected[2], testVal[2])

}

func TestGetIndexOfColumn(t *testing.T) {
	header := []string{
		"dingo",
		"boshi",
		"ting",
		"tong",
	}

	index, err := GetIndexOfColumn(header, "dingo")
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, index)

	index, err = GetIndexOfColumn(header, "tong")
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, index)

	index, err = GetIndexOfColumn(header, "captain")
	assert.Equal(t, false, err == nil)

}

func TestCountLinesInCSV(t *testing.T) {
	path := "test_count.csv"
	err = generate.GenerateCSV(path, 10000)
	defer Cleanup(path)
	assert.Equal(t, nil, err)

	c, err := CountLinesInCSV(path)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(10000), c)

}

func TestNewReader(t *testing.T) {
	testFile := "test.csv"
	testFileLines := int64(33)

	_, err = os.Stat(testFile)
	if err == nil {
		Cleanup(testFile)
	}

	generate.GenerateCSV(testFile, testFileLines)
	defer Cleanup(testFile)

	r, err := NewReader(testFile)
	assert.Equal(t, nil, err)

	for i := int64(1); i <= testFileLines; i++ {

		line, err := r.Cr.Read()

		assert.Equal(t, nil, err)

		assert.Equal(t, true, len(line) > 1)

	}

	_, err = r.Cr.Read()

	assert.Equal(t, io.EOF, err)

	err = r.File.Close()

	assert.Equal(t, nil, err)

}

func TestLine(t *testing.T) {
	testFile := "test.csv"
	testFileLines := int64(33)

	_, err = os.Stat(testFile)
	if err == nil {
		Cleanup(testFile)
	}

	generate.GenerateCSV(testFile, testFileLines)
	defer Cleanup(testFile)

	r, err := NewReader(testFile)
	assert.Equal(t, nil, err)

	// read header and first 5 lines
	for i := int64(1); i <= testFileLines; i++ {

		line, err := r.Line()

		assert.Equal(t, nil, err)

		assert.Equal(t, true, len(line) > 1)

	}

	_, err = r.Line()

	assert.Equal(t, io.EOF, err)

	err = r.File.Close()

	assert.Equal(t, nil, err)

}

func TestLineGz(t *testing.T) {
	testFile := "test.csv.gz"
	testFileLines := int64(33)

	_, err = os.Stat(testFile)
	if err == nil {
		Cleanup(testFile)
	}

	generate.GenerateCSV(testFile, testFileLines)
	defer Cleanup(testFile)

	r, err := NewReader(testFile)
	assert.Equal(t, nil, err)

	// read header and first 5 lines
	for i := int64(0); i <= testFileLines; i++ {

		line, err := r.Line()

		assert.Equal(t, nil, err)

		assert.Equal(t, true, len(line) > 1)

	}

	_, err = r.Line()

	assert.Equal(t, io.EOF, err)

	err = r.File.Close()

	assert.Equal(t, nil, err)

}

func TestCSVGZ(t *testing.T) {
	testFile := "test.csv.gz"
	testFileLines := int64(33)

	_, err = os.Stat(testFile)
	if err == nil {
		Cleanup(testFile)
	}

	generate.GenerateCSV(testFile, testFileLines)
	defer Cleanup(testFile)

	r, err := NewReader(testFile)

	assert.Equal(t, nil, err)

	for i := int64(1); i <= testFileLines; i++ {

		line, err := r.Cr.Read()
		assert.Equal(t, nil, err)
		assert.Equal(t, true, len(line) > 1)

	}

	_, err = r.Cr.Read()

	assert.Equal(t, io.EOF, err)

	err = r.File.Close()

	assert.Equal(t, nil, err)

}

func TestClose(t *testing.T) {
	testFile := "test.csv.gz"
	testFileLines := int64(33)

	_, err := os.Stat(testFile)
	if err == nil {
		Cleanup(testFile)
	}

	generate.GenerateCSV(testFile, testFileLines)
	defer Cleanup(testFile)

	r, err := NewReader(testFile)
	assert.Equal(t, nil, err)

	err = r.Close()

	assert.Equal(t, nil, err)

}
