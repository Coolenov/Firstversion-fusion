package api

import "NaxProject/lib"

type Album struct {
	sourse string    `json:"sourse_post"`
	post   *lib.Post `json:"obj_post"`
}
