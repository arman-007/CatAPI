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

        // Populate the carousel with images
        const carouselContainer = document.getElementById("carousel-container");
        carouselContainer.innerHTML = images
            .map(img => `<img src="${img.url}" alt="Cat Image">`)
            .join("");

        // Populate dot navigation
        const dotNavigation = document.getElementById("dot-navigation");
        dotNavigation.innerHTML = images
            .map((_, index) => `<span class="dot" onclick="navigateCarousel(${index})"></span>`)
            .join("");

        // Set the first image and first dot as active
        updateCarousel(0);

        // Populate breed info
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

let currentSlide = 0;

function updateCarousel(index) {
    const carouselContainer = document.getElementById("carousel-container");
    const dots = document.querySelectorAll(".dot");

    // Update the slide position
    carouselContainer.style.transform = `translateX(-${index * 100}%)`;

    // Update the active dot
    dots.forEach(dot => dot.classList.remove("active"));
    dots[index].classList.add("active");

    currentSlide = index;
}

function navigateCarousel(index) {
    updateCarousel(index);
}

let autoSlideInterval;

function startAutoSlide() {
    clearInterval(autoSlideInterval);
    autoSlideInterval = setInterval(() => {
        const dots = document.querySelectorAll(".dot");
        const nextSlide = (currentSlide + 1) % dots.length;
        updateCarousel(nextSlide);
    }, 3000); // Change slide every 3 seconds
}

document.addEventListener("DOMContentLoaded", () => {
    startAutoSlide();
});



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