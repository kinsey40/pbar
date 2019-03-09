# Pbar - A terminal progress bar for Go

[![Build Status](https://travis-ci.com/kinsey40/pbar.svg?branch=master)](https://travis-ci.com/kinsey40/pbar.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kinsey40/pbar)](https://goreportcard.com/report/github.com/kinsey40/pbar)
[![Coverage Status](https://coveralls.io/repos/github/kinsey40/pbar/badge.svg?branch=master)](https://coveralls.io/github/kinsey40/pbar?branch=master)
[![GoDoc](https://godoc.org/github.com/kinsey40/pbar?status.svg)](https://godoc.org/github.com/kinsey40/pbar)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Welcome to Pbar! A simple, easy-to-use, flexible terminal progress bar for the Go/Golang programming language! 

## Requirements
Pbar is tested to work on Go v1.10+, previous versions are not guaranteed to be compatible. 

## Installation
The Pbar repository can be installed via the standard Go package installation process:

```bash
$ go get github.com/kinsey40/pbar
```

## Usage
The file examples/example.go from the projects root directory highlights how the progress bar can be created in a variety of different circumstances. To create a progress bar from an array (as an example) the following is done:

```go
import "pbar"

x := []int{1, 2, 3}
p, err := pbar.Pbar(x)
if err != nil {
    panic(err)
}

p.Initialize()
for _, v := range x {
    // Do something...
    p.Update()
}
```

Generally, the object is first created via the function ```Pbar```. This can be altered as necessary (e.g. setting description for the progress bar). The pbar object must then be ```Initialized``` immediately before the for-loop and the ```Updates``` performed AFTER each iteration of the for loop. 

Hence, the Update function must be at the bottom of the for-loop. 

## Known Issues
Currently, there are two main issues relating to pbar. Firstly, pbar objects do not render correctly when called in seperate threads.  
Secondly, pbar has not been checked to work correctly on Windows OS; this may present problems due to pbars reliant on line 
ending functionality. 

* Progress bars in separate threads
* Windows OS

## Contributing
All Contributions to improving this project are welcome! Please examine the Contributing file for instructions on how to contribute. 

#### Authors
* Nicholas Kinsey (kinsey40)

## Feedback
All feedback regarding the quality, structure and maintainability of this code-base are welcome! If you discover an issue, or want an additional feature then please raise an issue.  
