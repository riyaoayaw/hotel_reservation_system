# Hotel Reservation System

This is a hotel reservation system program written in Go.

## Features

- **Room Management**: The hotel has five rooms.
- **User Menu**: Allows users to:
  - Reserve rooms
  - View room details
  - Cancel reservations
- **Concurrency Handling**: Utilizes goroutines to periodically check for vacant rooms and ensure thread-safe operations.

## Instructions to Run the Program

1. **Navigate to the Project Folder**:
   ```bash
   cd path/to/your/project
   ```

2. **Initialize Dependencies**:
    ```bash
    go mod tidy
    ```

3. **Run the Program**:
    ```bash
    go run main.go
    ```

### Testing

1. **Switch to the hotel folder inside the project**: 
    ```bash
    cd hotel
    ```

2. **To test the functionality of the program, including concurrency, use the following commands**:
    ```bash
    go test -v
    ```

3. **For detecting race conditions in your tests, run**:
    ```bash
    go test -race
    ```