# Golang generic list types - Stack and Queue v1.3.0

[![Go Reference](https://pkg.go.dev/badge/github.com/tanerius/lists.svg)](https://pkg.go.dev/github.com/tanerius/lists)

**Lists** is a simple O(1) implementation of a stack, a queue and a limited size queue in golang. 
The reson for this is that Golang as of yet does not have a standard implementation of a generic queue and stack.  
  
**Stack**, **Queue** and **LSQueue** and NOT thread safe.  
**Stack** and **Queue** have their thread safe counterparts, **SafeStack** and **SafeQueue** respectively.  

## Queue and Limited Size Queue

A queue is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence. By convention, the end of the sequence at which elements are added is called the back, tail, or rear of the queue, and the end at which elements are removed is called the head or front of the queue, analogously to the words used when people line up to wait for goods or services.  

A **Limited Size Queue** is simply a queue with a limited size and once its size is reached, each time a new element is Enqueued to the back one element is also Dequeued from the front of the list. This guarantees a maximum size.
  
A queue implements the Fifo interface.

## Stack 

A stack is an abstract data type that serves as a collection of elements with two main operations:

 - Push, which adds an element to the collection, and  
 - Pop, which removes the most recently added element.  
  
Additionally, a peek operation can, without modifying the stack, return the value of the last element added. The name stack is an analogy to a set of physical items stacked one atop another, such as a stack of plates.  
  
A stack implements the Fifo interface.

## Usage 
Below are examples of usage:  
  
```go
// Create a regular queue of integers
queue := NewQueue[int]()
// enqueue elements
queue.Enqueue(2)
queue.Enqueue(3)
queue.Enqueue(4)
// dequeue and print
fmt.Println(queue.Dequeue())
fmt.Println(queue.Dequeue())
fmt.Println(queue.Dequeue())
fmt.Println(queue.Dequeue())
/*
Output:
    2 <nil>
    3 <nil>
    4 <nil>
    0 empty list
*/

// Create a fixed size queue of 3 integers
fixedQueue := NewLSQueue[int](3)
// enqueue elements
fixedQueue.Enqueue(5)
fixedQueue.Enqueue(6)
fixedQueue.Enqueue(7)
fixedQueue.Enqueue(8)
fixedQueue.Enqueue(9)
// dequeue and print
fmt.Println(fixedQueue.Dequeue())
fmt.Println(fixedQueue.Dequeue())
fmt.Println(fixedQueue.Dequeue())
fmt.Println(fixedQueue.Dequeue())

/*
Output:
    7 <nil>
    8 <nil>
    9 <nil>
    0 empty list
*/

// Create a stack of integers
stack := NewStack[int]()
// push elements to the stack
stack.Push(2)
stack.Push(3)
stack.Push(4)
// dequeue and print
fmt.Println(stack.Pop())
fmt.Println(stack.Pop())
fmt.Println(stack.Pop())
fmt.Println(stack.Pop())
/*
Output:
    4 <nil>
    3 <nil>
    2 <nil>
    0 empty list
*/

```

## Benchmarks

The following are results from running a benchamark test with `go test -bench=.`

### Benchmark of v1.3.0

```
goos: linux
goos: linux
goarch: amd64
pkg: github.com/tanerius/lists/lists
cpu: 13th Gen Intel(R) Core(TM) i7-13700HX
BenchmarkRegularQueueEnqueue-24                 467376658                2.300 ns/op
BenchmarkRegularQueueDequeue-24                 1000000000               0.8321 ns/op
BenchmarkLimitedSizeQueueEnqueue-24             279022058                4.330 ns/op
BenchmarkLimitedSizeQueueDequeue-24             1000000000               0.2041 ns/op
BenchmarkRegularSafeQueueEnqueue-24             52909809                21.50 ns/op
BenchmarkRegularSafeQueueDequeue-24             56552139                20.49 ns/op
BenchmarkLimitedSizeSafeQueueEnqueue-24         58708641                19.93 ns/op
BenchmarkLimitedSizeSafeQueueDequeue-24         36732121                31.56 ns/op
BenchmarkSafeStackPush-24                       53247632                21.89 ns/op
BenchmarkSafeStackPop-24                        58177344                19.50 ns/op
BenchmarkStackPush-24                           539450361                2.269 ns/op
BenchmarkStackPop-24                            1000000000               0.7253 ns/op
PASS
ok      github.com/tanerius/lists/lists 26.910s
```

### Benchmark of v1.2.0

```
goos: linux
goarch: amd64
pkg: github.com/tanerius/lists/lists
cpu: 13th Gen Intel(R) Core(TM) i7-13700HX
BenchmarkRegularQueueEnqueue-24         501823215                2.501 ns/op
BenchmarkRegularQueueDequeue-24         1000000000               0.8231 ns/op
BenchmarkLimitedSizeQueueEnqueue-24     275466990                4.287 ns/op
BenchmarkLimitedSizeQueueDequeue-24     1000000000               0.2093 ns/op
BenchmarkRegularSafeQueueEnqueue-24     28228327                41.16 ns/op
BenchmarkRegularSafeQueueDequeue-24     54736119                22.42 ns/op
BenchmarkSafeStackPush-24               30719182                40.89 ns/op
BenchmarkSafeStackPop-24                52156732                22.62 ns/op
BenchmarkStackPush-24                   607518001                2.386 ns/op
BenchmarkStackPop-24                    1000000000               0.7070 ns/op
PASS
ok      github.com/tanerius/lists/lists 26.666s
```

### Benchmark of v1.1.1

```
goos: linux
goarch: amd64
pkg: github.com/tanerius/lists/lists
cpu: 13th Gen Intel(R) Core(TM) i7-13700HX
BenchmarkRegularQueueEnqueue-24                 43230771                24.15 ns/op
BenchmarkRegularQueueDequeue-24                 618508268                2.309 ns/op
BenchmarkLimitedSizeQueueEnqueue-24             68871850                16.00 ns/op
BenchmarkLimitedSizeQueueDequeue-24             1000000000               0.2043 ns/op
BenchmarkRegularSafeQueueEnqueue-24             28962522                39.99 ns/op
BenchmarkRegularSafeQueueDequeue-24             51305078                23.03 ns/op
BenchmarkLimitedSizeSafeQueueEnqueue-24         33558945                33.57 ns/op
BenchmarkLimitedSizeSafeQueueDequeue-24         31427727                33.64 ns/op
BenchmarkSafeStackPush-24                       28756356                40.16 ns/op
BenchmarkSafeStackPop-24                        55884932                21.52 ns/op
BenchmarkStackPush-24                           57936614                19.51 ns/op
BenchmarkStackPop-24                            903547963                1.585 ns/op
PASS
ok      github.com/tanerius/lists/lists 74.369s
```

## TODO

A list od things I plan to add:

- Implement Single and Double Linked Lists