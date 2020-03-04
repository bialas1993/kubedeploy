package apply

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/bialas1993/kubedeploy/pkg/fs"
	"github.com/bialas1993/kubedeploy/pkg/http"
	khttp "github.com/bialas1993/kubedeploy/pkg/k8s/http"
	"github.com/bialas1993/kubedeploy/pkg/k8s/template"
	log "github.com/sirupsen/logrus"
)

const (
	ErrorCodeCanNotReadFile = iota + 120
	ErrorCodeCanNotBuildTemplate
	ContentTypeYaml = "application/yaml"
)

var (
	ErrorCanNotReadFile      = errors.New("service: Can not read file.")
	ErrorCanNotReadDirectory = errors.New("service: Can not directory.")
	ErrorCanNotBuildTemplate = errors.New("service: Can not build template.")
	ErrorCanNotApplyConfig   = errors.New("service: Can not apply config.")
)

type apply struct {
	client *http.Client
}

func (d *apply) Do(ctx context.Context, directory string, filePath string) error {
	var result khttp.Response

	directory = strings.TrimRight(directory, "/") + "/"
	reader := fs.NewFsReader()
	files := []string{}

	if len(filePath) > 0 {
		files = append(files, filePath)
	}

	dirFiles, err := reader.ReadDir(directory)
	if err != nil {
		log.WithError(err).Error(ErrorCanNotReadDirectory)
		return err
	}

	for _, file := range dirFiles {
		files = append(files, strings.Join([]string{directory, file.Name()}, ""))
	}

	for _, file := range files {
		data, err := reader.ReadFile(file)
		if err != nil {
			log.WithError(err).Error(ErrorCanNotReadFile)
			return err
		}

		tpl, err := template.Build(data)
		if err != nil {
			log.WithError(err).Error(ErrorCanNotBuildTemplate)
			return err
		}

		log.Infof("Deploying %s", file)

		url := khttp.BuildUrl(ctx, tpl)
		resp, err := d.client.Do("POST", url, bytes.NewBuffer(data))
		if err != nil {
			log.WithError(err).Error(ErrorCanNotApplyConfig)
			return err
		}

		err = json.NewDecoder(resp.Body).Decode(&result)

		log.Debugf("%+v, %+v", err, result)

		if err = khttp.ValidResponse(err, &result); err != nil {
			return err
		}
	}

	return nil
}

func New(client *http.Client) *apply {
	return &apply{client}
}
