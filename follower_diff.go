package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// url_gh_test := "https://api.github.com/octocat"
	url_gh_followers := "https://api.github.com/user/followers"
	data, err := os.ReadFile(os.ExpandEnv("$HOME/Library/CloudStorage/Dropbox/env/keys/gh-PAT"))
	if err != nil {
		panic(err)
	}
	token := strings.TrimSpace(string(data))

	// GITHUB API REQ
	req, _ := http.NewRequest("GET", url_gh_followers, nil)
	// "GET" = HTTP method
	// "https://api.github.com/user/followers" = the URL you're hitting
	// nil = no request body (GET requests don't have bodies)

	// SET HEADER AUTHORIZATION
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token)) // "Hey, GH, here's my token"
	// ^ MUST BE: "token <token>"

	// SET HEADER ACCEPTANCE
	// Most REST APIs use: `req.Header.Set("Accept", "application/json")`
	// GitHub API uses custom media types to version their APIs:
	// req.Header.Set("Accept", "application/vnd.github.v3+json") <- github docs specify this
	// vnd = vendor-specific
	//   GitHub API Version History
	// 	- https://docs.github.com/en/rest
	//  - v1 & v2: Deprecated/removed years ago
	//  - v3: Current REST API (most used)
	//  - v4: GraphQL API (different syntax, very flexible/complex)
	//
	// MIME ("Content-Type") structuring = type/subtype
	// types: [application[json, xml, pdf, zip, octet-stream...], text[plain, html, css, csv],
	// more types:image[png, jpg, gif, svg], audio[mpeg, wav], video/[mpeg, mp4, webm], multipart/form-data]

	// Most common
	req.Header.Set("Accept", "application/json")

	// Create an HTTP client which is the actual requester
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // Tells to close the response body when done, but not immediately

	// Read the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
