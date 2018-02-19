package main

import (
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "os"
    )

func get_gist_contents(name string) string {
    response, err := http.Get("https://git.io/" + name)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
        return "b"
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            panic(err)
        }
        return string(contents)
    }
}


/*  def create_gist(name, code)
    uri = URI('https://api.github.com/gists')
    payload = {
      'public' => false,
      'files' => {
        name => {
          'content' => code
        }
      }
    }

    req = Net::HTTP::Post.new(uri.path)
    req.body = payload.to_json
    res = Net::HTTP.start(uri.hostname, uri.port, :use_ssl => true) do |http|
      http.request(req)
    end

    raw_url = JSON.parse(res.body)['files'][name]['raw_url']
    return raw_url
end*/

func create_gist(name string, code string) string {
    json := []byte(`{
            "public": "false",
            "files": {
                "` + name + `": {
                    "content":"` + code + `"
                    }
                }
            }`)

    req, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewBuffer(json))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body)
}

func main() {
    // name := "bla123321he"
    // contents := get_gist_contents(name)
    // fmt.Println(contents)

	// contents := create_gist("bla123321he", "this is sparta")
	// fmt.Println(contents)
}
