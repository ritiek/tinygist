package main

import (
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "github.com/tidwall/gjson"
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


func main() {
    // name := "bla123321he"
    // contents := get_gist_contents(name)
    // fmt.Println(contents)

	raw_url := create_gist("bla123321he", "this is sparta!")
	fmt.Println(raw_url)
}
