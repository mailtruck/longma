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
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

// generateSampleCmd represents the generateSample command
var generateSampleCmd = &cobra.Command{
	Use:   "generateSample",
	Short: "Generate a sample csv.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateSample called")
		generateSample()
	},
}

func init() {
	RootCmd.AddCommand(generateSampleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateSampleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateSampleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateSample() {

	numLines := 12345678
	filename := "sample.csv"

	header := []string{"Chien Qián", "Sun Xùn", "K'an", "Ken Gèn", "K'un Kūn", "Chen Zhèn", "Li", "Tui Duì"}

	out, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Fatal(err)
	}

	cw := csv.NewWriter(out)
	defer cw.Flush()

	cw.Write(header)

	for i := 0; i < numLines; i++ {
		cw.Write(randomRow())
	}

}
func randomRow() []string {
	ret := []string{}
	for i := 0; i < 8; i++ {
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
