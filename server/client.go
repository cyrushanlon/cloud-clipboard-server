package server

//Client holds a clients authentication information and clipboard
type Client struct {
	Username  string
	Password  string
	Clipboard []byte
}
