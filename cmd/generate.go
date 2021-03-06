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
	"fmt"
	"log"

	"github.com/mailtruck/longma/generate"
	"github.com/spf13/cobra"
)

// generateSampleCmd represents the generateSample command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a sample csv.",
	Long: `Generates a sample csv file. By default you will get a csv file named generated.csv with 8 rows and 3141 lines/rows

	longma generate

	although the output is random, there is no random seed used so the output will always be the same for files with the same number of rows

	to specify the number of lines use -r or --rows-per-file

	longma geneate -r 314159

	to specify the file name use the -f or --file-name flag

	longma generate -f "awesome.csv"

	to generate a csv with gzip compression, specify filename with .csv.gz

	longma generate -f "awesome.csv.gz"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")

		flags := cmd.Flags()
		rowsPerFile, err := flags.GetInt64("rows-per-file")
		if err != nil {
			log.Fatal(err)
		}
		// this is really weird, but im reusing the same flag
		// to specify the rows per file in a split and also
		// in a generated file (will move away from that soon)
		// so for now if you generate a file, then try to split
		// it the default is you will get a new directory with
		// 1 file in it because the generated and split file lines
		// are the same, so for now if the default of 3141
		// is parsed from flags we are guessing you are using the
		// default, so we are going to change it to 31415 so a split
		// with default flags will actually do something if you do a
		// longngma generate
		// longma split generated.csv
		if rowsPerFile == int64(3141) {
			rowsPerFile = int64(31415)
		}

		filename, err := flags.GetString("file-name")
		if err != nil {
			log.Fatal(err)
		}

		err = generate.GenerateCSV(filename, rowsPerFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringP("file-name", "f", "generated.csv", "Specify file name of generated csv.")
}
