# Setting in VSCode

Я использую VSCode для работы с Go.

Ниже приведены настройки:

```json
{
	"settings": {
		"go.delveConfig": {
			"debugAdapter": "dlv-dap"
		},
		"go.lintOnSave": "package",
		"go.vetOnSave": "package",
		"go.useLanguageServer": true,
		"go.lintTool": "golangci-lint",
		"go.lintFlags": [
			"--fast",
			"--print-issued-lines=false",
			"--disable-all",
			"--enable=deadcode",
			"--enable=errcheck",
			"--enable=gosimple",
			"--enable=govet",
			"--enable=ineffassign",
			"--enable=staticcheck",
			"--enable=structcheck",
			"--enable=stylecheck",
			"--enable=typecheck",
			"--enable=unused",
			"--enable=varcheck",
			"--enable=bodyclose",
			"--enable=contextcheck",
			"--enable=decorder",
			"--enable=dupl",
			"--enable=durationcheck",
			"--enable=errchkjson",
			"--enable=errname",
			"--enable=exportloopref",
			"--enable=goconst",
			"--enable=gocritic",
			"--enable=gocyclo",
			"--enable=gosec",
			"--enable=misspell",
			"--enable=revive",
			"--enable=rowserrcheck",
			"--enable=sqlclosecheck",
			"--enable=unconvert",
			"--enable=unparam",
			"--enable=gas",
		],
		"files.eol": "\n",
		"makefile.extensionOutputFolder": "./.vscode"
	},
	"extensions": {
		"recommendations": [
			"golang.go",
			"zxh404.vscode-proto3"
		]
	}
}
```
