# benchmarking tool in golang
`go test -bench=. -cpu=1,2,4`
# Profiling tool in golang
```
## install
go get -u github.com/google/pprof
brew install graphviz # on mac
sudo apt-get install graphviz # on linux
## run pprof tool
go tool pprof /path/to/file/.pprof
go tool pprof -http=:8080 /path/to/file/.pprof
```
