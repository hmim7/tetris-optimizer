# PRD - Tetris Optimizer
## 1. Problem Statement

We need a Go CLI program that receives a single text file containing a list of tetrominoes and arranges them in order to form the smallest possible square. The program must validate the input, handle errors gracefully, and output the solution in a readable format, identifying each tetromino with a unique uppercase letter. This tool aims to automate the process of tetromino placement while demonstrating algorithmic problem-solving in Go.

---

## 2. User / Use Case

- **Primary user:** Go developers or learners practicing algorithmic challenges.
- **Use case:** The user provides a text file containing one or more tetrominoes defined in a strict 4x4 grid format and expects the program to output the smallest possible square arrangement.  
  - The user runs the CLI tool: `go run . <inputFile>`  
  - The program validates the tetromino shapes, assembles them, and prints the resulting square to the console.  
  - If the file is incorrectly formatted or a solution cannot fully occupy a square, the program outputs `ERROR` or leaves spaces where necessary.  

---

## 3. CLI Contract

- **Command:** `go run . <inputFile>`
- **Inputs:**
  - `<inputFile>`: Path to a text file containing tetromino definitions.
- **Output:**
  - Prints the square arrangement to stdout.
  - Each tetromino is represented by a unique uppercase Latin letter (`A`, `B`, `C`, ...).
- **Error handling:**
  - If no file path is provided → print `ERROR` and exit.
  - If the file cannot be read → print `ERROR` and exit.
  - if the file is empty → print `ERROR` and exit.
  - If the file format is invalid → print `ERROR` and exit.
  - if the file has no reading rights → print `ERROR` and exit.
  - if more than the expected number of arguments provided (e.g. `go run . file1.txt file2.txt`) → print `ERROR` and exit.
  - if the user does not have read permissions for the input file. → print `ERROR` and exit.
  - If the tetrominoes cannot form a complete square → output partially filled square with spaces.

---

## 4. Functional Requirements
The program must strictly validate the input file. If any of the following conditions are met, the program must terminate and print exactly `ERROR` and nothing else.
### 4.1 Structural Validation
- **Grid Dimensions:** Each tetromino block must consist of exactly 4 lines, each containing exactly 4 characters.
- **Character Set:** Only the hash symbol (#), the period (.), and the newline character (\n) are permitted.
- **Separation:** Individual tetromino blocks must be separated by exactly one empty line.

**Example of Valid Structure:** 
```
....
..##
..##
....
```
→ Valid square tetromino.
```
....
####
....
....
```
→ Valid line tetromino.

---

### 4.2 Tetromino Integrity
- **Block Count:** Each 4x4 grid must contain exactly four # characters.
- **Connectivity:** All four # blocks must be connected to at least one other block by a shared edge (top, bottom, left, or right). Diagonal-only connections are invalid.

**Example of Invalid Tetromino (ERROR):**
```
##..
....
##..
....
```
*Reason: The blocks are not connected (Disconnected shape).*

---

### 4.3 Algorithmic Requirements
- **Order Preservation:** Tetrominoes must be identified by uppercase letters (`A`, `B`, `C`...) based on their order in the input file.
- **Minimal Square:** The result must be the smallest possible square that can contain all pieces.

**Example of Order Identification:**
- 1st Tetromino in file → Assigned letter A
- 2nd Tetromino in file → Assigned letter B

---

### 4.4 Square Assembly Rules
- **No Overlap:** Pieces cannot occupy the same coordinate on the board.
- **Full Usage:** Every tetromino provided in the input file must be included in the final square.
- **Dynamic Sizing:** If a solution is not found for a size $N \times N$, the program must increment to $(N+1) \times (N+1)$ and retry. Empty spaces in the final square are represented by `.`.

**Example of Smallest Square (3 pieces):** *3 pieces = 12 blocks. Smallest square possible is 4x4 (16 cells), leaving 4 empty spaces.*
```
AAAB
..AB
..BB
CCCC
```

---

### 4.5 Output Formatting
- Print the final square to stdout.
- Each tetromino labeled by a unique uppercase letter.
- Maintain consistent row/column alignment for readability.

---

### 4.6 Error Handling
- Input validation errors → print `ERROR`.
- File reading errors → print `ERROR`.
- Invalid tetromino shapes → print `ERROR`.

**Validation Error Cases**
| Case            | Trigger Condition                                           | Expected Output |
|-----------------|-------------------------------------------------------------|-----------------|
| Missing Argument| Running `go run .` without a file path                      | ERROR           |
| Empty File      | `os.Stat(file).Size() == 0`                                  | ERROR           |
| Invalid Char    | Any character other than `#`, `.` or `\n` found in the grid | ERROR           |
| Invalid Grid    | A block that is not exactly `4x4` characters                | ERROR           |
| Too Many `#`    | More than 4 hashes found in a single `4x4` block            | ERROR           |
| Too Few `#`     | Fewer than 4 hashes found in a single `4x4` block           | ERROR           |
| Bad Connection  | No shared edge between blocks (disconnected shape)          | ERROR           |
| Bad Separation  | Missing or multiple empty lines between tetrominoes         | ERROR           |

**Error Classification Table**
| Error Category     | Trigger Condition (Technical)                                             | Expected Output |
|--------------------|---------------------------------------------------------------------------|-----------------|
| CLI Arguments      | `len(os.Args) != 2`                                                       | ERROR           |
| File Accessibility | File does not exist, is a directory, or permission is denied              | ERROR           |
| File Content       | File is empty (0 bytes)                                                   | ERROR           |
| Character Set      | Any character found that is not `#`, `.` or `\n`                          | ERROR           |
| Grid Geometry      | A line is not exactly 4 characters long or a block is not 4 lines high    | ERROR           |
| Separation         | Missing empty line between blocks or multiple empty lines are present     | ERROR           |
| Block Integrity    | Number of `#` characters in a 4×4 grid is not equal to 4                  | ERROR           |
| Connectivity       | Total adjacencies between `#` blocks is less than 6

---

### 4.7 Constraints
- The program must be written entirely in Go.
- Only standard Go packages are allowed.
- Follow Go best practices (naming, modularity, error handling).
- Unit tests are recommended to validate algorithm correctness.

---

## 5. Non-Goals (Out of Scope)

- No graphical representation of the tetrominoes.
- No interactive mode; CLI only.
- No support for non-standard polyominoes (only tetrominoes of 4 blocks).
- No optimization for speed beyond correctness (algorithm efficiency is secondary).
- No external dependencies; standard library only.



---

## 6. Acceptance Criteria

According to the project requirements, the tool must pass **all** of the following test cases exactly as specified in order to be considered successful.

### 6.1 Basic Functional Cases

- [ ] `go run . samples/goodexample00.txt` produces 2x2 square with 0 empty spaces (all filled)
  - Input: Single square tetromino
  - Expected: 2x2 grid with all positions filled by letter A

- [ ] `go run . samples/goodexample01.txt` produces 5x5 square with exactly 9 empty spaces
  - Input: Four tetrominoes (I, line, L, Z shapes)
  - Expected: 5x5 grid with tetrominoes A-D and 9 dots

- [ ] `go run . samples/goodexample02.txt` produces 6x6 square with exactly 4 empty spaces
  - Input: Eight tetrominoes
  - Expected: 6x6 grid with tetrominoes A-H and 4 dots

- [ ] `go run . samples/goodexample03.txt` produces 7x7 square with exactly 5 empty spaces
  - Input: Eleven tetrominoes
  - Expected: 7x7 grid with tetrominoes A-K and 5 dots

- [ ] `go run . samples/hardexam.txt` produces 7x7 square with exactly 1 empty space
  - Input: Twelve tetrominoes (challenging configuration)
  - Expected: 7x7 grid with tetrominoes A-L and 1 dot

### 6.2 CLI Validation & Error Cases

- [ ] `go run .` prints `ERROR` and exits non-zero (missing argument)
  - Input: No command line arguments
  - Expected: Exactly "ERROR" output

- [ ] `go run . file1.txt file2.txt` prints `ERROR` and exits non-zero (too many arguments)
  - Input: Multiple file arguments
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/empty_file.txt` prints `ERROR` and exits non-zero
  - Input: File with 0 bytes
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/invalid_chars_tabs.txt` prints `ERROR` and exits non-zero
  - Input: File containing tab characters
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/invalid_chars_spaces.txt` prints `ERROR` and exits non-zero
  - Input: File containing space characters (leading whitespace)
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/invalid_chars_numbers.txt` prints `ERROR` and exits non-zero
  - Input: File containing numeric characters
  - Expected: Exactly "ERROR" output

### 6.3 Structural Format Error Cases

- [ ] `go run . samples/incorrect_line_length_long.txt` prints `ERROR` and exits non-zero
  - Input: Lines longer than 4 characters
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/incorrect_line_length_short.txt` prints `ERROR` and exits non-zero
  - Input: Lines shorter than 4 characters
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/incorrect_block_height_long.txt` prints `ERROR` and exits non-zero
  - Input: Tetromino blocks with more than 4 lines
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/incorrect_block_height_short.txt` prints `ERROR` and exits non-zero
  - Input: Tetromino blocks with fewer than 4 lines
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/missing_separator.txt` prints `ERROR` and exits non-zero
  - Input: No empty line between tetromino blocks
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/badformat.txt` prints `ERROR` and exits non-zero
  - Input: Multiple empty lines between tetromino definitions
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/leading_newline.txt` prints `ERROR` and exits non-zero
  - Input: File beginning with empty line
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/trailing_newlines.txt` prints `ERROR` and exits non-zero
  - Input: File ending with multiple empty lines
  - Expected: Exactly "ERROR" output

### 6.4 Tetromino Integrity Error Cases

- [ ] `go run . samples/badexample00.txt` prints `ERROR` and exits non-zero
  - Input: Tetromino with more than 4 blocks (5 blocks)
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/too_few_blocks.txt` prints `ERROR` and exits non-zero
  - Input: Tetromino with fewer than 4 blocks (3 blocks)
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/badexample03.txt` prints `ERROR` and exits non-zero
  - Input: Tetromino with zero blocks (all dots)
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/badexample01.txt` prints `ERROR` and exits non-zero
  - Input: Blocks connected only diagonally
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/badexample04.txt` prints `ERROR` and exits non-zero
  - Input: Disconnected blocks (separate pairs)
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/badexample02.txt` prints `ERROR` and exits non-zero
  - Input: Disconnected blocks with vertical gap
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/only_dots.txt` prints `ERROR` and exits non-zero
  - Input: File containing only dots (zero tetrominoes)
  - Expected: Exactly "ERROR" output

### 6.5 Edge Cases & Special Scenarios

- [ ] `go run . samples/non_standard_placement.txt` produces valid 2x2 square output
  - Input: Valid tetromino with blocks at right edge of 4x4 grid
  - Expected: 2x2 square with letter A (not ERROR)

- [ ] `go run . samples/half_block_eof.txt` prints `ERROR` and exits non-zero
  - Input: Incomplete tetromino definition at end of file
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/ghost_block_trailing.txt` prints `ERROR` and exits non-zero
  - Input: Extra lines at end that don't form complete 4x4 block
  - Expected: Exactly "ERROR" output

- [ ] `go run . samples/non_uniform_blocks.txt` prints `ERROR` and exits non-zero
  - Input: Mixed block sizes (some 4 lines, some 5 lines)
  - Expected: Exactly "ERROR" output

### 6.6 Performance & Time Limit Cases

- [ ] `time go run . samples/goodexample02.txt` completes in ≤ 1 second
  - Input: 8 tetrominoes
  - Expected: Valid output within time limit

- [ ] `time go run . samples/goodexample03.txt` completes in ≤ 3 seconds
  - Input: 11 tetrominoes
  - Expected: Valid output within time limit

- [ ] `time go run . samples/hardexam.txt` completes in ≤ 5 seconds
  - Input: 12 tetrominoes (challenging case)
  - Expected: Valid output within time limit

### 6.7 Output Format Validation Cases

- [ ] `go run . samples/goodexample01.txt` assigns letters in input file order (A, B, C, D)
  - Input: Four tetrominoes
  - Expected: First tetromino = A, second = B, third = C, fourth = D

- [ ] `go run . samples/goodexample01.txt | grep -o '\.' | wc -l` outputs exactly 9
  - Input: Four tetrominoes in 5x5 square
  - Expected: Exactly 9 empty space characters

- [ ] `go run . samples/goodexample00.txt | cat -A` shows no trailing spaces on lines
  - Input: Single square tetromino
  - Expected: Lines end immediately after last character

- [ ] `go run . samples/goodexample02.txt` produces exactly 6x6 grid dimensions
  - Input: Eight tetrominoes
  - Expected: Output is exactly 6 lines of 6 characters each

### 6.8 Integration & Build Verification

- [ ] `go list -deps ./... | grep -v "^tetris-optimizer"` shows only standard library packages
  - Expected: No external dependencies

- [ ] `ls samples/*.txt` confirms all test files exist
  - Expected: All referenced sample files are present

- [ ] `go test ./... -cover` achieves minimum 80% code coverage
  - Expected: Comprehensive test coverage

- [ ] `gofmt -l .` and `go vet ./...` produce no warnings
  - Expected: Clean code formatting and no static analysis issues

- [ ] `go build -o tetris-optimizer .` compiles without errors
  - Expected: Clean compilation with no build errors

---

## 7. Implementation Approach
### 7.1 Overall Strategy

The problem is solved using a **backtracking algorithm** that incrementally assembles tetrominoes into a square board of increasing size until a valid configuration is found.

The implementation follows a **modular and layered (pipeline) architecture**, where each layer has a single, well-defined responsibility. This approach improves readability, testability, and reduces coupling between components.

The solution prioritizes:
- correctness over premature optimization,
- deterministic behavior,
- clear separation between parsing, validation, domain modeling, and solving logic.

---

### 7.2 Architectural Overview

The program is structured as a linear pipeline with a clear execution flow:
```
    ┌───────────────┐
    │   CLI Args    │
    └──────┬────────┘
           ↓
    ┌───────────────┐
    │ Input Reader  │
    └──────┬────────┘
           ↓
    ┌───────────────┐
    │    Parser     │
    └──────┬────────┘
           ↓
    ┌───────────────┐
    │   Validator   │
    └──────┬────────┘
           ↓
┌─────────────────────────┐
│ Tetromino Normalization │
└──────────┬──────────────┘
           ↓
 ┌──────────────────────┐
 │ Solver (Backtracking)│
 └─────────┬────────────┘
           ↓
   ┌─────────────────┐
   │ Output Renderer │
   └─────────────────┘

```

Each stage either produces valid data for the next stage or fails fast with an error.

---

### 7.3 Module Responsibilities

The codebase is divided into the following logical modules:

#### Input Layer
- Reads and validates command-line arguments.
- Loads the input file contents into memory.
- Does not perform any domain validation.

#### Parsing & Validation Layer
- Splits the raw input into individual tetromino blocks.
- Validates format rules (dimensions, characters, block count).
- Ensures each tetromino is composed of exactly four connected blocks.
- Rejects malformed input early.

#### Domain Model
- Represents tetrominoes as collections of 2D coordinates.
- Normalizes each tetromino so that its top-left block aligns at origin (0,0).
- Provides size-related helpers (width, height).
- Contains no solver or board logic.

#### Solver Layer
- Computes the minimal possible square size based on the number of tetrominoes.
- Uses a depth-first backtracking algorithm to place tetrominoes sequentially.
- Attempts all valid positions for each tetromino before backtracking.
- Increases board size progressively if no valid configuration is found.
- Guarantees deterministic placement order.

#### Output Layer
- Converts the solved board into the required textual representation.
- Assigns uppercase Latin letters to tetrominoes based on input order.
- Handles only presentation concerns.

---

### 7.4 Backtracking Algorithm

The solver operates as follows:

1. Compute the initial board size using the square root of total occupied cells.
2. Create an empty board of that size.
3. Attempt to place the first tetromino at every valid position.
4. For each successful placement:
   - Recursively attempt to place the next tetromino.
5. If placement fails:
   - Remove the previously placed tetromino.
   - Continue searching.
6. If all tetrominoes are placed successfully:
   - Return the solved board.
7. If no solution exists at the current size:
   - Increase the board size and repeat the process.

This approach ensures that the smallest possible square is always found first.

---

### 7.5 Error Handling Philosophy

The program follows a **fail-fast strategy**:

- Any invalid input format or malformed tetromino results in immediate termination.
- Errors are propagated upward and handled centrally in the entry point.
- The program outputs only a single string (`ERROR`) in failure cases, as specified.

No partial output or debug information is printed.

---

### 7.6 Testing Strategy (High Level)

Although not mandatory, unit tests are recommended and planned for:

- Parsing and validation logic.
- Tetromino normalization.
- Solver behavior on small and edge-case inputs.

This ensures confidence in correctness and simplifies debugging during development. Performance optimization is critical to meet the specified time limits in the acceptance criteria.

---

### 7.7 Design Trade-offs

Accepted trade-offs include:
- Using backtracking with performance optimizations to meet time requirements while maintaining clarity.
- Avoiding concurrency to maintain deterministic behavior.
- Not supporting tetromino rotations, as they are not required by the specification.

These decisions align with the project scope and evaluation criteria.


---

## 8. Milestones (Implementation Roadmap)

The project is divided into incremental, verifiable milestones to ensure steady progress and testable outcomes.

### Milestone 1. Input Handling
   - Implement `input` module to read CLI arguments and load the text file.
   - Validate the presence of exactly one argument.
   - Fail fast with `ERROR` if the file cannot be read or arguments are invalid.
   - **Testable:** Run program with valid and invalid args and verify error handling.
---
### Milestone 2: Parsing & Tetromino Validation
   - Implement `parser` module to split raw input into tetromino blocks.
   - Validate format:
     - 4x4 grids
     - Allowed characters only (`.` and `#`)
     - Exactly 4 `#` per tetromino
     - Connectivity of blocks
   - **Testable:** Unit tests for valid and invalid tetrominoes, including edge cases (minimum, maximum, malformed inputs).
---
### Milestone 3: Tetromino Domain Modeling
   - Implement `model` module:
     - Represent tetrominoes as collections of coordinates.
     - Normalize each tetromino (align top-left to origin).
     - Provide helper methods (Width, Height).
   - **Testable:** Unit tests for normalization and size calculations.
---
### Milestone 4: Backtracking Solver Implementation
   - Implement `solver` module:
     - Compute minimal board size based on tetromino count.
     - Place tetrominoes sequentially using backtracking.
     - Increment board size if no solution is found.
     - Guarantee deterministic placement order.
   - **Testable:** Verify solver outputs correct smallest square for given test files.
---
### Milestone 5: Output Rendering
   - Implement `output` module:
     - Convert solved board to textual output.
     - Map tetrominoes to uppercase Latin letters in input order.
     - Maintain gaps if a perfect square is not possible.
   - **Testable:** Compare program output against expected board layouts for sample and custom test files.
---
### Milestone 6: Integration & End-to-End Testing
   - Integrate all modules.
   - Ensure correct CLI flow: `Read -> Parse -> Validate -> Solve -> Print`.
   - Test complete program end-to-end using multiple input scenarios, including edge cases.
   - **Testable:** Automated or manual verification against sample outputs and error cases.

---

## 9. Risks / Open Questions

### Risks

- **Risk: Input File Format Errors**  
  Tetromino input files may be malformed, missing lines, or contain invalid characters. The program must handle these gracefully by printing `ERROR` without crashing.

- **Risk: Large Number of Tetrominoes**  
  As the number of tetrominoes increases, the computation to find the smallest square may become computationally intensive. There is a risk of performance bottlenecks for very large input sets.

- **Risk: Non-Square Solutions**  
  In cases where a perfect square cannot be formed, the placement algorithm must handle spacing correctly. Incorrect handling could lead to overlapping or misaligned tetrominoes.

- **Risk: Go Language Constraints**  
  The project is restricted to standard Go packages only. Certain algorithms or utilities available in external libraries will need to be implemented manually, increasing development complexity.

### Open Questions

- **Question: Maximum Input Size**  
  What is the practical limit for the number of tetrominoes the program should support without performance degradation?

- **Question: Error Reporting Details**  
  Should the program provide detailed error messages indicating the type of input error, or is a generic `ERROR` sufficient?
  
  *Decision:* The program will only output the generic string ERROR for all cases listed in Section 4.6 to maintain compatibility with the audit requirements.

- **Question: Testing Requirements**  
  How extensive should unit testing be? Should it cover only correct inputs, or also include malformed and edge-case tetromino sets?

