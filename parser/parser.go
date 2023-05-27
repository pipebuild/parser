package parser

import (
	"context"

	"github.com/pipebuild/parser/config"
)

type Parser interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Config config.Config
}

type parser struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Parser {
	return &parser{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (p *parser) Init(_ context.Context) error {
	return nil
}

func (p *parser) Deinit(_ context.Context) error {
	return nil
}

func (p *parser) Run(ctx context.Context) error {
	return nil
}
