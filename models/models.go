// types created for the application
package models

type ParentDirectories struct {
	Name           string
	Size           float64
	SubDirectories []SubDirectories
}

type SubDirectories struct {
	Name string
	Size float64
}
