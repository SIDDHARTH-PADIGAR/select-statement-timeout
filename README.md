##  Use Case: Timeout While Waiting for a Channel

Imagine you're making a request to a microservice, doing a DB read, or waiting on some goroutine to send you data — but you don’t want to wait forever. If it doesn’t respond in 2 seconds, you move on.

---

##  What's Going On:

* `ch` is a channel that's supposed to get a message from a goroutine.
* The goroutine **sleeps for 3 seconds** and then sends something.
* But the `select` block is **only willing to wait for 2 seconds** (`time.After(2 * time.Second)` returns a channel that sends a signal after 2 seconds).
* So the timeout case hits first, and the program prints:
  **`Timeout: no response in 2 seconds`**

---

##  Why It Worked

* **`select` listens to multiple channels simultaneously.**
* **Whichever case is ready first gets executed.**
* `time.After(2 * time.Second)` is a built-in channel that sends a signal after 2s.
* The **faster goroutine** sends a value to `ch` in 1s, so it wins the race against the timeout.

> If both goroutines took longer than 2s, the timeout case would have executed instead.

### Output
```
Received: Response from fast goroutine
```
---

##  When To Use This Pattern

###  Timeouts for Slow Operations

* Waiting for API responses
* Reading from DB/cache/microservice with a fallback
* Making sure your system doesn't hang forever on blocked channels

###  Racing Between Fastest Response

* Fan-out: multiple workers doing the same task; take the first one done
* Load balancing: fastest responder wins
* Fall-back systems: try primary, fallback to secondary, all within a timeout
