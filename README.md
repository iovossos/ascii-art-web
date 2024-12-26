# 🎨 ASCII Art Web Project 🌐

## 📖 Description
Welcome to the **ASCII Art Web Project**, where creativity meets code! 🚀 This web application allows users to transform plain text into stunning **ASCII art** using banner styles like **Standard**, **Shadow**, and **Thinkertoy**. Built with Go 🐹, this project provides a sleek graphical interface for all your artistic needs. 🌟

---

## 👩‍💻 Authors
- **Andriana (astas)**  
  *"Through this project, Andy discovered that coding can be a canvas for creativity, blending logic with artistic expression!"* ✨
- **Dilhan (daslamac)**  
  *"Dilhan unlocked her creative spark, turning programming into an art form that merges technical skill and imagination."* 💡

---

## 🎮 Usage: How to Run
### 1. Clone the Repository to Your Local Machine:
```bash
git clone https://platform.zone01.gr/git/daslamac/ascii-art-web
cd ascii-art-web 
```

### 2. Set Up the Go Backend
Make sure you have Go installed on your machine.

In the project directory, run the Go server:
```bash
go run server.go
OR
go run .
```
This will start the server on http://localhost:8080.

### 3. Open the Application
Once the Go server is running, open your web browser and navigate to:  
http://localhost:8080  

### 4. Use the ASCII Art Generator
- Enter the text in the input field.
- Choose a font style from the dropdown menu (**Standard**, **Shadow**, **Thinkertoy**).
- Click on the **Generate** button to see the ASCII art result displayed on the page.

---

## 🔍 Implementation Details
### 🛠️ Algorithm
1. **Input Handling**:
   - Receive text and banner style via a **POST request**.
   - Validate text input to ensure only ASCII characters are allowed.

2. **ASCII Art Generation**:
   - Load the selected banner file from the `fonts` folder.
   - Map each character of the input text to its ASCII art equivalent.
   - Combine and format the ASCII art for display.

3. **Error Management**:
   - **400 Bad Request**: Invalid or missing text input.
   - **404 Not Found**: Requested resource or page is unavailable.
   - **500 Internal Server Error**: Unexpected issues during processing.

### 🗂️ Folder Structure
```
ascii-art-web/
│  server.go
│  README.md
│  go.mod
├── fonts/
│   ├─ shadow.txt
│   ├─ standard.txt
│   └─ thinkertoy.txt
├── templates/
│   ├─ fonts/
│   │   └─ Meow.ttf
│   ├─ stylings/
│   │   ├─ background.jpg
│   │   ├─ errors.css
│   │   ├─ favicon.png
│   │   └─ style.css
│   ├─ 400.html
│   ├─ 404.html
│   ├─ 500.html
│   ├─ about.html
│   └─ index.html
├── utils/
│   ├─ handler.go
│   ├─ asciiPrint.go

```
---

## 🛣️ HTTP Endpoints
### **GET /**
- Displays the **main page** with the input form and banner options.

### **POST /generate**
- Processes form input and generates ASCII art.
- Renders the result directly on the same page.

### 🛑 Status Codes
- **200 OK**: Successfully processed the request.
- **400 Bad Request**: Invalid or missing input.
- **404 Not Found**: Resource or page does not exist.
- **500 Internal Server Error**: Something went wrong on the server.

---

## 🔮 Future Improvements
- ✨ Add more creative banner styles.
- 🎨 Enhance the UI with animations and better visuals.
- 🕒 Implement a history feature to save and revisit past creations.

---

## 🎉 Get Started and Have Fun!
Dive into the world of ASCII art and let your creativity shine! 🌟🖌️ We hope you enjoy this fun and interactive tool as much as we loved building it. Happy ASCII-ing! 😄
