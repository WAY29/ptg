package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	ft "github.com/valyala/fasttemplate"

	"ptg/encoders"
)

// TEMPLATEPATH
var TEMPLATEPATH string
var TEMPLATEBINARYPATH string

func init() {
	pwd, _ := os.Getwd()
	exd, _ := os.Executable()
	TEMPLATEPATH = filepath.Join(pwd, "template")
	TEMPLATEBINARYPATH = filepath.Join(filepath.Dir(exd), "template")
}

// Total Template
var (
	TotalTemplate map[string]Template
)

// Template struct
type Template struct {
	Name    string `json:"name"`
	Payload map[string]struct {
		Number   int               `json:"number"`
		Template string            `json:"template"`
		Encoding map[string]string `json:"encoding"`
	} `json:"payload"`
}

// Load template from TEMPLAEPATH
func LoadTemplate() {
	var fp string
	TotalTemplate = make(map[string]Template)
	rd, err := ioutil.ReadDir(TEMPLATEPATH)
	if err != nil {
		rd, err = ioutil.ReadDir(TEMPLATEBINARYPATH)
		if err != nil {
			Error(fmt.Sprintf("Cannot read from %s and %s", TEMPLATEPATH, TEMPLATEBINARYPATH))
		} else {
			fp = TEMPLATEBINARYPATH
		}
	} else {
		fp = TEMPLATEPATH
	}
	for _, fi := range rd {
		fname := fi.Name()
		if !fi.IsDir() && strings.HasSuffix(fname, ".json") {
			bytes, err := ioutil.ReadFile(filepath.Join(fp, fname))
			if err != nil {
				Error("Cannot read file: " + fname)
			}
			template := new(Template)
			if err := json.Unmarshal([]byte(bytes), &template); err == nil {
				TotalTemplate[template.Name] = *template
			}
		}
	}
}

// Parse template and generate payload
func ParseAndGenerate(args []string) {
	totalName := args[0]
	if !strings.Contains(totalName, ".") {
		totalName += ".default"
	}
	parseArray := strings.Split(totalName, ".")
	templateName := parseArray[0]
	payloadName := parseArray[1]
	template, ok := TotalTemplate[templateName]
	if !ok {
		Error("No this template")
	}

	args = args[1:]

	payloadStruct, ok := template.Payload[payloadName]
	if !ok {
		fmt.Println("\n[+] Available payloads in " + templateName + ":")
		for k := range template.Payload {
			fmt.Println("    - " + k)
		}
		fmt.Println()
		Error("No this payload in template")
	}
	payloadTemplate := payloadStruct.Template
	payloadPlaceholderNumber := payloadStruct.Number
	if payloadPlaceholderNumber != len(args) {
		fmt.Printf("\n[+] Payload template %s.%s:\n  %s\n\n", templateName, payloadName, payloadTemplate)
		Error("Arguments number not same as payload template placeholder number")
	}
	argsMap := make(map[string]interface{}, len(args))
	encodingArgs := payloadStruct.Encoding
	if encodingArgs != nil { // Encoding arguments
		for i := 0; i < len(encodingArgs); i++ {
			if i >= len(args) {
				break
			}
			e, ok := encodingArgs[strconv.Itoa(i)]
			if !ok {
				break
			}
			eArray := strings.Split(e, ",")
			for _, e := range eArray {
				args[i] = encoders.GetEncoder(e).Encode(args[i])
			}
		}
	}
	// ? use fast template
	for i := 0; i < len(args); i++ {
		argsMap[strconv.Itoa(i)] = args[i]
	}
	payload := ft.ExecuteString(payloadTemplate, "{{", "}}", argsMap)
	fmt.Println("[+] Generate payload:")
	fmt.Println(payload)
	CopyPayloadToClipboard(payload)
}

// Copy payload to clipboard
func CopyPayloadToClipboard(payload string) error {
	return clipboard.WriteAll(payload)
}

// Print Error and exit
func Error(str string) {
	fmt.Println("[-] Error: " + str)
	os.Exit(1)
}

// Print Usage
func Usage() {
	fmt.Println(`Usage:
	ptg template.payload arguments...
	
Examples:
	ptg # show all available templates
	ptg pty # show all payload templates in pty
	ptg pty.bash
	ptg demo.demo test

Available templates:`)
	for k := range TotalTemplate {
		fmt.Println("	- " + k)
	}
	fmt.Println()
}

func main() {
	LoadTemplate()
	if len(os.Args) < 2 {
		Usage()
		return
	} else {
		ParseAndGenerate(os.Args[1:])
	}
}
