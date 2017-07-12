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
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// headerCmd represents the header command
var headerCmd = &cobra.Command{
	Use:   "header",
	Short: "Print the header and the first line of input file",
	Long: `Print out the header and the first line of the file with column indexes
	
			You may pass in a .csv or .csv.gz file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			return
		}
		path := args[0]
		fmt.Println("\nPrinting header for " + path + "\n")

		flags := cmd.Flags()
		lazy, err := flags.GetBool("lazy-quotes")
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
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

		cr.LazyQuotes = lazy
		cr.FieldsPerRecord = -1
		header, err := cr.Read()
		if err != nil {
			log.Fatal(err)
		}
		for i, field := range header {
			fmt.Println(strconv.Itoa(i) + ": " + field)
		}
		line2, err := cr.Read()
		if err != nil {
			log.Fatal(err)
		}
		for i, field := range line2 {
			fmt.Println(strconv.Itoa(i) + ": " + field)
		}
	},
}

func init() {
	RootCmd.AddCommand(headerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// headerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// headerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
