package main

import (
	"go-ls/domain"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func ListFiles(path string) ([]*domain.DirItem, []*domain.FileItem, error) {
	f, err := os.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	var dirs []*domain.DirItem
	var files []*domain.FileItem

	for _, item := range f {
		info, _ := item.Info()
		stat := info.Sys().(*syscall.Stat_t)
		u := strconv.FormatUint(uint64(stat.Uid), 10)
		g := strconv.FormatUint(uint64(stat.Gid), 10)
		usr, _ := user.LookupId(u)
		group, _ := user.LookupGroupId(g)

		if info.IsDir() {
			dirs = append(dirs, &domain.DirItem{
				Item: domain.Item{
					Name:       item.Name(),
					Info:       info,
					UserOwner:  usr.Username,
					GroupOwner: group.Name,
				},
			})
		} else {
			files = append(files, &domain.FileItem{
				Item: domain.Item{
					Name:       item.Name(),
					Info:       info,
					UserOwner:  usr.Username,
					GroupOwner: group.Name,
				},
			})
		}
	}

	return dirs, files, nil
}
