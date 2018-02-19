package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func getGistContents(identifier string) string {
	response, err := http.Get("https://git.io/" + identifier)
	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		return string(contents)
	}
}

func createGist(identifier string, code string) string {
	payload := []byte(`{
            "public": "false",
            "files": {
                "` + identifier + `": {
                    "content":"` + code + `"
                    }
                }
            }`)

	req, err := http.NewRequest("POST", "https://api.github.com/gists",
		bytes.NewBuffer(payload))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	value := gjson.Get(string(body), "files." + identifier + ".raw_url")
	return value.String()
}

func shortenURL(github_url string, code string) (string, int) {
	form := url.Values{
		"url":  {github_url},
		"code": {code},
	}
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post("https://git.io", "", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	return rsp.Header.Get("Location"), rsp.StatusCode
}

func main() {
	downloadFlag := flag.Bool("d", true,
		"Copy gist contents to clipboard")

	uploadFlag := flag.Bool("u", false,
		"Create gist from clipboard content")

	identifier := flag.String("i", "",
		"Unique Identifier")

	flag.Parse()

	if *identifier == "REQUIRED" {
		fmt.Println("Please pass a unique identifier: `-i <uniqueIdentifier>`")
		os.Exit(1)
	}

	if *uploadFlag {
		content, _ := clipboard.ReadAll()
		rawURL := createGist(*identifier, content)
		_, returnCode := shortenURL(rawURL, *identifier)
		if returnCode >= 200 && returnCode <= 299 {
			fmt.Println(rawURL)
		} else {
			fmt.Println("Identifier already used up!")
			fmt.Println("Try again with different unique identifer")
		}
	} else if *downloadFlag {
		content := getGistContents(*identifier)
		clipboard.WriteAll(content)
		fmt.Println(content)
	}
}
