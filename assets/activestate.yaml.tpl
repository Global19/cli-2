project: {{.Project}}
{{if ne .LanguageName ""}}
languages: # Please run 'state push' to create your language runtime, once you do the language entry here will be removed
  - name: {{.LanguageName}}
    version: {{.LanguageVersion}}
{{end}}
{{.Content}}