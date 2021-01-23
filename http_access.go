package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://api.github.com/users/defunkt"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	respBodyByteArray, _ := ioutil.ReadAll(resp.Body)
	respBodyStr := string(respBodyByteArray)

	outputText := fmt.Sprintf("%d\t%s\n", 42, respBodyStr)

	filename := "./output.txt"
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(outputText); err != nil {
		return
	}
	writer.Flush()
}
