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
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var err error

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split /path/to/file/to/split",
	Short: "split a file into chunks",
	Long: `Split a file into chunks
	
			use -r to specify rows per file
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Usage()
			return
		}
		flags := cmd.Flags()
		lazy, err := flags.GetBool("lazy-quotes")
		if err != nil {
			log.Fatal(err)
		}

		rowsPerFile, err := flags.GetInt64("rows-per-file")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("split called")
		fmt.Println(args[0])
		inPath := args[0]

		path, err := split(inPath, rowsPerFile, lazy)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(inPath + " split into " + path)

	},
}

func init() {
	RootCmd.AddCommand(splitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// splitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// splitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func getBaseBase(base string) (string, error) {
	baseBase := base[0 : len(base)-len(filepath.Ext(base))]
	ext := filepath.Ext(baseBase)
	switch ext {
	case ".csv":
		return getBaseBase(baseBase)
	case "":
		return baseBase, nil
	default:
		return "", errors.New("Unknown file ext. Known file exts are .csv and .csv.gz")
	}
}

func getFolderName(base string) string {

	baseBase, err := getBaseBase(base)
	if err != nil {
		log.Fatal(err)
	}

	ret := baseBase + "-split"
	if _, err := os.Stat(ret); err != nil {
		return ret

	}

	// this needs its own function, but tonight i need to go to seleep
	i := 1
	for {
		if _, err = os.Stat(ret + "_" + strconv.Itoa(i)); os.IsNotExist(err) {
			return ret + "_" + strconv.Itoa(i)
		}
		i++

	}
}

// split returns the path of folder that the files were split into
func split(path string, rowsPerFile int64, lazyQuotes bool) (string, error) {
	fmt.Println(path)

	base := filepath.Base(path)

	folderName := getFolderName(base)

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var cr *csv.Reader
	if filepath.Ext(filepath.Base(path)) == ".gz" {
		gr, err := gzip.NewReader(f)
		if err != nil {
			log.Fatal(err)
		}

		cr = csv.NewReader(gr)
	} else {
		cr = csv.NewReader(f)
	}

	cr.LazyQuotes = lazyQuotes
	header, err := cr.Read()
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.Mkdir(folderName, 0700)
	}

	i := 0
	for {

		fileIndex := strconv.Itoa(i + 1)
		outFilename := base[:len(base)-4] + "_" + fileIndex + ".csv"
		out, err := os.OpenFile(filepath.Join(folderName, outFilename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return "", err
		}

		cw := csv.NewWriter(out)

		// read before writing the header just incase we are about
		// to hit EOF so we don't end up with a header only file
		row1, err := cr.Read()
		if err == io.EOF {
			return "", nil
		} else if err != nil {
			return "", err
		}

		cw.Write(header)
		cw.Write(row1)

		for j := int64(0); j < rowsPerFile; j++ {
			row, err := cr.Read()
			if err == io.EOF {
				return folderName, nil
			} else if err != nil {
				return "", err
			}

			cw.Write(row)
		}
		// closing explicity because it's possible to have too many files open
		cw.Flush()
		err = out.Close()
		if err != nil {
			log.Fatal(errors.Wrap(err, "closing file"))
		}
		i++
	}

}
