# GoFFI - Golang FFI

A Minimal Plugin Framework using FFI.

## Why?

Coming from a Windows user perspective who's been a Windows XP, 7, 8.1 and 10 user I prefer pluggable and extensive tools.

I wrote this for reference code on how a plugin framework can be written.

## How it works?

It uses FFI as the name implies. It looks for valid plugin sig and calls the respective functions. We use Golang generics for callable signature ambiguity.

## How to use

1. Build Plugin using `build.bat`
2. Run Manager using `go run main.go`

## Credits
- Me (CypherpunkSamurai)
- Windows 7 and 8.1 Software Creators for creating the coolest shit ever