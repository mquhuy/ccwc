# CCWC - A Wc Clone written in Go

This is a clone of [the UNIX command tool `wc`](https://en.wikipedia.org/wiki/Wc_(Unix)) written in Golang. As a clone, it tries to mimic all behaviors from the origin; however, there are still some minor differences.

Currently, `ccwc` supports only one file input, and no `STDIN` input is supported. The support is not guaranteed to be added in the future.

## Usage

### Clone and build the binary
```
git clone <this repo>
cd ccwc
go build
```

### Running
```
./ccwc [FLAGS] [FILE]
```

The following flags are currently supported:
```
  -c, --bytes   Print the byte counts
  -m, --chars   Print the character counts
  -h, --help    help for ccwc
  -l, --lines   Print the newline counts
  -w, --words   Print the word counts

```

If no flags is provided, the outputs will be in format `[line_count] [word_count] [byte_count]`.

Whenever multiple outputs is required, the structure will be `[line_count` (if applicable) -> `[word_count]` (if applicable) -> `[char_count]` (if applicable -> `[byte_count]` (if applicable). This is also the output format used by `wc`.

This program was written thanks to the encouragement from [this coding challenge](https://codingchallenges.fyi/challenges/challenge-wc/). However, the exact implimentation didn't follow what was described there, but rather the behavior of `wc` tool version `(GNU coreutils) 8.32`.
