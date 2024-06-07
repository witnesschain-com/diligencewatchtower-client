/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func StartBlock(message string) {

	pc, file, line, ok := runtime.Caller(1)

	_ = file
	_ = line
	_ = ok

	from := runtime.FuncForPC(pc).Name()

	if strings.HasPrefix(from, "w") {
		from = strings.Replace(from, "github.com/witnesschain-com/diligencewatchtower-client/", "", 1)
	}

	from_length := len(from)
	if from_length > 20 {
		from = from[0:17] + " .."
	}

	now := time.Now().Format(time.ANSIC)
	r := regexp.MustCompile("^([^ ]*) (.*)$")
	now = r.ReplaceAllString(now, "${2}") // get rid of day

	fmt.Printf("%20s │ %s ├──────────────────────────── %s ────────\n", from, now, message)
}

func EndBlock(message string) {

	pc, file, line, ok := runtime.Caller(1)

	_ = file
	_ = line
	_ = ok

	from := runtime.FuncForPC(pc).Name()

	if strings.HasPrefix(from, "w") {
		from = strings.Replace(from, "github.com/witnesschain-com/diligencewatchtower-client/", "", 1)
	}

	from_length := len(from)
	if from_length > 20 {
		from = from[0:17] + " .."
	}

	now := time.Now().Format(time.ANSIC)
	r := regexp.MustCompile("^([^ ]*) (.*)$")
	now = r.ReplaceAllString(now, "${2}") // get rid of day

	fmt.Printf("%20s │ %s ├────────────────────────────────────────────────────────\n", from, now)
}

func Success(message any) bool {

	pc, file, line, ok := runtime.Caller(1)

	_ = file
	_ = line
	_ = ok

	from := runtime.FuncForPC(pc).Name()

	if strings.HasPrefix(from, "w") {
		from = strings.Replace(from, "github.com/witnesschain-com/diligencewatchtower-client/", "", 1)
	}

	from_length := len(from)
	if from_length > 20 {
		from = from[0:17] + " .."
	}

	now := time.Now().Format(time.ANSIC)
	r := regexp.MustCompile("^([^ ]*) (.*)$")
	now = r.ReplaceAllString(now, "${2}") // get rid of day

	fmt.Printf("%20s │ %s │ \x1B[32m✓ %v\x1B[0m\n", from, now, message)

	return true
}

func Info(message any) bool {

	pc, file, line, ok := runtime.Caller(1)

	_ = file
	_ = line
	_ = ok

	from := runtime.FuncForPC(pc).Name()

	if strings.HasPrefix(from, "w") {
		from = strings.Replace(from, "github.com/witnesschain-com/diligencewatchtower-client/", "", 1)
	}

	from_length := len(from)
	if from_length > 20 {
		from = from[0:17] + " .."
	}

	now := time.Now().Format(time.ANSIC)
	r := regexp.MustCompile("^([^ ]*) (.*)$")
	now = r.ReplaceAllString(now, "${2}") // get rid of day

	fmt.Printf("%20s │ %s │ \x1B[34m➤ %v\x1B[0m\n", from, now, message)

	return true
}

func Warning(message any) bool {

	pc, file, line, ok := runtime.Caller(1)

	_ = file
	_ = line
	_ = ok

	from := runtime.FuncForPC(pc).Name()

	if strings.HasPrefix(from, "w") {
		from = strings.Replace(from, "github.com/witnesschain-com/diligencewatchtower-client/", "", 1)
	}

	from_length := len(from)
	if from_length > 20 {
		from = from[0:17] + " .."
	}

	now := time.Now().Format(time.ANSIC)
	r := regexp.MustCompile("^([^ ]*) (.*)$")
	now = r.ReplaceAllString(now, "${2}") // get rid of day

	fmt.Printf("%20s │ %s │ \x1B[33m! %v\x1B[0m\n", from, now, message)

	return false
}

func Error(message any) bool {

	pc, file, line, ok := runtime.Caller(1)

	_ = file
	_ = line
	_ = ok

	from := runtime.FuncForPC(pc).Name()
	r := regexp.MustCompile("^(.*)?/src/(.*)$")
	file = r.ReplaceAllString(file, "${2}") // get rid of first ".*/src/"

	if strings.HasPrefix(from, "w") {
		from = strings.Replace(from, "github.com/witnesschain-com/diligencewatchtower-client/", "", 1)
	}

	from_length := len(from)
	if from_length > 20 {
		from = from[0:17] + " .."
	}

	now := time.Now().Format(time.ANSIC)
	r = regexp.MustCompile("^([^ ]*) (.*)$")
	now = r.ReplaceAllString(now, "${2}") // get rid of day

	fmt.Printf("%20s │ %s │ \x1B[31m✕ [@%s:%d] - %v\x1B[0m\n", from, now, file, line, message)

	return false
}

func Fatal(message any) bool {

	pc, file, line, ok := runtime.Caller(1)

	_ = file
	_ = line
	_ = ok

	split := strings.Split(file, "/src/")

	if len(split) > 1 {
		file = split[1]
	}

	for i := 2; i < len(split); i++ {
		file += "/src/" + split[i]
	}

	from := runtime.FuncForPC(pc).Name()

	if strings.HasPrefix(from, "w") {
		from = strings.Replace(from, "github.com/witnesschain-com/diligencewatchtower-client/", "", 1)
	}

	from_length := len(from)
	if from_length > 20 {
		from = from[0:17] + " .."
	}

	now := time.Now().Format(time.ANSIC)
	r := regexp.MustCompile("^([^ ]*) (.*)$")
	now = r.ReplaceAllString(now, "${2}") // get rid of day

	fatalErrorMessageString := fmt.Sprintf("%v", message)
	fmt.Printf("%20s │ %s │ \x1B[31m⚠ [@%s:%d] - %v\x1B[0m\n", from, now, file, line, message)

	// send the error if alert url is set
	currentConfig := LoadConfigFromJson()

	if currentConfig.AlertURL != "" {

		message := fmt.Sprintf("watchtower_id: %v\nfrom: %v\ntimestamp: %v\nfile: %v\nline: %vmessage: %v\n", currentConfig.WatchtowerAddress, from, now, file, line, fatalErrorMessageString)
		
		request, _ := json.Marshal(map[string] interface{}{
			"text": message,
		})

		requestBody := bytes.NewBuffer(request)
		resp, err := http.Post(currentConfig.AlertURL, "application/json", requestBody)
		if err != nil {
			Error(err)
		}
		if resp.StatusCode != 200 {
			fmt.Print(resp)
		}
	}

	os.Exit(123)

	return false
}
