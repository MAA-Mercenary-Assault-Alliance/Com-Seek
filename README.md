# Com-Seek

A web application for Computer Engineering students and alumni at Kasetsart University to find employments

## Requirement
- Go version 1.25.0 or higher
- Node.js 22+
- Docker (or local MySQL version 9.0.0 or higher)

## Installaion

### Back end

1. Navigate to the back end directory
    
    ```
    cd backend
    ```

2. Install Go dependencies

    ```
    go mod download
   ```

3. Set up environment variables

    ```
    # Windows
    copy sample.env .env

    # Linux/Unix/MacOS
    cp sample.env .env
    ```

    Change the values in .env to match your configuration

4. Set up database

    ```
    docker compose up -d db
    ```

    Or use your local MySQL database

### Front end

1. Navigate to the front end directory

    ```
    cd frontend
    ```

2. Install dependencies

    ```
    npm install
    ```

## Running the Application

### Back end

1. Navigate to the back end directory
    
    ```
    cd backend
    ```

2. Start the database (If using Docker)

    ```
    docker compose up -d db
    ```

3. Run the application

    ```
    go run main.go
    ```

### Front end

1. Navigate to the front end directory

    ```
    cd frontend
    ```

2. Run the application

    ```
    npm run dev
    ```
