# 🛠 **Category Delete Microservice**

The **Category Delete Service** is a Go-based microservice designed to delete category records via a RESTful API.

---
## 🐳 **Deployment Docker Image**

Visit the repository on Docker Hub [here](https://hub.docker.com/r/lxgonzalez/category-delete-service)

1. **Check if port 8080 is free**.
2. **Run the following command in your terminal**, replace the environment variables with your actual database credentials
   
    ```bash
    > docker pull lxgonzalez/category-delete-service
    > docker run -e DATASOURCE_URL=<DATASOURCE_URL> \
                       -e DATASOURCE_PORT=<DATASOURCE_PORT> \
                       -e DATABASE=<DATABASE> \
                       -e DATASOURCE_USERNAME=<DATASOURCE_USERNAME> \
                       -e DATASOURCE_PASSWORD=<DATASOURCE_PASSWORD> \
                       -d --name category-delete-service \
                       -p 8080:8080 lxgonzalez/category-delete-service:latest
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

    Once the application is running, you can access the service by opening your browser and navigating to:   http://localhost:8080


4. **Sending DELETE Requests**

    To delete a category, send a DELETE request to the following endpoint: http://localhost:8080/categories/{idCategory}

 5. **Example using curl:**
  
    ```bash
    curl --location --request DELETE 'http://localhost:8080/categories/1'
    ```

---

## 📽️ Evidence
![image](https://github.com/user-attachments/assets/571949bf-bd5c-4875-a4d7-38425b9f5cf6)

