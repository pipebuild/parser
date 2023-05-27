package lexer

import (
	"context"

	"github.com/pipebuild/parser/config"
)

type Lexer interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Config config.Config
}

type lexer struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Lexer {
	return &lexer{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (l *lexer) Init(_ context.Context) error {
	return nil
}

func (l *lexer) Deinit(_ context.Context) error {
	return nil
}

func (l *lexer) Run(ctx context.Context) error {
	return nil
}
