package cmd

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kingpin/v2"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/pipebuild/parser/ast"
	"github.com/pipebuild/parser/config"
	"github.com/pipebuild/parser/lexer"
	"github.com/pipebuild/parser/parser"
	"github.com/pipebuild/parser/tokenizer"
)

var (
	app        = kingpin.New("parser", "pipebuild parser").Version(config.Version + "-build-" + config.Build)
	configFile = app.Flag("config-file", "Config file (.yml)").String()
	inputPath  = app.Flag("input-path", "Input path").Required().String()
	outputFile = app.Flag("output-file", "Output file (.json)").Required().String()
)

func Run(ctx context.Context) error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	c, _ := initConfig(ctx, *configFile)

	a, err := initAst(ctx, c)
	if err != nil {
		return errors.Wrap(err, "failed to init ast")
	}

	l, err := initLexer(ctx, c)
	if err != nil {
		return errors.Wrap(err, "failed to init lexer")
	}

	t, err := initTokenizer(ctx, c)
	if err != nil {
		return errors.Wrap(err, "failed to init tokenizer")
	}

	p, err := initParser(ctx, c)
	if err != nil {
		return errors.Wrap(err, "failed to init parser")
	}

	if err := runParser(ctx, a, l, t, p); err != nil {
		return errors.Wrap(err, "failed to run parser")
	}

	return nil
}

func initConfig(_ context.Context, name string) (*config.Config, error) {
	c := config.New()

	fi, err := os.Open(name)
	if err != nil {
		return c, errors.Wrap(err, "failed to open")
	}

	defer func() {
		_ = fi.Close()
	}()

	buf, _ := io.ReadAll(fi)

	if err := yaml.Unmarshal(buf, c); err != nil {
		return c, errors.Wrap(err, "failed to unmarshal")
	}

	return c, nil
}

func initAst(ctx context.Context, _ *config.Config) (ast.Ast, error) {
	c := ast.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	return ast.New(ctx, c), nil
}

func initLexer(ctx context.Context, _ *config.Config) (lexer.Lexer, error) {
	c := lexer.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	return lexer.New(ctx, c), nil
}

func initTokenizer(ctx context.Context, _ *config.Config) (tokenizer.Tokenizer, error) {
	c := tokenizer.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	return tokenizer.New(ctx, c), nil
}

func initParser(ctx context.Context, _ *config.Config) (parser.Parser, error) {
	c := parser.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	return parser.New(ctx, c), nil
}

func runParser(ctx context.Context, _ ast.Ast, _ lexer.Lexer, _ tokenizer.Tokenizer, p parser.Parser) error {
	if err := p.Init(ctx); err != nil {
		return errors.New("failed to init")
	}

	go func(ctx context.Context, p parser.Parser) {
		if err := p.Run(ctx); err != nil {
			log.Fatalf("failed to run: %v", err)
		}
	}(ctx, p)

	s := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be caught, so don't need add it
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func(ctx context.Context, p parser.Parser, s chan os.Signal, done chan bool) {
		<-s
		_ = p.Deinit(ctx)
		done <- true
	}(ctx, p, s, done)

	<-done

	return nil
}
