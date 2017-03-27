// Copyright © 2017 bang <sqibang@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"os"

	"io/ioutil"

	"path/filepath"

	"github.com/bmatcuk/doublestar"
	"github.com/spf13/cobra"
)

var destination string
var pattern string

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "copy file to destinations",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		a, _ := doublestar.Glob(pattern)
		for _, f := range a {
			fi, err := os.Stat(f)
			d := fmt.Sprintf("%s/%s", destination, f)
			if err != nil {
				return
			}
			log("%s will be cp to: %s", f, d)
			switch mode := fi.Mode(); {
			case mode.IsDir():
				os.MkdirAll(d, fi.Mode())
			case mode.IsRegular():

				op := filepath.Dir(f)
				ops, err := os.Stat(op)
				if err != nil {
					log("check dir [%s] failure", op)
					return
				}
				// mkdir
				p := filepath.Dir(d)
				err = os.MkdirAll(p, ops.Mode())
				if err != nil {
					log("mkdir %s", p)
					return
				}
				fi, err := ioutil.ReadFile(f)
				if err != nil {
					log("Error", err)
					return
				}
				df, err := os.Create(d)
				_, err = df.Write(fi)
				if err != nil {
					log("cp file %s failure", d)
				}
				log("%s", f)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(cpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cpCmd.Flags().StringVarP(&destination, "destination", "d", "", "filesystem path to copy files to")
	cpCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "glob pattern for cp")
}
