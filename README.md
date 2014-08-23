golearnbenchmarks
=================

Benchmarking framework for machine learning libraries, and benchmark tests for the [`golearn`](https://github.com/sjwhitworth/golearn) library in particular.

Installation
============

Simple (ignore the warnings about no buildable source files): 

`go get -d github.com/amitkgupta/golearnbenchmarks`

Usage
=====

This library will provide a framework for writing benchmark tests for machine learning libraries written in Go.  Instructions on using it for this purpose are forthcoming.

This library will also contain concrete benchmark tests for algorithms implemented in the `golearn` project, and can be run as follows:

```
# For now, check out my fork of 'golearn' (won't be necessary in the future):
go get -d github.com/sjwhitworth/golearn/
cd </path/to/golearn>
git remote add amit https://github.com/amitkgupta/golearn.git
git fetch amit
git checkout amit/develop

# Run the benchmark tests:
cd </path/to/golearnbenchmarks>
go get -t ./...
go install github.com/onsi/ginkgo/ginkgo
ginkgo -r
```
