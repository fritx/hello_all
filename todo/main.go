package main

// Run extra commands while being `go get`
// https://pkg.go.dev/cmd/go@master#hdr-Generate_Go_files_by_processing_source
//go:generate go run github.com/steebchen/prisma-client-go db push

import (
	"github.com/fritx/hello_all/todo/cmd"
)

func main() {
	cmd.Execute()
}
