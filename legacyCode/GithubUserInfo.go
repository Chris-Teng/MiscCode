package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userInfos struct {
	Login               string `"login"`
	Id                  string `"id"`
	Node_id             string `"node_id"`
	Avatar_url          string `"avatar_url"`
	Gravatar_id         string `"gravatar_id"`
	Url                 string `"url"`
	Html_url            string `"html_url"`
	Followers_url       string `"followers_url"`
	Following_url       string `"following_url"`
	Gists_url           string `"gists_url"`
	Starred_url         string `"starred_url"`
	Subscriptions_url   string `"subscriptions_url"`
	Organizations_url   string `"organizations_url"`
	Repos_url           string `"repos_url"`
	Events_url          string `"events_url"`
	Received_events_url string `"received_events_url"`
	Type                string `"type"`
	Site_admin          string `"site_admin"`
	Name                string `"name"`
	Company             string `"company"`
	Blog                string `"blog"`
	Location            string `"location"`
	Email               string `"email"`
	Hireable            string `"hireable"`
	Bio                 string `"bio"`
	Twitter_username    string `"twitter_username"`
	Public_repos        string `"public_repos"`
	Public_gists        string `"public_gists"`
	Followers           string `"followers"`
	Following           string `"following"`
	Created_at          string `"created_at"`
	Updated_at          string `"updated_at"`
}

func GithubUserInfo(username string) userInfos {
	url := "https://api.github.com/users/Chris" + username
	request, _ := http.NewRequest("GET", url, nil)
	// request.Header.Add("User-Agent","")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return userInfos{}
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(result))
	var account userInfos
	if err := json.Unmarshal(result, &account); err != nil {
		fmt.Println(err)
		return userInfos{}
	}
	return account
}

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		if strings.Contains(arg, "help") {
			fmt.Println("You can use a=xxx to get the account info")
		} else if strings.Split(arg, "=")[0] == "a" {
			fmt.Println(GithubUserInfo(strings.Split(arg, "=")[1]))
		}
	} else {
		fmt.Println("You have to give me an argument,try 'help'")
	}
}
