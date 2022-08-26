package cli

import (
	"errors"
	"fmt"
	"kreator/internal/generator"
	"os"
	"path/filepath"
)

func ParseCommands(args []string) error {
	switch args[1] {
	case "project":
		err := handleProjectCommands(args)
		if err != nil {
			return err
		}
	case "module":
		err := handleModuleCommands(args)
		if err != nil {
			return err
		}
	case "help":
		// TODO: Print info about kreator
	default:
		fmt.Println("The command wasn't recognized!\n Use kreator help to see the complete list of commands!")
	}
	return nil
}

func handleModuleCommands(args []string) error {
	if args[2] == "create" {
		if len(args) != 3 {

			moduleName := args[3]
			generator.CreateModuleInProject("", moduleName)

			return nil
		} else {
			return errors.New("args: invalid number of arguments")
		}
	} else {
		return errors.New("project: invalid project option")
	}
}

func handleProjectCommands(args []string) error {
	if args[2] == "create" {
		if len(args) != 4 {
			location, err := getProjectLocation(args[4])
			if err != nil {
				return err
			}

			projectName := args[3]
			err = generator.CreateProjectAt(projectName, location)
			if err != nil {
				return err
			}

			return nil
		} else {
			return errors.New("args: invalid number of arguments")
		}
	} else {
		return errors.New("project: invalid project option")
	}
}

func getProjectLocation(argument string) (string, error) {
	if argument == "." {
		crtPath, err := os.Executable()
		if err != nil {
			return "", err
		}

		return filepath.Dir(crtPath), nil
	} else if _, err := os.Stat(argument); err == nil {
		return argument, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return "", err
	} else {
		return "", err
	}
}
