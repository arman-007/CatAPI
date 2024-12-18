<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Voting</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            padding: 20px;
        }
        .cat-card {
            border: 1px solid #ddd;
            border-radius: 8px;
            overflow: hidden;
            display: inline-block;
            width: 300px;
            text-align: center;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            margin: auto;
        }
        .cat-card img {
            width: 100%;
            height: auto;
        }
        .button-container {
            display: flex;
            justify-content: center;
            gap: 10px;
            margin: 15px 0;
        }
        button {
            padding: 10px 15px;
            border: none;
            border-radius: 5px;
            background-color: #007BFF;
            color: white;
            cursor: pointer;
            font-size: 16px;
        }
        button:hover {
            background-color: #0056b3;
        }
        .fav-btn {
            background-color: #FFC107;
        }
        .fav-btn:hover {
            background-color: #e0a800;
        }
    </style>
</head>
<body>
    <h1>Cat Voting</h1>
    <div class="cat-container">
        {{range .Votes}}
        <div class="cat-card" id="cat-card">
            <img id="cat-image" src="{{.url}}" alt="Cat Image">
            <div class="button-container">
                <button onclick="vote('up')">👍 Upvote</button>
                <button onclick="vote('down')">👎 Downvote</button>
                <button class="fav-btn" onclick="favorite()">⭐ Favorite</button>
            </div>
        </div>
        {{end}}
    </div>

    <script>
        // Function to handle voting
        function vote(voteType) {
            // Fetch the next cat image
            fetchNextCat();
        }

        // Function to handle favorite action
        function favorite() {
            console.log("Cat added to favorites!");
        }

        // Function to fetch the next cat image
        async function fetchNextCat() {
            try {
                const response = await fetch('/voting'); // Call backend API to get the next image
                if (response.ok) {
                    const html = await response.text();
                    const parser = new DOMParser();
                    const doc = parser.parseFromString(html, "text/html");

                    // Replace the cat card content
                    const newCatCard = doc.getElementById("cat-card");
                    const oldCatCard = document.getElementById("cat-card");
                    oldCatCard.innerHTML = newCatCard.innerHTML;
                } else {
                    console.error("Failed to fetch next cat image.");
                }
            } catch (error) {
                console.error("Error fetching next cat:", error);
            }
        }
    </script>
</body>
</html>
