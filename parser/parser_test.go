package parser

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	var err error
	var p parser

	defer goleak.VerifyNone(t)

	ctx := context.Background()

	err = p.Init(ctx)
	assert.Equal(t, nil, err)

	err = p.Run(ctx)
	assert.Equal(t, nil, err)

	err = p.Deinit(ctx)
	assert.Equal(t, nil, err)
}
