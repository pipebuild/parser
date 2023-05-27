package ast

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	var err error
	var a ast

	defer goleak.VerifyNone(t)

	ctx := context.Background()

	err = a.Init(ctx)
	assert.Equal(t, nil, err)

	err = a.Run(ctx)
	assert.Equal(t, nil, err)

	err = a.Deinit(ctx)
	assert.Equal(t, nil, err)
}
