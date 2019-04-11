package format

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatMethods(t *testing.T) {
	testString := []string{"GET", "POST", "DELETE"}
	methods := Methods(testString)
	assert.Equal(t, "[GET, POST, DELETE]", methods)
}

func TestFormatMiddleware(t *testing.T) {
	testMap := map[string]bool{"logging": true, "recovery": false}
	middleware := Middleware(testMap)
	assert.Equal(t, "[Logging]", middleware)
}

func TestSingleJoiningSlash(t *testing.T) {
	u1, err := url.Parse("https://jsonplaceholder.typicode.com/")
	if err != nil {
		t.Error(err)
	}

	p1 := SingleJoiningSlash(u1.Path, "/posts")

	assert.Equal(t, "/posts", p1)

	u2, err := url.Parse("https://reqres.in/api")
	if err != nil {
		t.Error(err)
	}

	p2 := SingleJoiningSlash(u2.Path, "/users/2")

	assert.Equal(t, "/api/users/2", p2)
}

func TestAddSlashes(t *testing.T) {
	prefix := AddSlashes("prefix")
	assert.Equal(t, "/prefix/", prefix)
}

func TestCamelToSnakeUnderscore(t *testing.T) {
	snakeUnderscore := CamelToSnakeUnderscore("ConfigPath")
	assert.Equal(t, "CONFIG_PATH", snakeUnderscore)
}
