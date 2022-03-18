## Software Design

- Kent Beck: https://tidyfirst.substack.com/
- System Design Interview: https://github.com/alex-xu-system/bytebytego

### Recurring

[UUID](https://www.cockroachlabs.com/blog/what-is-a-uuid/):
- [K-Sortable Unique IDentifier](https://github.com/segmentio/ksuid) 

[Sagas](https://microservices.io/patterns/data/saga.html)
- https://docs.microsoft.com/en-us/azure/architecture/reference-architectures/saga/saga
- https://docs.temporal.io/blog/reliable-crypto-transactions-at-coinbase/ / https://github.com/temporalio/samples-java/blob/main/src/main/java/io/temporal/samples/bookingsaga/TripBookingWorkflowImpl.java

CircuitBreaker
- 

Bulkhead
-

RateLimiter
-

Retry
-

**State:**

1. Key - arbitrary user defined 
2. 

**RFC Standards:**

| Header | Description  |
| ------ | ------------ |
| X-RateLimit-Limit | how many request the client can do |
| X-RateLimit-Remaining | how many request remain to the client in the timewindow |
| X-RateLimit-Reset | how many seconds must pass before the rate limit resets |
| Retry-After | if the max has been reached, the millisecond the client must wait before perform new requests |

**Samples:**
1. https://github.com/vladimir-bukhtoyarov/bucket4j
2. https://github.com/sethvargo/go-limiter
