# go-osx-version

Get OSX version info.

## Installation

```
$ go get github.com/petehouston/go-osx-version
```

## Usage

```go
import "github.com/petehouston/go-osx-version"

osx := new(osxversion.OSXVersion)
if err := osx.Query(); err != nil {
	fmt.Println(err)
	return
}

fmt.Println("Name: " + osx.Name)
fmt.Println("Version: " + osx.Version)
fmt.Println("Build Version: " + osx.BuildVersion)
```