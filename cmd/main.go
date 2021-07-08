package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/doc-contributor-generator/templates"
)

var s = flag.String("o", "contributor.md", "generate contributors")
var l = flag.String("l", "zh", "language")
var f = flag.String("f", "", "generate type")
var i = flag.String("i", "", "generate index")

func main() {
	flag.Parse()

	if len(*i) > 0 {
		fmt.Println(*i)
		generateIndex(*i)
	} else {
		if *f == "comment" {
			appendComment()
		} else {
			generateContributorPage(*l, *s)
		}
	}
	fmt.Println("Done.")
}

func generateIndex(path string) {
	err := ioutil.WriteFile(path, []byte(templates.Index), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func appendComment() {
	list, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range list {
		name := item.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}

		data, err := ioutil.ReadFile(name)
		if err != nil {
			fmt.Printf("%+v\n", err)
			continue
		}
		buffer := bytes.NewBuffer(data)
		buffer.WriteString(`
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/gitalk@1/dist/gitalk.css">
<script src="./javascript/gitalk.js"></script>
<div id="gitalk-container"></div>
<script>
var gitalk = new Gitalk({
  "clientID": "b9f61eecd453fb93cefc",
  "clientSecret": "8e70a3944fa689c13136aa56ea5bfb9945802bd2",
  "repo": "go-zero-doc",
  "owner": "zeromicro",
  "admin": ["anqiansong"],
  "id": window.location.pathname,      
  "distractionFreeMode": false  
});
gitalk.render("gitalk-container");
</script>
`)
		err = ioutil.WriteFile(name, buffer.Bytes(), fs.ModePerm)
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}
}

func generateContributorPage(language, output string) {
	var text = templates.Contributors
	if language == "en" {
		text = templates.Contributors_EN
	}
	t, err := template.New("contributors").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	goZeroContribotuors := getContributors("tal-tech", "go-zero")
	goZeroDocContributors := getContributors("zeromicro", "go-zero-doc")
	buffer := bytes.NewBuffer(nil)
	err = t.Execute(buffer, map[string]interface{}{
		"goZeroList":    goZeroContribotuors,
		"goZeroDocList": goZeroDocContributors,
	})
	if err != nil {
		log.Fatal(err)
	}

	filename, err := filepath.Abs(output)
	if err != nil {
		return
	}

	dir := filepath.Dir(filename)
	os.Mkdir(dir, 0777)

	err = ioutil.WriteFile(filename, buffer.Bytes(), fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

}

// Contributor describes the contributor of github
type Contributor struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Contributions     int64  `json:"contributions"`
}

func getContributors(owener, repo string) []*Contributor {
	var (
		list    []*Contributor
		perPage = 100
		page    = 1
	)
	for {
		resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/contributors?page=%d&per_page=%d", owener, repo, page, perPage))
		if err != nil {
			fmt.Printf("%+v\n", err)
			return nil
		}

		if resp.StatusCode != http.StatusOK {
			return nil
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%+v\n", err)
			return nil
		}

		if len(bytes) == 0 {
			break
		}

		var contributors []*Contributor
		err = json.Unmarshal(bytes, &contributors)
		if err != nil {
			fmt.Printf("%+v\n", err)
			return nil
		}

		if len(contributors) == 0 {
			break
		}

		list = append(list, contributors...)
		page++
	}
	return list
}
