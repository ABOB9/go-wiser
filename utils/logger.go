package utils

import (
	gl "log"
	"os"
)

var (
	Log = gl.New(os.Stdout, "[go-wiser]", gl.LstdFlags|gl.Lshortfile)
)
