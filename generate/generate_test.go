package generate

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"log"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCSV(t *testing.T) {

	filename := "testCSV.csv"

	err := GenerateCSV(filename, int64(100))

	assert.Equal(t, nil, err)

	info, err := os.Stat(filename)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, info.Size() > 0)

	file, err := os.Open(filename)
	assert.Equal(t, nil, err)
	defer file.Close()

	cr := csv.NewReader(file)

	line1, err := cr.Read()
	assert.Equal(t, nil, err)
	spew.Dump(len(line1))

	assert.Equal(t, true, len(line1) > 0)

	err = os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}

}

func TestGenerateLinesNumber(t *testing.T) {

	filename := "testCSV.csv"

	err := GenerateCSV(filename, int64(100))

	assert.Equal(t, nil, err)

	info, err := os.Stat(filename)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, info.Size() > 0)

	file, err := os.Open(filename)
	assert.Equal(t, nil, err)
	defer file.Close()

	cr := csv.NewReader(file)
	i := 0
	for {
		_, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		i++
	}
	assert.Equal(t, 100, i)

	err = os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}

}

func TestGenerateCSVGZ(t *testing.T) {

	filename := "testCSV.csv.gz"

	err := GenerateCSV(filename, int64(1000))

	assert.Equal(t, nil, err)

	info, err := os.Stat(filename)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, info.Size() > 0)

	file, err := os.Open(filename)
	assert.Equal(t, nil, err)
	defer file.Close()

	gr, err := gzip.NewReader(file)
	assert.Equal(t, nil, err)
	defer gr.Close()

	cr := csv.NewReader(gr)

	line1, err := cr.Read()
	assert.Equal(t, nil, err)
	spew.Dump(len(line1))

	assert.Equal(t, true, len(line1) > 0)

	err = os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}

}
