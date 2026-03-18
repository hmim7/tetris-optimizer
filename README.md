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
go run . sample.txt
```

**Input (`sample.txt`):**
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

## Algorithm

The solver uses **depth-first backtracking**:

1. Calculate minimum square size based on total blocks
2. Attempt to place each tetromino sequentially
3. Backtrack on conflicts
4. Increment board size if no solution exists
5. Return the smallest valid square

## Performance

The solver is designed to be efficient and has been tested against the following performance benchmarks:

- **8 tetrominoes:** Solves in ≤ 1 second.
- **11 tetrominoes:** Solves in ≤ 3 seconds.
- **12 tetrominoes:** Solves in ≤ 5 seconds.

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
go run . sample.txt
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

Expected behavior matches specifications in `docs/PRD.md`.
