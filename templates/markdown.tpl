# Justen Walker

![Keybase](https://keybase.io/favicon.ico) [justenwalker](https://keybase.io/justenwalker)

![Twitter](https://abs.twimg.com/favicon.ico) [@justenwalker](https://twitter.com/justenwalker)

He/Him

---
## Work Experience

This is my work history starting from most recent all the way back to college. 

TLDR; I started out making websites in PHP. I did a couple of jobs in Java. Eventually, I started working in Go and have not looked back.

{{ range $company := .WorkExperience }}

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

Talk: **[{{ .Name }}]({{ .Link }})**

{{ .Description }}

{{ if .Notes }}*Note:* {{ .Notes }}{{ end }}{{ end }}

---
## Writing

I've written a couple of blog posts explaining some of the research and work my team-mates and I have done.
These may give you a better idea of what I think about.

{{ range .Writings }}
### [{{ .Title }}]({{ .Link }})

{{ .Date }} - {{ if .PublicationLink }}[{{ .Publication }}]({{ .PublicationLink }}){{ else }}{{ .Publication }}{{ end }}

{{ .Description }}{{ end }}

---
## Education

This may not be as important anymore, but I paid for it so I get to put it here.

{{ range .HigherEducation }}
### **{{ .Degree }}** from [{{ .Name }}]({{ .Website }})
{{ .Start }} - {{ .End }}

GPA: **{{ .GPA }}**
{{ if .Honors }}
#### Honors{{ range .Honors }}
- {{ . }}{{ end }}{{ end}}{{ end }}