# Go Note App

A simple API built with Go and SQLite.

## Features

- Create, read, update, and delete notes
- RESTful API endpoints
- Built with Gin and GORM

## Getting Started

```bash
git clone https://github.com/ElanDeyan/go-note-app.git
cd go-note-app
go build
```

### Usage

```bash
./go-note-app
```

The server will start on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint     | Description         |
|--------|--------------|---------------------|
| GET    | /notes       | List all notes      |
| POST   | /notes       | Create a new note   |
| GET    | /notes/{id}  | Get a specific note |
| PATCH  | /notes/{id}  | Update a note       |
| DELETE | /notes/{id}  | Delete a note       |
