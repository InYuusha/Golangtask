package dto

type Value struct {
	Val     string
	Expiry  int64
	Created int64
	Push    chan string
	Pop     chan string
}
