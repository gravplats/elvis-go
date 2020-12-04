package debug

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	DEBUG            = "DEBUG"
	DEBUG_SESSION_ID = "DEBUG_SESSION_ID"
)

func DumpInput() {
	if !isDebug() {
		return
	}

	buf := []byte(strings.Join(os.Args, " "))

	d, ok := GetDebugDir()
	if !ok {
		return
	}

	f := filepath.Join(d, "input")
	err := ioutil.WriteFile(f, buf, 0644)
	if err != nil {
		log.Println(err)
	}
}

func DumpJson(v interface{}, filename string) {
	if !isDebug() {
		return
	}

	buf, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return
	}

	d, ok := GetDebugDir()
	if !ok {
		return
	}

	f := filepath.Join(d, filename)
	err = ioutil.WriteFile(f, buf, 0644)
	if err != nil {
		log.Println(err)
	}
}

func GetDebugDir() (string, bool) {
	if !isDebug() {
		return "", false
	}

	sessionId := os.Getenv(DEBUG_SESSION_ID)
	if sessionId == "" {
		log.Printf("%s is empty.\n", DEBUG_SESSION_ID)
		return "", false
	}

	dir := filepath.Join("debug", sessionId)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return "", false
	}

	return dir, true
}

func isDebug() bool {
	debug, err := strconv.ParseBool(os.Getenv(DEBUG))
	if err != nil {
		log.Println(err)
		return false
	}

	return debug
}
