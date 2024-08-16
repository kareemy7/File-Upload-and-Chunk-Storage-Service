# File Upload and Chunk Storage Service

This project is a Go-based web service that allows users to upload large files, split them into smaller chunks, and store these chunks in LevelDB. The service also supports reassembling the file from its chunks and downloading it upon request.

## Features

- **File Upload:** Upload large files that are automatically split into smaller chunks for efficient storage.
- **Chunk Storage:** Chunks are stored in LevelDB, a fast key-value storage library.
- **File Metadata Management:** Metadata associated with each uploaded file, including chunk IDs and file details, are stored and managed in LevelDB.
- **File Download:** Retrieve and reassemble the file from its chunks and download it as a complete file.
- **Efficient Resource Usage:** Handles large files by processing them in chunks, reducing memory overhead.

## Installation

### Prerequisites

- **Go 1.19+** installed on your system.
- **LevelDB** Go package (`github.com/syndtr/goleveldb/leveldb`).
- **Gin Web Framework** (`github.com/gin-gonic/gin`).

### Steps

1.  Clone the repository:

```bash
git clone https://github.com/kareemy7/File-Upload-and-Chunk-Storage-Service.git
cd File-Upload-and-Chunk-Storage-Service
```

2.  Install dependencies:
    
    ```bash
    go mod tidy
    ```
    
3.  Run the service:
    
    ```bash
    go run main.go
    ```
    
    Alternatively, use `CompileDaemon` for live reloading during development:
    
    ```bash
    CompileDaemon -command="./File-Upload-and-Chunk-Storage-Service"
    ```
    

## API Endpoints

### Upload File

- **Endpoint:** `/upload`
- **Method:** `POST`
- **Description:** Upload a file. The file is split into chunks and stored in LevelDB.
- **Request:**
    - `form-data` with a key `file_` containing the file to upload.
- **Response:**
    - `200 OK` with a JSON containing the `file_id`.

### Download File

- **Endpoint:** `/download/:file_id`
- **Method:** `GET`
- **Description:** Download a previously uploaded file by its `file_id`.
- **Response:**
    - `200 OK` with the file data.

## Usage

### Uploading a File

Use Postman or any other HTTP client to upload a file:

1.  **Set the URL:** `http://localhost:8080/upload`
2.  **Method:** `POST`
3.  **Body:**
    - Select `form-data`.
    - Use key `file_` and choose the file you want to upload.

### Downloading a File

To download a file:

1.  **Set the URL:** `http://localhost:8080/download/{file_id}`
2.  **Method:** `GET`
3.  The file will be reassembled and downloaded as a complete file.

## Contributing

Contributions are welcome! Please fork this repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

- **Author:** Karim Elbassiouny
- **Email:** <ins>karimelbassiouny@gmail.com</ins>
- **GitHub:** <ins>kareemy7</ins>