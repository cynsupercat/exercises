package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
)

func main() {
	now := time.Now()
	const (
		thisDir = "exercises"
	)

	home, _ := os.UserHomeDir()
	repoDir := filepath.Join(home, "code")
	err := os.Chdir(repoDir)
	if err != nil {
		panic(err)
	}

	files, err := os.ReadDir(repoDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("Processing %s\n", file.Name())
		if !file.IsDir() {
			fmt.Println("Not a directory. Skipping.")
			continue
		}

		if file.Name() == thisDir {
			fmt.Println("This program's directory. Ignoring.")
			continue
		}

		fileDir := filepath.Join(repoDir, file.Name())

		err = gitPull(fileDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	elapsed := time.Since(now)

	// Duration 22.9686205s
	fmt.Printf("Duration %s", elapsed)
}

func gitPull(repoDir string) error {
	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		if err.Error() == "repository does not exist" {
			fmt.Println(err.Error())
			return nil
		}

		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = worktree.Pull(&git.PullOptions{})
	if err != nil {
		if err.Error() == "already up-to-date" || err.Error() == "remote not found" {
			fmt.Println(err.Error())
			return nil
		}

		return err
	}

	ref, err := repo.Head()
	if err != nil {
		return err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	fmt.Println(commit)
	return nil
}
