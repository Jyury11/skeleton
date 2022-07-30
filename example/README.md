# skeleton example

以下の様にテンプレート用のフォルダと、変数用のyamlが用意されている。

```
.
├── template
│   ├── docker-compose.yaml.tmpl
│   ├── go.mod
│   ├── Makefile
│   ├── README.md
│   └── {{ .ServiceName }}
│       ├── cmd
│       │   └── cli
│       │       ├── di
│       │       │   └── di.go.tmpl
│       │       ├── Dockerfile.tmpl
│       │       └── main.go.tmpl
│       └── internal
│           ├── entity
│           │   └── {{ .model.name }}.go.tmpl
│           ├── infra
│           │   └── std_repository.go.tmpl
│           ├── repository
│           │   └── repository.go.tmpl
│           ├── ui
│           │   └── cli.go.tmpl
│           └── usecase
│               ├── create.go.tmpl
│               └── usecase.go.tmpl
└── values.yaml
```

テンプレートは`text/template`が使用できる。
たとえば、{{ .model.name }}.go.tmplの中身は次の通りである。

```go
// Code generated by skeleton; DO NOT EDIT.

package entity
{{- $model := .model }}
const (
 {{ range $event := $model.events -}}
 {{ $event.message | Title }} = "{{ $event.message | CamelToUpperSnake }}"
 {{ end -}}
)

// {{ $model.name | Title }} ...
type {{ $model.name | Title }} struct {
 id int
}

// New{{ $model.name | Title }} ...
func New{{ $model.name | Title }}(id int) *{{ $model.name | Title }} {
 u := &{{ $model.name | Title }}{id}
 return u
}

// Id ...
func (m *{{ $model.name | Title }}) Id() int {
 return m.id
}

```

`text/template`で使用している変数はvalues.yamlに記載されている。
values.yamlの中身は次の通りである。

```yaml
model:
  name: user
  events:
    - name: create
      message: userCreateEvent
    - name: update
      message: userUpdateEvent
    - name: delete
      message: userDeleteEvent
```

サービス名、テンプレートのパス、出力のパス、変数ファイルのパス、上書きオプションを指定して実行する。

```bash
go run ../cmd/cli/main.go create -s user --src template --dst dst --values values.yaml -f
```

実行するとスケルトンが出力される。

```
.
├── dst
│   ├── docker-compose.yaml
│   ├── go.mod
│   ├── Makefile
│   ├── README.md
│   └── user
│       ├── cmd
│       │   └── cli
│       │       ├── di
│       │       │   └── di.go
│       │       ├── Dockerfile
│       │       └── main.go
│       └── internal
│           ├── entity
│           │   └── user.go
│           ├── infra
│           │   └── std_repository.go
│           ├── repository
│           │   └── repository.go
│           ├── ui
│           │   └── cli.go
│           └── usecase
│               ├── create.go
│               └── usecase.go
├── README.md
├── template
│   │
│   :
└── values.yaml
```

生成したファイルの動作を確認する。

```bash
$ cd dst
$ make up

docker-compose up
Recreating dst_user_1 ... done
Attaching to dst_user_1
user_1  | 2022/07/30 02:22:19 USER_CREATE_EVENT
user_1  | user_id: 99999
dst_user_1 exited with code 0
```