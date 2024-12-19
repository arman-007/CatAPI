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
            width: 800px;
            height: 600px;
            text-align: center;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            margin: auto;
        }
        .cat-card img {
            width: 100%;
            height: 90%;
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
        async function vote(voteType) {
            const catImage = document.getElementById("cat-image");
            const imageId = catImage.src.split('/').pop().split('.')[0]; // Extract image ID from URL

            const payload = {
                image_id: imageId,               // The ID of the cat image
                sub_id: localStorage.getItem("sub_id") || `demo-${Math.random().toString(36).substr(2, 9)}`, // Generate or use stored sub_id
                value: voteType === "up"         // true for upvote, false for downvote
            };

            try {
                const response = await fetch("https://api.thecatapi.com/v1/votes", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        "x-api-key": "DEMO-API-KEY" // Use the provided API key
                    },
                    body: JSON.stringify(payload),
                });

                if (response.ok) {
                    console.log("Vote registered successfully with The Cat API.");
                    fetchNextCat(); // Load the next cat image
                } else {
                    const errorText = await response.text();
                    console.error("Failed to register vote with The Cat API:", errorText);
                }
            } catch (error) {
                console.error("Error sending vote to The Cat API:", error);
            }
        }

        // Function to handle favorite action
        async function favorite() {
            const catImage = document.getElementById("cat-image");
            const imageId = catImage.src.split('/').pop().split('.')[0]; // Extract image ID from URL

            const payload = {
                image_id: imageId, // ID of the currently displayed cat image
                sub_id: localStorage.getItem("sub_id") || `demo-${Math.random().toString(36).substr(2, 9)}` // Unique user/session identifier
            };

            try {
                const response = await fetch("https://api.thecatapi.com/v1/favourites", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        "x-api-key": "DEMO-API-KEY" // Replace with your actual API key
                    },
                    body: JSON.stringify(payload),
                });

                if (response.ok) {
                    const data = await response.json();
                    console.log("Image favorited successfully:", data);
                    fetchNextCat();
                } else {
                    const errorText = await response.text();
                    console.error("Failed to favorite image:", errorText);
                    alert("Failed to favorite the image. Please try again.");
                }
            } catch (error) {
                console.error("Error favoriting image:", error);
                alert("An error occurred while trying to favorite the image.");
            }
        }

        // Function to fetch the next cat image
        async function fetchNextCat() {
            try {
                const response = await fetch("https://api.thecatapi.com/v1/images/search?limit=1", {
                    headers: {
                        "x-api-key": "DEMO-API-KEY" // Use the same API key
                    }
                });

                if (response.ok) {
                    const data = await response.json();
                    const newImageUrl = data[0].url;

                    // Update the cat image in the DOM
                    const catImage = document.getElementById("cat-image");
                    catImage.src = newImageUrl;
                } else {
                    console.error("Failed to fetch the next cat image.");
                }
            } catch (error) {
                console.error("Error fetching the next cat image:", error);
            }
        }

    </script>
</body>
</html>
