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

{{- range $company := .WorkExperience }}

### {{ .Title }} at [{{ .Name }}]({{ .Website }})

{{ .Start }} - {{ if .End.Present }}present{{ else }}{{ .End }}{{end }}
{{ range .Roles }}
{{ if gt ($company.Roles | len) 1 }}#### Role: {{ .Name }}
{{ end -}}
{{ .Description }}
{{ if .Highlights }}{{ range .Highlights }}
- {{ . }}{{ end }}{{ end }}
{{ end }}
{{ if .Projects }}#### Projects
{{ range .Projects }}- **{{ .Name }}** : {{ .Description }}
{{ end }}{{ end }}{{ end }}
---
## Open Source Contributions

Below are a selection of open source projects I've been involved in, either as a creator, or as a contributor. I've done some other minor stuff in other languages, but the majority of my work has been in Go.
{{ range .OpenSource }}
### [{{ .Name }}]({{ .Website }}) ({{- if .Contribution | eq "creator" }}Creator{{end}} 
{{- if .Contribution | ne "creator" }}
{{- if .Link }}[{{ .Contribution | title }} Contribution]({{ .Link }})
{{- else }}{{ .Contribution | title }} Contribution{{ end }}
{{- end}})

{{ .Description }}
{{ end }}
---
## Speaking Events

I've been speaking at a couple of events. If that interests you, check out links to my talks. I've never watched any of these, so hopefully I didn't embarrass myself; And if I did, please don't tell me.
{{ range .SpeakingEvents }}
### ({{ .Date }}) {{ .Event }}

{{- if .Link }}
Talk: **[{{ .Name }}]({{ .Link }})**
{{- else }}
Talk: **{{ .Name }}**
{{- end }}

{{ .Description }}
{{- if .Notes }}

*Note:* {{ .Notes }}{{ end }}
{{ end }}
---
## Writing

I've written a couple of blog posts explaining some of the research and work my team-mates and I have done.
These may give you a better idea of what I think about.
{{ range .Writings }}
### [{{ .Title }}]({{ .Link }})

{{ .Date }} - {{ if .PublicationLink }}[{{ .Publication }}]({{ .PublicationLink }}){{ else }}{{ .Publication }}{{ end }}

{{ .Description }}
{{ end }}
---
## Education

This may not be as important anymore, but I paid for it and did OK; so I get to put it here.
{{ range .HigherEducation }}
### **{{ .Degree }}** from [{{ .Name }}]({{ .Website }})
{{ .Start }} - {{ .End }}

GPA: **{{ .GPA }}**
{{ if .Honors }}
#### Honors{{ range .Honors }}
- {{ . }}{{ end }}{{ end}}{{ end }}