EXERCISE 2:


A)
First some theory. What is:
 - An atomic operation?
	An operation is atomic if it appears to occur instantaneously
	to the rest of the system.
	Usually some features are present in hardware to make sure this
	appears instant in the program.
 - A semaphore?
	Semaphores are used to control the access to resources. Sort of
	like an abstract traffic light for different processess if they
	can access specific registers or not.
 - A mutex?
	"When I am having a big heated discussion at work, I use a
	rubber chicken which I keep in my desk for just such occasions.
	The person holding the chicken is the only person who is
	allowed to talk. If you don't hold the chicken you cannot speak
	. You can only indicate that you want the chicken and wait
	until you get it before you speak. Once you have finished
	speaking, you can hand the chicken back to the moderator who
	will hand it to the next person to speak. This ensures that
	people do not speak over each other, and also have their own
	space to talk.

	Replace Chicken with Mutex and person with thread and you
	basically have the concept of a mutex."
 - A critical section?
	In concurrent programming, concurrent accesses to shared resources can lead to unexpected or erroneous behavior, so parts of the program where the shared resource is accessed is protected. This protected section is the critical section or critical region. It cannot be executed by more than one process. Typically, the critical section accesses a shared resource, such as a data structure, a peripheral device, or a network connection, that would not operate correctly in the context of multiple concurrent accesses.

