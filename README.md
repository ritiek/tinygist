# tinygist

A simple tool to to share gist content from and to clipboard.

## Installation

```
$ gem install tinygist
```

## Usage

```
$ tinygist
usage: tinygist [-d | -u] <unique_identifier>
```

Create a new gist from your clipboard content:
```
$ tinygist -u <some_unique_keywords>
```

Get the gist content on another PC's clipboard:
```
$ tinygist -d <same_unique_keywords>
```

**BEWARE:** Make sure you do not have any sensitive data on your
clipboard before sharing your content. You won't be able to remove it later on!
(Even though we create a secret gist but prying eyes can reach anywhere!)

## Contributing

Bug reports and pull requests are welcome. This project is intended to be a safe
, welcoming space for collaboration, and contributors are expected to adhere to
the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The gem is available as open source under the terms of the
[MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the tinygist project’s codebases, issue trackers,
chat rooms and mailing lists is expected to follow
the [code of conduct](https://github.com/[USERNAME]/tinygist/blob/master/CODE_OF_CONDUCT.md).
