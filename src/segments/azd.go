package segments

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/jandedobbeleer/oh-my-posh/src/log"
)

type Azd struct {
	base

	azdConfig
}

type azdConfig struct {
	DefaultEnvironment string `json:"defaultEnvironment"`
	Version            int    `json:"version"`
}

func (t *Azd) Template() string {
	return " \uebd8 {{ .DefaultEnvironment }} "
}

func (t *Azd) Enabled() bool {
	var parentFilePath string

	folders := t.props.GetStringArray(LanguageFolders, []string{".azure"})
	for _, folder := range folders {
		if file, err := t.env.HasParentFilePath(folder, false); err == nil {
			parentFilePath = file.ParentFolder
			break
		}
	}

	if parentFilePath == "" {
		log.Debug("no .azure folder found in parent directories")
		return false
	}

	dotAzureFolder := filepath.Join(parentFilePath, ".azure")
	files := t.env.LsDir(dotAzureFolder)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.EqualFold(file.Name(), "config.json") {
			return t.TryReadConfigJSON(filepath.Join(dotAzureFolder, file.Name()))
		}
	}

	return false
}

func (t *Azd) TryReadConfigJSON(file string) bool {
	if file == "" {
		return false
	}

	content := t.env.FileContent(file)
	var config azdConfig
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		return false
	}

	t.azdConfig = config
	return true
}
