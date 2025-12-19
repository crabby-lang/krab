package terminal

import (
	"bufio"
	"io"
	"os/exec"
	"sync"
)

type Executor struct {
	buffer     *Buffer
	workingDir string
	cmd        *exec.Cmd
	mu         sync.Mutex
}

func NewExecutor(buffer *Buffer, workingDir string) *Executor {
	return &Executor{
		buffer:     buffer,
		workingDir: workingDir,
	}
}

func (e *Executor) Execute(command string, args ...string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.cmd = exec.Command(command, args...)
	e.cmd.Dir = e.workingDir

	stdout, err := e.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := e.cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := e.cmd.Start(); err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		e.readOutput(stdout)
	}()

	go func() {
		defer wg.Done()
		e.readOutput(stderr)
	}()

	wg.Wait()
	return e.cmd.Wait()
}

func (e *Executor) ExecuteShell(command string) error {
	return e.Execute("cmd", "/C", command)
}

func (e *Executor) readOutput(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		e.buffer.Append(line)
	}
}

func (e *Executor) IsRunning() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.cmd != nil && e.cmd.Process != nil && e.cmd.ProcessState == nil
}

func (e *Executor) Stop() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.cmd != nil && e.cmd.Process != nil {
		e.cmd.Process.Kill()
	}
}
