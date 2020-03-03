{{ range .Events }}
*{{ .Summary }}*
{{ .Description }}
{{ .Location }}
{{ end }}