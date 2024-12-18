{{define "content"}}
<h2>Breeds</h2>
{{range .Breeds}}
<div>
    <p><strong>Name:</strong> {{.name}}</p>
    <p><strong>Description:</strong> {{.description}}</p>
</div>
{{end}}
{{end}}
