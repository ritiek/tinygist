package main

import (
    "fmt"
    "net/http"
    "net/url"
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

func shorten_url(github_url string, code string) string {
    form := url.Values{
            "url"  : {github_url},
            "code" : {code},
    }
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post("https://git.io", "application/x-www-form-urlencoded", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

    return rsp.Header.Get("Location")
}


func main() {
    name := "hello_crazy_"
    raw_url := "https://gist.githubusercontent.com/anonymous/c7d4e1983288f318e3b732fce88bbd88/raw/df8c54286144c469de33ff7d75b5cefe8a9edf7d/hello_crazy_"
	//raw_url := create_gist(name, "this is sparta! :)")
	//fmt.Println(raw_url)

    //contents := get_gist_contents(name)
    //fmt.Println(contents)

    github_url := shorten_url(raw_url, name)
    fmt.Println(github_url)
}
