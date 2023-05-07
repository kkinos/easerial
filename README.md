# easerial

[![Go](https://github.com/kinpoko/easerial/actions/workflows/go.yml/badge.svg)](https://github.com/kinpoko/easerial/actions/workflows/go.yml)
![License](https://img.shields.io/github/license/kinpoko/easerial?color=blue)

A simple command line application for serial communication.

## Description

`easerial` is a simple command line application for sending data through a serial port. It provides two subcommands, `easerial string` and `easerial file`, which allow you to send hex data from a string or from a text file.

## Install

```bash
go install github.com/kinpoko/easerial@latest
```

## Usage

### `easerial string`

```bash
easerial string <string> [--port <port name>] [--baud <baud rate>] [--data-bits <data bits>] [--read-bytes <read bytes>] [--read-timeout-sec <read timeout sec>]
```

- `<string>`: Required. The string to send as hex data.
- `--port`: Optional. The serial port to use. Defaults to /dev/ttyUSB1.
- `--baud`: Optional. The baud rate to use. Defaults to 9600.
- `--data-bits`: Optional. The number of data bits to send. Defaults to 8.
- `--read-bytes`: Optional. The number of reading bytes after sending. Defaults to 4.
- `--read-timeout-sec`: Optional. The number of seconds of timeout for reading serial port after sending. Defaults to 1.

Example:

```bash
easerial string 68656c6c6f20776f726c64 --baud 115200 --port /dev/ttyUSB0
```

### `easerial file`

```bash
easerial file <file_path> [--port <port name>] [--baud <baud rate>] [--data-bits <data bits>] [--read-bytes <read bytes>][--read-timeout-sec <read timeout sec>]
```

- `<file_path>`: Required. The path to the file send as a hex data.
- `--port`: Optional. The serial port to use. Defaults to /dev/ttyUSB1.
- `--baud`: Optional. The baud rate to use. Defaults to 9600.
- `--data-bits`: Optional. The number of data bits to send. Defaults to 8.
- `--read-bytes`: Optional. The number of reading bytes after sending. Defaults to 4.
- `--read-timeout-sec`: Optional. The number of seconds of timeout for reading serial port after sending. Defaults to 1.

Example:

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
