# topology-go

## dependence manage

- golang 1.8+
- glide 0.12.3+

use glide for dependence manage

install [glide](https://github.com/Masterminds/glide)

## format
```bash
# see which code needs format
find . -path ./vendor -prune -o -name '*.go' -print | xargs gofmt -d | grep '^'
# format
glide nv | xargs go fmt
```

## before run
add $HOME/.glide/mirrors.yaml
and add the settings below(unless you can connect to https://golang.org)
```
repos:
- original: https://golang.org/x/crypto
  repo: https://github.com/golang/crypto
  vcs: git
- original: https://golang.org/x/image
  repo: https://github.com/golang/image
  vcs: git
- original: https://golang.org/x/mobile
  repo: https://github.com/golang/mobile
  vcs: git
- original: https://golang.org/x/net
  repo: https://github.com/golang/net
  vcs: git
- original: https://golang.org/x/sys
  repo: https://github.com/golang/sys
  vcs: git
- original: https://golang.org/x/sys/unix
  repo: https://github.com/golang/sys
  vcs: git
  base: golang.org/x/sys
- original: https://golang.org/x/text
  repo: https://github.com/golang/text
  vcs: git
- original: https://golang.org/x/tools
  repo: https://github.com/golang/tools
  vcs: git
- original: https://gopkg.in/yaml.v2
  repo: https://github.com/go-yaml/yaml
  vcs: git
- original: https://golang.org/x/text/transform
  repo: https://github.com/golang/text
  vcs: git
  base: golang.org/x/text
- original: https://golang.org/x/text/unicode/norm
  repo: https://github.com/golang/text
  vcs: git
  base: golang.org/x/text
```

## run
```bash
glide install
go run topology-go.go
```

## test

```bash
# run all test
glide nv | xargs go test
# show coverage
glide nv | xargs go test -cover
# for more information
glide nv | xargs go test -v
```
