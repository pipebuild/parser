package cmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	ctx := context.Background()

	_, err := initConfig(ctx, "invalid.yml")
	assert.NotEqual(t, nil, err)

	_, err = initConfig(ctx, "../test/config/invalid.yml")
	assert.NotEqual(t, nil, err)

	_, err = initConfig(ctx, "../test/config/config.yml")
	assert.Equal(t, nil, err)
}

func TestInitAst(t *testing.T) {
	ctx := context.Background()

	c, err := initConfig(ctx, "../test/config/config.yml")
	assert.Equal(t, nil, err)

	_, err = initAst(ctx, c)
	assert.Equal(t, nil, err)
}

func TestInitLexer(t *testing.T) {
	ctx := context.Background()

	c, err := initConfig(ctx, "../test/config/config.yml")
	assert.Equal(t, nil, err)

	_, err = initLexer(ctx, c)
	assert.Equal(t, nil, err)
}

func TestInitTokenizer(t *testing.T) {
	ctx := context.Background()

	c, err := initConfig(ctx, "../test/config/config.yml")
	assert.Equal(t, nil, err)

	_, err = initTokenizer(ctx, c)
	assert.Equal(t, nil, err)
}

func TestInitParser(t *testing.T) {
	ctx := context.Background()

	c, err := initConfig(ctx, "../test/config/config.yml")
	assert.Equal(t, nil, err)

	_, err = initParser(ctx, c)
	assert.Equal(t, nil, err)
}
