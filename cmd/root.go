/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
Copyright © 2023 Jun Nishimura <n.junjun0303@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/JunNishimura/konmari/internal/file"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "konmari",
	Short: "clean up source code comments",
	Long:  "clean up source code comments",
	RunE: func(cmd *cobra.Command, args []string) error {
		// check file existence
		for _, arg := range args {
			if _, err := os.Stat(arg); os.IsNotExist(err) {
				return fmt.Errorf("cannot find '%s'", arg)
			}
		}

		// change directory to files
		filePaths := make([]string, 0)
		for _, arg := range args {
			cleanedArg := filepath.Clean(arg)
			cleanedArg = strings.ReplaceAll(cleanedArg, `\`, "/")
			info, err := os.Stat(cleanedArg)
			if err != nil {
				return fmt.Errorf("fail to get '%s' info: %w", cleanedArg, err)
			}
			if info.IsDir() {
				extractedFiles, err := file.ExtractFileFromDirectory(cleanedArg)
				if err != nil {
					return fmt.Errorf("fail to extract files under '%s': %w", cleanedArg, err)
				}
				filePaths = append(filePaths, extractedFiles...)
			} else {
				filePaths = append(filePaths, cleanedArg)
			}
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
