package main

import (
    "fmt"
    "net/http"
    "net/url"
    "bytes"
    "io/ioutil"
    "flag"
    "os"
    "github.com/tidwall/gjson"
    "github.com/atotto/clipboard"
    )


func get_gist_contents(name string) string {
    response, err := http.Get("https://git.io/" + name)
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


func create_gist(name string, code string) string {
    payload := []byte(`{
            "public": "false",
            "files": {
                "` + name + `": {
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

    value := gjson.Get(string(body), "files." + name + ".raw_url")
    return value.String()
}


func shorten_url(github_url string, code string) (string, int) {
    form := url.Values {
            "url"  : {github_url},
            "code" : {code},
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
    download := flag.Bool("d", true,
                        "Copy gist contents to clipboard")

    upload := flag.Bool("u", false,
                        "Create gist from clipboard content")

    name := flag.String("i", "REQUIRED",
                        "Unique Identifier")

    flag.Parse()

    if *name == "REQUIRED" {
        fmt.Println("Please set a unique identifier: `-i <unique_identifier>`")
        os.Exit(1)
    }

    if *upload {
        content, _ := clipboard.ReadAll()
        raw_url := create_gist(*name, content)
        _, return_code := shorten_url(raw_url, *name)
        if return_code >= 200 && return_code <= 299 {
            fmt.Println(raw_url)
        } else {
            fmt.Println("Identifier already used up!")
            fmt.Println("Try again with different unique identifer")
        }
    } else if *download {
        content := get_gist_contents(*name)
        clipboard.WriteAll(content)
        fmt.Println(content)
    }
}
