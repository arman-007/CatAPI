<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
    <header>
        <h1>Cat App</h1>
        <nav>
            <a href="/voting">Voting</a> |
            <a href="/breeds">Breeds</a> |
            <a href="/favs">Favorites</a>
        </nav>
    </header>
    <main>
        {{.Content}}
    </main>
</body>
</html>
