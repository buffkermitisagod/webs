package webs

import (
	"fmt"
	"webs/mod/server"

	"github.com/buffkermitisagod/webs/mod/server"
)

// global data sets
type Path_static_template struct {
	Path string
	File string
}

var paths_static []Path_static_template

// debug/util stuff
var version string = "v0.0.1"

func Version() {
	fmt.Println(version)
}

func Print_paths() {
	fmt.Println("static paths -> ")

	for _, i := range paths_static {
		fmt.Print(" ", i, " ")
	}

	fmt.Println("")
}

// path setting data

func Path_static(path, file string) {
	// set a static path
	var tmp Path_static_template

	tmp.Path = path
	tmp.File = file

	// update the paths_static var
	paths_static = append(paths_static, tmp)
}

func Run(ip_port string) {
	server.Set_vars(paths_static)
	server.Start(":8080")

}
