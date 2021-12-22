# Open software design

> Open-source software has made it easier for software developers to study and learn programming by looking at real-world working software. But what about software design? Building architects (not software) study thousands of buildings in their training and career. Software developers, in contrast, only study a handful of other large software designs. This means we repeat mistakes that we would have been able to avoid, should we be able to learn from others.

-- https://notes.ceilfors.com/Open_software_design.html

## Sandbox

* [Cloud](https://notes.ceilfors.com/Cloud_sandbox.html)
* [Auth0](https://learn.sandcastle.cloud/)
* [Authzed](https://play.authzed.com/schema)
* [Metadata](https://sandbox.open-metadata.org/)
* [DataHub](https://demo.datahubproject.io/)
* [Energy](https://energy-models.com/what-is-energy-modeling-building-simulation)
* [Building](https://brickschema.org/get-started/)

## SaaS

[SaaS, the Past, Present and Future - A Continuum...](https://www.linkedin.com/pulse/saas-past-present-future-continuum-dr-srikanth-sundararajan/)

### Build

* The Case for ‘Developer Experience’ -- https://future.a16z.com/the-case-for-developer-experience/
> "There's a lot of hype around developer productivity platforms..."
-- https://twitter.com/martinfowler/status/1387041315196702720

[SaaS Community](https://www.youtube.com/channel/UCZuLNqvV4oUMVyNq70mFF0g/about) 
- videos: https://www.youtube.com/channel/UCZuLNqvV4oUMVyNq70mFF0g
- podcast: https://anchor.fm/saas-for-developers

#### Perspectives

**Cycles**

> History only rhymes, and unlike productivity software, SaaS is not a single horizontal market. But there are some big clusters that seem ripe for consolidation, as well as across verticals. **ERP** in fact was an early example of functional consolidation, one that started in manufacturing and spread to other industries. The question for SaaS Universe™ inhabitants is are you a **consolidator or the consolidated?** -- https://www.platformonomics.com/2021/05/saas-sprawl-and-the-potential-for-market-consolidation/

[Is the modern analytics stack unbundling, or consolidating?](https://sisudata.com/blog/modern-analytics-stack)

**API's**

> "Over the past decade we have seen nearly every industry unlock innovation at scale through a single platform API that abstracts away the historical complexities inherent to that industry. **Twilio** for communications, **Stripe** for payments, **Plaid** for financial records, **Segment** for customer records, and now Mapped for IoT." -- https://blog.mapped.com/ 

**"Open"**

> “Open standards are the best way to make sure that everybody in the industry is involved in setting the direction of technology in our industry.” --https://venturebeat.com/2021/08/09/why-major-fintech-players-are-already-embracing-open-api-standards-vb-live/

> Open protocols backed by proprietary implementations is the future of managed services. -- https://twitter.com/kelseyhightower/status/1428411566307766272

> Circa 2008 on Open Source vs Open Standards [1] & lack of portability between SaaS [2]
> Which one is more important open source or open standards? According to Opera Software's CEO Jon S. von Tetzchner it's open standards. A lot of people seem to agree:
> "open standards ARE more important than the licensing of an individual piece of software. Who cares what license the software someone uses is, as long as its always possible to replace it, and freely compete, which a free and open standard ensures."
> I disagree and I'm starting to feel like a lone voice here, maybe I'm wrong? Well let's go through my reasoning. I happen to agree with Bob Suter that "the easy availability of software can accelerate the adoption of a standard" and that adoption makes something a standard or not. For me, the fastest way of achieving such adoption is through providing an operational means of implementing an open standard - which means open source. So in my view, if you want portability you will need **BOTH open standards AND open source. BOTH are important**.
> 1. https://blog.gardeviance.org/2008/05/data-portability-radioactive-agpl-and.html
> 2. https://blog.gardeviance.org/2008/03/saas-utility-and-clouds.html 
> 3. https://blog.gardeviance.org/2008/02/market-forces-part-iv-open-source-vs.html

**Boost**

> Our mission is to create a community-driven suite of extensible building blocks for Software-as-a-Service (SaaS) builders. Our goal is to foster an open environment for developing and sharing reusable code that accelerates the ability to deliver and operate multi-tenant SaaS solutions on AWS. -- AWS SaaS Boost, https://aws.amazon.com/partners/saas-boost/ 

> One slide which says it all. Before anyone goes "serverless is the future" ... no, "serverless is the norm" - it's just that so many are yet to realise this. Give it time, another five years or so -- https://twitter.com/swardley/status/1412001525191356418

### Build vs Buy

> Apply appropriate purchasing methods. There is no magic one size fits all. Hence you break into small contracts with components that are similar (in terms of evolution). Don't mix commodity with genesis unless you like massive change control cost overruns. -- https://twitter.com/swardley/status/971029206510563328

> In mapping terms, you want to be winning in the orange (good) areas and extracting yourself from the blue (legacy and new legacy) areas. -- https://twitter.com/swardley/status/1428478888401936395

> And if something is not your core domain, why are you wasting time and money (extensively) customizing software (or SaaS for that matter). It's rare that you win the marketplace because you have the best payroll processes.
> Use generic processes that are supported by COTS/SaaS.-- https://twitter.com/crichardson/status/1422967303885389824

* [Wardley Mapping with the Value Flywheel](https://www.theserverlessedge.com/wardley-mapping-with-the-value-flywheel/)

> Why is it so hard to decide to buy? -- https://skamille.medium.com/why-is-it-so-hard-to-decide-to-buy-d86fee98e88e

### **Enterprise**
> Developing and running transformative applications in the cloud requires innovative technology, infrastructure, and services. It also requires a foundation of trust and stability, so organizations have the confidence to make long-term investments on our platform. -- https://cloud.google.com/blog/topics/inside-google-cloud/new-api-stability-tenets-govern-google-enterprise-apis

- **Architecture** - https://www.enterpriseintegrationpatterns.com/ramblings.html | https://microservices.io/patterns/microservices.html | https://www.api.expert/collection/enterprise-it | https://docs.microsoft.com/en-us/graph/overview | https://developer.salesforce.com/ |  

- **Patterns:**
  - Applications: OnPrem (Private Infra) -> SaaS? - [Oracle Cloud Applications](https://docs.oracle.com/en/cloud/saas/index.html) | 
  - Tech Stack:
    - https://stackshare.io
    - https://www.notion.so/Founders-Checklist-Startup-Stack-f24d0601ffc04abfb6b5faf86098371c
    - https://about.gitlab.com/handbook/business-technology/tech-stack/
  - Boost:
    - https://jonmeyers.io/series/build-a-saas-platform-with-next-js-prisma-auth0-and-stripe  
    - https://bedrock.mxstbr.com/
    - https://supabase.io/
    - https://retool.com/
    - https://www.onegraph.com/
    - https://wso2.com/choreo/
    - https://temporal.io/
  
### Components
  
  - **Services:**
    - **Customer** 
      - https://www.rocketlane.com/
      - https://workos.com/
      - https://heap.io/

    - **Auth**     
      - https://auth0.com/blog/introducing-auth0-organizations/
      - https://www.ory.sh/ | http://www.authzed.com | https://www.osohq.com/
      - https://docs.microsoft.com/en-us/azure/active-directory/develop/
      - https://oauth.io/ (120+ OAuth providers)

    - **Communications**
      - https://www.twilio.com/iot 
      - https://sendgrid.com/ ...

    - **Things** 
      - https://eagle.io/
      - https://www.mapped.com/
      - https://www.zensors.com/
    
    - **Security** 
      - https://www.prancer.io/
 
  - **Industry:**
    - **Finance**
      - https://stripe.com/use-cases/saas  (https://stripe.com/docs)
      - https://plaid.com/ ...

    - **Health**
      - https://azure.microsoft.com/en-us/services/azure-api-for-fhir/
      - https://cloud.google.com/healthcare/docs/concepts/fhir ...

### Catalog

- https://www.sastrify.com/
- Search & Browse: https://www.producthunt.com/ | https://saasblocks.io/ | https://welovenocode.com/nocodelist

## API

* https://graphql.stepzen.com/
* https://hopper.wundergraph.com/

## Data

> "For the longest time, we fixated on cycles, making things work faster—the processors, CPUs, GPUs, and more parallel servers. We ignored the data part. Now we have to fixate on data." -- https://www.technologyreview.com/2021/10/27/1037134/data-science-challenges-trustworthy-ai/

> "the teams that produce data are decoupled from the downstream analytics and ML use cases. Likewise, the people dealing with analytics use cases are typically unaware of operational systems that produce the data" -- https://m.subbu.org/broken-state-of-data-4c8a8a30a0c3

Visualization of Data Science workflow https://docs.metaflow.org/introduction/why-metaflow

- Citizen Data Science - https://future.a16z.com/solo-workers-software-stack/
  - https://databricks.com/blog/2021/10/06/bringing-lakehouse-to-the-citizen-data-scientist-announcing-the-acquisition-of-8080-labs.html

- https://www.nanalyze.com/2021/04/no-code-platforms-machine-learning/

- Open (Public?) Datasets Hub:
  - https://github.com/opendatadiscovery/awesome-data-catalogs
  - https://github.com/quiltdata/quilt 

Modern Data Stack - https://sisudata.com/blog/modern-analytics-stack

--

* Database / Warehouse:
  * https://www.firebolt.io/resources/cloud-data-warehouse-comparison
* Wrangling:
  * https://www.getdbt.com/
* Metadata:
  * </>
* Private data:
> Organizations have vast amounts of confidential data locked down due to privacy concerns. Opaque Systems makes confidential data useful by enabling secure analytics and machine learning on encrypted data that comes from one or more sources.
With Opaque Systems, organizations can analyze encrypted data in the cloud using popular tools like Apache Spark, while ensuring that their data is never exposed unencrypted to the cloud provider. -- https://opaque.co/

> Confidential computing is most commonly used in financial services and health care industries and by government agencies, but every industry can benefit from it. -- https://azure.microsoft.com/en-us/solutions/confidential-compute/#customer-stories

--

OpenAI:
* [Applications](https://beta.openai.com/examples)
* [Can Code](https://spectrum.ieee.org/openai-wont-replace-coders)
> "due to the cost of training, it wasn't feasible to retrain the model."
[Cost of Training](https://spectrum.ieee.org/deep-learning-computational-cost)
* [Cohere](https://venturebeat.com/2021/11/15/openai-rival-cohere-launches-language-model-api/)

Stanford:
  * https://sisudata.com/
  * https://ai.stanford.edu/blog/
  * https://hai.stanford.edu/
  * https://crfm.stanford.edu/
    * Foundation models - https://www.youtube.com/watch?v=RLrjKGN89Fc

University of Washington:
  * https://falx.cs.washington.edu/

UC Berkeley:
  * [After Databricks](https://rise.cs.berkeley.edu/blog/the-inside-story-of-how-uc-berkeley-became-the-incubator-for-red-hot-enterprise-startups-databricks-sifive-and-anyscale/)
    * [Anyscale](https://docs.ray.io/en/master/index.html)

CMU:
  * https://ottertune.com/
  * https://db.cs.cmu.edu/seminar2021-dose2/

https://incidentdatabase.ai/

## Infrastructure

“computation may someday be organized as a public utility, just as the telephone system is a public utility. We can envisage computer service companies whose subscribers are connected to them [...]. Each subscriber needs to pay only for the capacity that he actually uses, but he has access to all programming languages characteristic of a very large system.” -- https://sigops.org/s/conferences/hotos/2021/papers/hotos21-s02-stoica.pdf

### Cloud 
 * Providers - https://www.pulumi.com/docs/intro/cloud-providers
 * Universe - https://docs.dagger.io/reference/README

#### Identity, Authentication & Authorization 
  * OpenID
  * [GitHub](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#overview-of-openid-connect)
  * [GitLab](https://docs.gitlab.com/ee/administration/auth/README.html)
  * [Azure AD](https://docs.microsoft.com/en-us/azure/active-directory/external-identities/)
  * [Google Workspace (Formerly G Suite)](https://cloud.google.com/architecture/identity)
  * PingIdentity
  * OneLogin
  * SAML 2.0 Identity Providers ( System for Cross-domain Identity Management (SCIM) 2.0 - Manage users and groups centrally in your IdP and setup synchronization between Acuity and your SAML 2.0 identity providers )
    * Azure AD
    * Okta
    * Google Workspace (Formerly G Suite)

#### Database
  
Databricks 
| Redis Labs
| Neo4j
| Cockroach Labs
| Couchbase
| DataStax
| Firebolt
| MariaDB
| Imply
| PlanetScale
| Timescale 

- https://cloudwars.co/cloud-database-report/
- https://www.honeycomb.io/blog/time-series-database/

#### Queues
  * Kafka or Kafka compat
  
#### Containers

* [Why?](https://insights.sei.cmu.edu/blog/11-leading-practices-when-implementing-a-container-strategy/)

* GCP Cloud Run - https://cloud.google.com/run/
  * FAQ - https://github.com/ahmetb/cloud-run-faq 
  * Anthos - https://cloud.google.com/anthos/docs/setup/public-cloud
* Thousands of ways on AWS, one way -> [Introducing AWS App Runner](https://aws.amazon.com/blogs/containers/introducing-aws-app-runner/)
* Build better containerized apps with less friction - https://www.slim.ai/

### Security

* [Cloud Native Security Whitepaper](https://github.com/cncf/tag-security/tree/main/security-whitepaper)
  * https://control-plane.io/

### Testing
* [Talk About Testing](https://dannorth.net/2021/07/26/we-need-to-talk-about-testing/)

### Reliability
* SLOs, Error Budgets - https://nobl9.com/

* Status Dashboards - Introduced 2008, https://aws.amazon.com/blogs/aws/the-service-hea/ 
  - https://statuscast.com/cloud-app-integrations/
  - [AWS Service Health Dashboard](https://status.aws.amazon.com/) 
  - [Microsoft Azure Status](https://status.azure.com/en-us/status)
  - [Google Cloud Status Dashboard](https://status.cloud.google.com/)
  - [Oracle Cloud Infrastructure Status](https://ocistatus.oraclecloud.com/)
  - [Alibaba Cloud condition monitoring](https://status.alibabacloud.com/)
  - [IBM Cloud Status](https://cloud.ibm.com/status)
  - [Cloudflare System Status](https://www.cloudflarestatus.com/)
  - https://status.dev.azure.com/ | https://status.gitlab.com/ | https://www.gitpodstatus.com/ 

* Disaster Recovery - https://cloud.google.com/architecture/disaster-recovery

* https://sre.google/
  * https://thenewstack.io/google-sre-site-reliability-engineering-at-a-global-scale/
* [Single best repo of DevOps-related content](https://cloud.google.com/architecture/devops/capabilities) 

### 
* Infrastructure as Software - https://en.wikipedia.org/wiki/Infrastructure_as_Software
  * [Infrastructure as Software vs Infrastructure as Code](https://www.youtube.com/watch?v=rtng6GNQd4w)
  * https://github.com/kris-nova/public-speaking/blob/main/presentations/2021-infrastructure-as-software/a-1-iac-vs-ias.zomg
  * https://nivenly.com/lib/2021-10-20-gitops/
* Learning from Incidents - https://www.thevoid.community/

## Open Source

* https://github.com/hofstadter-io
* https://backstage.io/docs/auth/ 
* https://github.com/authzed
* https://github.com/ory
* https://github.com/argoproj/argo-cd/
* https://opentelemetry.io/
  * DataDog 
    * https://opstrace.com/blog/introducing-datadog-compatible-http-api
    * https://github.com/SigNoz/signoz
* https://artillery.io/

## Twitter

* [rereading 25 of my favorite computer science papers](https://twitter.com/funcOfJoe/status/1473745632061837313)
* [The technology adoption curve is much longer than people on the cutting edge assume](https://twitter.com/emollick/status/1473677758316355591)

## Things 

## Architecture

