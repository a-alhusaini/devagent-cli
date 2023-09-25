package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
)

func main() {
	args := os.Args
	if len(args) >= 2 {
		proccessPrompt(strings.Join(args[1:], " "))

	} else {
		var prompt string
		fmt.Print("prompt: ")
		_, err := fmt.Scan(&prompt)
		if err != nil {
			panic(err)
		}
		proccessPrompt(prompt)

	}
}

func highlightCode(codeText string) string {
	var buffer bytes.Buffer

	err := quick.Highlight(&buffer, codeText, "py", "terminal", "vim")
	if err != nil {
		return codeText
	}

	result := buffer.String()
	return result
}

func proccessPrompt(prompt string) {
	request, err := http.Post("https://devagent.io/api/chat", "application/json", strings.NewReader(fmt.Sprintf(`{"prompt":"%v"}`, prompt)))
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	content := string(body)
	if request.StatusCode != 200 {
		fmt.Println(request.Status)
		fmt.Println(content)
		return
	}


	sections := strings.Split(content,"```")
	response := ""
	for i,section := range sections {
		if i%2 == 1{
			response = response + highlightCode(section)
		} else {
			response = response + section
		}
	}

	// set value to 1 if true
	debug := os.Getenv("DEVAGENT_DEBUG")
	if debug == "1" {
		fmt.Println(content)
	}

	fmt.Println(response)
}
