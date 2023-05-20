# parser

[![Build Status](https://github.com/pipebuild/parser/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/pipebuild/parser/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/pipebuild/parser/branch/main/graph/badge.svg?token=JRENVWAT7J)](https://codecov.io/gh/pipebuild/parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/pipebuild/parser)](https://goreportcard.com/report/github.com/pipebuild/parser)
[![License](https://img.shields.io/github/license/pipebuild/parser.svg)](https://github.com/pipebuild/parser/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/pipebuild/parser.svg)](https://github.com/pipebuild/parser/tags)



## Introduction

*parser* is the parser of [pipebuild](https://github.com/pipebuild) written in Go.



## Prerequisites

- Go >= 1.18.0



## Run

```bash
version=latest make build
./bin/parser --config-file="$PWD"/test/config/config.yml
```



## Docker

```bash
version=latest make docker
docker run -v "$PWD"/test:/tmp ghcr.io/pipebuild/parser:latest --config-file=/tmp/config/config.yml
```



## Usage

```
usage: parser --config-file=CONFIG-FILE [<flags>]

pipebuild parser

Flags:
  --[no-]help                Show context-sensitive help (also try --help-long
                             and --help-man).
  --[no-]version             Show application version.
  --config-file=CONFIG-FILE  Config file (.yml)
```



## Settings

*parser* parameters can be set in the directory [config](https://github.com/pipebuild/parser/blob/main/config).

An example of configuration in [config.yml](https://github.com/pipebuild/parser/blob/main/config/config.yml):

```yaml
TBD
```



## License

Project License can be found [here](LICENSE).



## Reference
