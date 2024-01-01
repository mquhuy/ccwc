/*
Copyright © 2024 Huy Mai <me@huymai.fi>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/mquhuy/ccwc/wc"
)

var (
	countBytes *bool
	countChars *bool
	countLines *bool
	countWords *bool
)

var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "A clone of UNIX wc, written in Go",
	Long: `Print newline, word, and byte counts for each FILE, and a total line if
more than one FILE is specified.  A word is a non-zero-length sequence of
characters delimited by white space.`,
	Run: func(cmd *cobra.Command, args []string) {
		info, _ := os.Stdin.Stat()
		var reader io.Reader
		var err error
		var fileName string

		// If received a file, read it, else handle stdin
		if (info.Mode() & os.ModeCharDevice == os.ModeCharDevice) {
			fileName = args[0]
			reader, err = os.Open(fileName)
			if err != nil {
				os.Exit(1)
			}
		} else {
			reader = os.Stdin
		}
		counts := []int{}
		res, err := wc.Count(reader)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		if *countLines {
			counts = append(counts, res.LineCount)
		}
		if *countWords {
			counts = append(counts, res.WordCount)
		}
		if *countChars {
			counts = append(counts, res.CharCount)
		}
		if *countBytes {
			counts = append(counts, res.ByteCount)
		}
		if len(counts) == 0 {
			counts = []int{res.LineCount, res.WordCount, res.ByteCount}
		}
		resStr := ""
		for _, s := range counts {
			resStr = fmt.Sprintf("%s\t%d", resStr, s)
		}
		if fileName != "" {
			resStr = fmt.Sprintf("%s\t%s", resStr, fileName)
		}
		fmt.Println(resStr)
		os.Exit(0)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	countBytes = rootCmd.Flags().BoolP("bytes", "c", false, "Print the byte counts")
	countChars = rootCmd.Flags().BoolP("chars", "m", false, "Print the character counts")
	countLines = rootCmd.Flags().BoolP("lines", "l", false, "Print the newline counts")
	countWords = rootCmd.Flags().BoolP("words", "w", false, "Print the word counts")
}
