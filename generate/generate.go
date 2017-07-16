package generate

import (
	"bufio"
	"compress/gzip"
	"encoding/csv"
	"encoding/hex"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// GenerateCSV generates a csv. Lines includes header
func GenerateCSV(filename string, lines int64) error {

	header := []string{"Chien Qián", "Sun Xùn", "K'an", "Ken Gèn", "K'un Kūn", "Chen Zhèn", "Li", "Tui Duì"}

	out, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return errors.Wrap(err, "error opening file to write")
	}
	defer out.Close()

	var cw *csv.Writer
	switch filepath.Ext(filename) {
	case ".gz":

		gw := gzip.NewWriter(out)
		if err != nil {
			return errors.Wrap(err, "error creating gzip reader")
		}
		defer gw.Close()

		bw := bufio.NewWriter(gw)
		defer bw.Flush()

		cw = csv.NewWriter(bw)
		break
	case ".csv":
		cw = csv.NewWriter(out)
		break
	default:
		return errors.New("not a CSV or GZ file")
	}
	defer cw.Flush()

	cw.Write(header)

	for i := int64(1); i < lines; i++ {
		cw.Write(randomRow(header))
	}

	return nil

}
func randomRow(header []string) []string {
	ret := []string{}
	for i := 0; i < len(header); i++ {
		ret = append(ret, randomString(rand.Intn(10)))
	}

	return ret
}

func randomString(len int) string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b[:5])[:len]
}
