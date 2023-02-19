package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
)

type result struct {
	file string
	ok   bool
	err  error
}

var ignoreDirs = map[string]bool{
	"exercises": true,
}

func main() {
	now := time.Now()

	home, _ := os.UserHomeDir()
	repoDir := filepath.Join(home, "code")

	files, err := os.ReadDir(repoDir)
	if err != nil {
		log.Fatal(err)
	}

	results := processEntries(files)
	for res := range results {
		fmt.Println(res)
	}

	elapsed := time.Since(now)
	fmt.Printf("Duration %s", elapsed)
}

func processEntries(entries []os.DirEntry) <-chan result {
	results := make(chan result)

	go func() {
		for _, entry := range entries {
			results <- <-processEntry(entry)
		}
		close(results)
	}()

	return results
}

func processEntry(entry os.DirEntry) <-chan result {
	ch := make(chan result)

	go func() {
		defer close(ch)

		home, _ := os.UserHomeDir()
		repoDir := filepath.Join(home, "code")

		if !entry.IsDir() {
			fmt.Println("Not a directory. Skipping.")
			return
		}

		if _, ok := ignoreDirs[entry.Name()]; ok {
			return
		}

		fileDir := filepath.Join(repoDir, entry.Name())

		err := gitPull(fileDir)
		if err != nil {
			ch <- result{
				file: entry.Name(),
				ok:   false,
				err:  err,
			}
		}

	}()

	return ch
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
