package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"sync"
)

type Command struct {
	Command string
	Args []string
}

func NewCommand(command string, args ...string) *Command {
	var newCommand Command
	newCommand.Command = command
	newCommand.Args = args
	return &newCommand
}

func RunCommand(command* Command) error {
	cmd := exec.Command(command.Command, command.Args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	outScanner := bufio.NewScanner(stdout)
	errScanner := bufio.NewScanner(stderr)


	wg := sync.WaitGroup{}
	wg.Add(2)
	go readScanner(outScanner, &wg)
	go readScanner(errScanner, &wg)
	wg.Wait()

	if err := outScanner.Err(); err != nil {
		return err
	}

	if err := errScanner.Err(); err != nil {
		return err
	}

	return nil
}

func readScanner(s *bufio.Scanner, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for s.Scan() {
		fmt.Println(s.Text())
	}
}