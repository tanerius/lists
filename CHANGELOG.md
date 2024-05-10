# Changelog

All notable changes to this project will be documented in this file.
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [1.0.2] - 2024-05-09

### Added

- Countable interface
- Implemented Countable interface in Stack and Queue
- CHANGELOG.md

## [1.1.0] - 2024-05-10

### Added

- Sliceable interface
- SafeStack and SafeQueue thread safe versions of Stack and Queue respectively
- Benchmarks that can be run in the `lists` folder by running `go test -bench=.`
- Added benchmark results, and usage information to README.md

### Updated

- README.md

## [1.1.1] - 2024-05-10

### Updated

- Fifo and Lifo to generic, to actually be compatible with the structs allowing them to implement the interfaces