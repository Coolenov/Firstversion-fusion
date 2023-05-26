package api

import "NaxProject/lib"

type Contents struct {
	data []Content `json:"data"`
}

type Content struct {
	sourceName string     `json:"sourceName"`
	postData   []lib.Post `json:"postdata"`
}
