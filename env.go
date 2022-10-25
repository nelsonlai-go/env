package env

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func String(key string) string {
	return os.Getenv(key)
}

func StringList(key string) []string {
	return strings.Split(os.Getenv(key), ",")
}

func Int(key string) int {
	v := os.Getenv(key)
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("env: cannot convert %s to int (key: %s)", v, key)
		return 0
	}
	return i
}

func IntList(key string) []int {
	vs := strings.Split(os.Getenv(key), ",")
	result := make([]int, len(vs))
	for index, v := range vs {
		i, err := strconv.Atoi(v)
		if err != nil {
			result[index] = 0
			log.Printf("env: cannot convert %s to int (key: %s)", v, key)
			continue
		}
		result[index] = i
	}
	return result
}

func Int64(key string) int64 {
	v := os.Getenv(key)
	i64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Printf("env: cannot convert %s to int64 (key: %s)", v, key)
		return 0
	}
	return i64
}

func Int64List(key string) []int64 {
	vs := strings.Split(os.Getenv(key), ",")
	result := make([]int64, len(vs))
	for index, v := range vs {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			result[index] = 0
			log.Printf("env: cannot convert %s to int64 (key: %s", v, key)
			continue
		}
		result[index] = i64
	}
	return result
}

func Float64(key string) float64 {
	v := os.Getenv(key)
	f64, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Printf("env: cannot convert %s to float64 (key: %s)", v, key)
		return 0
	}
	return f64
}

func Float64List(key string) []float64 {
	vs := strings.Split(os.Getenv(key), ",")
	result := make([]float64, len(vs))
	for index, v := range vs {
		f64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			result[index] = 0
			log.Printf("env: cannot convert %s to float64 (key: %s)", v, key)
			continue
		}
		result[index] = f64
	}
	return result
}

func Bool(key string) bool {
	v := os.Getenv(key)
	return v == "true" || v == "1"
}

func BoolList(key string) []bool {
	vs := strings.Split(os.Getenv(key), ",")
	result := make([]bool, len(vs))
	for index, v := range vs {
		result[index] = v == "true" || v == "1"
	}
	return result
}
