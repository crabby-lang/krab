/**
* This uses both the crabby interpreter and the crab cli which acts like a builder, linter,
* formatter, and tester
* The compiler part however since crabby hasn't implemented a finished compiler yet it currently
* focuses on finishing the interpreter & VM bytecode first.
 */

package runner

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Executor interface {
	Execute(command string, args ...string) error
	ExecuteShell(command string, args ...string) error
	IsRunning() bool
	Stop()
}

type CrabRunner struct {
	interpreterPath string
	workingDir      string
}

func NewCrabRunner(interpreterPath, workingDir string) *CrabRunner {
	return &CrabRunner{
		interpreterPath: interpreterPath,
		workingDir:      workingDir,
	}
}

func (kr *CrabRunner) Run(executor Executor, filePath string) error {
	if executor.IsRunning() {
		return fmt.Errorf("another process is already running!")
	}

	// get absolute path
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("invalid file path: %w !", err)
	}

	// checks file extension
	if !strings.HasSuffix(strings.ToLower(absPath), ".crab") {
		return fmt.Errorf("invalid file extension, only .crab files are supported!")
	}

	if kr.interpreterPath != "" {
		return executor.Execute(kr.interpreterPath, absPath)
	}

	return executor.ExecuteShell("crabby", absPath)
}

func (kr *CrabRunner) RunWithArgs(executor Executor, filePath string, args []string) error {
	if executor.IsRunning() {
		return fmt.Errorf("another process is already running!")
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("invalid path path: %w !", err)
	}

	cmdArgs := append([]string{"run", absPath}, args...)

	if kr.interpreterPath != "" {
		return executor.Execute(kr.interpreterPath, cmdArgs...)
	}

	return executor.Execute("crabby", cmdArgs...)
}

func (kr *CrabRunner) Build(executor Executor, filePath, outputPath string) error {
	if executor.IsRunning() {
		return fmt.Errorf("another process is already running!")
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("invalid path path: %w !", err)
	}

	// crab build my_project (crab the cli, not the compiler, that'll be implemented sooner on)
	if kr.interpreterPath != "" {
		return executor.Execute(kr.interpreterPath, "build", absPath, "-o", outputPath)
	}

	return executor.Execute("crab", "build", absPath, "-o", outputPath)
}

func (kr *CrabRunner) Linter(executor Executor, filePath string) error {
	if executor.IsRunning() {
		return fmt.Errorf("another process is already running!")
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("invalid path path: %w !", err)
	}

	if kr.interpreterPath != "" {
		return executor.Execute(kr.interpreterPath, "lint", absPath)
	}

	return executor.Execute("crab", "lint", absPath)
}

func (kr *CrabRunner) Format(executor Executor, filePath string) error {
	if executor.IsRunning() {
		return fmt.Errorf("another process is already running!")
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("invalid path path: %w !", err)
	}

	if kr.interpreterPath != "" {
		return executor.Execute(kr.interpreterPath, "fmt", absPath)
	}

	return executor.Execute("crab", "fmt", absPath)
}

func (kr *CrabRunner) Test(executor Executor, filePath string) error {
	if executor.IsRunning() {
		return fmt.Errorf("another process is already running!")
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("invalid path path: %w !", err)
	}

	if kr.interpreterPath != "" {
		return executor.Execute(kr.interpreterPath, "test", absPath)
	}

	return executor.Execute("crab", "test", absPath)
}

func (kr *CrabRunner) GetVersionOfCrabby(executor Executor) error {
	if kr.interpreterPath != "" {
		return executor.Execute(kr.interpreterPath, "version") // `crabby version`
	}

	return executor.Execute("crab", "version")
}
