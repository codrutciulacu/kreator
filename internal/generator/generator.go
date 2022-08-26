package generator

import (
	"kreator/internal/templates"
	"os"
)

func CreateProjectAt(projectName, location string) error {
	err := createFolderStructure(projectName, location)
	if err != nil {
		return err
	}

	return nil
}

func CreateModuleInProject(projectName, moduleName string) {

}

func createFolderStructure(projectName, location string) error {
	projectRoot := location + "/" + projectName
	err := os.Mkdir(projectRoot, os.ModeDir)
	if err != nil {
		return err
	}

	sourceFolder := projectRoot + "/src"
	err = os.Mkdir(sourceFolder, os.ModeDir)
	if err != nil {
		return err
	}

	testFolder := projectRoot + "/test"
	err = os.Mkdir(testFolder, os.ModeDir)
	if err != nil {
		return err
	}

	err = createMainFile(sourceFolder)
	if err != nil {
		return err
	}

	//TODO: Create test file

	return nil
}

func createMainFile(sourceFolder string) error {
	file, err := os.OpenFile(sourceFolder+"/main.c", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	templateContent := templates.SIMPLE_MAIN_C_TEMPLATE
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(templateContent))
	if err != nil {
		return err
	}

	return nil
}
