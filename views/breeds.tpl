<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Breeds</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            padding: 20px;
        }
        .cat-container {
            border: 1px solid #ddd;
            border-radius: 8px;
            overflow: hidden;
            display: inline-block;
            width: 800px;
            text-align: center;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            margin: auto;
        }
        .dropdown {
            margin: 20px 0;
        }
        .carousel-container {
            position: relative;
            max-width: 500px;
            margin: 0 auto;
        }
        .carousel img {
            max-width: 100%;
            border-radius: 8px;
            display: none; /* Hide all images by default */
        }
        .carousel img.active {
            display: block; /* Show only the active image */
        }
        .dots {
            display: flex;
            justify-content: center;
            margin-top: 10px;
            gap: 5px;
        }
        .dot {
            width: 10px;
            height: 10px;
            background-color: #bbb;
            border-radius: 50%;
            cursor: pointer;
        }
        .dot.active {
            background-color: #007BFF;
        }
        .breed-info {
            margin-top: 20px;
            text-align: left;
            max-width: 600px;
            margin-left: auto;
            margin-right: auto;
        }
        .breed-info h2 {
            margin-bottom: 10px;
        }
        .breed-info p {
            margin: 5px 0;
        }
        .breed-info a {
            color: #007BFF;
            text-decoration: none;
        }
        .breed-info a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <h1>Cat Breeds</h1>
    <div class="cat-container">
        <div class="dropdown">
            <select id="breed-selector" onchange="fetchBreedImages()">
                {{range $index, $breed := .Breeds}}
                <option value="{{.id}}" {{if eq $index 0}}selected{{end}}>{{.name}}</option>
                {{end}}
            </select>
        </div>

        <div class="carousel-container">
            <div class="carousel" id="carousel-container">
                <!-- Images will be dynamically loaded here -->
            </div>
            <div class="dots" id="dots-container">
                <!-- Dots will be dynamically loaded here -->
            </div>
        </div>

        <div class="breed-info" id="breed-info">
            <!-- Breed details will be dynamically loaded here -->
        </div>
    </div>

    <script>
        let currentIndex = 0;
        let images = [];
        let timer;

        // Fetch images for the selected breed
        async function fetchBreedImages() {
            const breedSelector = document.getElementById("breed-selector");
            const breedId = breedSelector.value;

            if (!breedId) {
                document.getElementById("carousel-container").innerHTML = "";
                document.getElementById("breed-info").innerHTML = "";
                return;
            }

            try {
                // Fetch images for the selected breed
                const response = await fetch(`https://api.thecatapi.com/v1/images/search?limit=8&size=med&sub_id=demo-0.256737525313076634&breed_id=${breedId}`, {
                    headers: { "x-api-key": "DEMO-API-KEY" },
                });

                if (response.ok) {
                    images = await response.json();
                    const carouselContainer = document.getElementById("carousel-container");
                    const dotsContainer = document.getElementById("dots-container");
                    const breedInfo = document.getElementById("breed-info");

                    // Populate the carousel
                    carouselContainer.innerHTML = images
                        .map((img, index) => `<img src="${img.url}" alt="Cat Image" class="${index === 0 ? "active" : ""}">`)
                        .join("");

                    // Populate dots
                    dotsContainer.innerHTML = images
                        .map((_, index) => `<div class="dot ${index === 0 ? "active" : ""}" onclick="showImage(${index})"></div>`)
                        .join("");

                    // Display breed info
                    const breed = images[0].breeds[0];
                    breedInfo.innerHTML = `
                        <h2>${breed.name}</h2>
                        <p><strong>Origin:</strong> ${breed.origin}</p>
                        <p><strong>ID:</strong> ${breed.id}</p>
                        <p><strong>Description:</strong> ${breed.description}</p>
                        <p><a href="${breed.wikipedia_url}" target="_blank">Learn more on Wikipedia</a></p>
                    `;

                    // Start the carousel rotation
                    clearInterval(timer);
                    timer = setInterval(nextImage, 3000);
                } else {
                    console.error("Failed to fetch breed images:", await response.text());
                }
            } catch (error) {
                console.error("Error fetching breed images:", error);
            }
        }

        // Show a specific image
        function showImage(index) {
            const carouselImages = document.querySelectorAll(".carousel img");
            const dots = document.querySelectorAll(".dot");

            carouselImages[currentIndex].classList.remove("active");
            dots[currentIndex].classList.remove("active");

            currentIndex = index;

            carouselImages[currentIndex].classList.add("active");
            dots[currentIndex].classList.add("active");
        }

        // Show the next image in the carousel
        function nextImage() {
            const nextIndex = (currentIndex + 1) % images.length;
            showImage(nextIndex);
        }

        // Automatically fetch images for the first breed on page load
        document.addEventListener("DOMContentLoaded", fetchBreedImages);
    </script>
</body>
</html>
