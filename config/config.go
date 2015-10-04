package config

import (
	"log"
	"os"
	"path"

	"github.com/libgit2/git2go"

	"github.com/ilkka/seita/ui"
)

var repository *git.Repository

// GetSeitaPath returns the full path of the directory where all
// seita files are kept.
func GetSeitaPath() string {
	return path.Join(os.ExpandEnv("$HOME"), ".seita")
}

// GetRepoURL returns the URL to the repository.
func GetRepoURL() string {
	return "/Users/ilau/src/seitarepo"
}

// GetRepoPath returns the directory where the repo should be checked out.
func GetRepoPath() string {
	return path.Join(GetSeitaPath(), "repo")
}

// Make sure a basic configuration exists.
func init() {
	if !repoIsCloned() {
		repository = cloneRepo()
	} else {
		repository = openRepo()
		updateRepo()
	}
}

func repoIsCloned() bool {
	fileinfo, err := os.Stat(path.Join(GetRepoPath(), ".git"))
	if err == nil && fileinfo.IsDir() {
		return true
	}
	return false
}

func openRepo() *git.Repository {
	repo, err := git.OpenRepository(GetRepoPath())
	if err != nil {
		log.Fatalf("Could not open seita repo: %s", err)
	}
	return repo
}

func updateRepo() {
	ui.Printf("Updating seita repo...\n")
	origin, err := repository.Remotes.Lookup("origin")
	if err != nil {
		log.Fatalf("Could not find origin for repo: %s", err)
	}
	err = origin.Fetch([]string{}, &git.FetchOptions{}, "")
}

func cloneRepo() *git.Repository {
	ui.Printf("Cloning seita repo...\n")
	repo, err := git.Clone(GetRepoURL(), GetRepoPath(), &git.CloneOptions{})
	if err != nil {
		log.Fatalf("Could not clone seita repo: %s", err)
	}
	return repo
}
