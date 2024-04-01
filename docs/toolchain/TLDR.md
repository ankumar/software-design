## TLDR

References & Evolving Context, Exponential growth of Nodes/Edges - UI/UX Web, Mobile, Data, APIs, Desktop, Console, Voice, AR/VR, ...

1. **Connected:**
- Apple, Stripe, Meta <-> +[AWS](https://aws.amazon.com/solutions/case-studies) 
- Walmart <-> +[Azure](https://azure.microsoft.com/en-us/resources/customer-stories/), +[GCP](https://cloud.google.com/customers)
- Shopify, Twitter <-> +GCP
- ...
- ...
- IoT <-> +\[AWS, Azure, GCP\]
- SpaceX/STARLINK <-> +GCP

2. **Infrastructure:**
- Stateless & Stateful Apps
- Synchronous, Serving System 
- Asynchronous, Event Driven System
- Global Load Balancing, DNS-based & SDN Anycast
- Caching & CDN
- CAP Theorem
- Edge & Cloud Computing
- Availability, Regions & Zones

**1. As a service:**
![](https://github.com/ankumar/Architecture/blob/main/images/Cumulative%20CAPEX.jpg)

Source: https://www.platformonomics.com/2021/02/follow-the-capex-cloud-table-stakes-2020-retrospective/

> hyperscalers (Amazon, Google and Microsoft) have moved far outside their initial niche of standard IT infrastructure services. In addition to providing services for developers to reduce time-to-market, they have built specialist services targeted at the major technology trends such as blockchain, 5G, machine learning, artificial intelligence and digital identity. Moreover, they are adding many industry-focused solutions.

* **Cloud Native DB's**
  * **Distributed SQL** - GCP Spanner,  YugaByte DB (PostgreSQL), MariaDB SkySQL, CockroachDB, ...
  * **Key/Value** - Migrating Amazon retail services to NoSQL, **DynamoDB** ; re:Invent & Twitch Talks from Rick Houlihan 

* **IAM** 
  * https://www.okta.com/ / https://auth0.com/  
  * https://www.authorizon.com/ 
  * https://authzed.com/ 
  
* **Kubernetes**
  * [Google Cloud Run - FAQ](https://github.com/ahmetb/cloud-run-faq)
  * https://www.slim.ai/

**2. Building:**

- The bottom half of the OSI reference model?
- Compute, Network, Storage
- Connected with a kernel
Example Infrastructure:
- (Compute) Server, Operating System, Container Runtime, Process
- (Storage) 💾Disk, Database, Memory
- (Network) Load Balancer, Firewall, Network Switch, Router

-- https://github.com/kris-nova/public-speaking/blob/main/presentations/2021-infrastructure-as-software/5-what-is.zomg

[@scale](https://atscaleconference.com/)

* [LogDevice](https://logdevice.io/)
* [Mcrouter (pronounced mc router)](https://github.com/facebook/mcrouter)
* https://db.cs.cmu.edu/archived-events/

---

**1. Publications:**

* [Deployment Archetypes for Cloud Applications](https://arxiv.org/abs/2105.00560)
* [Considerations for Mandating Open Interfaces](https://www.internetsociety.org/resources/doc/2020/white-paper-considerations-for-mandating-open-interfaces/)
* [100 Year Study on AI released its 2021 Report](https://ai100.stanford.edu/)
* [Our Journey towards Data-Centric AI: A Retrospective](https://ai.stanford.edu/blog/data-centric-ai-retrospective/)
* [Architecting the Future of Software Engineering: A National Agenda for Software Engineering Research & Development](https://resources.sei.cmu.edu/library/asset-view.cfm?assetID=741193)
* [Why AI is Harder Than We Think](https://arxiv.org/abs/2104.12871)
* [Cloud Infrastructure Services: An analysis of potentially anti-competitive practices](https://www.fairsoftwarestudy.com/)
* [From Cloud Computing to Sky Computing](https://sigops.org/s/conferences/hotos/2021/papers/hotos21-s02-stoica.pdf)
* [Why Google Stores Billions of Lines of Code in a Single Repository](https://cacm.acm.org/magazines/2016/7/204032-why-google-stores-billions-of-lines-of-code-in-a-single-repository/fulltext)
* [The Datacenter as a Computer](http://bnrg.eecs.berkeley.edu/~randy/Courses/CS294.F09/wharehousesizedcomputers.pdf)
* [Fail at Scale](https://queue.acm.org/detail.cfm?id=2839461)
* [Zero downtime deployments at Facebook](https://dl.acm.org/doi/abs/10.1145/3387514.3405885)
* [Pitfalls of Anycast](https://www.usenix.org/sites/default/files/conference/protected-files/srecon17emea_slides_murali_suriar.pdf)
* [Influence of heavy-tailed distributions on load balancing](http://www.cs.cmu.edu/~harchol/ISCA15show.pdf)
* [Akamai DNS: Providing Authoritative Answers to the World’s Queries](https://groups.cs.umass.edu/ramesh/wp-content/uploads/sites/3/2020/07/sigcomm2020-final289.pdf)
* [Monarch: Google’s Planet-Scale In-Memory Time Series Database](http://www.vldb.org/pvldb/vol13/p3181-adams.pdf)
* [Read-Write Quorum Systems Made Practical](https://mwhittaker.github.io/publications/quoracle.html)
* [Zanzibar: Google’s Consistent, Global Authorization System](https://news.ycombinator.com/item?id=26980254)
* [Colossus - Google’s Storage Foundation](https://www.infoq.com/news/2021/04/google-colossus/)
* [Designing a Serverless Cloud-Native Database-as-a-Service Based on Apache Cassandra™](https://www.datastax.com/sites/default/files/content/whitepaper/files/2021-06/Astra%20Serverless%20Whitepaper_0.pdf)
* [Maglev: A Fast and Reliable Software Network Load Balancer](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/44824.pdf)
* [Sizeless: Predicting the Optimal Size of Serverless Functions](https://se.informatik.uni-wuerzburg.de/software-engineering-group/staff/simon-eismann/?tx_extbibsonomycsl_publicationlist[action]=download&tx_extbibsonomycsl_publicationlist[controller]=Document&tx_extbibsonomycsl_publicationlist[fileName]=MIDDLEWARE2021___Sizeless.pdf&tx_extbibsonomycsl_publicationlist[intraHash]=ececf0bdbba4e0517d68cb7dba7423e6&tx_extbibsonomycsl_publicationlist[userName]=simon.eismann&cHash=dd84cf187c8303334c8f4fde69bed2a2)

---

**2. Tech Talks:**

* [Quarantine Database Tech Talks (2020)](https://www.youtube.com/playlist?list=PLSE8ODhjZXjagqlf1NxuBQwaMkrHXi-iz)
* [New Directions in Cloud Programming](https://www.youtube.com/watch?v=FeRg-7Sr1L8)

**3. More Artifacts:**

* [An Introduction to Distributed Systems](https://github.com/aphyr/distsys-class)
* [the morning paper](https://blog.acolyer.org/)
* [Marc's Blog](https://brooker.co.za/blog/)
* [Client-side load balancing technique](https://blog.twitter.com/engineering/en_us/topics/infrastructure/2019/daperture-load-balancer.html)
* [DNS based GSLB](https://dropbox.tech/infrastructure/intelligent-dns-based-load-balancing-at-dropbox)
* [Shard Manager](https://engineering.fb.com/production-engineering/scaling-services-with-shard-manager/)
* [Consensus Algorithms at Scale](https://www.planetscale.com/blog/blog-series-consensus-algorithms-at-scale-1)
* [Steve Yegge’s Google Platforms Rant](https://gist.github.com/chitchcock/1281611)
* [Excel is still going strong, just became a Turing-complete programming language!](https://www.microsoft.com/en-us/research/blog/lambda-the-ultimatae-excel-worksheet-function/)
* [The Distributed Operating System Void](https://nivenly.com/lib/2021-04-02-operating-system-interface/)
* [COSI - The Common Operating System Interface](https://docs.google.com/document/d/1OuwTSsSsIPefDViheK-nzaF9xSOg1Mn62mwR2FmGPu8/edit#heading=h.1grxkjflkt7d)
* [Micah Lerner](https://www.micahlerner.com/)
* [Metadata](http://muratbuffalo.blogspot.com/)
* [TLA+](http://lamport.azurewebsites.net/tla/tla.html)
* [The P programming language](https://github.com/p-org/P)
* [Bento - Interactive Notebook that Empowers Development Collaboration & Best Practices]( https://developers.facebook.com/blog/post/2021/09/20/eli5-bento-interactive-notebook-empowers-development-collaboration-best-practices/)

-- 

**Large-scale distributed infrastructure:**

Apps:
1. Amazon.com Prime Days - https://aws.amazon.com/blogs/aws/prime-day-2021-two-chart-topping-days/
 
1. **Authenticate & Authorize**

* AWS Identity - [500 million API calls EVERY SECOND](https://aws.amazon.com/blogs/aws/happy-10th-birthday-aws-identity-and-access-management/)
* Azure AD - [Manages more than 1.2 billion identities and processes over 8 billion authentications every day](https://azure.microsoft.com/en-us/services/active-directory/)

2. **"The Log"**

* [Event-Driven Architectures - The Queue vs The Log](https://jack-vanlightly.com/blog/2018/5/20/event-driven-architectures-the-queue-vs-the-log)

* Kafka:
  * Benchmarking - https://www.confluent.io/blog/kafka-fastest-messaging-system/
  * [Kafka Improvement Proposals](https://cwiki.apache.org/confluence/display/KAFKA/Kafka+Improvement+Proposals)
    * Tiered Storage
    * [Replace Zookeeper - A Raft Protocol for the Metadata Quorum](https://cwiki.apache.org/confluence/display/KAFKA/KIP-595%3A+A+Raft+Protocol+for+the+Metadata+Quorum)
  * [Vectorized](https://vectorized.io/) **Redpanda** - Kafka® API compatible, C++ Implementation
  * Write Ahead Log - https://wepay.github.io/waltz/docs/introduction
* Queue:
  * https://engineering.fb.com/2021/02/22/production-engineering/foqs-scaling-a-distributed-priority-queue/

3. **Purpose-built Databases**

Managing state - one of the hardest problem in distributed systems
* Key/Value - Cassandra; ScyllaDB, Apache Cassandra in C++
* Distributed SQL 
 
 * AI/ML:
    * Anyscale / Ray - https://www.youtube.com/watch?v=dn8hu2sgRWU
 
* Database:
  * [The PlanetScale Workflow](https://docs.planetscale.com/concepts/planetscale-workflow)
  * [DynamoDB - Single-Table Modeling](https://amazondynamodbofficehrs.splashthat.com/)
  * [Prisma - "application models"](https://www.prisma.io/docs/concepts/overview/what-is-prisma/data-modeling)
  
* Metadata: 
  * [LinkedIn GMA](https://github.com/linkedin/datahub-gma/blob/master/docs/how/metadata-modelling.md)
  * [OpenMetadata](https://docs.open-metadata.org/openmetadata/schemas)

* Things & More:
  * [The first open core platform for building and running distributed, continuous intelligence applications](https://www.swim.ai/products/)
  * [Azure opendigitaltwins](https://github.com/Azure/opendigitaltwins-building)
  * [AWS IoT Things Graph](https://docs.aws.amazon.com/thingsgraph/latest/ug/iot-tg-whatis.html)
  * [Google Digital Buildings](https://google.github.io/digitalbuildings/)
  * [Smart Data Models](https://smartdatamodels.org/)

4. **Data Warehouse / Data Lake**

* Storage - [Cloud, S3](https://www.allthingsdistributed.com/2021/03/happy-15th-birthday-amazon-s3.html) / Open Source https://delta.io/
* Data Processing - </>

5. **Networking**

* Kubernetes - https://cloud.google.com/blog/products/containers-kubernetes/new-gke-gateway-controller-implements-kubernetes-gateway-api / https://smi-spec.io/#ecosystem

6. **Telemetry**
* eBPF - https://github.com/cloudflare/ebpf_exporter / https://engineering.fb.com/2021/04/27/developer-tools/reverse-debugging/ / https://cilium.io/
  * [How eBPF will solve Service Mesh - Goodbye Sidecars](https://isovalent.com/blog/post/2021-12-08-ebpf-servicemesh)

7. **Linux / Unix**
* Evolution of the Unix System Architecture ("Besides historical details, it also has a number of interesting lessons for software architecture evolution") - https://ieeexplore.ieee.org/stamp/stamp.jsp?tp=&arnumber=8704965


> Software is eating the world. But progress in software technology itself largely stalled around 1996.
> [The Great Software Stagnation](https://alarmingdevelopment.org/?p=1475)
>
  - [Structure Eats Strategy](https://janbosch.com/blog/index.php/2017/11/25/structure-eats-strategy/)
  - [Coordination Headwind](https://komoroske.com/slime-mold/)
  - [The next “next” generation](https://swardley.medium.com/how-organisations-are-changing-cf80f3e2300)
  - ["Open & Public"](https://en.wikipedia.org/wiki/Open_collaboration), Few examples:
    - [Sourcegraph](https://about.sourcegraph.com/about/)
    - [GitLab](https://about.gitlab.com/company/strategy/#why-is-this-page-public)

* Product Strategy - Transparent Roadmap
  * https://github.com/aws?q=roadmap
  * https://github.com/github/roadmap

* Journey - Celebrate & Share
  * [A Conversation with Werner Vogels](https://queue.acm.org/detail.cfm?id=1142065)
  * [A Second Conversation with Werner Vogels](https://queue.acm.org/detail.cfm?id=3434573)
  * [The Amazon Builders' Library](https://aws.amazon.com/builders-library/)
   
* Platform Strategy - "Leverage"
  * [Open Software Design](https://notes.ceilfors.com/Open_software_design.html) -> https://upmo.com/dev/
    * [Architecture](https://www.mediawiki.org/wiki/Wikimedia_Architecture_Team#The_practice) 
    * [Tech Radar](https://opensource.zalando.com/tech-radar/)
  * [Open Source Culture](https://opensource.zalando.com/)
![](https://github.com/ankumar/Architecture/blob/main/images/open%20source.jpeg)
-- https://twitter.com/mjasay 

![](https://github.com/ankumar/Architecture/blob/main/images/open%20source%20%26%20large%20company.png)
-- [Open Source & Community / Kelsey Hightower, Developer & Open Source Advocate, Google Cloud](https://www.youtube.com/watch?v=jiaLsxjBeOQ) 

https://cacm.acm.org/magazines/2021/7/253459-why-computing-students-should-contribute-to-open-source-software-projects/fulltext

> Developers are the creative workforce who can solve critical business problems and create hit products for customers — not just "code monkeys" who grind through rote tasks. Companies that bring software developers in as partners are winning.
> -- https://twitter.com/jeffiel/status/1418682712559398916

Scaling engineering teams [Circa 1994, Backyard at Yahoo!](https://github.com/ankumar/Architecture/blob/main/Patterns/History.md#platform--tools--yahoo), Now Backstage open source project Incubated by Spotify, Adopted as Runway@American Airlines, 

* https://engineering.atspotify.com/2021/03/16/happy-birthday-backstage-spotifys-biggest-open-source-project-grows-up-fast/
  * https://engineering.atspotify.com/category/backstage/
  * https://roadie.io/blog/developer-portals-are-a-superpower/

Likely termed **"Internal Platforms"** @Hundreds of organizations ...

* Culture
<img src="https://github.com/ankumar/Architecture/blob/main/images/Hack%20The%20Culture.jpg" width="800">

> "culture is not a static ‘thing’ but something which everyone is constantly creating, affirming and expressing" -- Mary Douglas

**Culture is not a set of beliefs, a set of actions**
 
> Companies & Culture: What You Do Is Who You Are - a16z editor in chief Sonal Chokshi interviews a16z co-founder Ben Horowitz -- author of the book What You Do Is > Who You Are -- on whether companies and people can change; how the very thing that is your strength can also be your weakness; how startups evolve from pirates to > the navy; actions vs words and values; and more.
 
https://a16z.simplecast.com/episodes/what-you-do-is-who-you-are-companies-culture-book-computer-history-museum-6JfGIJJs

**Culture is "Code"**

What kind of development organization are you a part of: pathological, bureaucratic, or generative? https://qualitysafety.bmj.com/content/qhc/13/suppl_2/ii22.full.pdf

> I am reading a book that was recommended by Kumar, Anil called "[Accelerate - Building and Scaling High Performing Technology Organizations by Nicole Forsgren, Jez Humble, and Gene Kim](https://www.amazon.com/Accelerate-Software-Performing-Technology-Organizations/dp/1942788339/)".  I am about 100 pages in and finally starting to understand the vision around our "destination" culture and honestly the "different cultures" that exist among our divisions.  
> On pg30 the book describes 3 types of organizational cultures.
> 1. **Pathological** (power oriented) organizations are characterized by large amounts of fear and threat.  People often hoard information or withhold it for political reasons, or distort it to make themselves look better.
> 2. **Bureaucratic** (rule-oriented) organizations protect departments.  Those in the department want to maintain their "turf", insist on their own rules, and generally do things by the book - their book.
> 3. **Generative** (performance-oriented) organizations focus on the mission.  How do we accomplish our goal?  Everything is subordinated to good performance, to doing what we are supposed to do.
> Moving forward in the book I'm learning about changes that need to be made around testing, version control, automation, CICD, communication, etc...  I am starting to be able to tie ideas that are nicely articulated in the book to statements I've heard in meetings by folks like Kumar, Anil, Raghavendra, Vijay, and Charlton, Paul.  Wow.  This is so cool...  It's like a giant light bulb turned on this week!
> I believe our destination culture is Generative.  I also feel that "we all kinda want it".  So, I'm getting excited about what the future holds...
> I recommend that all Tech managers & Architects get a copy of this book and read through it.
> Cheers!  ~https://www.linkedin.com/in/timothycdahl/
> 

* Non-Functional

> "We start by looking at which **“-ilities”** are most important to architects. A software architect is responsible for the cross-cutting concerns and making sure that individual components of a large system can work together seamlessly to meet overall objectives. In 2021, four areas we feel architects are concerned with are designing for resilience, designing for observability, designing for portability, and designing for sustainability." 

-- https://www.infoq.com/articles/architecture-trends-2021/

![Reliability](https://github.com/ankumar/Architecture/blob/main/images/Reliability.png)

-- [From book Implementing Service Level Objectives](https://www.amazon.com/Implementing-Service-Level-Objectives-Practical/dp/1492076813)

* Functional

> "if building a software-intensive system, these are the forces we must weigh" https://twitter.com/ruthmalan/status/989206552044294145
![Software-Intensive](https://github.com/ankumar/Architecture/blob/main/images/software-intensive.jpeg)

* "On The Criteria for Decomposing Software" - Dr. David Parnas 1972 Paper:
  * [On the criteria to be used in decomposing systems into modules](https://blog.acolyer.org/2016/09/05/on-the-criteria-to-be-used-in-decomposing-systems-into-modules/)
  * [45m podcast of selected readings from this paper](https://podcasts.apple.com/us/podcast/on-the-criteria-to-be-used-in-decomposing-systems/id1364166414?i=1000527256187)
* [Revisiting Information Hiding](https://link.springer.com/chapter/10.1007%2F978-3-642-22655-7_8)
* [The theory of graceful extensibility](https://link.springer.com/article/10.1007/s10669-018-9708-3)

* Table stakes

![](https://github.com/ankumar/Architecture/blob/main/images/oxide.computer.png)

-- [importance of lowest-level details in secure systems and the pressing need for open firmware down to the boot ROMs!](https://twitter.com/bcantrill/status/1388181932068929538)

- [Privacy](https://dbpedia.org/page/General_Data_Protection_Regulation), Collaboration, Security & Ethics
- Applications - Interoperability with [Open APIs](https://www.youtube.com/watch?v=LzMp6uQbmns)
  - [Amazons](https://apievangelist.com/2012/01/12/the-secret-to-amazons-success-internal-apis/), [OpenID](https://openid.net/), [Open Banking](https://en.wikipedia.org/wiki/Open_banking), [Open Banking](https://blog.stoplight.io/open-banking-guide), [Healthcare](https://www.hl7.org/fhir/), IoT, ... / [Cloud](https://github.com/kcp-dev/kcp), ...

* SocioTechnical
  * https://twitter.com/jessmartin/status/1383121887585312769
  * https://esilva.net/sociotechnical
  * https://github.com/matthewskelton/sociotechnical-architecture
  * [DevOps, Observability, and the need to tear down organizational boundaries](https://medium.com/lightstephq/devops-observability-and-the-need-to-tear-down-organizational-boundaries-f5d25755ff3a)
  * [System Thinking: Approaches and Methodologies](https://www.burgehugheswalsh.co.uk/Uploaded/1/Documents/Soft-Systems-Methodology.pdf)
  * https://itrevolution.com/the-idealcast-podcast/ / https://itrevolution.com/westrums-organizational-model-in-tech-orgs/ /  
  * [Staff-plus](https://staffeng.com/guides)
  * [Sample of Career ladders](https://github.com/sdras/career-ladders)
  * [Shades of Conway's Law](https://thinkinglabs.io/articles/2021/05/07/shades-of-conways-law.html)

* Domain-Driven	Design (TripleD :stuck_out_tongue_winking_eye: https://en.wikipedia.org/wiki/Diners,_Drive-Ins_and_Dives)
> Build a shared understanding, Use ubiquitous language, Express in domain models, Separate by bounded contexts
  * https://verraes.net/2021/09/what-is-domain-driven-design-ddd/
  * https://github.com/ddd-crew/ddd-starter-modelling-process
  * https://www.domainlanguage.com/wp-content/uploads/2016/05/DDD_Reference_2015-03.pdf
  * https://twitter.com/ntcoding/status/1381249293315682307 

* Applications
  * https://martinfowler.com/tags/application%20architecture.html
  * https://www.postman.com/postman/workspace/postman-open-technologies-specifications/overview
  * https://codersociety.com/blog/articles/twelve-factor-app-methodology

* Decisions - Build In-house vs Buy/ use off-the-shelf vs outsource
> "Getting people to record decisions takes commitment, practice...and support in doing so. This isn't magic. It's technical leadership." https://www.agilealliance.org/resources/experience-reports/distribute-design-authority-with-architecture-decision-records/

  * https://www.susannekaiser.net/conferences/slides/yow-adaptive-systems-2020.pdf 
  * https://leadingedgeforum.com/insights/mapping-1-a-practical-guide-to-avoiding-outsourcing-problems/

* Architecture-Qualities - https://github.com/mtnygard/architecture-qualities
  * Why microservices?
    * https://martinfowler.com/tags/microservices.html
    * https://chrisrichardson.net/post/microservices/2020/05/21/why-microservices-part-4.html
  
* Characterizing Architecturally Significant Requirements
  * https://ulir.ul.ie/bitstream/handle/10344/3061/Chen_2013_characterising.pdf;jsessionid=47B707698DAEE64A10746731ADEE08F9?sequence=2
  * https://twitter.com/ersiner/status/1376438100801503237

* Heuristics
  * [Software Design Heuristics](https://github.com/stanislaw/SoftwareDesignHeuristics)
  * [Hints and Principles for Computer System Design](https://www.microsoft.com/en-us/research/uploads/prod/2019/09/Hints-137-short.pdf)
  * [Domain-Driven Design Heuristics](https://www.dddheuristics.com/)

* OpenAPI https://www.openapis.org/

https://github.com/OAI/OpenAPI-Specification
> The OpenAPI Specification (OAS) defines a standard, programming language-agnostic interface description for HTTP APIs, which allows both humans and computers to discover and understand the capabilities of a service without requiring access to source code, additional documentation, or inspection of network traffic. When properly defined via OpenAPI, a consumer can understand and interact with the remote service with a minimal amount of implementation logic. Similar to what interface descriptions have done for lower-level programming, the OpenAPI Specification removes guesswork in calling a service.

* Remote development, collaboration, & productivity, SPACE Framework - https://www.infoq.com/news/2021/03/space-developer-productivity/
  * https://github.blog/2021-05-25-octoverse-spotlight-good-day-project/
  * https://speakerdeck.com/charity/cd
  * https://future.a16z.com/on-workplace-productivity/

* Open Source

  * https://github.com/LappleApple/zalando-howto-open-source
  * Sustainable ecosystems 
    * Case Study Presto vs Facebook - https://trino.io/blog/2020/12/27/announcing-trino.html
  * List of system quality attributes - https://en.wikipedia.org/wiki/List_of_system_quality_attributes
  * Emerging Reference Architectural Patterns - https://github.com/wso2/reference-architecture
  * https://boards.greenhouse.io/wikimedia/jobs/2976623
    * https://www.mediawiki.org/wiki/Wikimedia_Architecture_Team#The_practice
    * https://techblog.wikimedia.org/2020/10/29/wikipedia-as-a-castle-in-the-wilderness-modernization-in-the-dynamic-world-of-the-internet/ 
    * https://wikimediafoundation.org/about/values/
  * Strategy
    * https://about.gitlab.com/company/strategy/
    * https://about.sourcegraph.com/about/
  * Design
    * https://upmo.com/dev/ 
  * License
    * https://martin.kleppmann.com/2021/04/14/goodbye-gpl.html
  * Operate
    * https://www.operate-first.cloud/
  * Experience / UI - https://github.com/pinterest/querybook
  * Looker ( Open Source) - https://github.com/lightdash/lightdash
  * Metadata (From LinkedIn Team) - https://datahubproject.io/
  * Metadata (From Uber Team) - https://open-metadata.org/
  * https://github.com/authzed
  * https://github.com/ory 
  * https://www.keycloak.org/ 
  * https://www.authorizon.com/
  * https://github.com/porter-dev/dex
  * https://github.com/uc-cdis/fence   
  * https://www.haproxy.com/blog/haproxy-enterprise-offers-saml-based-single-sign-on/
  * https://github.com/yahoo/k8s-athenz-istio-auth / https://www.athenz.io/comparison.html#general
  * https://www.athenz.io/
  * "extracted our common patterns around OIDC and JWT into a shared Go library that is now used between HashiCorp products. Really nice way to integrate auth and includes wonderful testing helpers. Much easier than direct low-level lib integration." - https://github.com/hashicorp/cap
  * Kubernetes
    * https://kubernetes.io/docs/reference/access-authn-authz/authentication/  
    * https://backstage.io/docs/auth/ 
    * https://clutch.sh/docs/advanced/auth 
  * OpenTelemetry
    * [Well Instrumented](https://redmonk.com/jgovernor/2021/06/08/aws-well-instrumented-on-instrumentation-observability-and-ecosystems/)
    * DataDog 
      * https://opstrace.com/blog/introducing-datadog-compatible-http-api
      * https://github.com/SigNoz/signoz

* Platform 
  * https://leanpub.com/platformstrategy
  * https://stories.platformdesigntoolkit.com/design-apis-for-disobedience-7894f930e2cc

* Cloud     
    
    * Cloud Deployment Model - https://www.sciencedirect.com/topics/computer-science/cloud-deployment-model
      * The NIST Definition of Cloud Computing - https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-145.pdf

    * Separation of Providers & Applications
      * https://www.pulumi.com/cloud-engineering/
 
    * Cloud Carbon Footprint - https://www.cloudcarbonfootprint.org/
    
    * New Directions in Cloud Programming - https://www.youtube.com/watch?v=FeRg-7Sr1L8 
        * https://www.arnnet.com.au/article/687460/inside-decline-heroku/ 
    
    * https://www.container-solutions.com/wtf-is-cloud-native/resources
  
    * "Well Architected" - Ideally need set of Abstractions, Crisp boundaries, Clear separation of concerns, and a balanced set of responsibilities...

      1.[AWS](https://aws.amazon.com/architecture/) - 5 pillars
        * Operational Excellence
        * Security
        * Reliability
        * Performance Efficiency
        * Cost Optimization

      2.[Azure](https://azure.microsoft.com/en-us/blog/introducing-the-microsoft-azure-wellarchitected-framework/) - Five pillars
        * Cost Optimization	- Managing costs to maximize the value delivered.
        * Operational Excellence	- Operations processes that keep a system running in production.
        * Performance Efficiency	- The ability of a system to adapt to changes in load.
        * Reliability	- The ability of a system to recover from failures and continue to function.
        * Security	- Protecting applications and data from threats.

      3.[GCP](https://cloud.google.com/architecture/framework) - 4 principles
        * Operational excellence
        * Security, privacy, and compliance
        * Reliability
        * Performance and cost optimization

      4.[IBM](https://www.ibm.com/cloud/architecture/architectures)

    * On GitHub
        * Systems:
          * [Systems and failure reading list](https://github.com/lorin/systems-reading)
          * [Resilience engineering](https://github.com/lorin/resilience-engineering/blob/master/intro.md) 
       * Awesome:
          * [CTO](https://github.com/kuchin/awesome-cto)
          * [Leading-and-Managing](https://github.com/LappleApple/awesome-leading-and-managing)
          * [Bug-bounty](https://github.com/djadmin/awesome-bug-bounty)
          * [DDD](https://github.com/heynickc/awesome-ddd)
          * [Safety-critical](https://github.com/stanislaw/awesome-safety-critical)
          * [Cold-showers](https://github.com/hwayne/awesome-cold-showers)
          * [Cloud-native](https://github.com/senthilrch/awesome-cloud-native)
          * [Kubernetes](https://github.com/senthilrch/awesome-kubernetes)
          * [Serverless](https://github.com/senthilrch/awesome-serverless)
          * [Scalability](https://github.com/binhnguyennus/awesome-scalability)

* Compliance & Paperwork
  * SOC2 - </>
  * GDPR/CCPA/ - </>
  * HIPAA - </>

* Complexity & Cognitive Work 
  * https://www.adaptivecapacitylabs.com/blog/2019/01/30/human-cognitive-work-happens-above-the-line/

* Amazon shareholder letters, 1997 to 2021 - https://ir.aboutamazon.com/annual-reports-proxies-and-shareholder-letters/default.aspx
  * https://www.oreilly.com/radar/checking-jeff-bezoss-math/

* ACM - www.acm.org/acm-focus

* InfoQ - https://www.infoq.com/

- **Cloud Computing**
  - [The NIST Definition of Cloud Computing](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-145.pdf)
> **Software as a Service (SaaS)** - The capability provided to the consumer is to use the provider’s applications running on a cloud infrastructure. The applications are accessible from various client devices through either a thin client interface, such as a web browser (e.g., web-based email), or a program interface. The consumer does not manage or control the underlying cloud infrastructure including network, servers, operating systems, storage, or even individual application capabilities, with the possible exception of limited userspecific application configuration settings 
  - **SaaS** - [What is SaaS?](https://www.oracle.com/applications/what-is-saas/)
    - SaaS Community - https://github.com/saas-community/saas-links
    
- **Identity** - [When the internet was built it, wasn’t built with essential piece of plumbing, that’s identity](https://openid.net/2016/09/27/the-foundation-of-internet-identity/)
  - [Okta Identity 101](https://www.okta.com/identity-101/)
    - Single Sign-on (SSO) Experience - https://auth0.com/intro-to-iam/what-is-single-sign-on-sso/
  - [Open Standards](https://www.forgerock.com/platform/common-services/open-standards)
  - [Implementing OpenID Connect and OAuth 2.0 – Tips from the Trenches - Dominick Baier](https://www.youtube.com/watch?v=QpkVnB-N20c)
    - OpenID Connect - https://openid.net/connect/ (https://en.wikipedia.org/wiki/List_of_OAuth_providers)
      - Finance - https://fapi.openid.net/
    - [OAuth 2 Simplified](https://aaronparecki.com/oauth-2-simplified/)
    - [OAuth 2.0 Flows (Simplified)](https://dev.to/hem/oauth-2-0-flows-explained-in-gifs-2o7a)
    - [OpenID Connect](https://developer.yahoo.com/oauth2/guide/openid_connect/#key-concepts)
    - [OpenID Connect - id_token, what they are, how they work](https://www.youtube.com/watch?v=3u1el6f6mdE)
    - [Difference Between OAuth, OpenID Connect, and SAML](https://www.okta.com/identity-101/whats-the-difference-between-oauth-openid-connect-and-saml/)

- **AuthZ**, Access control models like ACL, RBAC, ABAC - https://www.okta.com/identity-101/authentication-vs-authorization/
  - [All About AuthZ](https://speakerdeck.com/dschenkelman/all-about-authz) | [Video](https://www.youtube.com/channel/UCZuLNqvV4oUMVyNq70mFF0g)
  - [Azure AD Authentication vs. Authorization](https://docs.microsoft.com/en-us/azure/active-directory/develop/authentication-vs-authorization)
  - https://learn.sandcastle.cloud/intro/authz-and-sandcastle#what-is-authorization-authz
  - [Authorization FAQ](https://docs.authzed.com/concepts/authz) 
  - [Zanzibar-based implementations](https://authzed.com/blog/zanzibar-implementations/)
  
- **Infrastructure Facets**
  - Data - [Warehouse, Lake & Lakehouse](https://databricks.com/blog/2021/08/30/frequently-asked-questions-about-the-data-lakehouse.html)
  - Serverless - [“managed services that scale to zero”](https://redmonk.com/rstephens/2018/12/14/serverless-more-than-just-functions/) 
    - Serverless - [Well-Architected Serverless](https://www.youtube.com/watch?v=NFtZSvywRew&t=3071s)
  - As Software - https://www.pulumi.com/docs/intro/concepts/
  - As Platform - **Kubernetes Resource Model**
    - https://twitter.com/kelseyhightower/status/935252923721793536?lang=en
    - https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources
    - https://github.com/kcp-dev/kcp
      - https://github.com/askmeegs/learn-krm
  - As "Edge" - 

* Glossary - https://github.com/cncf/glossary
* Open Glossary of Edge Computing - https://github.com/State-of-the-Edge/glossary 

--

- **Radar**

A technology radar is an opinionated guide to a set of emerging technologies. The popular format originated at [Thoughtworks](https://www.thoughtworks.com/radar) and has been adopted by dozens of companies including Spotify, Zalando, AOE, Porsche and Intuit. 

The key idea is to place solutions at one of three levels:

**ADOPT** — We feel strongly that the industry should be adopting these items. We use them when appropriate in our projects. 

Technologies we have high confidence in to serve our purpose, also in large scale. Technologies with a usage culture in production environment, low risk and recommended to be widely used.

**TRIAL** —  Worth pursuing. It’s important to understand how to build up this capability. Enterprises can try this technology on a project that can handle the risk. 

Technologies that we have seen work with success in project work to solve a real problem; first serious usage experience that confirm benefits and can uncover limitations. TRIAL technologies are slightly more risky; some engineers in our organization walked this path and will share knowledge and experiences.

**ASSESS** — Worth exploring with the goal of understanding how it will affect your enterprise. 

Technologies that are promising and have clear potential value-add for us; technologies worth to invest some research and prototyping efforts in to see if it has impact. ASSESS technologies have higher risks; they are often brand new and highly unproven in our organization. You will find some engineers that have knowledge in the technology and promote it, you may even find teams that have started a prototyping effort.

**HOLD** — Proceed with caution. 

Technologies not recommended to be used for new projects. Technologies that we think are not (yet) worth to (further) invest in. HOLD technologies should not be used for new projects, but usually can be continued for existing projects.

TRIAL

1. SRE - https://www.datadoghq.com/ | 
2. Collaboration - https://miro.com/ | https://www.smartsheet.com/ | https://www.productplan.com/ | 
3. Design - https://zeplin.io/ | https://www.adobe.com/ | 
4. Security - https://www.blackducksoftware.com/ | https://neuvector.com/ | 

ASSESS

1. [Platform Ecosystem](https://github.com/gneisstech/bedrock/issues/33)
2. https://www.microsoft.com/en-us/microsoft-viva/overview | 
3. https://www.notion.so/ | https://www.atlassian.com/ | 
4. https://www.figma.com/ | 
5. https://lightstep.com/ | https://www.honeycomb.io/ | 
6. https://octopus.com/ | https://www.jetbrains.com/space/ | https://github.com/features | https://about.gitlab.com/releases/gitlab-com/#container-scanning-engine-clair

----

* Public Technology Radars: 
  * https://radar.cncf.io/ - CNCF is home to over 70 of the fastest growing open source infrastructure projects, including Kubernetes, Prometheus and Envoy, as part of the Linux Foundation.
    * https://www.oicheryl.com/2021/03/23/10-predictions-for-cloud-native-in-2021-the-devops-conference/ 
  * https://upmo.com/dev/radar/
  * https://github.com/zalando/tech-radar
  * https://storage.googleapis.com/wf-techradar-prod/index.html
  
- APIs
 * [History of APIs](https://history.apievangelist.com/) 
 * API 360 - **“your data model is not your object model is not your resource model is not your message model”**
    * [API design 101](https://cloud.google.com/blog/products/api-management/api-design-101-links-our-most-popular-posts)
    * [Level 3 REST](https://level3.rest/)
    * [API Design Systems](https://apidesign.systems/)
  * Postman Public workspaces, Ex: https://www.postman.com/salesforce-developers 
    * Products API Blueprint - https://www.postman.com/api-evangelist/workspace/products-api-blueprint/overview 
    * AWS Workspace - https://www.postman.com/api-evangelist/workspace/amazon-web-services-aws/overview 
    * SOAP APIs - https://www.postman.com/cs-demo/workspace/public-soap-apis/overview 
    * REST APIs - https://www.postman.com/cs-demo/workspace/public-rest-apis/overview 
    * GraphQL APIs - https://www.postman.com/cs-demo/workspace/public-graphql-apis/overview 
 
  * RBAC - https://puppet.com/docs/pe/2021.2/rbac-api.html
 
- Auth

-- Identity Platforms
  * [AWS Cognito](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation.html)
  * [Azure AD](https://docs.microsoft.com/en-us/azure/active-directory/)
  * [Google Cloud](https://cloud.google.com/identity-platform)
  * [Okta / Auth0](https://auth0.com/intro-to-iam/)

- Data

-- IoT
Smart Data Models - Digital Twins represent **connected things** like tools, cars, machines, sensors, and other web-enabled things in the cloud in a reusable and abstracted way, an explorer https://github.com/Azure-Samples/digital-twins-explorer

* Personal
  * [IoT devices, cloud platforms, Machine Learning & Office API's](https://www.youtube.com/watch?v=u5oTz1e5qqE) 

* Autonomous
  * https://www.youtube.com/watch?v=E-JHrt73R-4 (https://github.com/carla-simulator/carla)

* Buildings
  * [Smart Building - 6 real buildings, Brick Schema Open-source ontology for building assets, subsystems and data](https://brickschema.org/get-started) 
  * [Smart Building - Bosch Singapore campus: smart building concept turned reality](https://www.youtube.com/watch?v=n5aTeQkBBew)
  * [Smart Building - Google Using Cloud IoT and AI ](https://www.youtube.com/watch?v=Zz6jkLYkzSQ)
  * [Smart Retail - How Walmart Leverages IoT to Keep Your Ice Cream Frozen](https://corporate.walmart.com/newsroom/2021/01/14/how-walmart-leverages-iot-to-keep-your-ice-cream-frozen)

* Robotics
  * https://aws.amazon.com/blogs/opensource/aws-deepracer-is-now-open-source-and-ready-to-hit-the-road-with-ros-2/

* Planet
  * https://www.overstory.ai/ | https://www.amazon.science/latest-news/creating-a-digital-twin-of-the-earth-with-computer-vision  

* https://brickschema.org/
* https://project-haystack.org/
* https://rise.cs.berkeley.edu/projects/xbos/     
* https://github.com/Azure/opendigitaltwins-dtdl
* https://docs.aws.amazon.com/thingsgraph/latest/ug/iot-tg-whatis.html

* Buildings
  * Google - https://cloud.google.com/iot-core 
    * https://github.com/google/digitalbuildings 
    * https://linkedbuildingdata.net/ldac2020/files/presentations/LDAC2020_IND6_CharbelKaed.pdf
  * Azure - https://docs.microsoft.com/en-us/azure/digital-twins/
    * https://techcommunity.microsoft.com/t5/internet-of-things/realestatecore-a-smart-building-ontology-for-digital-twins-is/ba-p/1914794
    * https://azure.microsoft.com/en-us/blog/azure-digital-twins-now-generally-available-create-iot-solutions-that-model-the-real-world/ 
    * https://github.com/Azure?q=digital&type=&language=&sort=
 
* Carnegie Mellon University (CMU) Software Engineering Institute (SEI) key emerging technologies 2020 https://resources.sei.cmu.edu/library/asset-view.cfm?assetid=650822
1. Advanced computing
2. The Smarter edge
3. Digital twins
4. AI
5. Extended reality
6. Data privacy, trust, ethics

* Curated Digital Twin Knowledge Repository - https://digitaltwin.io/

* https://www.cloudflight.io/expert-views/learnings-from-the-digital-twins-data-architecture-of-tesla-40097/
  
* https://anylog.co/ | http://dbmsmusings.blogspot.com/2019/12/its-time-to-rethink-how-we-share-data.html

![Linux Foundation](https://www.edgexfoundry.org/cmsfiles/image/Platform/edgex-focused-at-the-iot-edge-2.jpg)
* [Open Source EdgeX Foundry](https://docs.edgexfoundry.org/1.3/)
  * 5G - https://en.wikipedia.org/wiki/5G
  * MQTT - https://mqtt.org/
  * BACNET - http://www.bacnet.org/
    * https://www.automatedbuildings.com/news/apr20/articles/S4group/200324105101s4.html
  * BLE - https://en.wikipedia.org/wiki/Bluetooth_Low_Energy
  * ZIGBEE - https://zigbeealliance.org/
  * MODBUS - https://modbus.org/ 
* Sensors for Web - https://web.dev/generic-sensor/

- Cloud & Edge

* Nvidia is building a giant virtual ‘metaverse’ of the world, with ‘digital twins’ of cars, cities, and people - https://finance.yahoo.com/news/nvidia-building-giant-virtual-metaverse-115216046.html

> "distinction between physical and digital will disappear" 
> https://www.admcloudtech.com/2021/04/06/amazon-cto-aws-will-erase-the-line-between-physical-and-digital/ 

- https://www.cncf.io/blog/

- Linux Foundation - [LF Edge](https://www.linuxfoundation.org/en/press-release/the-linux-foundation-launches-new-lf-edge-to-establish-a-unified-open-source-framework-for-the-edge/)
  * https://stateoftheedge.com/
    * https://wiki.lfedge.org/display/LE/LF+Edge
    * https://github.com/lf-edge/edge-containers
    * https://www.edgexfoundry.org/software/platform/

  * [LF AI & Data](https://lfaidata.foundation/)
  
- Kube
  * https://docs.google.com/document/d/1We-pRDV9LDFo-vd9DURCPC5-Bum2FvjHUGZ1tacGmk8/edit#

- Telemetry

* What is OpenTelemetry? - https://opentelemetry.lightstep.com/
* Specification - https://github.com/open-telemetry/opentelemetry-specification
* Microservices - https://github.com/lightstep/hipster-shop#see-telemetry-data-in-lightstep

- SRE

- [Public Post-mortems](https://github.com/danluu/post-mortems)

- [Template](https://sre.google/sre-book/example-postmortem/)

- [Status Pages](https://marquez.co/statusfy?ref=aceforth)

- https://cloud.google.com/blog/products/operations/on-the-road-to-sre-with-cloud-operations-sandbox

- Developer Experience Products

* **Application centric** approach- https://kubevela.io/
* **Open Software Design** - [Upmo Dev](https://upmo.com/dev/)
-> Suggest reading [full thread from Upmo CTO](https://twitter.com/ceilfors/status/1313783840029192193)
* **Data**, a lot of going on - https://www.amundsen.io/ ; https://dagster.io/ ; </>
* Backstage, An open platform for building **developer portals** - https://engineering.atspotify.com/2020/09/24/cloud-native-computing-foundation-accepts-backstage-as-a-sandbox-project/ 
Backstage, needs TypeScript
  * 1 yr. Anniversary https://engineering.atspotify.com/2021/03/16/happy-birthday-backstage-spotifys-biggest-open-source-project-grows-up-fast/
  * Adopting -> https://backstage.io/blog/2021/05/20/adopting-backstage
  * Adoption -> https://github.com/backstage/backstage/blob/master/ADOPTERS.md
  * Plugins -> https://backstage.io/plugins
  * SaaS Platform -> https://roadie.io/

* GitHub - https://githubuniverse.com/ & https://github.com/features
* GitLab - https://about.gitlab.com/company/#about-us
* Sourcegraph, universal **code search** - https://about.sourcegraph.com/about/

---

- Signal:

- https://backstage.io/ 
- https://kubevela.io/#/ 

- Noise:

iOS & Android for IaC Frameworks - https://devops.com/iac-frameworks-vendor-specific-or-multi-cloud/ 
- https://aws.amazon.com/proton/  
- https://github.com/Azure/bicep 

Microsoft Bedrock
1. https://github.com/microsoft/bedrock
2. https://github.com/microsoft/spektate
3. https://github.com/microsoft/bedrock-cli
- https://github.com/magicaks

- https://clutch.sh/ 
  * https://www.moment.dev/, Any relation to Open Source [Clutch?](https://clutch.sh/docs/about/what-is-clutch?)

- https://hub.tekton.dev/ 
- https://gitlab.comwork.io/oss/tekton-task-images 
 
**Security Policy** 

**Environments**

Dev Like https://dev.to/jgoux/preview-environments-per-pull-request-using-aws-cdk-and-github-actions-bfi | https://www.gitpod.io/ ?

**Database** 

1. https://schemahero.io/ 

**Cloud IAM / AuthN / AuthZ / RBAC / ABAC** 

https://www.youtube.com/watch?v=K65e5QRQ1tc 

**1. Applications** 

API - I love this definition, **API is a promise**  from video series **What is an API** 

1. It’s a Promise! - https://www.youtube.com/watch?v=gFu7CQKmBeQ 
2. The Fundamentals - https://www.youtube.com/watch?v=_nAfNxhzJy0 
3. APIs & Products - https://www.youtube.com/watch?v=VNwAYyfQ8Bw 
4. It’s a Language! - https://www.youtube.com/watch?v=wDYeU2T4v2o 

* Protocol Buffers - https://buf.build/
* Docs & Changes - https://useoptic.com/
* Viz - https://www.akitasoftware.com/

**GraphQL:** 

- A Thread about coupling concerns around GraphQL implementations https://twitter.com/samnewman/status/1346749556583780352 
- https://the-guild.dev/
- https://thegraph.com/ 

**Protobuf:** 

https://docs.buf.build/ 
 
**Database:** 

- https://www.liquibase.org/ 
- https://medium.com/google-cloud/liquibase-spanner-extension-6df18d3829f3 

**Cassandra:** 

- https://stargate.io/ 

**Containers:** 

[Buildpacks](https://buildpacks.io/)
- [Azure](https://docs.microsoft.com/en-us/cli/azure/acr/pack?view=azure-cli-latest)
- [GCP](https://github.com/GoogleCloudPlatform/buildpacks)
- [Comparing containerization methods: Buildpacks, Jib, and Dockerfile](https://cloud.google.com/blog/topics/developers-practitioners/comparing-containerization-methods-buildpacks-jib-and-dockerfile)

**Registry**

With [Helm 3](https://helm.sh/blog/helm-3-released/) the storage backend is [changing](https://helm.sh/blog/helm-3-preview-pt3/) to container registries. 

- [Artifact Hub](https://artifacthub.io/) is the official Helm charts repository.
- [Azure container registry](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-helm-repos) can be used as a host for Helm chart repositories.
- GitHub has added beta support for storing and managing public and private Helm Charts in the GitHub Container Registry to their roadmap for [Q2 2021](https://github.com/github/roadmap/projects/1).
  * https://gitlab.com/gitlab-org/charts/gitlab/-/blob/master/doc/architecture/architecture.md
  * https://charts.gitlab.io/ | https://docs.gitlab.com/charts/  
 
Security
- https://github.com/aquasecurity/trivy#table-of-contents 
- https://github.com/aquasecurity/starboard 

**2. Kubernetes** 

- https://www.cncf.io/sandbox-projects/
- https://www.platformer.com/products/platformer-console/overview
- https://github.com/kubernetes-sigs  
- https://github.com/Azure/azure-service-operator 
- https://techcommunity.microsoft.com/t5/azure-arc/azure-arc-enabled-kubernetes-is-now-generally-available/ba-p/2174859  
- https://github.com/aws-controllers-k8s/community  
- https://cloud.google.com/config-connector/docs/overview 
- https://twitter.com/embano1/status/1346437160031969282 
- https://cdk8s.io/ 
- https://about.gitlab.com/direction/package/helm_chart_registry/ 
- https://danielmangum.com/posts/crossplane-infrastructure-llvm/
- https://github.com/crossplane/crossplane/tree/master/design 
- https://www.jetstack.io/ 
- https://www.powertools.dev/ 
- https://tilt.dev/ 
- https://github.com/fluxcd/helm-operator-get-started 
- https://www.gitpod.io/docs/ 
- https://www.forbes.com/sites/janakirammsv/2021/02/20/microsoft-aims-to-simplify-cloud-native-development-with-the-dapr-project 
- https://github.com/deislabs/akri
- https://keptn.sh/ 
- https://www.shipa.io/
- https://draft.sh/

**3. DevOps / SRE** 

- https://statusfy.co/ 
- https://optimyze.cloud/about/  
> Full-system lightweight continuous profiling for Linux Kernel, C/C++, Rust, Golang, Python, JVM, PHP (with Ruby and Node planned for the future) 
- https://cuelang.org/ 
- https://github.com/GoogleCloudPlatform/fourkeys 
- https://gtoolkit.com/
 
**4. Low-Code & No-Code** 

- https://cloudmaker.ai/ 
- https://tufan.io/  
- https://steampipe.io/

**5. Views** 

* [Explore assets and their relationships across your technical infrastructure.](https://github.com/lyft/cartography)

Going back to a movie clip [**The Matrix**](https://www.youtube.com/watch?v=aVLexf_dyCM) (prepared for education & learning purposes)
> Neo: This...this isn't real?
> Morpheus: What is real. How do you define real? If you're talking about what you can feel, what you can smell, what you can taste and see, then real is simply electrical signals interpreted by your brain. This is the world that you know. The world as it was at the end of the twentieth century. It exists now only as part of a neural-interactive simulation that we call the Matrix. You've been living in a dream world, Neo. This is the world as it exists today....

**Dev**
  - IDE - https://blog.jetbrains.com/blog/2021/03/11/projector-is-out/
  
  - Auth 
      - https://casbin.org/en/editor
      - https://learn.sandcastle.cloud/
  
  - Cloud - https://notes.ceilfors.com/Cloud_sandbox.html

---

**Data**
  - [Metadata](https://sandbox.open-metadata.org/) | [datahub](https://demo.datahubproject.io/)

  - [Financial Data](https://setu.co/data/account-aggregator) 

-- https://yourstory.com/2021/08/setu-launches-sandbox-account-aggregator-api-fintech-developers-startups/amp
  
  - [DBpedia](https://www.dbpedia.org/)
  
  - [Wikidata](https://www.wikidata.org/wiki/Wikidata:Project_chat#Wikidata:Query_Builder_has_been_deployed)

---

**Things, Physical & Digital**
  - Energy - https://energy-models.com/what-is-energy-modeling-building-simulation
  
  - Building - https://brickschema.org/get-started/

---

