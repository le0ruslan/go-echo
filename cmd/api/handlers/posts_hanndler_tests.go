package handlers

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPostHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndToEndSuite) TestPostNoREsult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/post/5654")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status":"ok", "data":[]}`, string(b))
}
