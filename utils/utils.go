package utils

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/pkg/errors"
)

var err error

// StringMap takes a slice of strings and a function and returns
// a new slice with the function aplied to each member of the slice
func StringMap(ss []string, f func(string) string) []string {
	ret := make([]string, len(ss), len(ss))

	for i, s := range ss {
		ret[i] = f(s)
	}

	return ret
}

func Cleanup(path string) {

	err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}

}

func GetIndexOfColumn(header []string, column string) (int, error) {
	for i, field := range header {
		if field == column {
			return i, nil
		}
	}
	return 0, errors.New("column not found in header")
}

// StringMapToSlice returns a sorted slice of strings from given map
func StringMapToSlice(in map[string]string) []string {
	ret := []string{}

	for _, val := range in {
		ret = append(ret, val)
	}

	sort.Strings(ret)

	return ret

}

// SliceToString returns a comma separate list as a string when given a slice of strings
func SliceToString(imploded []string) string {
	ret := ""
	for i, item := range imploded {
		if i > 0 {
			ret = ret + ","
		}
		ret = ret + item
	}
	return ret
}

func CountLinesInCSV(path string) (int64, error) {

	r, err := NewReader(path)
	if err != nil {
		return 0, err
	}

	i := int64(0)
	for {
		_, err = r.Read()
		if err == io.EOF {
			break

		} else if err != nil {
			log.Fatal(err)
		}

		i++

	}
	return i, nil

}

type CSVReader struct {
	File *os.File
	Gr   *gzip.Reader
	Cr   *csv.Reader
}

func NewReader(filePath string) (CSVReader, error) {
	ret := CSVReader{}
	ret.File, err = os.Open(filePath)
	if err != nil {
		return ret, err
	}

	switch filepath.Ext(filePath) {
	// assuming .csv.gz
	case ".gz":
		ret.Gr, err = gzip.NewReader(ret.File)

		if err != nil {
			// not sure if its good to close this if theres an error
			ret.File.Close()
			return ret, errors.Wrap(err, "couldn't open our file to read gzip")
		}
		ret.Cr = csv.NewReader(ret.Gr)
		break

	case ".csv":
		ret.Cr = csv.NewReader(ret.File)
		break

	default:
		ret.File.Close()
		return ret, errors.New("invalid file extension")
		break
	}

	return ret, nil

}

func (r CSVReader) Line() ([]string, error) {
	ret, err := r.Cr.Read()
	if err != nil {
		return []string{}, err
	}
	return ret, nil
}

func (r CSVReader) Read() ([]string, error) {
	return r.Line()
}

func (r CSVReader) Close() error {
	if r.Gr != nil {
		err = r.Gr.Close()
		if err != nil {
			return err
		}
	}
	err = r.File.Close()
	if err != nil {
		return err
	}

	return nil
}
