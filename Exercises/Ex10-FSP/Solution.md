Exercise 10 : FSP
=======================


----------
**Safety/Deadlock Assignment**

The absence of deadlocks is a built-in safety criterion in LTSA. Create a system with a deadlock.

The standard deadlock with semaphores - As shown on the blackboard during the lectures:

    #+BEGIN_SRC C
    T1 = (t1wa -> t1wb -> t1sb -> t1sa -> T1).
    T2 = (t2wb -> t2wa -> t2sa -> t2sb -> T2).
    SA = (t1wa -> t1sa -> SA
       |  t2wa -> t2sa -> SA).
    SB = (t1wb -> t1sb -> SB
       |  t2wb -> t2sb -> SB).

    ||SYSTEM = (T1 || T2 || SA || SB).
    #+END_SRC

Does LTSA detect it?

Yes :-)

It is easy to see from the state model that we have a deadlock, because there will be a state in the diagram that there is no way to get out from. How can you detect a deadlock from the FSP model itself?

A very good and interesting question! There are no obvious, easy ways,
is it? We are left to "simulating" the system mentally, as with
traditional "staring at the code to see if it works".

**Dining Philosophers**

The code for a "dining philosophers" system can look like this if we
model forks with semaphores:

    phil1(){
      wait(fork1)
      wait(fork2)
      // eat
      signal(fork2)
      signal(fork1)
    }

*Assingment*

If you look closely at this it has the same structure as the "deadlock
with semaphores" over (and from the lectures).

*Assignment*

    #+BEGIN_SRC C
    Phil1 = (phil1StartsUsingFork1 -> phil1StartsUsingFork2 -> phil1StopsUsingFork2 -> phil1StopsUsingFork1 -> Phil1).
    Phil2 = (phil1StartsUsingFork2 -> phil1StartsUsingFork3 -> phil1StopsUsingFork3 -> phil1StopsUsingFork2 -> Phil2).
    Phil3 = (phil1StartsUsingFork3 -> phil1StartsUsingFork1 -> phil1StopsUsingFork1 -> phil1StopsUsingFork3 -> Phil3).
    Fork1 = (phil1StartsUsingFork1 -> phil1StopsUsingFork1 -> Fork1
           | phil3StartsUsingFork1 -> phil3StopsUsingFork1 -> Fork1).
    Fork2 = (phil2StartsUsingFork2 -> phil2StopsUsingFork2 -> Fork2
           | phil1StartsUsingFork2 -> phil1StopsUsingFork2 -> Fork2).
    Fork3 = (phil3StartsUsingFork3 -> phil3StopsUsingFork3 -> Fork3
           | phil2StartsUsingFork3 -> phil2StopsUsingFork3 -> Fork3).

    ||SYSTEM = (Phil1 || Phil2 || Phil3 || Fork1 || Fork2 || Fork3).
    #+END_SRC C

*Optional Assignment*

I find this model among the ltsa example programs. Very short, and
uses 'forall' and 'menu' (ignore this) that was not mentioned in the
lectures. It also uses the trick of indexing the philosopher /in the
sharing set/ of the fork.

    #+BEGIN_SRC C
    /** Concurrency: State Models and Java Programs
     *             Jeff Magee and Jeff Kramer
     */

    PHIL = (sitdown->right.get->left.get
              ->eat->left.put->right.put
              ->arise->PHIL).

    FORK = (get -> put -> FORK).

    ||DINERS(N=5)=
       forall [i:0..N-1]
       (phil[i]:PHIL
       ||{phil[i].left,phil[((i-1)+N)%N].right}::FORK).

    menu RUN = {phil[0..4].{sitdown,eat}}
    #+END_SRC C

*Optional Assignment*

See the ltsa exampleprogram; chapter 6, "DeadlockFreePhilosophers"

*Optional Assignment*

See also the little book of semaphores.

Communication
-------------------------
*Assignment*

Synchronous communication is an event - with two participants. You
should not differ between sending and receiving; The "transfer of
data", "the communication" - or even "the channel" are metaphores for
the event.

Our system becomes:

    T1 = (ch1->ch2->T1).
    T2 = (ch1->ch2->T2).

How many states will there be in the transition diagram of this system?

2

*Assignment*

A bounded buffer would be our starting point. Here we do need to
differ between send and receive ('put' and 'get'.)









