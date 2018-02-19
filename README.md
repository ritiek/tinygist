# tinygist

A simple tool to to share gist content from and to clipboard.

## Installation

```
$ go get github.com/ritiek/tinygist
```

## Usage

```
$ tinygist -h
Usage of tinygist:
  -d    Copy gist contents to clipboard (default true)
  -i string
        Unique Identifier
  -u    Create gist from clipboard content
```

Create a new gist from your clipboard content:
```
$ tinygist -u <some_unique_keywords>
```

Get the gist content on another PC's clipboard:
```
$ tinygist -d <same_unique_keywords>
```

**BEWARE:** Make sure you do not have any sensitive data on your clipboard
before sharing your content. You won't be able to remove it later on!
(Even though we create a secret gist but prying eyes can reach anywhere!)

## How it Works?

- This tool creates a secret Gist from your clipboard content and then
shortens the generated raw URL to https://git.io/unique_identifer.
- When somebody tries to access this unique identifier, the tool
expands this shortened URL and places the Gist content on the clipboard.

## Contributing

Ideas, bug reports and pull requests are welcome but remember, we want
to keep this tool plain and simple.

## License

[![License](https://img.shields.io/github/license/ritiek/tinygist.svg)](https://github.com/ritiek/tinygist/blob/golang/LICENSE)
