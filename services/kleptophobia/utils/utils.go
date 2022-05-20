package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"golang.org/x/term"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
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

func ReadValueWithValidation(prompt string, regexp *regexp.Regexp) string {
	for {
		fmt.Print(prompt)
		reader := bufio.NewReader(os.Stdin)

		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		value := strings.Trim(text, "\n ")
		if !regexp.MatchString(value) {
			fmt.Printf("Value must match regexp: %v\n", regexp)
			continue
		}
		return value
	}
}

func ReadUIntValue(prompt string) uint32 {
	for {
		value := ReadValue(prompt)
		intVal, err := strconv.ParseInt(value, 10, 32)
		if err == nil {
			if intVal < 0 {
				log.Println("Room number must be positive")
			} else {
				return uint32(intVal)
			}
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

func InitConfig[T proto.Message](filename string, config T) {
	rawConfig, err := ioutil.ReadFile(filename)
	FailOnError(err)

	err = protojson.Unmarshal(rawConfig, config)
	FailOnError(err)
}

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandString(ln int) string {
	res := make([]rune, ln)
	for i := range res {
		res[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(res)
}

type SimpleSet map[byte]bool

func makeSimpleSet(data []byte) *SimpleSet {
	res := SimpleSet{}
	for _, v := range data {
		res[v] = true
	}
	return &res
}

func equals(first, second *SimpleSet) bool {
	if len(*first) != len(*second) {
		return false
	}
	for k, _ := range *first {
		_, ok := (*second)[k]
		if !ok {
			return false
		}
	}
	return true
}

func EqualAsSets(first, second []byte) bool {
	return equals(makeSimpleSet(first), makeSimpleSet(second))
}
