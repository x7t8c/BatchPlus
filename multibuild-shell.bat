:: Delete the old builds
DEL batch-plus-shell-*
:: Build for darwin/amd64
SET GOOS=darwin
SET GOARCH=amd64
go build shell.go batch.go release.go cmd.go
RENAME shell batch-plus-shell-darwin-amd64
:: Build for darwin/arm64
SET GOOS=darwin
SET GOARCH=arm64
go build shell.go batch.go release.go cmd.go
RENAME shell batch-plus-shell-darwin-arm64
:: Build for freebsd/386
SET GOOS=freebsd
SET GOARCH=386
go build shell.go batch.go release.go cmd.go
RENAME shell batch-plus-shell-freebsd-386
:: Build for freebsd/amd64
SET GOOS=freebsd
SET GOARCH=amd64
go build shell.go batch.go release.go cmd.go
RENAME shell batch-plus-shell-freebsd-amd64
:: Build for linux/amd64
SET GOOS=linux
SET GOARCH=amd64
go build shell.go batch.go release.go cmd.go
RENAME shell batch-plus-shell-linux-amd64
:: Build for linux/386
SET GOOS=linux
SET GOARCH=386
go build shell.go batch.go release.go cmd.go
RENAME shell batch-plus-shell-linux-386
:: Build for windows/386
SET GOOS=windows
SET GOARCH=386
go build shell.go batch.go release.go cmd.go
RENAME shell.exe batch-plus-shell-windows-386.exe
:: Build for windows/amd64
SET GOOS=windows
SET GOARCH=amd64
go build shell.go batch.go release.go cmd.go
RENAME shell.exe batch-plus-shell-windows-amd64.exe
