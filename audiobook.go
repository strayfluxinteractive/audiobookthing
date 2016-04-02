package main

//Audiobook represents an audiobook
type Audiobook struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

//Audiobooks holds the audiobooks
type Audiobooks []Audiobook
