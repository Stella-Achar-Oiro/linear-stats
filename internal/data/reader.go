package data

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

// ReadData reads data from a file and returns a slice of float64.
func ReadData(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		values := strings.Split(line, ",")
		for _, v := range values {
			value, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
			if err != nil {
				return nil, errors.New("invalid number format in data file")
			}
			if value >= math.MaxInt64 || value <= -math.MaxInt64 {
				return nil, errors.New("number out of range in data file")
			}
			data = append(data, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
