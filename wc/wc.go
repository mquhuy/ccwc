/*
Copyright © 2024 Huy Mai <me@huymai.fi>
*/
package wc

import (
	"bufio"
	"os"
	"strconv"
)

func count(filePath string, splitFunc bufio.SplitFunc) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	var result int
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)

	for scanner.Scan() {
		result++
		if err := scanner.Err(); err != nil {
			return 0, err
		}
	}

	return result, nil
}

func CountAll(filePath string, splitFuncs []bufio.SplitFunc) (*[]string, error) {
	res := []string{}
	for _, f := range splitFuncs {
		c, err := count(filePath, f)
		if err != nil {
			return nil, err
		}
		res = append(res, strconv.Itoa(c))
	}
	res = append(res, filePath)
	return &res, nil
}
