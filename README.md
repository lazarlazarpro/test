# Pack Service

This service provides an API that calculates the optimal pack sizes based on given item quantities. It ensures that the fewest packs are used while satisfying the total item count.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (if building locally)
- [Docker](https://docs.docker.com/get-docker/)

## Building and Running the Service

### With Go

1. Navigate to the project directory:

   ```bash
   cd path/to/project
   ```

2. Build and run:

   ```bash
   go build
   ./main
   ```

3. The service will be accessible at `http://localhost:8080`.

### With Docker

1. Navigate to the project directory:

   ```bash
   cd path/to/project
   ```

2. Build the Docker image:

   ```bash
   docker build -t pack-service .
   ```

3. Run the Docker container:

   ```bash
   docker run -p 8080:8080 pack-service
   ```

4. The service will be accessible at `http://localhost:8080`.

## API Usage

To calculate the packs, use the `/pack` endpoint with the `items` query parameter:

```plaintext
http://localhost:8080/pack?items=<NUMBER_OF_ITEMS>
```

Replace `<NUMBER_OF_ITEMS>` with the desired item quantity.
