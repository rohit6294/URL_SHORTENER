# URL Shortener

A simple URL shortener application built with Go and HTML. This project allows users to shorten long URLs and provides a short URL that redirects to the original URL.

## Features

- Shorten long URLs
- Redirect short URLs to the original URLs
- Simple and clean web interface

## Getting Started

### Prerequisites

- Go (https://golang.org/doc/install)
- Git (https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/url-shortener.git
    cd url-shortener
    ```

2. Initialize the Go module:

    ```bash
    go mod init url-shortener
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

4. Open a web browser and go to `http://localhost:3000`.

## Deployment

### Deploying on Heroku

1. Login to Heroku:

    ```bash
    heroku login
    ```

2. Create a new Heroku app:

    ```bash
    heroku create
    ```

3. Add a `Procfile` to the root of your project directory with the following content:

    ```text
    web: go run main.go
    ```

4. Commit your changes and push to Heroku:

    ```bash
    git add .
    git commit -m "Initial commit"
    git push heroku master
    ```

5. Open your app on Heroku:

    ```bash
    heroku open
    ```

## Usage

1. Enter a long URL in the input field and click the "Shorten" button.
2. The shortened URL will be displayed below the input field.
3. Use the shortened URL to redirect to the original URL.

## Project Structure

```plaintext
url-shortener/
├── main.go         # Go application code
├── index.html      # HTML file for the web interface
├── go.mod          # Go module file
└── Procfile        # Heroku process file
