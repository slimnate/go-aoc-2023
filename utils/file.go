package utils

import (
	"bufio"
	"os"
	"strconv"
)

func Lines(path string) (lines []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func StrToIntSlice(in []string) (out []int, err error) {
	for _, s := range in {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		out = append(out, n)
	}

	return out, nil
}
