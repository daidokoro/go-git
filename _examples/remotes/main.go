package main

import (
	"fmt"

	"github.com/daidokoro/go-git"
	. "github.com/daidokoro/go-git/_examples"
	"github.com/daidokoro/go-git/config"
	"github.com/daidokoro/go-git/plumbing"
	"github.com/daidokoro/go-git/storage/memory"
)

// Example of how to:
// - Create a new in-memory repository
// - Create a new remote named "example"
// - List remotes and print them
// - Pull using the new remote "example"
// - Iterate the references again, but only showing hash references, not symbolic ones
// - Remove remote "example"
func main() {
	// Create a new repository
	Info("git init")
	r, err := git.Init(memory.NewStorage(), nil)
	CheckIfError(err)

	// Add a new remote, with the default fetch refspec
	Info("git remote add example https://github.com/git-fixtures/basic.git")
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "example",
		URL:  "https://github.com/git-fixtures/basic.git",
	})

	CheckIfError(err)

	// List remotes from a repository
	Info("git remotes -v")

	list, err := r.Remotes()
	CheckIfError(err)

	for _, r := range list {
		fmt.Println(r)
	}

	// Pull using the create repository
	Info("git pull example")
	err = r.Pull(&git.PullOptions{
		RemoteName: "example",
	})

	CheckIfError(err)

	// List the branches
	// > git show-ref
	Info("git show-ref")

	refs, err := r.References()
	CheckIfError(err)

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		// The HEAD is omitted in a `git show-ref` so we ignore the symbolic
		// references, the HEAD
		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}

		fmt.Println(ref)
		return nil
	})

	CheckIfError(err)

	// Delete the example remote
	Info("git remote rm example")

	err = r.DeleteRemote("example")
	CheckIfError(err)
}
