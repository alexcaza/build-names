# build-names

> Heavily inspired by this repository: https://github.com/fnichol/names

A simple cli tool that generates random hyphenated names in an `adjective-noun` format that can be used to identify builds.

Example output after running `build-names`:

```
$ build-names
> abject-hands
```

## Back story

I originally wrote a version [similar to this in Clojure](https://github.com/graingiant/build-names), and found that as a toy project it taught me a surprising amount about the language. It touches most of the basic things you'll encounter while building a program, and gives an enjoyable output. So, I've decided this is going to become my go-to project to start with when playing around with a new language.

## Building

Assuming you have `go` already installed.

`git clone git@github.com:alexcaza/build-names.git && cd build-names && sh build.sh`

## Options

```
-r "Include a random number in the name"
-rl "Include a random number in the name (max 6; default 3)"
-d "Include today's date (yyyy-mm-dd) in the name"
-n "How many names to generate (default 1)"
-a "Generate alliterative names (ex: abject-animal)"
```

## License

MIT License

Copyright (c) 2022 Alex Caza

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
