# 🛠 **Category List Microservice**

The **Category List Service** is a Go-based microservice using GORM and Gorilla Mux, designed to list categories from a MySQL database via a RESTful API.

---

## 🐳 **Deployment Docker Image**

Visit the repository on Docker Hub [here](https://hub.docker.com/r/lxgonzalez/category-list-service)

1. **Check if port 8080 is free.**
2. **Run the following command in your terminal:**
   
   ```bash
    docker pull lxgonzalez/category-list-service

    docker run -e DATASOURCE_URL=<DATASOURCE_URL> \
           -e DATASOURCE_PORT=<DATASOURCE_PORT> \
           -e DATABASE=<DATABASE> \
           -e DATASOURCE_USERNAME=<DATASOURCE_USERNAME> \
           -e DATASOURCE_PASSWORD=<DATASOURCE_PASSWORD> \
           -d --name category-list-service \
           -p 8080:8080 lxgonzalez/category-list-service:latest
    ```
   
   ---
## 🚀**Deployment Locally**

Follow these steps to run the API on your local machine:

1. **Clone the Repository**

    ```bash
    git clone: https://github.com/lxgonzalez/category-service/
    ```

2. **Install Dependencies**
   
   Ensure you have Go installed. Run the following commands to install dependencies and build the project:
   
    ```bash
    > go mod download
    > go build
    ```
    
3. **Connecting to the Service**

    Once the application is running, you can access the service by opening your browser and navigating to:
    http://localhost:8080


4. **Sending DELETE Requests**
   
   To retrieve the list of categories or a specific category by ID, use the following cURL commands:

   **- Get all categories:**
   
   To get the list of all categories, use the following cURL command:
   
    ```bash
    curl --location --request GET 'http://localhost:8080/categories'
    ```

    **- Get one category by ID:**
   
   To get a specific category by its ID, use the following cURL command, replacing {idCategory} with the actual ID of the category you want to retrieve:
    ```bash
    curl --location --request GET 'http://localhost:8080/categories/{idCategory}'
     ```
   
   

    
   
