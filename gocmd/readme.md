---
author: "" 
paging: ""
date: ""
---

# Go command and tools
Command/tools that can make your life more easier on golang.

---

## go env
list all the env related to go
for example
```sh
	go env GOPATH
```

### GOPATH
under the GOPATH, you should see 3 main directories.
- src: all Go source code lying here
- bin: if you use `go install` or `go get` the binary will be stored here.
- pkg: cache of the compiled code.

so this is the reason why when you install something via `go install` you will have to bind by doing this. `export PATH="$PATH:$GOPATH/bin"`  

---

## go get
Common flag that we used with go get
- `go get -d`: instructs get to stop after downloading the packages, only download but not compile
- `go get -u`: use the network to update the named packages
and their dependencies. By default, get uses the network to check out
missing packages but does not use it to look for updates to existing packages

try `go install github.com/rakyll/gotest@latest` 
and run `gotest -v ./...` instead of plain `go test`

---

### go get (cont)
just like I used this tool (https://github.com/maaslalani/slides) by install via
```sh
go install github.com/maaslalani/slides@latest
```

---

## go list
This is rarely used but we can use this to list all of the package inside the sourcecode.
`go list ./...`

---

## gofmt
Go community want to simplify and standard the format of the language so they come up with tool to format the code so every body in will be using the same style guide.
```sh
gofmt format/main.go # will print the formatted version on stdout
gofmt -d format/main.go # print the diff
gofmt -w format/main.go # automatically rewrite the file
# or you can use
go fmt .
```

---
## gofmt-simplify
gofmt has -s flag that will tell go to try simplify the code.
For example
```go
A slice expression of the form:
	s[a:len(s)]
will be simplified to:
	s[a:]

A range of the form:
	for x, _ = range v {...}
will be simplified to:
	for x = range v {...}
```
try run `gofmt -s simplify/main.go` to see the result.

---

## gomodifytags
Add tag in struct is tedious task, I know we all have been there so let me introduce you to `gomodifytags` command
Let's say you have Car struct
```go
type Car struct {
	Brand string
	Engine string
}
``` 
and you want to tag json tag so json.marshal can convert it instead of doing it manually you can use.
`gomodifytags -file modify/main.go -add-tags json -add-options json=omitempty -struct Car package main`
and as the same with gofmt you can pass `-w` to rewrite the file.

---

## go vet
Go vet help detect mistakes in a code by doing static analysis on the code. For more details on what go vet can do, I advise you to run another command `go doc` by simply run `go doc cmd/vet`.
`go vet vet/main.go`

---

## errcheck
You know that go treat error as a value so sometimes you will find a function like this
```go
func Hey(name string) error {
	return errors.New("error is here")
}
```
You can call this function like this `Hey("Arthur")` but you might not be aware the you silently omitting the error
luckily we have this errcheck 

---

## errcheck (cont)
simply install this tool via
`go install github.com/kisielk/errcheck@latest`
then `errcheck errchk/main.go`
or `errcheck ./...`

---

## staticcheck 
Another version of go lint provided by 3rd party
To install `go install honnef.co/go/tools/cmd/staticcheck@latest`
then `staticcheck ./...` 

---
