<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat API App</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            padding: 20px;
        }
        .tabs {
            display: flex;
            justify-content: center;
            gap: 20px;
            margin-bottom: 20px;
        }
        .tabs button {
            padding: 10px 20px;
            border: none;
            border-radius: 5px;
            background-color: #007BFF;
            color: white;
            cursor: pointer;
        }
        .tabs button.active {
            background-color: #0056b3;
        }
        .tabs button:hover {
            background-color: #0056b3;
        }
        .tab-content {
            display: none;
        }
        .tab-content.active {
            display: block;
        }
        .gallery img, .carousel img {
            max-width: 100%;
            border-radius: 8px;
        }
        .gallery {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
            gap: 20px;
        }
    </style>
</head>
<body>
    <h1>Cat API App</h1>
    <div class="tabs">
        <button id="tab-voting" onclick="showTab('voting')" class="active">Voting</button>
        <button id="tab-breeds" onclick="showTab('breeds')">Breeds</button>
        <button id="tab-favs" onclick="showTab('favs')">Favorites</button>
    </div>

    <!-- Voting Section -->
    <div id="voting" class="tab-content active">
        <h2>Voting</h2>
        <div id="voting-content">
            <img id="cat-image" src="" alt="Vote for this cat">
            <div>
                <button onclick="vote('up')">👍 Upvote</button>
                <button onclick="vote('down')">👎 Downvote</button>
                <button onclick="favorite()">⭐ Favorite</button>
            </div>
        </div>
    </div>

    <!-- Breeds Section -->
    <div id="breeds" class="tab-content">
        <h2>Breeds</h2>
        <div>
            <select id="breed-selector" onchange="fetchBreedImages()">
                <!-- Options will be dynamically populated -->
            </select>
        </div>
        <div class="carousel" id="carousel-container"></div>
        <div id="breed-info"></div>
    </div>

    <!-- Favorites Section -->
    <div id="favs" class="tab-content">
        <h2>Favorites</h2>
        <div class="gallery" id="favorites-gallery">
            <!-- Favorites will be dynamically populated -->
        </div>
    </div>

    <script>
        // Initialize
        if (!localStorage.getItem("sub_id")) {
            const subId = `demo-${Math.random().toString(36).substr(2, 9)}`;
            localStorage.setItem("sub_id", subId);
        }

        // Manage Tabs
        function showTab(tab) {
            const tabs = document.querySelectorAll('.tab-content');
            const buttons = document.querySelectorAll('.tabs button');

            tabs.forEach(t => t.classList.remove('active'));
            buttons.forEach(b => b.classList.remove('active'));

            document.getElementById(tab).classList.add('active');
            document.getElementById(`tab-${tab}`).classList.add('active');

            // Load content for the selected tab
            if (tab === 'voting') {
                fetchCatForVoting();
            } else if (tab === 'breeds') {
                fetchBreedList();
            } else if (tab === 'favs') {
                fetchFavorites();
            }
        }

        // Voting Logic
        async function fetchCatForVoting() {
            try {
                const response = await fetch("/api/voting/cat");
                const data = await response.json();
                const catImage = document.getElementById("cat-image");
                catImage.src = data[0].url;
            } catch (error) {
                console.error("Error fetching cat for voting:", error);
            }
        }

        async function vote(type) {
            const subId = localStorage.getItem("sub_id");
            const catImage = document.getElementById("cat-image");
            const imageId = catImage.src.split('/').pop().split('.')[0];

            const payload = { image_id: imageId, sub_id: subId, value: type === "up" };

            try {
                const response = await fetch("/api/voting/vote", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify(payload),
                });
                if (response.ok) {
                    console.log("Vote registered!");
                    fetchCatForVoting(); // Fetch next cat
                }
            } catch (error) {
                console.error("Error voting:", error);
            }
        }

        async function favorite() {
            const subId = localStorage.getItem("sub_id");
            const catImage = document.getElementById("cat-image");
            const imageId = catImage.src.split('/').pop().split('.')[0];

            const payload = { image_id: imageId, sub_id: subId };

            try {
                const response = await fetch("/api/favorites", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify(payload),
                });
                if (response.ok) {
                    console.log("Favorite added!");
                }
            } catch (error) {
                console.error("Error favoriting:", error);
            }
        }

        // Breeds Logic
        async function fetchBreedList() {
            const breedSelector = document.getElementById("breed-selector");
            if (breedSelector.options.length > 0) return; // Prevent duplicate fetching

            try {
                const response = await fetch("/api/breeds");
                const data = await response.json();

                breedSelector.innerHTML = data.map(breed => `<option value="${breed.id}">${breed.name}</option>`).join("");
                fetchBreedImages();
            } catch (error) {
                console.error("Error fetching breeds:", error);
            }
        }

        async function fetchBreedImages() {
            const breedSelector = document.getElementById("breed-selector");
            const breedId = breedSelector.value;

            try {
                const response = await fetch(`/api/breeds/images?breed_id=${breedId}`);
                const images = await response.json();

                const carouselContainer = document.getElementById("carousel-container");
                carouselContainer.innerHTML = images
                    .map(img => `<img src="${img.url}" alt="Cat Image">`)
                    .join("");

                const breed = images[0].breeds[0];
                const breedInfo = document.getElementById("breed-info");
                breedInfo.innerHTML = `
                    <h2>${breed.name}</h2>
                    <p><strong>Origin:</strong> ${breed.origin}</p>
                    <p><strong>Description:</strong> ${breed.description}</p>
                `;
            } catch (error) {
                console.error("Error fetching breed images:", error);
            }
        }

        // Favorites Logic
        async function fetchFavorites() {
            const subId = localStorage.getItem("sub_id");
            const gallery = document.getElementById("favorites-gallery");

            try {
                const response = await fetch(`/api/favorites?sub_id=${subId}`);
                const data = await response.json();
                gallery.innerHTML = data.map(fav => `<img src="${fav.image.url}" alt="Favorite Cat">`).join("");
            } catch (error) {
                console.error("Error fetching favorites:", error);
            }
        }

        document.addEventListener("DOMContentLoaded", fetchCatForVoting);
    </script>
</body>
</html>
