package filesystem

import (
	"github.com/daidokoro/go-git/storage"
	"github.com/daidokoro/go-git/storage/filesystem/internal/dotgit"
)

type ModuleStorage struct {
	dir *dotgit.DotGit
}

func (s *ModuleStorage) Module(name string) (storage.Storer, error) {
	return NewStorage(s.dir.Module(name))
}
