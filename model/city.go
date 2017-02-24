package model

type City struct {
	Id	  int    `json:"id"`
	Zip	  int    `json:"zip"`
	State string `json:"name"`
	City  string `json:"timezone"`
	Lat	  int `json:"lat"`
	Lng	  int `json:"long"`
}
