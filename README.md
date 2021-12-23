# Open software design

> Open-source software has made it easier for software developers to study and learn programming by looking at real-world working software. But what about software design? Building architects (not software) study thousands of buildings in their training and career. Software developers, in contrast, only study a handful of other large software designs. This means we repeat mistakes that we would have been able to avoid, should we be able to learn from others.
> Open software design is the idea that software design must be available for public use. It's the idea that we should Work with the garage door up. The main audience of the design is developers who built the software, and not external people. The contents are living artefacts, they naturally evolve as the software evolve. In the world where the majority of large real-world programs are closed-source, opening up our software design documentation is the least we can do to contribute back to the community.

-- https://notes.ceilfors.com/Open_software_design.html

## Twitter

* [rereading 25 of my favorite computer science papers](https://twitter.com/funcOfJoe/status/1473745632061837313)
* [The technology adoption curve is much longer than people on the cutting edge assume](https://twitter.com/emollick/status/1473677758316355591)
* [most lasting achievements of the early internet - Linux & open source, Wikipedia - were often driven by intrinsic incentives](https://twitter.com/emollick/status/1473721323688017927)
* [I hope @cue_lang gains more popularity this year. It's been the correct fit for so many new greenfield projects lately.](https://twitter.com/rakyll/status/1473198774914748417)
* [At the end of the day there wasn't one way to DevOps. The only real requirement was to be intentional about your engineering culture and continuously work to improve it.](https://twitter.com/kelseyhightower/status/1473370398498693120)

## Sandbox

* [Cloud](https://notes.ceilfors.com/Cloud_sandbox.html)
* [Auth0](https://learn.sandcastle.cloud/)
* [Authzed](https://play.authzed.com/schema)
* [Metadata](https://sandbox.open-metadata.org/)
* [DataHub](https://demo.datahubproject.io/)
* [Energy](https://energy-models.com/what-is-energy-modeling-building-simulation)
* [Building](https://brickschema.org/get-started/)

## API

* https://graphql.stepzen.com/
* https://hopper.wundergraph.com/

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

## Things 

## Architecture

https://martinfowler.com/articles/scaling-architecture-conversationally.html
1. A thinking and recording tool: **Decision Records**
2. A time and place for conversations: **The Architecture Advisory Forum**
3. </>
4. </>

