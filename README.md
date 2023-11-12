# Assignment3

## Summary

This app is a Go-based application designed to average the image by calculating average color of its squares of given size.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.20+ installed on your machine.

## Dependencies

This project uses the following dependency:

- **[Gioui](https://github.com/gioui/gio)** - For visualization of image through GUI.

## Setting up

Go is going to install missing dependencies when you build the project. Navigate to the project's folder and build it with the following command.

    $ go build cmd/main.go

It will create a build file in project's foder. For windows it's going to be main.exe. Now you can run program through the terminal.

    $ main.exe filename.jpg 5 S
    Saved result in result.jpg file.

## Usage

Running: main.exe filename square_size, mode

- `filename`: name of the image file to be processed.
- `square_size`: Must be an integer.
- `mode`: S | M (Here s stands for singlethreaded and M for multithreaded.)
