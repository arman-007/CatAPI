const originalFetch = window.fetch;

window.fetch = async function(url, options) {
    if (!url || url === "undefined") {
        console.error("Attempted to fetch an undefined URL");
        return Promise.reject(new Error("Undefined URL in fetch"));
    }
    return originalFetch.apply(this, arguments);
};

// Initialize sub_id if not already set
if (!localStorage.getItem("sub_id")) {
    const subId = `demo-${Math.random().toString(36).substr(2, 9)}`;
    localStorage.setItem("sub_id", subId);
}

// Manage Tabs
//checked
function showTab(tab) {
    const tabs = document.querySelectorAll(".tab-content");
    const buttons = document.querySelectorAll(".nav button");

    // Hide all tabs and deactivate all buttons
    tabs.forEach(t => t.classList.add("hidden"));
    buttons.forEach(b => b.classList.remove("active"));

    // Show the selected tab and activate its button
    document.getElementById(tab).classList.remove("hidden");
    document.getElementById(`tab-${tab}`).classList.add("active");

    // Manage auto-slide for breeds
    // if (tab !== "breeds") {
    //     clearInterval(autoSlideInterval);
    // } else {
    //     startAutoSlide();
    // }
}

// Voting Tab
//checked
function renderVotingTab(data) {
    const catImage = document.getElementById("cat-image");

    // Log the received data
    console.log("Data received in renderVotingTab:", data);

    // Parse JSON if data is a string
    if (typeof data === "string") {
        try {
            data = JSON.parse(data);
        } catch (error) {
            console.error("Failed to parse JSON string:", data);
            return;
        }
    }

    // Check if valid data is available
    if (Array.isArray(data) && data.length > 0 && data[0].url) {
        catImage.src = data[0].url; // Use the first object's `url`
    } else {
        console.error("Voting data is not available or invalid format:", data);
        catImage.src = "placeholder.png"; // Path to a default placeholder image
    }
}

//checked
async function vote(type) {
    const subId = localStorage.getItem("sub_id");
    console.log(subId)
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
            renderVotingTab(preloadedData.voting); // Render the next preloaded data
        }
    } catch (error) {
        console.error("Error voting:", error);
    }
}

//checked
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

//checked
function renderBreedsTab(data) {
    const breedSelector = document.getElementById("breed-selector");

    // Parse JSON if data is a string
    if (typeof data === "string") {
        try {
            data = JSON.parse(data);
        } catch (error) {
            console.error("Failed to parse JSON string:", data);
            return;
        }
    }

    // Ensure `data` is an array
    if (!Array.isArray(data)) {
        console.error("Breeds data is not an array:", data);
        breedSelector.innerHTML = `<option value="">No breeds available</option>`;
        return;
    }

    // Populate the dropdown with breed options
    breedSelector.innerHTML = data
        .map(breed => `<option value="${breed.id}">${breed.name}</option>`)
        .join("");

    // Automatically render the first breed's images and info
    if (data.length > 0) {
        renderBreedImagesAndInfo(data[0]);
    }
    startAutoSlide();
}

//checked
async function renderBreedImagesAndInfo(breed) {
    const carouselContainer = document.getElementById("carousel-container");
    const dotNavigation = document.getElementById("dot-navigation");
    const breedInfo = document.getElementById("breed-info");

    try {
        // Fetch images for the breed
        const response = await fetch(`/api/breeds/images?breed_id=${breed.id}`);
        if (!response.ok) {
            throw new Error(`Failed to fetch images for breed: ${breed.id}`);
        }

        const images = await response.json();
        console.log("Fetched breed images:", images);

        // Populate the carousel with images
        if (images.length > 0) {
            carouselContainer.innerHTML = images
                .map(img => `<img src="${img.url}" alt="Cat Image">`)
                .join("");

            // Populate dot navigation
            dotNavigation.innerHTML = images
                .map((_, index) => `<span class="dot" onclick="navigateCarousel(${index})"></span>`)
                .join("");

            // Set the first image and first dot as active
            updateCarousel(0);
        } else {
            console.error("No images available for the selected breed.");
            carouselContainer.innerHTML = "<p>No images available</p>";
            dotNavigation.innerHTML = "";
        }
    } catch (error) {
        console.error("Error fetching breed images:", error);
        carouselContainer.innerHTML = "<p>Error loading images</p>";
        dotNavigation.innerHTML = "";
    }

    // Populate breed info
    breedInfo.innerHTML = `
        <h2>${breed.name}</h2>
        <p><strong>Origin:</strong> ${breed.origin}</p>
        <p><strong>Description:</strong> ${breed.description}</p>
        <p><strong>Wikipedia URL:</strong> <a href="${breed.wikipedia_url}" target="_blank">${breed.wikipedia_url}</a></p>
    `;
}

//checked
async function fetchBreedImages(breedId) {
    // console.log(breedId)
    try {
        if (breedId){
            const response = await fetch(`/api/breeds/images?breed_id=${breedId}`);
            const data = await response.json();
            if (!Array.isArray(data)) throw new Error("Invalid data format");
            // console.log("Fetched breed images:", data);
            return data;
        }
    } catch (error) {
        console.error("Failed to fetch breed images:", error);
        return [];
    }
}

// Carousel Logic
let currentSlide = 0;

//checked ; needs to be reviewed for the images width and not showing the images issue
//checked ; needs to be reviewed for the images width and not showing the images issue
//checked ; needs to be reviewed for the images width and not showing the images issue
function updateCarousel(index) {
    const carouselContainer = document.getElementById("carousel-container");
    const dots = document.querySelectorAll(".dot");
    const images = carouselContainer.querySelectorAll("img");
    // console.log(images)

    if (!images.length) {
        console.error("No images found in the carousel.");
        return;
    }

    if (index < 0 || index >= images.length) {
        console.error("Invalid carousel index:", index);
        return;
    }

    currentSlide = index;

    //here are the issues i guess
    const slideWidth = images[0].getBoundingClientRect().width;
    console.log(slideWidth)
    carouselContainer.style.transform = `translateX(-${index * slideWidth}px)`;
    carouselContainer.style.transition = "transform 0.5s ease-in-out";

    dots.forEach(dot => dot.classList.remove("active"));
    if (dots[index]) {
        dots[index].classList.add("active");
    }
}

//checked ; images are being pulled correctly
function updateBreedCarousel(images) {
    const carouselContainer = document.getElementById("carousel-container");
    const dotNavigation = document.getElementById("dot-navigation");

    // Ensure valid images
    if (!images.length) {
        console.error("No images available for the carousel.");
        return;
    }

    // Populate the carousel with images
    carouselContainer.innerHTML = images
        .map(img => `<img src="${img.url}" alt="Cat Image">`)
        .join("");

    // Populate dots for navigation
    dotNavigation.innerHTML = images
        .map((_, index) => `<span class="dot cursor-pointer" onclick="navigateCarousel(${index})"></span>`)
        .join("");

    // Activate the first slide and dot
    updateCarousel(0);
}

//checked
function navigateCarousel(index) {
    updateCarousel(index);
}

let autoSlideInterval;

//checked
function startAutoSlide() {
    clearInterval(autoSlideInterval);
    autoSlideInterval = setInterval(() => {
        const dots = document.querySelectorAll(".dot");
        const nextSlide = (currentSlide + 1) % dots.length;
        updateCarousel(nextSlide);
    }, 3000);
}

// document.addEventListener("DOMContentLoaded", () => {
//     const initialImages = document.querySelectorAll("#carousel-container img");
//     if (initialImages.length > 0) {
//         updateCarousel(0); // Initialize the first slide
//         startAutoSlide();
//     } else {
//         console.error("No initial images to display in the carousel.");
//     }
// });

// Favorites Tab
function renderFavoritesTab(data) {
    const gallery = document.getElementById("favorites-gallery");

    // Parse JSON if data is a string
    if (typeof data === "string") {
        try {
            data = JSON.parse(data);
        } catch (error) {
            console.error("Failed to parse JSON string:", data);
            return;
        }
    }

    // Validate that data is an array
    if (!Array.isArray(data)) {
        console.error("Expected an array for favorites but got:", data);
        gallery.innerHTML = `<p>No favorites available</p>`;
        return;
    }

    // Populate the gallery with favorite images
    gallery.innerHTML = data
        .map(fav => {
            const url = fav?.image?.url || "placeholder.jpg"; // Fallback to placeholder if URL is missing
            return `<img src="${url}" alt="Favorite Cat">`;
        })
        .join("");
}

// Load preloaded data
document.addEventListener("DOMContentLoaded", () => {
    // Show the default tab
    showTab("voting");

    // Render preloaded data
    renderVotingTab(preloadedData.voting);
    // renderBreedsTab(preloadedData.breeds);
    renderFavoritesTab(preloadedData.favorites);
});

document.addEventListener("DOMContentLoaded", () => {
    const preloadedBreeds = (window.preloadedData && window.preloadedData.breeds) || [];
    if (preloadedBreeds.length === 0) {
        console.error("No breeds data available.");
        return;
    }

    renderBreedsTab(preloadedBreeds);

    const breedSelector = document.getElementById("breed-selector");
    breedSelector.addEventListener("change", async () => {
        const breedId = breedSelector.value;
        if (breedId) {
            // console.log("fetching breed images for ", breedId)
            const images = await fetchBreedImages(breedId);
            // console.log("Fetched breed images:", images);
            updateBreedCarousel(images); // Ensures dots are populated
            startAutoSlide();
        }
    });
});