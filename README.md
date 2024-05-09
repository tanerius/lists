# Golang generic list types - Stack and Queue v1.0.0

**Lists** is a simple O(1) implementation of a stack, a queue and a limited size queue in golang. 
The reson for this is that Golang as of yet does not have a standard implementation of a generic queue and stack.

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