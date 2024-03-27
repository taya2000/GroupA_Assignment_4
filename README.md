Weather API Server

Overview
This Go-based API server provides current weather information for a specified city. It handles both GET and POST requests at the /city endpoint and returns weather data in JSON format. Integration with the OpenWeather external API enables the retrieval of real-time weather data.

Objectives
Develop a RESTful API Server in Go to handle GET and POST requests.
Integrate the OpenWeather external API to fetch current weather data.
Containerize the application using Docker for easy deployment.
Manage the codebase and track changes using GitHub.

API Specification
GET /city: Accepts a city name as a query parameter and returns the current weather in JSON format.
POST /city: Accepts a JSON body with a city name and returns the current weather in JSON format.


Docker
Dockerfile: Contains instructions to build a Docker image for the application.
Build and Test: Build the Docker image and test it locally.
Push to Docker Hub: Push the Docker image to Docker Hub for deployment.


GitHub
Repository Initialization: Initialize a new repository for the project.
Project Files: Add Go source code, Dockerfile, and README.md to the repository.

Running the Application
1.Ensure you have Go and Docker installed on your system.
2.Clone the repository from GitHub.
3.Build the Docker image using the provided Dockerfile.
docker build -t weather-api-server .

Run the Docker container.
docker run -p 3001:3001 weather-api-server
Access the API endpoints using http://localhost:3001/city.
External API Used
The application integrates with the OpenWeather external API to fetch current weather data for specified cities.

GitHub Repository
https://github.com/taya2000/GroupA_Assignment_4