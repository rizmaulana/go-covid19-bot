package main

type HereMapsResponse struct {
	Response Response `json:"Response"`
}
type Response struct {
	View []View `json:"View"`
}

type View struct {
	Result []Result `json:"Result"`
}

type Result struct {
	Location Location `json:"Location"`
}

type Location struct {
	Address Address `json:"Address"`
}

type Address struct {
	City string `json:"City"`
}
