package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)
//https://gist.github.com/delsner/64e79da93a77aa364e79013d3baeaa3e
type Client struct {
	*http.Server
}

const port  =  ":8080"

func( s * Client) addr()string{
	s.Addr = port
	return s.Addr
}
func(s *Client) Serve(c context.Context, handler http.Handler){
	fmt.Println("Create a sample server at port %s",s.addr())
	s.Handler = handler
	s.ListenAndServe()
}

//NewClient return a new instance of client
func NewClient() *Client{
	return &Client{
		Server: &http.Server{
			ReadTimeout:    600 * time.Second,
			WriteTimeout:   600 * time.Second,
		},

	}
}

