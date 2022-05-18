package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"golang.org/x/term"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func GetHash(s string) []byte {
	res := md5.Sum([]byte(s))
	return res[:]
}

func FailOnError(err error) {
	if err != nil {
		panic(err)
	}
}

type Closable interface {
	Close() error
}

func ReadValue(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.Trim(text, "\n ")
}

func ReadIntValue(prompt string) int64 {
	for {
		value := ReadValue(prompt)
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return intVal
		} else {
			log.Println("Can not parse int: " + err.Error())
		}
	}
}

func ReadHiddenValue(prompt string) string {
	fmt.Print(prompt)
	password, err := term.ReadPassword(syscall.Stdin)
	FailOnError(err)
	return string(password)
}
