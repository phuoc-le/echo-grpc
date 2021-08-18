package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
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

func RandomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(RandInt(65, 90)) != temp {
			temp = string(RandInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}