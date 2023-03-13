{{ range $item := .Items -}}
{{ if ne $item.Count 0 -}}
Remaining {{ $item.Count }}{{ Quantifier $item.Name }} {{ $item.Name }}
{{ end -}}
{{ end -}}