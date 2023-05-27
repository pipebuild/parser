package ast

import (
	"context"

	"github.com/pipebuild/parser/config"
)

type Ast interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Config config.Config
}

type ast struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Ast {
	return &ast{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (a *ast) Init(_ context.Context) error {
	return nil
}

func (a *ast) Deinit(_ context.Context) error {
	return nil
}

func (a *ast) Run(ctx context.Context) error {
	return nil
}
