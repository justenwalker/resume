package data

func init() {
	MyCV = CV{
		WorkExperience: []Company{
			Company{
				Name:    "Jet.com, Inc.",
				Website: "https://jet.com/",
				Title:   "Staff Software Engineer",
				Start:   YM("2015-09"),
				End:     Present(),
				Roles: []Role{
					Role{
						Name:        "Software Engineer, DevOps",
						Description: `Designing and maintaining the cloud services, site reliability, and monitoring stacks`,
						Highlights: []string{
							"Initially explored HashiCorp Nomad, and eventually replaced our home-grown microservices scheduler",
							"Introduced HashiCorp Vault for secret storage and documented it's usage for developers",
							"Develop Microservice Platform using Docker and Nomad",
							"Create utilities to automate common maintenance tasks",
							"Create tooling in Go to support cloud platform provisioning and operation",
							"Automate provisioning of VMs with Ansible",
							"On-call support for platform issues and traffic failovers",
							"Work on projects to improve site reliability",
							"Documented and presented new solutions at Design Reviews and Production Readiness Reviews",
						},
					},
					Role{
						Name:        "Tech Lead, Microservices Platform Squad",
						Description: `Provided technical guidance to the Microservice platform team, assigned work to the squadmembers, and ran planning meetings`,
						Highlights: []string{
							"Coached a team of around 8 developers.",
							"Worked with local Hoboken and Remote Dublin members",
							"Met with other Tech leaders to advocate platform changes and best practices",
							"Ran stand-up, planning meetings to decide Sprint work, and lead end-of-sprint retrospective",
							"Advocated squads concerns and provided updates during weekly sync-up meetings with the rest of the squads",
						},
					},
					Role{
						Name:        "Go Language Club Organizer",
						Description: `Ran bi-weekly book club meetings to learn and advocate the Go language`,
						Highlights: []string{
							"Worked through exercises on gophercises.com with members of the club",
							"Facilitated presentations of exorcizes during the meeting",
							"Developed and presented alternative solutions and identified pros and cons",
							"Provided advice and assistance to anyone needing help",
						},
					},
				},
				Projects: []Project{
					Project{"Cable", "Dynamic inventory caching system to allow Ansible to have current Azure inventory up to date without exceeding our API quotas. Written mostly in Go."},
					Project{"Groot", "Failover automation technology using state machine, and desired state configuration to plan failover steps and ensure safety. Written in Go and some Angular for the front-end."},
					Project{"Damon", "Supervisor process for Windows executables to provide resource constraints without using Windows containers. Used Go as well as the Win32 API."},
					Project{"Mantis", "Shared library for common functionality that was used across many utilities such as: calling the Azure API, Retry logic, Getting and refreshing HashiCorp Vault Secrets."},
					Project{"Interstellar", "A Go library for interacting with Cosmos DB Rest API"},
					Project{"UpDog", "A server and client that keeps NGINX configurations in sync with Consul service discovery to allow NGINX to serve dynamic sets of vHosts. Written in Go, Replaced our earlier use of consul-template which ceased to scale well for our use. Also, an endless source of amusement in Slack when people asked what it was."},
					Project{"OhNo", "Slack chat-bot that can be used to manually page teams, look up when you are going to be on-call next, and who is currently on call, and going to be on call next. It also shared random comics from webcomicname.com if you asked it nicely."},
					Project{"Gizmo", "A utility to translate JSON descriptions of services into Nomad Job Manifests. This was used in order to maintain standards of how we stored and released code to Nomad as well as abstract away implementation details about platform features like proxies and alerting."},
					Project{"Ansible Auto-Provisioner", "A combination of Cable + Ansible + Jenkins and Azure Custom Data that allowed our VMs to automatically check in with Cable, and run a Jenkins provisioning script to run Ansible against newly provisioned hardware. Enabled the use of Virtual Machine Scale-Sets and Auto-scaling systems"},
					Project{"Nomad Agent Auto-Scaler", "Service that measured total allocated resources on a scale set versus and added/removed VMs to ensure adequate resources to accommodate new services and scale-out events."},
				},
				Tags: []string{
					"azure",
					"devops",
					"ansible",
					"git",
					"golang",
					"security",
					"observability",
					"opensource",
					"shell",
					"leadership",
				},
			},
			Company{
				Name:    "Skubana, Inc.",
				Website: "https://www.skubana.com/",
				Title:   "DevOps Engineer",
				Start:   YM("2014-12"),
				End:     YM("2015-09"),
				Roles: []Role{
					Role{
						Name:        "DevOps Engineer",
						Description: `Developed and maintained initial cloud infrastructure and release pipelines to deploy code on AWS.`,
						Highlights: []string{
							"Maintained most SaaS solutions including RDS, ELB, and EC2 instances",
							"Installed and maintained Jenkins infrastructure and build pipelines",
							"Developed Site-Reliability measures including alerting on application exceptions and operating the ELK stack.",
						},
					},
					Role{
						Name:        "Backend Software Engineer",
						Description: `Developed application logic for the Backend.`,
						Highlights: []string{
							"Participated in Agile Sprint planning and estimation",
							"Worked with team members to design and implement new features and fix issues",
						},
					},
				},
				Projects: []Project{
					Project{"Invoicing", "Created invoicing system to charge customers using Stripe API every month."},
					Project{"Batch Service", "Worked on Batch processing service to provide asynchronous task execution"},
					Project{"Order Split/Merge", "Developed order splitting and merging logic to facilitate consolidation of shipments to customers"},
				},
				Tags: []string{
					"aws",
					"devops",
					"ansible",
					"git",
					"java",
					"shell",
				},
			},
			Company{
				Name:    "The Cools, Inc.",
				Website: "https://cools.com/",
				Title:   "Software Engineer",
				Start:   YM("2012-01"),
				End:     YM("2014-11"),
				Roles: []Role{
					Role{
						Name:        "DevOps Engineer",
						Description: `Developed and maintained initial cloud and release pipelines to deploy infrastructure and code on AWS.`,
						Highlights: []string{
							"Worked in PHP, Python, Ruby and Bash",
							"Maintained most SaaS solutions including RDS, ELB, and EC2 instances",
							"Installed and maintained Jenkins infrastructure and build pipelines",
							"Developed Site-Reliability measures including alerting on application exceptions and operating the ELK stack.",
							"Developed practice of infrastructure-as-code using Puppet",
						},
					},
					Role{
						Name:        "Backend Software Engineer",
						Description: `Developed application logic for the Backend.`,
						Highlights: []string{
							"Primarily worked in Java, and some Groovy",
							"Participated in Agile Sprint planning and estimation",
							"Worked with team members to design and implement new features and fix issues",
							"Developed scripts to do video conversion and uploads to support some CMS features.",
						},
					},
				},
				Projects: []Project{
					Project{"Apache OFB", "Worked with contractors to extract back-end order processing into its own system (Apache Open-for-business)"},
					Project{"CMS Integration", "Integrated WordPress into custom CMS solution integrated into the main site. to give editors control over content without exposing WordPress to the internet"},
					Project{"Gradle Migration", "Helped move from Maven to Gradle build system"},
				},
				Tags: []string{
					"aws",
					"devops",
					"puppet",
					"git",
					"java",
					"ruby",
					"groovy",
					"shell",
				},
			},
			Company{
				Name:    "Credit Suisse, AG",
				Website: "https://www.credit-suisse.com/",
				Title:   "Technical Analyst",
				Start:   YM("2009-01"),
				End:     YM("2012-01"),
				Roles: []Role{
					Role{
						Name:        "Technical Analyst",
						Description: `Worked in Corporate Application Systems to manage infrastructure projects and assign work to contractors`,
						Highlights: []string{
							"Standardized a packaging method for deploying a business intelligence application and customizations.",
							"Created a process for tracking deployments to environments",
							"Documented many processes in an effort to hand over responsibilities to other team members",
						},
					},
				},
				Projects: []Project{
					Project{"CASIM", "An inventory management and support system in C# and ASP. Used to track hardware inventory and allocation to specific business projects."},
					Project{"SVN Migration", "Helped support developer migration from ClearCase to SVN"},
				},
				Tags: []string{
					"projectmanagement",
					"svn",
					"java",
					"shell",
				},
			},
			Company{
				Name:    "NJIT Residence Life",
				Website: "https://www.njit.edu/reslife/",
				Title:   "Technical Consultant",
				Start:   YM("2006-10"),
				End:     YM("2009-06"),
				Roles: []Role{
					Role{
						Name:        "Technical Analyst",
						Description: `Consulted with Residence life on technology projects to enhance student quality of life`,
						Highlights: []string{
							"Maintained the databases and assisted in room selection process",
							"Developed PHP Web applications for use by students and faculty",
						},
					},
				},
				Projects: []Project{
					Project{"Room selector", "Application to allow students to select rooms during semester room selection process"},
					Project{"Desk Manager", "Created desk management system to allow residence hall desk-staff to manage work schedule."},
					Project{"Payroll Tracker", "Web application to assist with payroll reporting for hourly employees."},
				},
			},
		},
		OpenSource: []OpenSource{
			OpenSource{
				Name:         "Damon",
				Website:      "https://github.com/jet/damon",
				Description:  `Supervisor program to constrain Windows executables running under Nomad's raw_exec driver.`,
				Contribution: Creator,
				Tags: []string{
					"nomad",
					"windows",
					"golang",
					"containers",
				},
			},
			OpenSource{
				Name:         "Interstellar",
				Website:      "https://github.com/jet/go-interstellar",
				Description:  `A Go client for interacting with the REST/SQL API of CosmosDB.`,
				Contribution: Creator,
				Tags: []string{
					"cosmosdb",
					"azure",
					"golang",
					"documentdb",
					"library",
				},
			},
			OpenSource{
				Name:         "Mantis",
				Website:      "https://github.com/jet/go-mantis",
				Description:  `A "standard library" for Jet's Golang codebase.`,
				Contribution: Creator,
				Tags: []string{
					"golang",
					"library",
				},
			},
			OpenSource{
				Name:         "Container Pilot",
				Website:      "https://www.joyent.com/containerpilot",
				Link:         "https://github.com/joyent/containerpilot/pulls?utf8=%E2%9C%93&q=author%3Ajustenwalker",
				Description:  `A service for autodiscovery and configuration of applications running in containers.`,
				Contribution: MajorContributor,
				Tags: []string{
					"consul",
					"orchestration",
					"containers",
					"docker",
					"joyent",
					"service-discovery",
					"triton",
				},
			},
			OpenSource{
				Name:         "Azure Terraform Provider",
				Website:      "https://www.terraform.io/docs/providers/azurerm/index.html",
				Link:         "https://github.com/terraform-providers/terraform-provider-azurerm/pulls?utf8=%E2%9C%93&q=author%3Ajustenwalker",
				Description:  "Terraform provider for Azure Resource Manager",
				Contribution: MinorContributor,
				Tags: []string{
					"consul",
					"orchestration",
					"containers",
					"docker",
					"joyent",
					"service-discovery",
					"triton",
				},
			},
			OpenSource{
				Name:         "Nomad",
				Website:      "https://nomadproject.io/",
				Description:  `Nomad is an easy-to-use, flexible, and performant workload orchestrator.`,
				Link:         "https://github.com/hashicorp/nomad/pulls?utf8=%E2%9C%93&q=author%3Ajustenwalker",
				Contribution: MinorContributor,
				Tags: []string{
					"golang",
					"docker",
					"scheduler",
				},
			},
		},
		SpeakingEvents: []SpeakingEvent{
			SpeakingEvent{
				Name:        "Managing Moderate-Scale Multi-Tenant Micro-Services @Jet.com",
				Event:       "QCon NYC",
				Description: `I presented how we manage our Microservices at scale and our journey from our legacy in-house system to using HashiCorp Nomad and the custom tooling we built.`,
				Date:        YMD("2018-06-19"),
				Link:        "https://qconnewyork.com/ny2018/presentation/managing-moderate-scale-multi-tenant-micro-services-jetcom",
				Notes:       `This talk was in the "Sponsored Solutions" track. Unfortunately, the video was never published.`,
			},
			SpeakingEvent{
				Name:        "Nomad auto-proxying with Consul-template and NGINX",
				Event:       "HashiDays NYC",
				Description: `I presented our (then current) method of using Consul-Template to dynamically write NGINX templates for our Nomad services.`,
				Date:        YMD("2017-05-15"),
				Link:        "https://www.youtube.com/watch?v=75vF92Vue2U",
			},
			SpeakingEvent{
				Name:        "Ansible at Jet - Managing Azure Cloud Deployments",
				Event:       "AnsibleFest Brooklyn",
				Description: `I gave a talk about how we use a combination of some custom go utilities, jenkins, and ansible to automatically provision our infrastructure when it gets created for the first time, and how we organize our repository structure to support a growing team.`,
				Date:        YMD("2016-10-11"),
				Link:        "https://www.ansible.com/ansible-at-jet-managing-azure-cloud-deployments",
			},
			SpeakingEvent{
				Name:        "Containers & ContainerPilot at Jet.com",
				Event:       "Container Summit - City Series: Chicago",
				Description: `As a major contributor to Container Pilot (at the time); Joyent invited me to have a "fire-side chat" with Bryan Cantrill (CTO, Joyent) about our use of Containers and Container Pilot at Jet.com.`,
				Date:        YMD("2016-06-14"),
				Link:        "https://containersummit.io/city-series/2016/chicago/videos/containers-and-containerpilot-at-jet-com",
			},
			SpeakingEvent{
				Name:        "Containers are Eating the Enterprise (Panel)",
				Event:       "Container Summit NYC (2016)",
				Description: `I was on a panel moderated by the venerable James Turnbull (Kickstarter) with three awesome people: Michael Hamrah (Uber), Larry Glenn (HBC), and Jeff Ashton (Canadian Tire). We talked about Docker containers in the enterprise and some of our experiences either running containers in production, or the road we're taking to get it there. This was pretty much my first time speaking in a professional capacity, so please be kind.`,
				Date:        YMD("2016-02-10"),
				Link:        "https://containersummit.io/events/nyc-2016/videos/containers-are-eating-the-enterprise",
			},
		},
		Writings: []Writing{
			Writing{
				Title: "Breaking all the rules: Using Go to call Windows API",
				Link: "https://medium.com/jettech/breaking-all-the-rules-using-go-to-call-windows-api-2cbfd8c79724",
				Date:  YMD("2019-01-15"),
				Description: "I go into detail about how to call the Win32 API through Go and some of the traps to avoid when doing so.",
				Publication: "Jet Tech Blog",
				PublicationLink: "https://medium.com/jettech",
			},
			Writing{
				Title: "Containing Windows Executables with Damon",
				Link: "https://medium.com/jettech/containing-windows-executables-with-damon-898ab4a31ea4",
				Date:  YMD("2018-12-11"),
				Description: "I describe an alternative method to constraining Windows executables without using Containers.",
				Publication: "Jet Tech Blog",
				PublicationLink: "https://medium.com/jettech",
			},
			Writing{
				Title: "A short introduction to Windows Containers",
				Link: "https://medium.com/jettech/a-short-introduction-to-windows-containers-db5adc0db536",
				Date:  YMD("2018-12-11"),
				Description: "An overview of how Windows Containers work and some of their features and drawbacks.",
				Publication: "Jet Tech Blog",
				PublicationLink: "https://medium.com/jettech",
			},
			Writing{
				Title: "Using Ansible to Manage Jet Infrastructure",
				Link: "https://medium.com/jettech/using-ansible-to-manage-jet-infrastructure-a3cddd4c6440",
				Date:  YMD("2016-10-07"),
				Description: "A short introduction to Ansible and why we at Jet chose it.",
				Publication: "Jet Tech Blog",
				PublicationLink: "https://medium.com/jettech",
			},

		},
		HigherEducation: []Institution{
			Institution{
				Name:    "New Jersey Institute of Technology",
				Website: "https://www.njit.edu/",
				Start:   YM("2005-09"),
				End:     YM("2009-06"),
				Degree:  "B.S., Computer Engineering",
				GPA:     "3.95/4.0",
				Honors: []string{
					"Member, Phi Eta Sigma honors society",
					"Albert Dorman Honors College (ADHC) Scholar",
					"Summa Cum Laude",
				},
			},
		},
	}
}

// MyCV encapsulates all of the CV data
var MyCV CV

// CV contains all the different experience sections
type CV struct {
	// WorkExperience lists all companies I've worked for and any relevant details:
	// - The role's I've had and what I've done in them
	// - Any notable projects and their descriptions
	WorkExperience []Company `json:"work_experience"`
	// OpenSource lists notable open source projects that I've provided patches to or published either personally or through my employer
	OpenSource []OpenSource `json:"open_source"`
	// HigherEducation lists contains a list of higher-education institutions and degree achieved
	HigherEducation []Institution `json:"higher_education"`
	// SpeakingEvents lists all of the public speaking engagements in which I've been a participant
	SpeakingEvents []SpeakingEvent `json:"speaking_events"`
	// Writings lists some of the publicly available articles and blog posts I've authored
	Writings []Writing `json:"writings"`
}