# Color4Tilix

This is a library to convert existing terminal color schemes in [Tilix](https://gnunn1.github.io/tilix-web/) compatibly color schemes.


## How

Tilix uses json in to save color schemes. This library provides a struct `TilixColor` which is responsible to the translation.

The library also provides an interface, `Paletter`, to allow the creation of parsers. Parses will used by `NewTilixColor` function which is in charge to provide a correct `TilixColor` struct.

Generally a conversion process uses a list of input files, an output directory where to save results and a set of instruction to translate input files to output ones. For this reason the library provides `CreateBatch` function which given a `FncTransformer` transform function, translates files for Tilix color schemes.

## Example

The library provides a parser ([itermcolors_parser.go](itermcolors_parser.go)) in order to translate [iTerm](https://iterm2.com/) color schemes in Tilix color schemes.

The CLI tool ([convert.go](iterm2tilix/convert.go)) provided by the library allow you to convert Iterm color schemes and can be used an example to create your own color scheme translator.

To get cli tool for conver .itermcolors schema file into Tilix color scheme install using:

```
go get -u github.com/tux-eithel/color4tilix/...
```

and then run `iterm2tilix -d dir_with_itermcolors_files`.

## How to transform \*term\* color schemes into Tilix color schemes

If you'd like to transform another terminal schema into Tilix color schemes, (and you're fine with `CreateBatch` workflow) you should do:
 
 - define a new struct which implements `Paletter` interface
 - provide a compatible `FncTransformer` function to transform input files
 - use the ([convert.go](iterm2tilix/convert.go)) as base
