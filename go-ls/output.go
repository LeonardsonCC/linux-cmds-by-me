package main

import (
	"fmt"
	"go-ls/domain"
	"strings"
)

type Output struct {
	Files          []*domain.FileItem
	Dirs           []*domain.DirItem
	ShowHidden     bool
	ShowUserOwner  bool
	ShowGroupOwner bool
}

func NewOutput(dirs []*domain.DirItem, files []*domain.FileItem) *Output {
	return &Output{
		Dirs:  dirs,
		Files: files,
	}
}

func (o *Output) Print() {
	s := ""

	for _, dir := range o.Dirs {
		s += o.FormatDir(dir)
	}

	for _, file := range o.Files {
		s += o.FormatFile(file)
	}

	fmt.Print(s)
}

func (o *Output) FormatFile(file *domain.FileItem) string {
	name := fmt.Sprintf("%s%s%s", White, file.Name, Reset)

	if file.Name[0] == '.' && !o.ShowHidden {
		return ""
	}
	if strings.Contains(file.Info.Mode().String(), "x") {
		name = fmt.Sprintf("%s%s%s", BoldGreen, file.Name, Reset)
	}
	return fmt.Sprintf("%s %s %s %s\n", file.Info.Mode().String(), showIf(o.ShowGroupOwner, file.GroupOwner), showIf(o.ShowUserOwner, file.UserOwner), name)
}

func (o *Output) FormatDir(dir *domain.DirItem) string {
	name := fmt.Sprintf("%s%s%s", Blue, dir.Name, Reset)

	if dir.Name[0] == '.' && !o.ShowHidden {
		return ""
	}
	return fmt.Sprintf("%s %s %s %s\n", dir.Info.Mode().String(), showIf(o.ShowGroupOwner, dir.GroupOwner), showIf(o.ShowUserOwner, dir.UserOwner), name)
}

func showIf(show bool, s string) string {
	if show {
		return s
	}
	return ""
}
