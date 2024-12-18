{{define "content"}}
<h2>Favorites</h2>
{{range .Favorites}}
<div>
    <p><strong>Favorite ID:</strong> {{.id}}</p>
    <p><strong>Image ID:</strong> {{.image_id}}</p>
</div>
{{end}}
{{end}}
