package integration

import (
	"testing"

	"github.com/naggie/dstask"
	"github.com/stretchr/testify/assert"
)

func TestShowPaused(t *testing.T) {
	repo, cleanup := makeDstaskRepo(t)
	defer cleanup()

	program := testCmd(repo)

	output, exiterr, success := program("add", "+one", "one")
	assertProgramResult(t, output, exiterr, success)

	// "Paused" means tasks that were started, then stopped.
	output, exiterr, success = program("start", "1")
	assertProgramResult(t, output, exiterr, success)

	output, exiterr, success = program("stop", "1")
	assertProgramResult(t, output, exiterr, success)

	output, exiterr, success = program("show-paused")
	assertProgramResult(t, output, exiterr, success)

	var tasks []dstask.Task

	tasks = unmarshalTaskArray(t, output)
	assert.Equal(t, tasks[0].Summary, "one")
}
