go mod init meeting
go mod tidy
go run internal\server\main.go
go build meeting/internal
go mod download github.com/davecgh/go-spew