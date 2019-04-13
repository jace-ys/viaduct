package proxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/jace-ys/viaduct/pkg/domain"
	"github.com/jace-ys/viaduct/pkg/utils/format"
)

type Proxy struct {
	service      *domain.Service
	reverseProxy *httputil.ReverseProxy
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p.service.AllowCrossOrigin {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	p.reverseProxy.ServeHTTP(w, r)
}

func New(service *domain.Service) *Proxy {
	target := service.UpstreamUrl
	reverseProxy := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.URL.Path = format.SingleJoiningSlash(target.Path, r.URL.Path)
			r.Host = target.Host
		},
	}

	return &Proxy{service: service, reverseProxy: &reverseProxy}
}

func StripPrefix(prefix string, proxy *Proxy) http.Handler {
	return http.StripPrefix(format.AddSlashes(prefix), proxy)
}
