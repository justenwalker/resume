# Justen Walker

**Mastodon** <a title="Eugen Rochko &amp; other Mastodon contributors, AGPL &lt;http://www.gnu.org/licenses/agpl.html&gt;, via Wikimedia Commons" href="https://commons.wikimedia.org/wiki/File:Mastodon_logotype_(simple)_new_hue.svg"><img width="16" alt="Mastodon logotype (simple) new hue" src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/d5/Mastodon_logotype_%28simple%29_new_hue.svg/16px-Mastodon_logotype_%28simple%29_new_hue.svg.png" style="margin-bottom: -0.25em;"></a> [@jwalk@mastodon.social](https://mastodon.social/@jwalk")

**matrix.org:** <a title="™/®Matrix.org, Public domain, via Wikimedia Commons" href="https://commons.wikimedia.org/wiki/File:Matrix_icon.svg"><img width="16" alt="Matrix icon" src="https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/Matrix_icon.svg/32px-Matrix_icon.svg.png" style="background-color: #fff; padding: 2px; border-radius: 2px; margin-bottom: -0.5em;"></a> [@justenwalker:matrix.org](https://matrix.to/#/@justenwalker:matrix.org)

**BlueSky** <a href="https://en.wikipedia.org/wiki/File:Bluesky_App_Icon.png#/media/File:Bluesky_App_Icon.png"><img src="https://upload.wikimedia.org/wikipedia/en/2/2b/Bluesky_App_Icon.png" alt="Bluesky App Icon" height="16" width="16" style="padding: 2px; border-radius: 2px; margin-bottom: -0.5em;"></a> [@justenwalker.bsky.social](https://bsky.app/profile/justenwalker.bsky.social)

He/Him

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