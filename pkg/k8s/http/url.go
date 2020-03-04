package http

import (
	"context"
	"fmt"
	"strings"

	"github.com/bialas1993/kubedeploy/pkg/k8s/template"
	log "github.com/sirupsen/logrus"
)

func BuildUrl(ctx context.Context, source *template.Source) string {
	url := fmt.Sprintf("%s", ctx.Value("url")) + version(strings.ToLower(source.ApiVersion)) +
		"/namespaces/" + fmt.Sprintf("%s", ctx.Value("namespace")) + "/" + kind(strings.ToLower(source.Kind))
	log.Debug(url)

	return url
}

func version(version string) string {
	if version == "v1" {
		return "api/v1"
	}

	return "apis/" + version
}

func kind(kind string) string {
	if strings.HasSuffix(kind, "s") {
		return kind + "es"
	}

	return kind + "s"
}
