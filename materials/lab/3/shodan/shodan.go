// Do not use this file directly, do not attemp to compile this source file directly
// Go To lab/3/shodan/main/main.go

// DO NOT MODIFY THIS FILE WITH YOUR KEY!
package shodan

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}