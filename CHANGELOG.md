# Changelog

All notable changes to this project will be documented in this file.
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [v1.2.0] - 2024-05-16

### Added

- LSQueue struct representing a Limited Size queue. 
- Sliceable and Limitable interfaces
- `arrnode` to `core` which is a struct helping to implement clustered containers
- Usage examples under the `examples` directory

### Updated

- NewQueue function which returns a new Queue pointer doesn't take a parameter anymore. This is a BREAKING CHANGE
- Optimized access time and write time performance of limited size queues (LSQueue) in comparison to previous versions.
- Optimized access time and write time performance of regular queues. Check Readme.MD for comparison.
- Optimized access times and write time performance of Stack in comparison to previous versions
- README.md with the benchmark results of this versio vs pre v1.2.0 numbers
- The ordering of updates in CHANGELOG.md

## [v1.1.1] - 2024-05-10

### Updated

- Fifo and Lifo to generic, to actually be compatible with the structs allowing them to implement the interfaces

## [v1.1.0] - 2024-05-10

### Added

- Sliceable interface
- SafeStack and SafeQueue thread safe versions of Stack and Queue respectively
- Benchmarks that can be run in the `lists` folder by running `go test -bench=.`
- Added benchmark results, and usage information to README.md

### Updated

- README.md

## [v1.0.2] - 2024-05-09

### Added

- Countable interface
- Implemented Countable interface in Stack and Queue
- CHANGELOG.md