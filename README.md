# ticktock
A visualization of Lamports ME on a binary tree. 
**Credit: The following algorithm is based on the answer from the collaborated exam prep document; OMSCS Advanced Operating System Fall 2023 batch.**

## Lock request
### Format
```
{
  lock_id: int
  requester_id: int
  timestamp: int
}
```

### Sending
1. Create a lock request.
2. Send to all children and parent nodes.
3. Wait for implicit and explicit acknowledgements.

### Receiving
1. Add the lock request to a priority queue (prioritised on timestamp).
2. Send acknowledgement back to the sender.
3. Forward lock request to all children and parent nodes; except the sender; without modifying the timestamp.

## Ack
### Format
```
{
  lock_id: int
  requester_id: int
  acker_id: int
  timestamp: int
}
```

### Sending
1. Send as part of `step 2` on receiving a lock.

### Receiving
1. Check if the ack is for a lock that we own.
2. If yes:
    2.1 Update the acknowledgement status for the lock.
    3.1 If the lock request is acked by all healthy nodes then safely acquire the lock.
3. No:
    3.1 Forward the ack to all children and parent nodes; except the sender, without any modifications.

## Unlock
### Format
```
{
  lock_id: int
  requester_id: int
}
```

### Sending
1. Create an unlock request and send to all children and parent nodes.

### Receiving
1. Remove the lock from own queue
2. Forward the unlock request to all children and parent nodes; except the sender.
