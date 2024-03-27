# Group A - Assignment 4 - Weather API Server

## Overview
This Go-based API server provides current weather information for a specified city. It handles both GET and POST requests at the /city endpoint and returns weather data in JSON format. Integration with the OpenWeather external API enables the retrieval of real-time weather data.

## Objectives
* Develop a RESTful API Server in Go to handle GET and POST requests.
* Integrate the OpenWeather external API to fetch current weather data.
* Containerize the application using Docker for easy deployment.
* Manage the codebase and track changes using GitHub.
* Team Collaboration.
* Merging the code from different team members and resolving the conflicts before pushing to the repository.

## Components in the project
* `main.go`: This is the application's starting point. It includes the main() method, which is where any Go program begins to run and the logic which is used for the API server.
* `go.mod`: Go modules offer a dependency management framework.The project's information, including the versions of its dependencies, is contained in the go.mod file.
* `DockerFile`: This file includes instructions on how to build a Docker image. It outlines the requirements and environment needed to launch a program.
* `README.md`: The file which contains the primary documentation and introduction of the project.

## API Specification
* Get Request 
  * Method: GET
  * Endpoint: /city
  * Description: Accepts a city name as a query parameter and returns the current weather in JSON format
  * Request Parameter: `name` The name of the city
  * Example request: `http://localhost:3001/city?name=Markham`
  * Example response:
    ```
    {
    	"city":"Markham",
    	"temperature":"6.9°C",
    	"feelslike":"4.0°C",
    	"weather":"Clouds"
    }
* POST
  * Method: POST
  * Endpoint: /city
  * Description: Accepts a JSON body with a city name and returns the current weather in JSON format
  * Request Body Parameter:
    ```
    {
    	"name":"Markham"
    }
  * Example request: `http://localhost:3001/city` with body `{"name": "Markham"}`
  * Example response:
    ```
    {
    	"city":"Markham",
    	"temperature":"6.9°C",
    	"feelslike":"4.0°C",
    	"weather":"Clouds"
    }

## Steps Used to Create Docker Image
* Create a Dockerfile which contains the instructions to build a Docker image for the application.
* Build and Test: Build the Docker image and test it locally.
  `docker build -t <image_name> . `
* Tag the docker image.
  `docker tag <image_name> <dockerhub_username>/<repository_name>:<tag>`

## Steps Used to Push the Docker Image to Docker Hub
* Login to docker hub.
  `docker login`
* Push the docker image to Docker Hub for deployment.
  `docker push <dockerhub_username>/<repository_name>:<tag>`

## Getting Started 
* Ensure you have Go and Docker installed on your system.
* Clone the repository from GitHub.
```
git clone https://github.com/taya2000/GroupA_Assignment_4.git
```
* Build the Docker image using the provided Dockerfile.
```
docker build -t weather-api-server .
```
* Run the Docker container.
```
docker run -p 3001:3001 weather-api-server
```
* Curl commands to access the APIs

```
curl -X GET http://localhost:3001/city?name=Markham
curl -X POST http://localhost:3001/city -H "Content-Type: application/json" -d '{"name":"Markham"}'
```

## OpenWeather API Integration

* Purpose: We are able to obtain current weather data for several areas worldwide with the OpenWeather API integration in our project.
* Features: A variety of endpoints are available via the OpenWeather API to retrieve different kinds of weather information, including predictions, historical data, and the present conditions.
* API Key: In order to access the API endpoints, users need to receive an API key from OpenWeather. This key is used to monitor usage limits and authenticate requests.
* Response: The JSON format that the API typically delivers data in makes it simple to read and include into apps.
* Integration: HTTP queries are used to query the relevant API endpoints in order to implement the OpenWeather API integration.
* Data Handling: Our program parses and processes the weather data that has been retrieved to extract pertinent information for display or additional analysis.

We improve our project's usefulness by giving customers useful weather information that is catered to their requirements and preferences by integrating the OpenWeather API into it.

## Version History

* Initial Version - V1

## Contributors

We collaborated on this project as a team. Following are our members:

* Tejaswi Cheripally
* Pamodi Bhagya Rathnayake
* Samhitha Dubbaka
* Abhisheik Yadla
* Mohammed Abdul Bari waseem
* Rohit
* Syed Abdul Qadeer
* Shubham Bathla
* Mandeep Bajwa
* Kamalpreet Kaur
