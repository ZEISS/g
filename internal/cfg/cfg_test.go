package cfg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeiss/g/internal/cfg"
)

func TestNew(t *testing.T) {
	t.Parallel()

	c := cfg.New()
	assert.NotNil(t, c)
	assert.False(t, c.Verbose)
	assert.Equal(t, "", c.Template)
	assert.False(t, c.Force)
}

func TestDefault(t *testing.T) {
	t.Parallel()

	c := cfg.Default()
	assert.NotNil(t, c)
	assert.False(t, c.Verbose)
	assert.Equal(t, "", c.Template)
	assert.False(t, c.Force)
}
