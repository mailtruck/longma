// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split /path/to/file/to/split",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			return
		}

		rowsPerFile := 10000

		fmt.Println("split called")
		fmt.Println(args[0])
		inPath := args[0]

		err := split(inPath, rowsPerFile)
		if err != nil {
			log.Fatal(err)
		}

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

func split(path string, rowsPerFile int) error {
	fmt.Println(path)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	cr := csv.NewReader(f)
	cr.LazyQuotes = true
	header, err := cr.Read()
	if err != nil {
		return err
	}

	i := 0
	for {

		fileIndex := strconv.Itoa(i + 1)
		outFilename := path[:len(path)-4] + "_" + fileIndex + ".csv"
		out, err := os.OpenFile(outFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 666)
		if err != nil {
			return err
		}
		defer out.Close()

		cw := csv.NewWriter(out)
		defer cw.Flush()

		// read before writing the header just incase we are about
		// to hit EOF so we don't end up with a header only file
		row1, err := cr.Read()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		cw.Write(header)
		cw.Write(row1)

		for j := 0; j < rowsPerFile; j++ {
			row, err := cr.Read()
			if err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}

			cw.Write(row)
		}

		i++
	}
}
