- Design Patterns and Principles That Support Large Scale Systems: https://medium.com/everything-full-stack/design-patterns-and-principles-that-support-large-scale-systems-f3c9adf89ad
- Software Architecture: I wish I had known about this earlier… - : https://medium.com/@olgamitroshyna/software-architecture-i-wish-i-had-known-about-this-earlier-4df43eae57db
- System Design Interview: https://github.com/alex-xu-system/bytebytego
- Kent Beck: https://tidyfirst.substack.com/

--

UUID:

"While preparing this specification the following 16 different implementations were analyzed for trends in total ID length, bit Layout, lexical formatting/encoding, timestamp type, timestamp format, timestamp accuracy, node format/components, collision handling and multi-timestamp tick generation sequencing." -- https://www.ietf.org/id/draft-peabody-dispatch-new-uuid-format-03.html

[UUID](https://www.cockroachlabs.com/blog/what-is-a-uuid/):
- [K-Sortable Unique IDentifier](https://github.com/segmentio/ksuid) 
  - https://news.ycombinator.com/item?id=24645786 
- [Distributed, Twitter Snowflake](https://developer.twitter.com/en/docs/twitter-ids) / Code: https://github.com/sony/sonyflake
- https://encore.dev/blog/go-1.18-generic-identifiers
- https://planetscale.com/blog/why-we-chose-nanoids-for-planetscales-api
- https://github.com/mplanchard/cuid-rust

[Sagas](https://microservices.io/patterns/data/saga.html)
- https://docs.microsoft.com/en-us/azure/architecture/reference-architectures/saga/saga
- https://docs.temporal.io/blog/reliable-crypto-transactions-at-coinbase/ / https://github.com/temporalio/samples-java/blob/main/src/main/java/io/temporal/samples/bookingsaga/TripBookingWorkflowImpl.java

CircuitBreaker: </>

Bulkhead: </>

RateLimiter: </>

Retry: </>

**State:**

1. Key - arbitrary user defined 
2. 

Napkin Math - https://github.com/sirupsen/napkin-math

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

--

