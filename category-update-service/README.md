# 🛠 **Category Update Microservice**

The **Category Update Service** is a Go-based microservice that allows updating existing categories in a MySQL database via a RESTful API.

---

## 🐳 **Deployment Docker Image**

Visit the repository on Docker Hub [here](https://hub.docker.com/r/lxgonzalez/category-update-service)

1. **Check if port 8080 is free**.
2. **Run the following command in your terminal:**
    ```bash
    > docker pull lxgonzalez/category-update-service

    > docker run -e DATASOURCE_URL=<DATASOURCE_URL> \
                   -e DATASOURCE_PORT=<DATASOURCE_PORT> \
                   -e DATABASE=<DATABASE> \
                   -e DATASOURCE_USERNAME=<DATASOURCE_USERNAME> \
                   -e DATASOURCE_PASSWORD=<DATASOURCE_PASSWORD> \
                   -d --name category-update-service \
                   -p 8080:8080 lxgonzalez/category-update-service:latest
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

3. **Configure the Environment Variables**

   Create a `.env` file in the root of the project and add the following environment variables, replacing with your database credentials:

    ```bash
    DATASOURCE_URL=localhost
    DATASOURCE_PORT=3306
    DATABASE=your_database_name
    DATASOURCE_USERNAME=your_database_user
    DATASOURCE_PASSWORD=your_database_password
    ```

4. **Run the Application**

    After configuring your environment variables, run the application:

    ```bash
    > go run main.go
    ```
5. **Sending PUT Requests**

    To update an existing category, send a PUT request with the updated category data to the following endpoint: [http://localhost:8080/categories](http://localhost:8080/categories)

    **Request Body Example:**
    ```json
   {
      "idCategory": 1,
      "name": "Updated Category Name"
   }
    ```

    **Example using curl:**
   ```bash
   curl --location --request PUT 'http://localhost:8080/categories' \
    --header 'Content-Type: application/json' \
    --data-raw '{
     "idCategory": 1,
     "name": "Updated Category Name"
      }'
    ```

   ---
   ## 📽️ Evidence
   ![image](https://github.com/user-attachments/assets/42c081dc-16f0-4961-a42e-c582cd62555a)


   
