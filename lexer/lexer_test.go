package lexer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	var err error
	var l lexer

	defer goleak.VerifyNone(t)

	ctx := context.Background()

	err = l.Init(ctx)
	assert.Equal(t, nil, err)

	err = l.Run(ctx)
	assert.Equal(t, nil, err)

	err = l.Deinit(ctx)
	assert.Equal(t, nil, err)
}
