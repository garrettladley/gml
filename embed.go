package gml

import _ "embed"

type Path string

const (
	CIGitHubWorkflowPath Path = ".github/workflows/ci.yml"
	GitIgnorePath        Path = ".gitignore"
	GolangCIPath         Path = ".golangci.yml"
	MakefilePath         Path = "Makefile"
)

//go:embed .github/workflows/ci.yml
var CIGitHubWorkflow []byte

//go:embed .gitignore
var GitIgnore []byte

//go:embed .golangci.yml
var GolangCI []byte

//go:embed: Makefile
var Makefile []byte

var ScaffoldFiles = map[Path][]byte{
	CIGitHubWorkflowPath: CIGitHubWorkflow,
	GitIgnorePath:        GitIgnore,
	GolangCIPath:         GolangCI,
	MakefilePath:         Makefile,
}
