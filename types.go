package main

// dhRequest api as described here: http://bit.ly/2MHs8UU
type dhRequest struct {
	url    string
	cmd    string
	format string
	key    string
}
