                 _____    _____   _____   _____                       _____    _______           __          __  ______   ____   
        /\      / ____|  / ____| |_   _| |_   _|              /\     |  __ \  |__   __|          \ \        / / |  ____| |  _ \  
       /  \    | (___   | |        | |     | |    ______     /  \    | |__) |    | |     ______   \ \  /\  / /  | |__    | |_) | 
      / /\ \    \___ \  | |        | |     | |   |______|   / /\ \   |  _  /     | |    |______|   \ \/  \/ /   |  __|   |  _ <  
     / ____ \   ____) | | |____   _| |_   _| |_            / ____ \  | | \ \     | |                \  /\  /    | |____  | |_) | 
    /_/    \_\ |_____/   \_____| |_____| |_____|          /_/    \_\ |_|  \_\    |_|                 \/  \/     |______| |____/

# ASCII Art Generator - v0.1.0

Looking for a bit of nostalgia? This ASCII Art Generator brings back the classic feel of ASCII-styled text, letting users transform plain text into fun, stylized art. Built as a lightweight Go web server, this project showcases the collaborative efforts of **imichalop**, **ivossos**, and **ustikker**, and is perfect for anyone interested in how web servers process input and render output in creative ways.


# Description

This project is a simple web server that transforms text into ASCII art, bringing a bit of retro charm to any input text. Users can enter text on the main page, choose a banner style, and generate an ASCII representation of their input.

The server is built in Go and handles requests to generate ASCII art, serve static assets, and render HTML pages. It uses pre-defined ASCII banners that map characters to stylized art, and it ensures compatibility with standard ASCII inputs.

This project is a fun exercise in handling web requests, file serving, and basic ASCII art manipulation using only the standard Go libraries.

## Usage

 - **Start the Server**

    Run the server by executing the following command in your terminal:
	                          
	                           go run .

	This will start the server on `http://localhost:8080`.



 -   **Access the ASCII Art Generator**  
    Open your web browser and go to `http://localhost:8080`. You’ll see a simple form where you can enter text, choose an ASCII banner style, and submit it to generate ASCII art.
    
-   **Enter Text**  
    Type in the text you’d like to convert into ASCII art. Only ASCII printable characters are supported.
    
-   **Choose a Banner Style**  
    Select one of the available banner styles ( **standard**, **shadow**, **thinkertoy** ). Each style has a unique look, and the generated ASCII art will follow that design.
    
-   **Generate the Art**  
    Click the submit button to generate and display your ASCII art on the page.

-   **Run The Tester**  
    Run 'go test' from the terminal, to use the tester program that tests if the HTTP error codes are returned properly.


The ASCII Art Generator is organized to handle different routes, serve necessary static files, and render ASCII art in response to user input. Here’s a breakdown of how it all works:

1.  **Server Setup and Routing**  
    The main server is initialized in `main.go`, where it sets up routes for:
    
    -   The homepage (`/`), serving the main interface.
    -   The ASCII art generation endpoint (`/ascii-art`), handling user input and generating ASCII art.
    -   Error handling pages for `400`, `404`, and `500` status codes.
    -   Static assets like CSS, images, and templates, allowing the site to have a styled interface.
    
    Once routes are set up, the server is started on `localhost:8080`, ready to respond to requests.
    
2.  **Request Handling and Templating**  
    In `handlers.go`, each route is associated with a specific handler function:
    
    -   The main page loads `index.html`, pulling in necessary stylesheets and initializing the user interface.
    -   For ASCII art generation (`/ascii-art`), the server processes the input text, maps each character to its ASCII art representation, and displays the result on the same page.
    -   If a user encounters an error (such as providing invalid input or an unknown route), the server renders appropriate error pages (`400`, `404`, or `500`), each styled with `styles.css` and any relevant error images (like `404.png`).
3.  **ASCII Art Processing**  
    The core ASCII art generation happens in `asciiart.go`. Here’s how it works:
    
    -   **Text Input Mapping**: The server reads the text input from the user and checks that it contains only ASCII characters.
    -   **Character Conversion**: Each valid character is mapped to its corresponding ASCII art representation using pre-loaded banner files.
    -   **Line Assembly**: The ASCII characters are arranged line by line, creating the final ASCII art output.
    -   **Result Display**: The result is then rendered in `index.html`, showing the ASCII art to the user on the same page.
4.  **Static and Template Files**
    
    -   Templates, such as `index.html`, are used to render the homepage and result page, with `400.html`, `404.html`, and `500.html` for error responses.
    -   Static assets (like CSS and error images) ensure consistent styling and branding across all pages, including error pages.


## Roadmap

1. Add more ASCII fonts.
2. Add the options for alignment (it's already in the program but not implemented in the UI - to change alignment, simply set the class of pre in index.html to either left, center or right. Center is selected by default.)
3. Support live view and multiple output areas at the same time.
4. Button that helps you directly copy / share the ASCII art.
5. Remove the hidden messages from the source code of index.html.

## Authors

1. [iovossos](https://github.com/iovossos)
2. [stikkeruip](https://github.com/stikkeruip)
