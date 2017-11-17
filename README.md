# embedfiles
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/embedfiles)](https://goreportcard.com/report/github.com/leighmcculloch/embedfiles)

Embedfiles is a tool for embedding files into Go code.

Files are gzip compressed and stored in a map of filenames to file data.

## Install

### Source

```
go get 4d63.com/embedfiles
```

## Usage

```
$ embedfiles
Embedfiles embeds files into a map in a go file.

Usage:

  embedfile -out=files.go -pkg=main <paths>

Flags:

  -out file
        output go file (default "files.go")
  -pkg package
        package name of the go file (default "main")
```