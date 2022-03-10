## Infrastructure

“Will AWS and Azure become "dumb pipes" of compute, storage, and other fundamental data center primitives? Or will they continue to compete and beat third party providers at high level services?” -- https://matt-rickard.com/aws-is-not-a-dumb-pipe/

“computation may someday be organized as a public utility, just as the telephone system is a public utility. We can envisage computer service companies whose subscribers are connected to them [...]. Each subscriber needs to pay only for the capacity that he actually uses, but he has access to all programming languages characteristic of a very large system.” -- https://sigops.org/s/conferences/hotos/2021/papers/hotos21-s02-stoica.pdf

### Cloud 
 * Cloud Native Glossary - https://glossary.cncf.io/
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
  * https://authzed.com/blog/zanzibar-implementations/
  * https://docs.permit.io/overview/authorization_concepts

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
