package dstask

// an interface to the filesystem/git based database -- loading, saving, committing

import (
	"path"
	"os"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func MustGetRepoDirectory(directory ...string) string {
	root := MustExpandHome(GIT_REPO)
	return path.Join(append([]string{root}, directory...)...)
}

func LoadTaskSetFromDisk(statuses []string) *TaskSet {
	gitDotGitLocation := MustGetRepoDirectory(".git")

	if _, err := os.Stat(gitDotGitLocation); os.IsNotExist(err) {
		ExitFail("Could not find git repository at "+GIT_REPO+", please clone or create")
	}

	for _, status := range ALL_STATUSES {
		dir := MustGetRepoDirectory(status)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.Mkdir(dir, 0700)
			if err != nil {
				ExitFail("Failed to create directory in git repository")
			}
		}
	}

	return &TaskSet{
		knownUuids: make(map[string]bool),
	}
}

func (ts *TaskSet) SaveToDisk() {
	for _, task := range(ts.Tasks) {
		filepath := MustGetRepoDirectory(task.status, task.uuid+".yml")
		fmt.Println(filepath)
		d, err := yaml.Marshal(&task)
		fmt.Println(string(d), err)

		err = ioutil.WriteFile(filepath, d, 0600)
		if (err != nil) {
			ExitFail("Failed to write task")
		}
	}
}