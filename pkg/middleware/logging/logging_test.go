package logging

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"

	"github.com/jace-ys/viaduct/pkg/config"
	"github.com/jace-ys/viaduct/pkg/utils/log"
)

func TestLoggingMiddleware(t *testing.T) {
	var buff bytes.Buffer

	apiRegistry, err := config.RegisterApis("../../../config/config.sample.yaml")
	if err != nil {
		t.Error(err)
	}

	log.WithLevels(log.Options{
		Prefix: "TestLogger",
		Out:    &buff,
	})

	logging := CreateMiddleware(log.Request(), &apiRegistry)

	n := negroni.New()
	n.UseFunc(logging)

	n.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test Ok"))
	}))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://google.com/", nil)

	n.ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, res.Body.String(), "Test Ok")
	assert.Equal(t, true, buff.Len() > 0)
}

func TestApiContext(t *testing.T) {
	var buff bytes.Buffer

	apiRegistry, err := config.RegisterApis("../../../config/config.sample.yaml")
	if err != nil {
		t.Error(err)
	}

	log.WithLevels(log.Options{
		Prefix: "TestLogger",
		Out:    &buff,
	})

	logging := CreateMiddleware(log.Request(), &apiRegistry)

	n := negroni.New()
	n.UseFunc(logging)

	n.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test Ok"))
	}))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://google.com/", nil)

	n.ServeHTTP(res, req)

	apiContext := getApiContext(req, &apiRegistry)

	assert.Equal(t, true, strings.Contains(buff.String(), apiContext.Name))
	assert.Equal(t, true, strings.Contains(buff.String(), apiContext.Host))
	assert.Equal(t, "GET", apiContext.Method)
}
