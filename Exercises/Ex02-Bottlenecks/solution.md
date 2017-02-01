Mutex and Channel basics
========================

Atomic operation
----------------
 - "Indivisible" or "un-interruptable" operation
 - Appears (to the rest of the system) to occur instantaneously
 
<!-- -->
 - What we need:
   - Some variable that can signify the ownership/usage of "resources" (eg shared data structures)
   - Some variable that can be incremented (to show we are taking the resource) and decremented (to show we are done) "atomically" (without the i++ read-modify-write problem)
   - => We need atomic operations *in hardware*
     - (How this is done (assembly, etc) is not important for this course)

Semaphore
---------

 - An integer flag with atomic increment/decrement
 - The flag value is always greater than or equal to `0`
 
<!-- -->
 - Supports two operations:
   - Decrement (or "try to decrement"):
     - aka `wait()`, P, "Prolaag", Probeer te verlagen (dutch for "try to decrement")
   - Increment
     - aka `signal()`, V, Verhogen
    
<!-- -->
 - Trying to decrement when the flag is `0` will block the current thread
 - That thread will be awoken when *someone else* increments the flag

Mutex
-----

 - A binary semaphore (can only be `0` or `1`)
   - ie. unavailable / available
   - ie. locked / unlocked
 - Has ownership, unlike semaphores
   - Only the thread that called `wait()` can call `signal()`
     - ie. only the one that called `lock()` can call `unlock()`
   - (With ownership we can also provide priority inheritance, but this is a later topic in the course)
 - Use mutexes when you want to convey exclusive ownership.
   - From a code quality perspective: Don't use semaphores when you mean a mutex (even though "they work fine")

Critical section
----------------

 - Any piece of code that
   - Must be protected from concurrent execution
     - ie. should only have one process executing it at a time
   - Therefore requires mutual exclusion
 - Commonly implemented using:
   - Mutexes, if your OS supports it
   - Disabling all interrupts, on microcontrollers running without an OS
 


