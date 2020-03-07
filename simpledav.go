// Program simpledav is a simple webdav server suitable for offensive use
package main

/*
 * simpledav.go
 * Small offensive webdav server
 * By J. Stuart McMurray
 * Created 20200306
 * Last Modified 20200306
 */

import (
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

var (
	network    = "unix"      /* Network for net.Listen */
	address    = "/tmp/.wds" /* Address for net.Listen */
	root       = "/"         /* Root of the served filesystem */
	allowWrite = ""          /* Non-empty to allow writes */
)

// Handler handles WebDAV requests
type Handler struct {
	w *webdav.Handler
}

// ServeHTTP implements http.Handler
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/* If we're allowing writes, any method works */
	if "" != allowWrite {
		h.w.ServeHTTP(w, r)
		return
	}

	/* Check if we have an allowed, read-only method */
	var ok bool
	for _, m := range []string{
		"OPTIONS",
		"GET",
		"HEAD",
		"PROPFIND",
	} {
		if r.Method == m {
			ok = true
			break
		}
	}

	/* If we don't have an allowed method, pretend everything's ok */
	if !ok {
		return
	}

	/* Do the read */
	h.w.ServeHTTP(w, r)
}

func main() {
	/* Register the webdav handler */
	http.Handle("/", Handler{w: &webdav.Handler{
		FileSystem: webdav.Dir(root),
		LockSystem: webdav.NewMemLS(),
	}})

	/* Listen and serve */
	if "unix" == network {
		if err := os.RemoveAll(address); nil != err {
			log.Fatalf("RemoveAll: %v", err)
		}
	}
	l, err := net.Listen(network, address)
	if nil != err {
		log.Fatalf("Listen: %v", err)
	}
	if "unix" == network {
		if err := os.Chmod(l.Addr().String(), 0777); nil != err {
			log.Printf("Chmod: %v", err)
		}
	}
	if err := http.Serve(l, nil); nil != err {
		log.Fatalf("Error: %v", err)
	}
}
