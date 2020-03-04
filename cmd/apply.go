package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/bialas1993/kubedeploy/pkg/command/apply"
	"github.com/bialas1993/kubedeploy/pkg/http"
	"github.com/bialas1993/kubedeploy/pkg/k8s"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ApplyCommand = &cobra.Command{
	Use:          "apply",
	Short:        "apply",
	Long:         `apply`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		namespace := os.Getenv("KUBE_NAMESPACE")
		token := os.Getenv("KUBE_TOKEN")
		host := strings.TrimRight(os.Getenv("KUBE_HOST"), "/") + "/"

		c := k8s.New(host, namespace, token, apply.New(http.NewClient(token)))

		ctx := context.Background()
		ctx = context.WithValue(ctx, "namespace", namespace)
		ctx = context.WithValue(ctx, "url", host)
		ctx = context.WithValue(ctx, "token", token)

		directory, _ := cmd.Flags().GetString("directory")
		file, _ := cmd.Flags().GetString("file")

		if err := c.Apply().Do(ctx, directory, file); err != nil {
			log.Error(err)
		}

		return nil
	},
}
