package validate

import (
	"check-chart/constants"
	"check-chart/validate/image"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"

	"check-chart/models"
)

func baseChartFolderCheck(folder string) (*models.Chart, *models.AppConfiguration, string, error) {
	folderName := path.Base(folder)
	if !isValidFolderName(folderName) {
		return nil, nil, "", fmt.Errorf(constants.InvalidFolderName, folder)
	}

	if !dirExists(folder) {
		return nil, nil, "", fmt.Errorf(constants.FolderNotExist, folder)
	}

	chartFile := filepath.Join(folder, "Chart.yaml")
	if !fileExists(chartFile) {
		return nil, nil, "", fmt.Errorf(constants.MissingChartYaml, folder)
	}

	chartContent, err := os.ReadFile(chartFile)
	if err != nil {
		return nil, nil, "", fmt.Errorf(constants.ReadChartYamlFailed, folder, err)
	}
	var chart models.Chart
	if err := yaml.Unmarshal(chartContent, &chart); err != nil {
		return nil, nil, "", fmt.Errorf(constants.ParseChartYamlFailed, folder, err)
	}

	if err := isValidChartFields(chart); err != nil {
		return nil, nil, "", err
	}

	valuesFile := filepath.Join(folder, "values.yaml")
	if !fileExists(valuesFile) {
		return nil, nil, "", fmt.Errorf(constants.MissingValuesYaml, folder)
	}

	templatesDir := filepath.Join(folder, "templates")
	if !dirExists(templatesDir) {
		return nil, nil, "", fmt.Errorf(constants.MissingTemplatesFolder, folder)
	}

	appCfgFile := filepath.Join(folder, "TerminusManifest.yaml")
	if !fileExists(appCfgFile) {
		return nil, nil, "", fmt.Errorf(constants.MissingAppCfg, folder)
	}

	appCfgContent, err := os.ReadFile(appCfgFile)
	if err != nil {
		return nil, nil, "", fmt.Errorf(constants.ReadAppCfgFailed, folder, err)
	}
	var appConf models.AppConfiguration
	if err := yaml.Unmarshal(appCfgContent, &appConf); err != nil {
		return nil, nil, "", fmt.Errorf(constants.ParseAppCfgFailed, folder, err)
	}

	return &chart, &appConf, folderName, nil
}

func CheckChartFolder(folder string) error { // todo extract func
	chart, appConf, folderName, err := baseChartFolderCheck(folder)
	if err != nil {
		return err
	}

	if err := isValidMetadataFields(appConf.Metadata, chart, folderName); err != nil {
		return err
	}

	//if err := image.CheckAppConfigImages(&appConf); err != nil {
	//	return err
	//}

	return nil
}

func CheckChartFolderWithTitle(folder string, titleInfo models.TitleInfo) error {
	chart, appConf, folderName, err := baseChartFolderCheck(folder)
	if err != nil {
		return err
	}

	if err = isValidMetadataFieldsWithTitle(appConf.Metadata, chart, folderName, titleInfo); err != nil {
		return err
	}

	if !checkCategories(appConf.Metadata.Categories) {
		return fmt.Errorf(constants.InvalidCategories, appConf.Metadata.Categories, validCategoriesSlice)
	}

	if checkReservedWord(folderName) {
		return fmt.Errorf(constants.FolderNameInvalid, folderName)
	}

	if err = image.CheckAppConfigImages(appConf); err != nil {
		return err
	}

	return nil
}

func checkReservedWord(str string) bool {
	reservedWords := []string{
		"user", "system", "space", "default", "os", "kubesphere", "kube",
		"kubekey", "kubernetes", "gpu", "tapr", "bfl", "bytetrade",
		"project", "pod",
	}

	for _, word := range reservedWords {
		if strings.EqualFold(str, word) {
			return true
		}
	}

	return false
}

func isValidFolderName(name string) bool {
	match, _ := regexp.MatchString("^[a-z0-9]{1,30}$", name)
	return match
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return (err == nil || os.IsExist(err)) && !info.IsDir()
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return (err == nil || os.IsExist(err)) && info.IsDir()
}

func isValidChartFields(chart models.Chart) error {
	if chart.APIVersion == "" {
		return fmt.Errorf(constants.ApiVersionFieldEmptyInAppCfg, chart)
	}

	if chart.Name == "" {
		return fmt.Errorf(constants.NameFieldEmptyInAppCfg, chart)
	}

	if chart.Version == "" {
		return fmt.Errorf(constants.VersionFieldEmptyInAppCfg, chart)
	}

	return nil
}

func isValidMetadataFieldsWithTitle(metadata models.AppMetaData, chart *models.Chart, folder string, titleInfo models.TitleInfo) error {
	if chart.Name != folder || titleInfo.Folder != folder || metadata.Name != folder {
		return fmt.Errorf(constants.NameMustSame2,
			chart.Name, folder, titleInfo.Folder, metadata.Name)
	}

	if metadata.Version != chart.Version || titleInfo.Version != chart.Version {
		return fmt.Errorf(constants.VersionMustSame2, metadata.Version, chart.Version, titleInfo.Version)
	}

	return nil
}

func isValidMetadataFields(metadata models.AppMetaData, chart *models.Chart, folder string) error {
	if chart.Name != folder || metadata.Name != folder {
		return fmt.Errorf(constants.NameMustSame1,
			chart.Name, folder, metadata.Name)
	}

	if metadata.Version != chart.Version {
		return fmt.Errorf(constants.VersionMustSame1, metadata.Version, chart.Version)
	}

	return nil
}
