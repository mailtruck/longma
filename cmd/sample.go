// Copyright © 2017 Brian Danowski <briandanowski@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// sampleCmd represents the split command
var sampleCmd = &cobra.Command{
	Use:   "sample /path/to/file/to/split",
	Short: "make a smaller dataset",
	Long: `get the first chunk of a directory of files

	       use -r to specify rows per file
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Usage()
			return
		}
		flags := cmd.Flags()
		rowsPerFile, err := flags.GetInt64("rows-per-file")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("sample called")
		fmt.Println(args[0])
		inPath := args[0]

		_, err = sample(inPath, rowsPerFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(sampleCmd)
}

func sampleFile(sampleFile, sampleFolder string, rowsPerFile int64) (string, error) {
	sampleOut := filepath.Join(sampleFolder, filepath.Base(sampleFile))

	f, err := os.Open(sampleFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var cr *csv.Reader
	if filepath.Ext(filepath.Base(sampleFile)) == ".gz" {
		gr, err := gzip.NewReader(f)
		if err != nil {
			log.Fatal(err)
		}

		cr = csv.NewReader(gr)
	} else {
		cr = csv.NewReader(f)
	}

	cr.LazyQuotes = true
	out, err := os.OpenFile(sampleOut, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer out.Close()

	cw := csv.NewWriter(out)
	defer cw.Flush()

	for i := int64(0); i < rowsPerFile; i++ {
		row, err := cr.Read()
		if err == io.EOF {
			return sampleFolder, nil
		} else if err != nil {
			return "", err
		}

		cw.Write(row)
	}

	return sampleFolder, nil

}

func sample(path string, rowsPerFile int64) (string, error) {
	fmt.Println(path)

	base := filepath.Base(path)

	fmt.Println(base)

	abs, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(abs)

	sampleFolder := filepath.Join(abs[:len(abs)-len(base)], base+"-sample")

	fmt.Println(sampleFolder)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(sampleFolder); os.IsNotExist(err) {
		os.Mkdir(sampleFolder, 0700)
	}

	for _, file := range files {
		full := filepath.Join(abs, file.Name())
		sampleFile(full, sampleFolder, rowsPerFile)
	}

	return sampleFolder, nil
}
