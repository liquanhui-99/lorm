package lorm

import "testing"

func TestHTTPServer(t *testing.T) {
	var s Server = &HTTPServer{}
	err := s.Start(":8080")
	if err != nil {
		return
	}
}
