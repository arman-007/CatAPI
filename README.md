# **CatAPI**

A Beego-based web application for interacting with TheCatAPI, enabling users to vote, favorite, and explore cat breeds with carousel support.

## **Features**

- **Voting**: Submit votes for your favorite cat images
- **Breeds**: Explore different cat breeds with images and detailed information
- **Favorites**: Mark cat images as favorites and view them in a gallery
- **Carousel**: Smooth sliding image carousel for breed images
- **Testing**: Comprehensive unit tests with mock HTTP client integration

## **Prerequisites**

- Go (>= 1.19)
- Beego Framework (v2)
- Docker (optional)

## **Steps to Set Up the Project**

1. **Clone the repository:**
```bash
git clone https://github.com/arman-007/CatAPI.git
cd CatAPI
```

2. **Install Dependencies** Use `go mod` to install required dependencies:
```bash
go mod tidy
```

3. **Set Up Environment Variables** Create a `app.conf` file in the `conf` folder and configure the following:
```bash
CAT_API_KEY=your-cat-api-key
PORT=8080
```

## **Testing**
Generate and view a test coverage report:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## **API Endpoints**

### **Breeds**

- **GET /api/breeds**
  - Retrieves all cat breeds.
- **GET /api/breeds/images?breed_id={id}**
  - Retrieves images for a specific breed.

### **Voting**

- **POST /api/voting/submit**
  - Submits a vote for a cat image.

### **Favorites**

- **POST /api/favorites**
  - Adds an image to the favorites list.
- **GET /api/favorites**
  - Retrieves all favorite images.


## **Project Structure**


```plaintext
CatAPI/
├── conf/              # Configuration for the project
├── controllers/       # API controllers (e.g., VotingController, BreedsController)
├── models/            # Data models (if applicable)
├── static/            # Static files (CSS, JS, Images)
├── views/             # HTML templates
├── tests/             # Unit test files
├── main.go            # Entry point of the application
├── go.mod             # Go module file
├── README.md          # Project documentation
└── .env               # Environment configuration