package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestRequestCommand_implement(t *testing.T) {
	var _ cli.Command = &RequestCommand{}
}
