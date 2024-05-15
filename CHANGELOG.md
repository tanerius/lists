# Changelog

All notable changes to this project will be documented in this file.
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [v1.2.0] - 2024-05-15

### Added

- LSQueue struct representing a Limited Size queue. 
- Sliceable and Limitable interfaces

### Updated

- NewQueue function which returns a new Queue pointer doesn't take a parameter anymore. This is a BREAKING CHANGE
- Drastically improved access time and write time performance on limites size queues (LSQueue) in comparison to previous versions.
- README.md
- The ordering of updates in 

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