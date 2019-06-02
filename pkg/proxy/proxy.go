package proxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/jace-ys/viaduct/pkg/api"
	"github.com/jace-ys/viaduct/pkg/utils/format"
)

type Proxy struct {
	api          *api.Api
	reverseProxy *httputil.ReverseProxy
}

// Add header to enable CORS if allow_cross_origin is set to true
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p.api.AllowCrossOrigin {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	p.reverseProxy.ServeHTTP(w, r)
}

// Create new reverse proxy using API definition
func New(api *api.Api) *Proxy {
	target := api.UpstreamUrl
	reverseProxy := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.URL.Path = format.SingleJoiningSlash(target.Path, r.URL.Path)
			r.Host = target.Host
		},
	}

	return &Proxy{api: api, reverseProxy: &reverseProxy}
}

func StripPrefix(prefix string, proxy *Proxy) http.Handler {
	return http.StripPrefix(format.AddSlashes(prefix), proxy)
}
