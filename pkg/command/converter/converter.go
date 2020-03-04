package converter

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/bialas1993/kubedeploy/pkg/fs"
	log "github.com/sirupsen/logrus"
)

const Extension = "kube.yaml"

type Converter interface {
	Do() error
}

type converter struct {
	output    string
	templates string
	env       string
	version   string
}

type converterParams struct {
	Output    string
	Templates string
	Env       string
	Version   string
}

var (
	ErrorTemplateDirectoryNotFound = errors.New("converter: Templates directory doesn't exists.")
)

func New(output string, templates string, env string, version string) Converter {
	return &converter{output, templates, env, version}
}

func (c *converter) valid() error {
	if !fs.NewFsReader().ExistsDir(c.templates) {
		return ErrorTemplateDirectoryNotFound
	}

	return nil
}

func (c *converter) Do() error {
	err := c.valid()
	if err == nil {
		outputDir := fmt.Sprintf("%s", c.output)
		reader := fs.NewFsReader()
		reader.CreateDir(outputDir)

		files, _ := reader.ReadDir(c.templates)
		files = func(files []os.FileInfo) []os.FileInfo {
			data := []os.FileInfo{}

			for _, file := range files {
				if strings.HasSuffix(file.Name(), Extension) {
					data = append(data, file)
				}
			}
			return data
		}(files)

		if len(files) == 0 {
			log.Info("Not found files to convert.")
			return nil
		}

		for _, file := range files {
			tpl := template.Must(template.New(file.Name()).ParseFiles(c.templates + "/" + file.Name()))

			log.Infof("Converting %s template...", file.Name())
			f, err := os.Create(outputDir + "/" + strings.Replace(file.Name(), Extension, "yaml", 1))
			if err != nil {
				return err
			}

			err = tpl.Execute(f, converterParams{
				Output:    c.output,
				Templates: c.templates,
				Env:       c.env,
				Version:   c.version,
			})

			f.Close()

			if err != nil {
				return err
			}
		}
	}

	return err
}
