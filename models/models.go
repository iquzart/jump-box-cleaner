// types created for the application
package models

import (
	"github.com/inhies/go-bytesize"
)

type Data struct {
	Hostname    string
	Date        string
	Path        string
	Directories []ParentDirectories
}

type ParentDirectories struct {
	Name           string
	Size           bytesize.ByteSize
	SubDirectories []SubDirectories
}

type SubDirectories struct {
	Name string
	Size bytesize.ByteSize
}
