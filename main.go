package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/okteto/cnd/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	cmd.Execute()
}
