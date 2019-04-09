package format

import (
	"log"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleJoiningSlash(t *testing.T) {
	u1, err := url.Parse("https://jsonplaceholder.typicode.com/")
	if err != nil {
		log.Fatal(err)
	}

	p1 := SingleJoiningSlash(u1.Path, "/posts")

	assert.Equal(t, "/posts", p1)

	u2, err := url.Parse("https://reqres.in/api")
	if err != nil {
		log.Fatal(err)
	}

	p2 := SingleJoiningSlash(u2.Path, "/users/2")

	assert.Equal(t, "/api/users/2", p2)
}

func TestAddSlashes(t *testing.T) {
	prefix := AddSlashes("prefix")
	assert.Equal(t, "/prefix/", prefix)
}
