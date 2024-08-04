```markdown
# Mars Rover Go

Mars Rover Go is a project simulating a Mars rover explorer using Go programming language and hexagonal architecture (Ports and Adapters). The project focuses on controlling the rover's movements and direction within a grid, showcasing modular software design. Key technologies include Go, CLI interface, and principles of clean architecture.

## Learning Competencies
- Understand and implement object-oriented design, DDD, patterns, and best practices.
- Implement tests using TDD.
- Properly handle input and output operations.

## The Challenge
NASA has deployed a fleet of robotic rovers on a rectangular plateau on Mars. Each rover's state consists of its position (X, Y) and the compass direction (N, S, E, W). Your task is to develop an API that controls the rovers' movements across a 10x10 grid, interpreting commands (`L`, `R`, `M`), and handling possible obstacles.

## Input
- The plateau is represented as a 10x10 grid.
- Initial positions and commands for each rover.

## Output
- Final coordinates and direction of each rover.
- If an obstacle is encountered, prefix the output with `O:`.

## Example Input
```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

## Expected Output
```
1 3 N
5 1 E
```

## Mars Rover Project Execution

### 1. **Preparation**

1. Clone the repository:
   ```sh
   git clone https://github.com/alejandrososa/mars-rover-go.git
   cd mars-rover-go
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

### 2. **Running the CLI**

1. Build the CLI binary:
   ```sh
   go build -o bin/mars-rover-cli ./cmd/mars-rover-cli
   ```

2. Execute the CLI:
   ```sh
   ./bin/mars-rover-cli
   ```

### 3. **Running the HTTP Server**

1. Build the HTTP server binary:
   ```sh
   go build -o bin/mars-rover-http ./cmd/mars-rover-http
   ```

2. Start the server:
   ```sh
   ./bin/mars-rover-http
   ```

3. Access the API at [http://localhost:8080](http://localhost:8080).

### 4. **Usage**

- **HTTP API**: Send a POST request to `/platform` with a JSON payload to create a platform.
- **CLI**: Follow the console prompts to create a platform and add rovers.

## Running the Tests

1. **Run the tests**:
   ```sh
   go test ./...
   ```

2. **Run the application**:
   ```sh
   go run cmd/mars-rover-http/main.go
   ```

## Assumptions
- The grid starts at (0, 0) in the bottom-left corner.
- Valid movements are: `L`, `R`, `M`.
- Valid compass directions are: `N`, `S`, `E`, `W`.
- Exception handling for moves outside the grid bounds or invalid inputs.

## Using the Makefile

To simplify the setup and execution of the project, you can use the provided Makefile:

### 1. **Setup the environment**

To install all necessary Go libraries and prepare the environment, run:
   ```sh
   make setup
   ```

### 2. **Build the project**

To compile the HTTP server, use:
   ```sh
   make build
   ```

### 3. **Start the server**

To start the HTTP server on port 8080, execute:
   ```sh
   make start
   ```

### 4. **Help**

For more information about available commands and usage, run:
   ```sh
   make help
   ```

Please refer to this README file for detailed instructions on how to use the project and understand its features.
```