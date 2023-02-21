# easerial

[![Go](https://github.com/kinpoko/easerial/actions/workflows/go.yml/badge.svg)](https://github.com/kinpoko/grvemu/actions/workflows/go.yml)
![License](https://img.shields.io/github/license/kinpoko/easerial?color=blue)

A simple command line application for serial communication.

## Install

```bash
go install github.com/kinpoko/easerial@latest
```

## Usage

easerial has two subcommand, `easerial string` and `easerial file`.

## Send data by hex string

use `easerial string`

```bash
easerial string 68656c6c6f20776f726c64 --baud 115200 --port /dev/ttyUSB0
```

## Send data by hex text file

use `easerial file`

```bash
easerial file helloworld.txt --baud 115200 --port /dev/ttyUSB0
```

`helloworld.txt`

```text
68
65
6c
6c
6f
20
77
6f
72
6c
64
```

or

```text
68,65,6c,6c,6f,20,77,6f,72,6c,64
```
