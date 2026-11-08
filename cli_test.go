package cli

import (
	"testing"
)

func TestAllEmpty(t *testing.T) {
	args := []string{""}

	root := New()
	cmd, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if cmd != root {
		t.Errorf("Should be the root command")
	}
}

func TestNoSubCommandNoFlag(t *testing.T) {
	args := []string{""}

	root := New()

	root.AddSubCommand("help")

	cmd, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if cmd != root {
		t.Errorf("Should be the root command")
	}
}

func TestSubCommandOneRegistered(t *testing.T) {
	args := []string{"", "help"}

	root := New()

	helpCmd := root.AddSubCommand("help")

	selCmd, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if selCmd != helpCmd {
		t.Errorf("Commands do not match")
	}
}

func TestSubCommandMultipleRegisteredSelectFirst(t *testing.T) {
	args := []string{"", "version"}

	root := New()

	versionCmd := root.AddSubCommand("version")
	root.AddSubCommand("help")

	selCmd, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if selCmd != versionCmd {
		t.Errorf("Commands do not match")
	}
}

func TestSubCommandMultipleRegisteredSelectSecond(t *testing.T) {
	args := []string{"", "help"}

	root := New()

	root.AddSubCommand("version")
	helpCmd := root.AddSubCommand("help")

	selCmd, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if selCmd != helpCmd {
		t.Errorf("Commands do not match")
	}
}

func TestSubCommandMultipleRegistered5SelectSecond(t *testing.T) {
	args := []string{"", "help"}

	root := New()

	root.AddSubCommand("version")
	helpCmd := root.AddSubCommand("help")
	root.AddSubCommand("run")
	root.AddSubCommand("init")
	root.AddSubCommand("test")

	selCmd, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if selCmd != helpCmd {
		t.Errorf("Commands do not match")
	}
}

func TestSubCommandMultipleRegisteredSelectLast(t *testing.T) {
	args := []string{"", "test"}

	root := New()

	root.AddSubCommand("version")
	root.AddSubCommand("help")
	root.AddSubCommand("run")
	root.AddSubCommand("init")
	testCmd := root.AddSubCommand("test")

	selCmd, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if selCmd != testCmd {
		t.Errorf("Commands do not match")
	}
}

func TestNestedSubCommand(t *testing.T) {
	args := []string{"", "run", "release"}

	root := New()

	root.AddSubCommand("version")
	root.AddSubCommand("help")
	runCmd := root.AddSubCommand("run")
	root.AddSubCommand("init")
	root.AddSubCommand("test")

	runCmd.AddSubCommand("debug")
	runReleaseCmd := runCmd.AddSubCommand("release")

	selCmd, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if selCmd != runReleaseCmd {
		t.Errorf("Commands do not match")
	}
}

func TestSubCommandInvalidCommand(t *testing.T) {
	args := []string{"", "test"}

	root := New()

	root.AddSubCommand("version")
	root.AddSubCommand("help")
	root.AddSubCommand("run")

	_, err := root.Parse(args)
	if err == nil {
		t.Errorf("There should be an error.")
	}
}

func TestPresenceFlag(t *testing.T) {
	args := []string{"", "-flag"}

	root := New()

	seen := false
	root.BoolVar("flag", &seen, false, "")

	_, err := root.Parse(args)
	if err != nil {
		t.Error(err)
	}

	if !seen {
		t.Errorf("Should be true")
	}
}
