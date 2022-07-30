# skeleton

## About

独自のスケルトンコードを生成します。

goの`text/template`を使用して、ファイルを生成します。

yamlファイルで変数を宣言し、`text/template`で使用し自由にカスタマイズ出来ます。

## Installation

go

```bash
go install -v github.com/jyury11/skeleton/cmd/skeleton@latest
```

linux

```bash
curl -sfL https://raw.githubusercontent.com/d-tsuji/qiisync/master/install.sh | sudo sh -s -- -b /usr/local/bin
```

[gihub releases](<https://github.com/jyury11/skeleton/releases>)

## Usage

```bash
skeleton create --service user --src ./template --dst ./dst --values values.yaml
```

## Example

[example](https://github.com/jyury11/skeleton/tree/main/example)

## Description

スケルトンの生成は、以下のルールに従います。

1. `text/template`に準拠したテンプレートファイルが使用出来る。
    - テンプレートはファイルの中身だけでなく、ディレクトリ名・ファイル名にも使用出来る。
2. 拡張子が `.tmpl`で終わるファイルは、最後からひとつ前の拡張子に変更される。
    - `user.go.tmpl` -> `user.go`
    - `docker-compose.yaml` -> `docker-compose.yaml`
3. yamlで宣言された変数ファイルを作成し、--values引数で指定しテンプレートで使用出来る。
4. 生成先にファイルが既にある場合、基本的には上書きをしない。ただし以下の場合、ファイルの上書きがされる。
    1. テンプレート・生成先ファイルの最初が、正規表現`^// Code generated .*; DO NOT EDIT.\n`にマッチするファイルは再生成される。
    2. --forceオプションを使用している場合、全てのファイルが再生成される。
5. `text/template`では次の追加の関数が使用出来る。
    1. `Echo`             : [`fmt.Println`](https://pkg.go.dev/fmt#Println)
    2. `CamelToSnake`     : キャメルケースをスネークケースに変換
    3. `CamelToUpperSnake`: キャメルケースをアッパースネークケースに変換
    4. `SnakeToCamel`     : スネークケースをキャメルケースに変換
    5. `ToUpper`          : [`strings.ToUpper`](https://golang.org/pkg/strings/#ToUpper)
    6. `ToLower`          : [`strings.ToLower`](https://golang.org/pkg/strings/#ToLower)
    7. `ToTitle`          : [`strings.ToTitle`](https://golang.org/pkg/strings/#ToTitle)
    8. `Title`            : [`cases.Title(language.Und, cases.NoLower).String`](https://pkg.go.dev/golang.org/x/text/cases#Title)

## Options

```bash
$ skeleton help create
create skeleton by template

Usage:
  skeleton create [flags]

Flags:
      --dst string       destination path (required)
  -f, --force            always overwrite files
  -h, --help             help for create
  -s, --service string   service name (required)
      --src string       source path (required)
  -v, --values string    values yaml path
```

## Advanced

ライブラリとして組み込んで使用する事も可能です。
その場合は`cmd/lib/main_test.go`を参考にしてください。

```go
package lib_test

import (
 "os"
 "path/filepath"
 "testing"

 "github.com/jyury11/skeleton/cmd/lib"
)

func TestMain(t *testing.T) {
 t.Run("main_lib", func(t *testing.T) {
  p, _ := os.Getwd()
  root := filepath.Join(p, "..", "..")
  src := filepath.Join(root, "example", "template")
  dst := filepath.Join(root, "example", "dst")
  val := filepath.Join(root, "example", "values.yaml")

  args := lib.CreateArgs{
   ServiceName: "user",
   Src:         src,
   Dst:         dst,
   Values:      val,
  }
  if err := lib.Create(args); err != nil {
   panic(err)
  }
 })
}


```

## License

[MIT](https://github.com/budougumi0617/lgen/blob/master/LICENSE)

## Author

[jyury11](https://github.com/jyury11)
