package main

type Course struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Discussion struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
