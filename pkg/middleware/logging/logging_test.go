package logging

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"

	"github.com/jace-ys/viaduct/pkg/config"
	"github.com/jace-ys/viaduct/pkg/utils/log"
)

func TestLogger(t *testing.T) {
	var buff bytes.Buffer

	serviceRegistry, err := config.RegisterServices("../../../config/config.sample.yml")
	if err != nil {
		t.Error(err)
	}

	log.WithLevels(log.Options{
		Prefix: "TestLogger",
		Out:    io.MultiWriter(&buff, os.Stdout),
	})

	logging := CreateMiddleware(log.Request(), &serviceRegistry)

	n := negroni.New()
	n.UseFunc(logging)

	n.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test Ok"))
	}))

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://testing.com/invalid/url", nil)
	if err != nil {
		t.Error(err)
	}

	n.ServeHTTP(res, req)

	serviceName := strings.Split(buff.String(), " | ")[1]

	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, res.Body.String(), "Test Ok")
	assert.Equal(t, true, strings.Contains(serviceName, "Testing"))
	assert.Equal(t, true, buff.Len() > 0)
}
