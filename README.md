golearnbenchmarks
=================

Benchmarking framework for machine learning libraries, and benchmark tests for the golearn library in particular.

Installation
============

`go get github.com/amitkgupta/golearnbenchmarks`

Usage
=====

This library will provide a framework for writing benchmark tests for machine learning libraries written in Go.  Instructions on using it for this purpose are forthcoming.

This library will also contain concrete benchmark tests for algorithms implemented in the [golearn](https://github.com/sjwhitworth/golearn) project.  In order to run these benchmark tests, the following should suffice:

```
cd <path/to/this/repo>
go get -t ./...
ginkgo -r
```
