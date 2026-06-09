package main

import "os"

type Server struct {
	Name string
	path string
	process os.Process
}

func (s *Server) Start() {

}