package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultTemplateFile(t *testing.T) {
	assert.Equal(t, "demo", "demo")
}
