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

    Change the values in .env to match your configuration.
    
    For email config, refer to [this](./backend/email_config.md).
    
    For reCAPTCHA config, please create one [here](https://www.google.com/recaptcha/admin/create) if you don't have them. Select V2 checkbox when creating them.

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
3. Set up environment variables

    ```
    # Windows
    copy sample.env .env

    # Linux/Unix/MacOS
    cp sample.env .env
    ```

    Change the values in .env to match your configuration.
    
    For reCAPTCHA config, please create one [here](https://www.google.com/recaptcha/admin/create) if you don't have them. Select V2 checkbox when creating them.

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

## Interacting with docker databases

1. go in your running docker container
2. go to "exec" tab
3. type the command

    ```
    mysql -u root -p
    ```
4. type your password which is in your .env file
```
(DB_ROOT_PASSWORD=???)
default is "mysqlrootpass"
```
5. type the command to use your database
```
show databases; # you should see "mysqldb" or your custom database name in .env file
(DB_NAME=mysqldb)
```

```
use mysqldb;
```
6. now you can interact with your database
```
show tables;

for example
select * from companies;
insert into users (email, password) values ('test@email', 'password');
```
