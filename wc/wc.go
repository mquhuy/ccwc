/*
Copyright © 2024 Huy Mai <me@huymai.fi>
*/
package wc

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// CountResult stores the result of counting a file input
type CountResult struct {
	ByteCount int
	CharCount int
	WordCount int
	LineCount int
}

func Count(reader io.Reader) (*CountResult, error) {
	result := &CountResult{}
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		result.LineCount++
		line := scanner.Text()
		result.WordCount += len(strings.Fields(line))
		lineBytes := scanner.Bytes()
		// Add 2 as the newline characters were stripped off in
		// bufio.ScanLines()
		result.ByteCount += len(lineBytes) + 2
		result.CharCount += len(bytes.Runes(lineBytes)) + 2
	}

	return result, nil
}
