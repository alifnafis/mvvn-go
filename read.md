
# Music Playlist Management System

A RESTful API service built with **Go** and **PostgreSQL** for managing users, music tracks, and personalized playlists. Containerized with Docker for seamless deployment and scalability.

---

## Features ✨

### 🔑 **User Management**

* Secure user registration and login using **JWT authentication**
* Store user details such as name, username, position, and salary

### 🎵 **Music Catalog**

* Add and manage music tracks with metadata including:
  * Title
  * Artist
  * Album
  * Genre
  * Duration

### 🎶 **Playlist System**

* Create and manage personalized playlists
* Associate playlists with users
* Add multiple music tracks to a playlist

### 🐳 **Dockerized Deployment**

* Fully containerized for hassle-free setup
* PostgreSQL integration for reliable database management

---

## Prerequisites 📋

Before starting, ensure you have the following installed:

* **Docker** & **Docker Compose**
* **Go** (version 1.19 or later) for local development
* An API client such as **Postman** or **Insomnia** for testing

---

## Quick Start 🚀

### 1. Clone the Repository

Clone the project and navigate to its directory:

```
git clone [your-repository-url]
cd project-directory
```

### 2. Build and Run the Docker Containers

Build and start the application with Docker Compose:

```
docker compose build
docker compose up
```

The application will be available at `<span>http://localhost:8080</span>` by default.

### 3. Configure the Database

Run the following SQL commands to set up the required database tables:

```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    position VARCHAR(255) NOT NULL,
    salary INT NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE music (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    artist VARCHAR NOT NULL,
    album VARCHAR,
    genre VARCHAR,
    duration INT
);

CREATE TABLE playlists (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE playlist_music (
    playlist_id INT NOT NULL,
    music_id INT NOT NULL,
    PRIMARY KEY (playlist_id, music_id),
    FOREIGN KEY (playlist_id) REFERENCES playlists(id),
    FOREIGN KEY (music_id) REFERENCES music(id)
);
```

---

## API Endpoints 🔧

### User Management

* **POST /signup** - Register a new user

  ```
  {
      "name": "John Doe",
      "username": "john",
      "position": "Software Engineer",
      "salary": 80000,
      "password": "1234"
  }

  ```
* **POST /login** - Authenticate user and retrieve JWT token

  ```
  {
      "username": "rty",
      "password": "1234"
  }

  ```

### Music Catalog

* **POST /music** - Add a new music track

  ```
  {
    "title": "music",
    "artist": "music",
    "album": "music",
    "genre": "music",
    "duration": 533
  }

  ```
* **GET /music** - Retrieve all music tracks

  ```
  {
      "app": "go-crud",
      "version": "1.0.0",
      "date": "2025-01-24T22:26:08+07:00",
      "data": [
          {
              "id": 1,
              "title": "music",
              "artist": "music",
              "album": "music",
              "genre": "music",
              "duration": 533
          }
      ],
      "message": "Success"
  }
  ```
* **PUT /music/{id}** - Update music metadata

  ```
  {
    "title": "musics",
    "artist": "musics",
    "album": "musics",
    "genre": "musics",
    "duration": 533
  }

  ```
* **DELETE /music/{id}** - Delete a music track

### Playlists

* **POST /playlist** - Create a new playlist

  ```
  {
    "name": "Playlist",
    "user_id": 1
  }

  ```
* **GET /playlist/{user_id}** - Retrieve my playlist

  ```
  {
      "app": "go-crud",
      "version": "1.0.0",
      "date": "2025-01-24T16:24:30Z",
      "data": [
          {
              "id": 1,
              "name": "Code Playlist",
              "user_id": 1,
              "music": [
                  {
                      "id": 1,
                      "title": "music",
                      "artist": "music",
                      "album": "music",
                      "genre": "music",
                      "duration": 533
                  }
              ]
          }
      ],
      "message": "Success"
  }
  ```
* **POST /playlists/add-music** - Add a music track to a playlist

  ```
  {
      "playlist_id": 1,
      "music_id": 1
  }

  ```
* **DELETE /playlists/{id}**- Remove a playlist

---

## License 📚

This project is licensed under the **STADITEK**.

---

Happy coding! 🚀
