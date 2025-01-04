# Step 1: Use the official Go image
FROM golang:1.23 

# Step 2: Set destination for COPY
WORKDIR /app

# Step 3: Download Go modules
COPY go.mod ./
RUN go mod download

# Step 4: Copy the source code 
COPY . ./

# Step 5: Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Step 6: Expose port 8080
EXPOSE 8080

# Step 7: Run the executable
CMD ["/docker-gs-ping"]