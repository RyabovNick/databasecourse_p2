# Setting in VSCode

Я использую VSCode для работы с Go.

Ниже приведены настройки:

```json
{
	"settings": {
		"go.lintOnSave": "package",
		"go.vetOnSave": "package",
		"go.useLanguageServer": true,
		"go.lintTool": "golangci-lint",
		"go.lintFlags": [
			"--print-issued-lines=false",
			"--no-config",
			"--disable-all",
			"--enable=govet",
			"--enable=stylecheck",
			"--enable=staticcheck",
			"--enable=dupl",
			"--enable=unconvert",
			"--enable=ineffassign",
			"--enable=goconst",
			"--enable=gosec",
			"--enable=unparam",
			"--enable=interfacer",
			"--enable=gocyclo",
			"--enable=errcheck",
			"--enable=deadcode",
			"--enable=varcheck",
			"--enable=gosimple",
			"--enable=structcheck",
			"--enable=typecheck",
			"--enable=unused",
			"--enable=misspell",
			"--enable=golint"
		],
		"files.eol": "\n"
	},
	"extensions": {
		"recommendations": ["golang.go", "zxh404.vscode-proto3"]
	}
}
```