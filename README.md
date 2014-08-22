golearnbenchmarks
=================

Benchmarking framework for machine learning libraries, and benchmark tests for the golearn library in particular.

Installation
============

`go get github.com/amitkgupta/golearnbenchmarks`

Usage
=====

This library will provide a framework for writing benchmark tests for machine learning libraries written in Go.  Instructions on using it for this purpose are forthcoming.

This library will also contain concrete benchmark tests for algorithms implemented in the [golearn](https://github.com/sjwhitworth/golearn) project.  Currently it works against a fork of that project, so the instructions below ask you to check out the fork of the `golearn` project.  If you have problems installing `golearn`, see its repo's README for a pointer to installation instructions.  Here are instructions for this project:

```
go get github.com/sjwhitworth/golearn/...
cd </path/to/golearn>
git remote add amit git@github.com:amitkgupta/golearn.git
git fetch amit
git checkout amit/develop

go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
go install github.com/onsi/ginkgo/ginkgo

cd </path/to/golearnbenchmarks>
ginkgo -r
```
