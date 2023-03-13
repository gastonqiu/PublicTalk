{{- range $item := .Item -}} 
Remaining {{ $item.Count }} {{ $item.Name }}
{{ end -}}