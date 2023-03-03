package services

import (
	"bufio"
	"os"
	"strconv"
)

func ReadPlace() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	if "h" == s {
		return 0, nil
	}
	n, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}
