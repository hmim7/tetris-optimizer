# Tetris Optimizer

A Go CLI program that arranges tetrominoes into the smallest possible square.

## Overview

Tetris Optimizer reads a text file containing tetromino definitions and computes the minimal square arrangement using a backtracking algorithm. Each tetromino is labeled with a unique uppercase letter based on input order.

## Usage

```bash
go run . <inputFile>
```

### Example

```go
go run . samples/sample.txt
```

**Input (`samples/sample.txt`):**
```
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....

....
....
##..
.##.

##..
.#..
.#..
....

....
###.
.#..
....
```

**Expected Output:**
```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

## Input Format

- Each tetromino is defined in a **4x4 grid**
- Use `#` for blocks and `.` for empty spaces
- Separate tetrominoes with **exactly one empty line**
- Each tetromino must contain **exactly 4 blocks** that are **connected by shared edges**

### Valid Example
```
..##
..##
....
....
```

### Invalid Examples
```
##..
....
##..
....
```
*Disconnected blocks*

```
.###
..##
```
*Invalid grid dimensions*

## Error Handling

The program outputs `ERROR` and exits for:

- Missing or invalid command-line arguments
- File not found or permission denied
- Empty file
- Invalid characters (only `#`, `.`, and newline allowed)
- Incorrect grid dimensions (must be 4x4)
- Invalid block count (must be exactly 4 `#` per tetromino)
- Disconnected blocks (diagonal-only connections invalid)
- Missing or multiple separator lines between tetrominoes

## Algorithm & Optimizations

The core engine relies on a **depth-first backtracking** algorithm, heavily optimized to solve even the most complex configurations (up to 26 pieces) well within the required time limits. Key techniques include:

1. **Iterative Deepening:** The board size starts at the absolute mathematical minimum (`⌈√(N × 4)⌉`) and dynamically grows only when a layout is mathematically proven impossible.
2. **Bitmasking (Fast Collision Detection):** Tetrominoes and the board state are represented using `uint16` bitmasks. Checking if a piece fits is reduced to a single, lightning-fast CPU Bitwise AND operation.
3. **Flood-Fill Pruning (Dead-end Detection):** During recursion, the algorithm uses Depth-First Search (DFS) to evaluate empty spaces. If a pocket of empty space is created that cannot be filled (e.g., modulo 4 check), the entire branch is instantly pruned.
4. **Identical Piece Pruning:** The program detects duplicate tetromino shapes. By forcing a strict placement order for identical pieces, it avoids exploring millions of redundant symmetrical states.

## Performance

The solver is designed to be efficient and has been tested against the following performance benchmarks:

- **8 tetrominoes:** Solves in ≤ 1 second.
- **11 tetrominoes:** Solves in ≤ 3 seconds.
- **12 tetrominoes:** Solves in ≤ 5 seconds.


### Performance Timing

To manually validate that the program meets the strict execution time limits (e.g., ≤ 5 seconds for 12 pieces), prefix the run command with the Unix `time` utility. Check the `real` time output:
```bash
time go run . samples/hardexam.txt
```

## Implementation

The program follows a modular pipeline architecture:

```
CLI Args -> Input Reader -> Parser -> Validator -> Solver -> Output
```

### Key Modules

- **Input Layer**: Validates arguments and reads file
- **Parser**: Splits input into tetromino blocks
- **Validator**: Ensures format and connectivity rules
- **Domain Model**: Normalizes tetrominoes to coordinate sets
- **Solver**: Backtracking algorithm for placement
- **Output**: Renders final square with letter labels

## Constraints

- Written in Go using standard library only
- No tetromino rotations
- Deterministic placement order
- Maximum 26 tetrominoes (A-Z labeling)

## Testing

Run the program against provided test cases:

```go
go run . samples/sample.txt
go run . samples/hardexam.txt
```
- ***Errored cases:***

    ```go
    go run . samples/empty_file.txt
    go run . samples/invalid_chars_tabs.txt
    go run . samples/invalid_chars_spaces.txt
    go run . samples/invalid_chars_numbers.txt
    go run . samples/incorrect_line_length_long.txt
    go run . samples/incorrect_line_length_short.txt
    go run . samples/incorrect_block_height_long.txt
    go run . samples/incorrect_block_height_short.txt
    go run . samples/missing_separator.txt
    go run . samples/leading_newline.txt
    go run . samples/trailing_newlines.txt
    go run . samples/half_block_eof.txt
    go run . samples/ghost_block_trailing.txt
    go run . samples/non_uniform_blocks.txt
    ```
- ***Good example cases***
    ```go
    go run . samples/goodexample04.txt
    go run . samples/goodexample05.txt
    go run . samples/goodexample06.txt
    go run . samples/goodexample07.txt
    go run . samples/goodexample08.txt
    go run . samples/goodexample09.txt
    go run . samples/goodexample10.txt
    go run . samples/goodexample11.txt
    go run . samples/goodexample12.txt
    go run . samples/goodexample13.txt
    go run . samples/goodexample14.txt
    ```
Expected behavior matches specifications in `docs/PRD.md`.

### Validation & Test Coverage

To verify the automated test coverage of the project across all packages (aiming for ≥ 80%), use the built-in Go test tool:
```bash
go test ./... -coverpkg=./...
```

## License

This project is licensed under the MIT License.     Part of the Zone01 School curriculum.
