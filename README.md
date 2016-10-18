#SWIM

This is an attempt to implement SWIM in GO.

## What is SWIM

I recommend you to read the paper in References sector first, i don't 
think i can explain better than the paper itself.

Here is the TL;DR version:

In distributed computing, a failure detector is an application or a 
subsystem that is responsible for detection of node failures or crashes 
in a distributed system.

A Traditional way to implement it is using heartbeat protocol: **which 
either impose network loads that grow quadratically with group size, or 
compromise response times or false positive frequency w.r.t. detecting 
process crashes**

The new system, called SWIM, provides a membership substrate that:
```
(1) imposes a constant message load per group member;
(2) detects a process failure in an (expected) constant time
at some non-faulty process in the group;
(3) provides a deterministic bound (as a function of group
size) on the local time that a non-faulty process takes to detect
failure of another process;
(4) propagates membership updates, including information
about failures, in infection-style (also gossip-style or
epidemic-style); the dissemination latency in the
group grows slowly (logarithmically) with the number of
members;
(5) provides a mechanism to reduce the rate of false positives
by “suspecting” a process before “declaring” it as
failed within the group 
```


SWIM has two components:
(1) a Failure Detector Component, that detects failures of members, and
(2) a Dissemination Component, that disseminates information
about members that have recently either joined or left
the group, or failed.

### Failure Detector
Given we have a cluster of n nodes. Each node has information about m 
nodes in the cluster (m <= n).
 
Every T' time unit (which called a protocol period), at node Mi

1) Increase the period `pr`
2) Select a random node in m nodes, called it Mj and ping it `ping(Mi, Mj , pr)`. 
Wait for the worst-case message round-trip for an `ack(Mi, Mj , pr)`. 

    - If the ack message come back, that means Mj is still alive.
    - If not, move to step 3
3) Select `k` nodes in m nodes, ask them to ping `Mj` `ping-req(Mi, Mj , pr)`

    - If one of them receive the ack message `ack(Mi, Mj , pr)`, the node is still alive
    - If no one receive the ack until the end of the period, declared Mj as failed.

4) To reduce the false positive in step 3, instead of mark `Mj` as failed
immediately, we will mark it as `suspected`. After a prespecified time-out, it
will be declared as `failed`. But if it response within the timeout, it will be
declared as `alive` again.


As any given time, at node `Mi`

On receipt of `ping-req(Mm, Mj , pr)` message (`Mj`!= `Mi`), send a `ping(Mi, Mj , Mm, pr)` message to `Mj`
On receipt of `ack(Mi, Mj , Mm, pr)` message from `Mj`, send an `ack(Mm, Mj , pr)` message to received to `Mm`
On receipt of `ping(Mm, Mi, Ml, pr)` message from `Mm`, reply with an `ack(Mm, Mi, Ml, pr)` message to `Mm`
On receipt of `ping(Mm, Mi, pr)` message from `Mm`, reply with an `ack(Mm, Mi, pr)` message to `Mm`

## Dissemination Component (cont.)

## References

1. [On scalable and efficient distributed failure detectors.](https://www.cs.cornell.edu/projects/quicksilver/public_pdfs/On%20Scalable.pdf)
2. [SWIM: Scalable Weakly-consistent Infection-style Process Group Membership Protocol](https://www.cs.cornell.edu/~asdas/research/dsn02-swim.pdf)
