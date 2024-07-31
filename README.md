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

## Running the Tests

1. **Install dependencies** (if any):
   ```sh
   go mod tidy
   ```

2. **Run the tests**:
   ```sh
   go test ./...
   ```

3. **Run the application**:
   ```sh
   go run cmd/mars-rover-go/main.go
   ```

## Assumptions
- The grid starts at (0, 0) in the bottom-left corner.
- Valid movements are: `L`, `R`, `M`.
- Valid compass directions are: `N`, `S`, `E`, `W`.
- Exception handling for moves outside the grid bounds or invalid inputs.
