package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetIntEnv(key string) int {
	val := os.Getenv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("some error"))
	}
	return ret
}

func GetStrEnv(key string) string {
	val := os.Getenv(key)
	return val
}

func GetBoolEnv(key string) bool {
	val := os.Getenv(key)
	ret, err := strconv.ParseBool(val)
	if err != nil {
		panic(fmt.Sprintf("some error"))
	}
	return ret
}