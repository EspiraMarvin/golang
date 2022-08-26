## Channels
- Used to pass data between different goroutines, in a safe way to prevent race conditions and memory sharing problems
### Channel basics
- Create a channel with make command
 -----
   make(chan int)
 -----

- Send message into channel
--------
    ch <- val
--------

- Receive message from channel
--------
    val := <- ch
--------

- Can have multiple senders and receivers


### Restricting data flow
- Channels can be cast into send-only or receive only channels
- Send-only channel
----------
    chan <- int
----------
- Receive-only channel
----------
     <-chan int
----------

### Buffered Channels
- Channels block send side till receiver is available
- Block receiver side till message is available
- Can decouple sender and receiver with buffered channels
 ---------
    make(chan int, 50)
 --------- 
- Use buffered channels when send and receiver have assymmetric loading / if we can generate messages faster than we can receive them in channels
### Closing channels

### For-range loops with channels
- Use to monitor channel and process messages as they arrive
- Loop exits when channel is closed

### Select statements
- Allow goroutines to monitor several channels at once
- Blocks if all channels block
- If multiple channels receive value simultaneously, behaviour is undefined

