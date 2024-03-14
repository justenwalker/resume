# Justen Walker

- [<img src="assets/mastodon-logo.png" alt="Mastodon" style="width: 1em; height: 1em; vertical-align:middle;"></img> @jwalk@hachyderm.io](https://hachyderm.io/@jwalk)
- [<img src="assets/matrix-icon.png" alt="matrix.org" style="width: 1em; height: 1em; vertical-align:middle;"></img> @justenwalker:matrix.org](https://matrix.to/#/@justenwalker:matrix.org)
- [<img src="assets/bluesky-logo.png" alt="BlueSky" style="width: 1em; height: 1em; vertical-align:middle;"></img> @justenwalker.bsky.social](https://bsky.app/profile/justenwalker.bsky.social)
- [<img src="assets/linkedin-logo.png" alt="LinkedIn" style="width: 1em; height: 1em; vertical-align:middle;"></img> justenwalker](https://www.linkedin.com/in/justenwalker/)
- [<img src="assets/document-icon.png" alt="Resume" style="width: 1em; height: 1em; vertical-align:middle;"></img> resume](assets/justen-walker-resume.pdf)
- He/Him

---
## Work Experience

I'm primarily a Go developer with a heavy focus on cloud infrastructure and platform development. I have experience with Kubernetes, AWS and Azure Clouds, Terraform, Ansible.

### Sr. Staff Software Engineer at [Stoke, LLC.](https://www.stoke.com)

2022-12 - present

Cloud Infrastructure and Application Programmer

- Authored over 10 OSS contributions to the AirByte Platform to ensure platform stability.
- Built and managed back-end infrastructure to support the platform; ensuring reliable and scalable performance.
- Continuously improved Go code quality and performance; Achieved 60% increase in code coverage, and near 100% across critical paths.
- Optimized cloud computing costs. Reduced cloud spend by 50% using Spot instances and auto-scaling.
- Implemented a new workflow engine using Temporal to create a scalable and robust data ingestion pipeline. Reduced cycle time by 50% when adding new data streams.
- Worked with and managed team members across time zones.

#### Projects
- **Airbyte Platform** : Managed the backbone of our data ingestion pipeline, ensuring the system could scale automatically; responding to outages and bugs; and upgrades.
- **AWS EKS** : Many of the back-end systems were hosted on EKS. Ensured timely upgrades of EKS and all core components including monitoring and autoscaling with Karpenter.
- **Stoke Workflows** : A workflow engine for scheduling and monitoring customer data-stream ingestion; built using Temporal and AirByte.


### Sr. Staff Software Engineer at [Artica, Inc.](https://www.linkedin.com/company/articaofficial/)

2021-10 - 2022-11

Cloud Infrastructure and Application Programmer

- Created initial AWS Cloud Infrastructure and Organization Structure.
- Implemented secure access to VPC resources via HashiCorp Boundary and TailScale - securing our access to sensitive data and services.
- Introduced Terraform to describe AWS infrastructure requirements as code (IaC), applied via Terraform Cloud - prevented manual errors from causing outages.
- Created libraries and documentation to enable End-to-end observability via OpenTelemetry and DataDog.
- Automated provisioning of group membership and access via AzureAD and Terraform - reducing onboarding times from days into minutes.

#### Projects
- **gsd** : A cli application that supports getting access to production systems securely by correctly configuring AWS configuration and calling boundary and aws apis to construct secure tunnels.
- **go-observe** : A library that encodes our observability best-practices and provides a simple interface to including observability across our applications.


### Principal Software Engineer at [Walmart](https://www.walmartlabs.com)

2019-10 - 2021-10

Creating controllers and tooling for Kubernetes to support Walmart's Kubernetes Platform

- Created many CLI utilities in Go
- Created a process for fully-automated end-to-end testing of controllers using Kind, Ansible, and Argo Workflows
- Investigated and benchmarked performance characteristics of controllers and databases used by the platform
- Created a pod resource usage-collector with minute precision, which collected data to be used for charge-backs on multi-tenant clusters.
- Contributed in small-ish ways to Open source projects that we used

#### Projects
- **mco-env** : A fully dockerized development environment to improve the onboarding experience, and standardize our tool-set.
- **Sledge** : A Core CLI for interacting with Walmart's Kubernetes platform. Including a plugin distribution system similar to Krew.
- **Infomox** : A mock application that mirrored Infoblox APIs to help with end-to-end testing.


### Principal Software Engineer at [Jet.com, Inc.](https://jet.com/)

2015-09 - 2019-10

#### Role: Software Engineer, DevOps
Designing and maintaining the cloud services, site reliability, and monitoring stacks

- Initially explored HashiCorp Nomad, and eventually replaced our home-grown microservices scheduler
- Introduced HashiCorp Vault for secret storage and documented it's usage for developers
- Develop Microservice Platform using Docker and Nomad
- Create utilities to automate common maintenance tasks
- Create tooling in Go to support cloud platform provisioning and operation
- Automate provisioning of VMs with Ansible
- On-call support for platform issues and traffic fail-overs
- Work on projects to improve site reliability
- Documented and presented new solutions at Design Reviews and Production Readiness Reviews

#### Role: Tech Lead, Microservices Platform Squad
Provided technical guidance to the Microservice platform team, assigned work to the squad-members, and ran planning meetings

- Coached a team of around 8 developers.
- Worked with local Hoboken and Remote Dublin members
- Met with other Tech leaders to advocate platform changes and best practices
- Ran stand-up, planning meetings to decide Sprint work, and lead end-of-sprint retrospective
- Advocated squads concerns and provided updates during weekly sync-up meetings with the rest of the squads

#### Role: Go Language Club Organizer
Ran bi-weekly book club meetings to learn and advocate the Go language

- Worked through exercises on gophercises.com with members of the club
- Facilitated presentations of exorcizes during the meeting
- Developed and presented alternative solutions and identified pros and cons
- Provided advice and assistance to anyone needing help

#### Projects
- **Cable** : Dynamic inventory caching system to allow Ansible to have current Azure inventory up to date without exceeding our API quotas. Written mostly in Go.
- **Groot** : fail-over automation technology using state machine, and desired state configuration to plan fail-over steps and ensure safety. Written in Go and some Angular for the front-end.
- **Damon** : Supervisor process for Windows executables to provide resource constraints without using Windows containers. Used Go as well as the Win32 API.
- **Mantis** : Shared library for common functionality that was used across many utilities such as: calling the Azure API, Retry logic, Getting and refreshing HashiCorp Vault Secrets.
- **Interstellar** : A Go library for interacting with Cosmos DB Rest API
- **UpDog** : A server and client that keeps NGINX configurations in sync with Consul service discovery to allow NGINX to serve dynamic sets of vHosts. Written in Go, Replaced our earlier use of consul-template which ceased to scale well for our use. Also, an endless source of amusement in Slack when people asked what it was.
- **OhNo** : Slack chat-bot that can be used to manually page teams, look up when you are going to be on-call next, and who is currently on call, and going to be on call next. It also shared random comics from webcomicname.com if you asked it nicely.
- **Gizmo** : A utility to translate JSON descriptions of services into Nomad Job Manifests. This was used in order to maintain standards of how we stored and released code to Nomad as well as abstract away implementation details about platform features like proxies and alerting.
- **Ansible Auto-Provisioner** : A combination of Cable + Ansible + Jenkins and Azure Custom Data that allowed our VMs to automatically check in with Cable, and run a Jenkins provisioning script to run Ansible against newly provisioned hardware. Enabled the use of Virtual Machine Scale-Sets and Auto-scaling systems
- **Nomad Agent Auto-Scaler** : Service that measured total allocated resources on a scale set versus and added/removed VMs to ensure adequate resources to accommodate new services and scale-out events.


### DevOps Engineer at [Skubana, Inc.](https://www.skubana.com/)

2014-12 - 2015-09

#### Role: DevOps Engineer
Developed and maintained initial cloud infrastructure and release pipelines to deploy code on AWS.

- Maintained most SaaS solutions including RDS, ELB, and EC2 instances
- Installed and maintained Jenkins infrastructure and build pipelines
- Developed Site-Reliability measures including alerting on application exceptions and operating the ELK stack.

#### Role: Backend Software Engineer
Developed application logic for the Backend.

- Participated in Agile Sprint planning and estimation
- Worked with team members to design and implement new features and fix issues

#### Projects
- **Invoicing** : Created invoicing system to charge customers using Stripe API every month.
- **Batch Service** : Worked on Batch processing service to provide asynchronous task execution
- **Order Split/Merge** : Developed order splitting and merging logic to facilitate consolidation of shipments to customers


### Software Engineer at [The Cools, Inc.](https://cools.com/)

2012-01 - 2014-11

#### Role: DevOps Engineer
Developed and maintained initial cloud and release pipelines to deploy infrastructure and code on AWS.

- Worked in PHP, Python, Ruby and Bash
- Maintained most SaaS solutions including RDS, ELB, and EC2 instances
- Installed and maintained Jenkins infrastructure and build pipelines
- Developed Site-Reliability measures including alerting on application exceptions and operating the ELK stack.
- Developed practice of infrastructure-as-code using Puppet

#### Role: Backend Software Engineer
Developed application logic for the Backend.

- Primarily worked in Java, and some Groovy
- Participated in Agile Sprint planning and estimation
- Worked with team members to design and implement new features and fix issues
- Developed scripts to do video conversion and uploads to support some CMS features.

#### Projects
- **Apache OFB** : Worked with contractors to extract back-end order processing into its own system (Apache Open-for-business)
- **CMS Integration** : Integrated WordPress into custom CMS solution integrated into the main site. to give editors control over content without exposing WordPress to the internet
- **Gradle Migration** : Helped move from Maven to Gradle build system


### Technical Analyst at [Credit Suisse, AG](https://www.credit-suisse.com/)

2009-01 - 2012-01

Worked in Corporate Application Systems to manage infrastructure projects and assign work to contractors

- Standardized a packaging method for deploying a business intelligence application and customizations.
- Created a process for tracking deployments to environments
- Documented many processes in an effort to hand over responsibilities to other team members

#### Projects
- **CASIM** : An inventory management and support system in C# and ASP. Used to track hardware inventory and allocation to specific business projects.
- **SVN Migration** : Helped support developer migration from ClearCase to SVN


### Technical Consultant at [NJIT Residence Life](https://www.njit.edu/reslife/)

2006-10 - 2009-06

Consulted with Residence life on technology projects to enhance student quality of life

- Maintained the databases and assisted in room selection process
- Developed PHP Web applications for use by students and faculty

#### Projects
- **Room selector** : Application to allow students to select rooms during semester room selection process
- **Desk Manager** : Created desk management system to allow residence hall desk-staff to manage work schedule.
- **Payroll Tracker** : Web application to assist with payroll reporting for hourly employees.

---
## Open Source Contributions

Below are a selection of open source projects I've been involved in, either as a creator, or as a contributor. I've done some other minor stuff in other languages, but the majority of my work has been in Go.

### [Goodwill](https://github.com/justenwalker/goodwill) (Creator)

Plugin for Walmart Labs Concord that provides a way to write tasks in Go.

### [Squiggly](https://github.com/justenwalker/squiggly) (Creator)

A Forwarding proxy with support for upstream Proxy Auto Config (PAC) written in Go.

### [Damon](https://github.com/jet/damon) (Creator)

Supervisor program to constrain Windows executables running under Nomad's raw_exec driver.

### [Interstellar](https://github.com/jet/go-interstellar) (Creator)

A Go client for interacting with the REST/SQL API of CosmosDB.

### [Mantis](https://github.com/jet/go-mantis) (Creator)

A "standard library" for Jet's Golang codebase.

### [Container Pilot](https://www.joyent.com/containerpilot) ([Major Contribution](https://github.com/joyent/containerpilot/pulls?utf8=%E2%9C%93&q=author%3Ajustenwalker))

A service for auto-discovery and configuration of applications running in containers.

### [AirByte Platform](https://github.com/airbytehq/airbyte-platform) ([Minor Contribution](https://github.com/airbytehq/airbyte-platform/pulls?q=is%3Apr+author%3Ajustenwalker))

Airbyte is an open-source data integration engine that helps you consolidate your data in your data warehouses, lakes and databases.

### [Helm](https://helm.sh/) ([Minor Contribution](https://github.com/helm/helm/pulls?utf8=%E2%9C%93&q=author%3Ajustenwalker))

The Kubernetes Package Manager

### [HashiCorp Boundary](https://www.boundaryproject.io/) ([Minor Contribution](https://github.com/hashicorp/boundary/pulls?utf8=%E2%9C%93&q=author%3Ajustenwalker))

Boundary enables identity-based access management for dynamic infrastructure. 

### [Kuberlr](https://github.com/flavio/kuberlr) ([Minor Contribution](https://github.com/flavio/kuberlr/pulls?q=is%3Apr+author%3Ajustenwalker+is%3Aclosed))

A CLI that wraps kubectl and auto-downloads the appropriate version for the target cluster based on the API version.

### [Termenv](https://github.com/muesli/termenv) ([Minor Contribution](https://github.com/muesli/termenv/pulls?q=is%3Apr+is%3Aclosed+author%3Ajustenwalker))

Go Library to support Advanced ANSI style & color support for your terminal applications

### [Azure Terraform Provider](https://www.terraform.io/docs/providers/azurerm/index.html) ([Minor Contribution](https://github.com/terraform-providers/terraform-provider-azurerm/pulls?utf8=%E2%9C%93&q=author%3Ajustenwalker))

Terraform provider for Azure Resource Manager

### [HashiCorp Nomad](https://nomadproject.io/) ([Minor Contribution](https://github.com/hashicorp/nomad/pulls?utf8=%E2%9C%93&q=author%3Ajustenwalker))

Nomad is an easy-to-use, flexible, and performant workload orchestrator.

---
## Speaking Events

I've been speaking at a couple of events. If that interests you, check out links to my talks. I've never watched any of these, so hopefully I didn't embarrass myself; And if I did, please don't tell me.

### (2020-11-11) GopherCon 2020
Talk: **[Safety Not Guaranteed - Using unsafe to syscall Windows APIs without CGO](https://www.youtube.com/watch?v=EsPcKkESYPA)**

I gave an in-depth talk about how you can use Go to call Win32 APIs without CGO. This has examples, caveats, and techniques for memory management..

### (2019-09-09) HashiConf US 2019
Talk: **[Containment Without Containers - Running Windows Microservices on Nomad](https://www.youtube.com/watch?v=SJgW1bVmj2c)**

I gave a talk about the state of Windows Containers, how Jet uses the Win32 APIs via Damon to constraint windows microservices on Nomad, and gave a crash course on calling the Win32 API from Go without needing CGO.

### (2018-06-19) QCon NYC
Talk: **Managing Moderate-Scale Multi-Tenant Micro-Services @Jet.com**

I presented how we manage our Microservices at scale and our journey from our legacy in-house system to using HashiCorp Nomad and the custom tooling we built.

*Note:* This talk was in the "Sponsored Solutions" track. Unfortunately, the video was never published. The 2018 Conference site is no longer available

### (2017-05-15) HashiDays NYC
Talk: **[Nomad auto-proxying with Consul-template and NGINX](https://www.youtube.com/watch?v=75vF92Vue2U)**

I presented our (then current) method of using Consul-Template to dynamically write NGINX templates for our Nomad services.

### (2016-10-11) AnsibleFest Brooklyn
Talk: **[Ansible at Jet - Managing Azure Cloud Deployments](https://www.ansible.com/ansible-at-jet-managing-azure-cloud-deployments)**

I gave a talk about how we use a combination of some custom go utilities, jenkins, and ansible to automatically provision our infrastructure when it gets created for the first time, and how we organize our repository structure to support a growing team.

*Note:* Video no longer available.

### (2016-06-14) Container Summit - City Series: Chicago
Talk: **Containers & ContainerPilot at Jet.com**

As a major contributor to Container Pilot (at the time); Joyent invited me to have a "fire-side chat" with Bryan Cantrill (Then CTO, Joyent) about our use of Containers and Container Pilot at Jet.com.

*Note:* Conference site and video no longer available.

### (2016-02-10) Container Summit NYC (2016)
Talk: **Containers are Eating the Enterprise (Panel)**

I was on a panel moderated by the venerable James Turnbull (Kickstarter) with three awesome people: Michael Hamrah (Uber), Larry Glenn (HBC), and Jeff Ashton (Canadian Tire). We talked about Docker containers in the enterprise and some of our experiences either running containers in production, or the road we're taking to get it there. This was pretty much my first time speaking in a professional capacity, so please be kind.

*Note:* Conference site and video no longer available.

---
## Writing

I've written a couple of blog posts explaining some of the research and work my team-mates and I have done.
These may give you a better idea of what I think about.

### [Breaking all the rules: Using Go to call Windows API](https://medium.com/jettech/breaking-all-the-rules-using-go-to-call-windows-api-2cbfd8c79724)

2019-01-15 - [Jet Tech Blog](https://medium.com/jettech)

I go into detail about how to call the Win32 API through Go and some of the traps to avoid when doing so.

### [Containing Windows Executables with Damon](https://medium.com/jettech/containing-windows-executables-with-damon-898ab4a31ea4)

2018-12-11 - [Jet Tech Blog](https://medium.com/jettech)

I describe an alternative method to constraining Windows executables without using Containers.

### [A short introduction to Windows Containers](https://medium.com/jettech/a-short-introduction-to-windows-containers-db5adc0db536)

2018-12-11 - [Jet Tech Blog](https://medium.com/jettech)

An overview of how Windows Containers work and some of their features and drawbacks.

### [Using Ansible to Manage Jet Infrastructure](https://medium.com/jettech/using-ansible-to-manage-jet-infrastructure-a3cddd4c6440)

2016-10-07 - [Jet Tech Blog](https://medium.com/jettech)

A short introduction to Ansible and why we at Jet chose it.

---
## Education

This may not be as important anymore, but I paid for it and did OK; so I get to put it here.

### **B.S., Computer Engineering** from [New Jersey Institute of Technology](https://www.njit.edu/)
2005-09 - 2009-06

GPA: **3.95/4.0**

#### Honors
- Member, Phi Eta Sigma honors society
- Albert Dorman Honors College (ADHC) Scholar
- Summa Cum Laude