# 🛠 **Category Add Microservice**

The **Category Add Service** is a Go-based microservice using GORM and Gorilla Mux, designed to add new categories to a MySQL database via a RESTful API.

---
## 🐳 **Deployment Docker Image**

Visit the repository on Docker Hub [here](https://hub.docker.com/r/lxgonzalez/category-add-service)

1. **Check if port 80 is free**.
2. **Run the following command in your terminal:**
    ```bash
    > docker pull lxgonzalez/category-add-service

    > docker run -d --name category-add-service -p 8080:8080 lxgonzalez/category-add-service:latest
    ```
## 🚀 **Deployment Locally**

Follow these steps to run the API on your local machine:

1. **Clone the Repository**

   Clone this repository to your local machine:

    ```bash
    git clone https://github.com/lxgonzalez/category-service/
    ```
2. **Install Dependencies**

   Ensure you have Go installed. Run the following commands to install dependencies and build the project:

    ```bash
    > go mod download
    > go build
    ```
 **Run the Application**

3. After configuring your environment variables, run the application:

    ```bash
    > go run main.go
    ```
4. **Sending POST Requests**
To create a new category, send a POST request with JSON data to the following endpoint: [http://localhost:8080/categories](http://localhost:8080/categories)

    **Request Body Example:**
     ```json
     {
        "name": "Electronics"
     }
     ```

    **Example using curl:**
     ``` bash
      curl --location --request POST 'http://localhost:8080/categories' \
      --header 'Content-Type: application/json' \
      --data-raw '{
         "name": "Electronics"
    }'
     ```

---

## 📽️ Evidence






    
   
