package server

import (
	"log"
	"net/http"
	"webs/mod/util"
)

// handles the true back-end of the server

// data types needed
type Path_static_template struct {
	Path string
	File string
}

// data holders
var paths_static []Path_static_template

func Set_vars(stats []Path_static_template) {
	// set all public data vers needed
	paths_static = stats

}

func Start(ip_run string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		// check if url is in the static pages
		for _, static := range paths_static {
			if static.Path == path {
				if util.Verify_path(static.File) {
					http.ServeFile(w, r, static.File)
				} else {
					log.Fatal("[404] path -> ", static.Path, " not found! and no 404 file found!")
				}
			}
		}
	})

	http.ListenAndServe(ip_run, nil)
}
