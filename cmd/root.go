/*
Copyright © 2024 Huy Mai <me@huymai.fi>
*/
package cmd

import (
	"bufio"
	"fmt"
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
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		splitFuncs := []bufio.SplitFunc{}
		if *countLines {
			splitFuncs = append(splitFuncs, bufio.ScanLines)
		}
		if *countWords {
			splitFuncs = append(splitFuncs, bufio.ScanWords)
		}
		if *countChars {
			splitFuncs = append(splitFuncs, bufio.ScanRunes)
		}
		if *countBytes {
			splitFuncs = append(splitFuncs, bufio.ScanBytes)
		}
		if len(splitFuncs) == 0 {
			splitFuncs = []bufio.SplitFunc{bufio.ScanLines, bufio.ScanWords, bufio.ScanBytes}
		}
		res, err := wc.CountAll(filePath, splitFuncs)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		for _, s := range *res {
			fmt.Printf("%*s\t", 1, s)
		}
		fmt.Print("\n")
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
