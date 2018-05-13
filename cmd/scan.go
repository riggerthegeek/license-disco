// Copyright © 2018 Simon Emms <simon@simonemms.com>
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

	"github.com/spf13/cobra"

	"github.com/riggerthegeek/license-disco/scanner"
	"github.com/riggerthegeek/license-disco/languages/javascript"
)

//var d scanner.Scanner

var d = scanner.Scanner{}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "This will scan the project for license data",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scan called")
		//fmt.Println(cmd.Flags().Lookup("detect.npm").Value)
		fmt.Println(d)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	fmt.Println(languages.Javascript{})

	//d.RegisterLang(languages.Javascript{})

	//scanner.RegisterLang()

	//f := scanner.GetFlags()

	//fmt.Println(f)
	//fmt.Println(b)

	//var s string = "Bool"

	// Languages

	// JavaScript
	scanCmd.Flags().Bool("detect.npm", false, "Search for npm dependencies")
	scanCmd.Flags().Bool("detect.bower", false, "Search for Bower dependencies")

	// Python
	scanCmd.Flags().Bool("detect.python", false, "Search for Python dependencies")
	scanCmd.Flags().String("detect.python.requirements.path", "./requirements.txt", "The path of the requirements.txt file")
}