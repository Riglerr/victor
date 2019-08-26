package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	switch args[1] {
	case "init":
		initializeVictorProject()
	case "build":
		err := RunCommand(NewCommand("docker-compose", "build", "--force-rm"))
		if err != nil {
			log.Fatal(err)
		}
	case "provision":
		err := RunCommand(NewCommand("terraform", "apply", "--auto-approve"))
		if err != nil {
			log.Fatal(err)
		}
	case "destroy":
		err := RunCommand(NewCommand("terraform", "destroy", "--auto-approve"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func initializeVictorProject() {
	err := RunCommand(NewCommand("touch", "Dockerfile"))
	if err != nil {
		log.Fatal(err)
	}
	err = RunCommand(NewCommand("touch", "Victorfile"))
	if err != nil {
		log.Fatal(err)
	}
	err = RunCommand(NewCommand("touch", "Jenkinsfile"))
	if err != nil {
		log.Fatal(err)
	}
}

