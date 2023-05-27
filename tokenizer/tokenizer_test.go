package tokenizer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	var err error
	var token tokenizer

	defer goleak.VerifyNone(t)

	ctx := context.Background()

	err = token.Init(ctx)
	assert.Equal(t, nil, err)

	err = token.Run(ctx)
	assert.Equal(t, nil, err)

	err = token.Deinit(ctx)
	assert.Equal(t, nil, err)
}
