// types created for the application
package models

import "github.com/inhies/go-bytesize"

type ParentDirectories struct {
	Name           string
	Size           bytesize.ByteSize
	SubDirectories []SubDirectories
}

type SubDirectories struct {
	Name string
	Size bytesize.ByteSize
}
