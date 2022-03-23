package domain

import "os"

type Item struct {
	Name       string
	Info       os.FileInfo
	IsHidden   bool
	UserOwner  string
	GroupOwner string
}

type FileItem struct {
	Item
}

type DirItem struct {
	Item
}
