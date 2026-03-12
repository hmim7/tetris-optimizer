# Golden Test Suite: Tetris-Optimizer

This document records the mandatory test cases used to verify the functional requirements of the `tetris-optimizer` tool, ensuring robust error handling, accurate tetromino placement, and optimal square generation.

---

# 1. Core Functionality Cases

These cases focus on the core functionality and common usage patterns.

## Verification Notes:

This verification confirms that the solver correctly parses tetromino input, validates connectivity, normalizes coordinates, and applies backtracking to find the smallest possible square. Accuracy is determined by minimal square size, correct letter assignment, and deterministic placement order.

- Correct tetromino parsing and validation
- Exact 4x4 grid structure enforcement
- Proper connectivity checking (edge-to-edge)
- Minimal square size calculation
- Deterministic letter assignment (A-Z based on input order)

---

## Summary Table of Core Functionality Cases

| ID  | Description              | Input File        | Expected Behavior / Output |
|-----|--------------------------|-------------------|----------------------------|
| 01  | Single square tetromino  | `goodexample00.txt` | 2x2 square with 0 empty spaces (all filled). |
| 02  | Four tetrominoes         | `goodexample01.txt` | 5x5 square with 9 empty spaces. |
| 03  | Eight tetrominoes        | `goodexample02.txt` | 6x6 square with 4 empty spaces. |
| 04  | Eleven tetrominoes       | `goodexample03.txt` | 7x7 square with 5 empty spaces. |
| 05  | Twelve tetrominoes (hard)| `hardexam.txt`      | 7x7 square with 1 empty space. |

### Command Bundle

```bash
go run . samples/goodexample00.txt
```
```bash
go run . samples/goodexample01.txt
```
```bash
go run . samples/goodexample02.txt
```
```bash
go run . samples/goodexample03.txt
```
```bash
go run . samples/hardexam.txt
```

---

# 2. Error & Validation Cases

These cases define how the CLI should handle incorrect usage and invalid input formats.

## Verification Notes:

This phase validates CLI robustness and strict enforcement of the input contract. The program must gracefully intercept invalid arguments, malformed files, invalid characters, incorrect grid dimensions, and connectivity violations, outputting exactly `ERROR` for all failure cases.

- Strict single-argument contract enforcement
- Immediate error on invalid file format
- Character set validation (only `#`, `.`, `\n`)
- Grid dimension enforcement (4x4 per tetromino)
- Block count validation (exactly 4 `#` per tetromino)
- Connectivity validation (edge-to-edge, no diagonal-only)

---

## Summary Table of Error Cases

| ID  | Description           | Input File            | Expected Behavior / Output |
|-----|-----------------------|-----------------------|----------------------------|
| 11  | Missing argument      | *(No args)*           | Print `ERROR`; exit non-zero. |
| 12  | Too many arguments    | `file1.txt file2.txt` | Print `ERROR`; exit non-zero. |
| 13  | Empty file            | `empty_file.txt`      | Print `ERROR`; exit non-zero. |
| 14  | Invalid chars - tabs  | `invalid_chars_tabs.txt` | Print `ERROR`; exit non-zero. |
| 15  | Invalid chars - spaces| `invalid_chars_spaces.txt` | Print `ERROR`; exit non-zero. |
| 16  | Invalid chars - numbers| `invalid_chars_numbers.txt` | Print `ERROR`; exit non-zero. |
| 17  | Line too long         | `incorrect_line_length_long.txt` | Print `ERROR`; exit non-zero. |
| 18  | Line too short        | `incorrect_line_length_short.txt` | Print `ERROR`; exit non-zero. |
| 19  | Block too tall        | `incorrect_block_height_long.txt` | Print `ERROR`; exit non-zero. |
| 20  | Block too short       | `incorrect_block_height_short.txt` | Print `ERROR`; exit non-zero. |

### Command Bundle

```bash
go run .
```
```bash
go run . samples/file1.txt samples/file2.txt
```
```bash
go run . samples/empty_file.txt
```
```bash
go run . samples/invalid_chars_tabs.txt
```
```bash
go run . samples/invalid_chars_spaces.txt
```

---

# 3. Separator & Formatting Error Cases

These cases test the robustness of the parser for separator and formatting violations.

## Verification Notes:

These tests stress the parser under malformed input to ensure strict adherence to the separator rules and file structure requirements. The system must reject files with missing separators, multiple separators, leading/trailing newlines, and incomplete blocks.

- Exactly one empty line between tetrominoes
- No leading newlines at file start
- No multiple trailing newlines at file end
- Complete 4x4 blocks only (no partial definitions)
- Uniform block structure throughout file

---

## Summary Table of Separator & Formatting Cases

| ID  | Description            | Input File              | Expected Behavior / Output |
|-----|------------------------|-------------------------|----------------------------|
| 21  | Missing separator      | `missing_separator.txt` | Print `ERROR`; exit non-zero. |
| 22  | Multiple separators    | `badformat.txt`         | Print `ERROR`; exit non-zero. |
| 23  | Leading newline        | `leading_newline.txt`   | Print `ERROR`; exit non-zero. |
| 24  | Trailing newlines      | `trailing_newlines.txt` | Print `ERROR`; exit non-zero. |
| 25  | Non-uniform blocks     | `non_uniform_blocks.txt`| Print `ERROR`; exit non-zero. |
| 26  | Half-block at EOF      | `half_block_eof.txt`    | Print `ERROR`; exit non-zero. |
| 27  | Ghost block trailing   | `ghost_block_trailing.txt` | Print `ERROR`; exit non-zero. |

### Command Bundle

```bash
go run . samples/missing_separator.txt
```
```bash
go run . samples/badformat.txt
```
```bash
go run . samples/leading_newline.txt
```
```bash
go run . samples/trailing_newlines.txt
```

---

# 4. Tetromino Integrity Error Cases

These cases validate the geometric integrity of tetromino shapes.

## Verification Notes:

This phase validates tetromino block count and connectivity rules. Each tetromino must contain exactly 4 blocks that are connected edge-to-edge (not diagonal-only). The validator must use flood fill or adjacency counting (minimum 6 shared edges) to ensure proper connectivity.

- Exactly 4 `#` blocks per tetromino
- All blocks connected via shared edges (top/bottom/left/right)
- Diagonal-only connections are invalid
- Disconnected shapes are rejected
- Zero-block tetrominoes are rejected

---

## Summary Table of Integrity Cases

| ID  | Description            | Input File              | Expected Behavior / Output |
|-----|------------------------|-------------------------|----------------------------|
| 31  | Too many blocks (>4)   | `badexample00.txt`      | Print `ERROR`; exit non-zero. |
| 32  | Too few blocks (<4)    | `too_few_blocks.txt`    | Print `ERROR`; exit non-zero. |
| 33  | Zero blocks (all dots) | `badexample03.txt`      | Print `ERROR`; exit non-zero. |
| 34  | Diagonal only          | `badexample01.txt`      | Print `ERROR`; exit non-zero. |
| 35  | Disconnected pairs     | `badexample04.txt`      | Print `ERROR`; exit non-zero. |
| 36  | Vertical gap           | `badexample02.txt`      | Print `ERROR`; exit non-zero. |
| 37  | Only dots (zero tetros)| `only_dots.txt`         | Print `ERROR`; exit non-zero. |

### Command Bundle

```bash
go run . samples/badexample00.txt
```
```bash
go run . samples/too_few_blocks.txt
```
```bash
go run . samples/badexample03.txt
```
```bash
go run . samples/badexample01.txt
```
```bash
go run . samples/badexample04.txt
```
```bash
go run . samples/badexample02.txt
```

---

# 5. Edge Cases & Special Scenarios

These cases verify correct handling of valid edge cases and boundary conditions.

## Verification Notes:

- Valid tetrominoes with non-standard placement (blocks at grid edges)
- Proper normalization of tetromino coordinates
- Correct handling of minimal and maximal input sizes
- Performance within acceptable time limits

---

## Summary Table of Edge Cases

| ID  | Description            | Input File                  | Expected Behavior / Output |
|-----|------------------------|-----------------------------|----------------------------|
| 41  | Non-standard placement | `non_standard_placement.txt`| Valid 2x2 square output. |
| 42  | Maximum tetrominoes    | *(26 tetrominoes)*          | Valid square with A-Z labels. |
| 43  | Single I-tetromino     | *(1 vertical bar)*          | 2x4 or 4x2 minimal rectangle. |
| 44  | All same shape         | *(Multiple identical)*      | Valid minimal square. |

### Command Bundle

```bash
go run . samples/non_standard_placement.txt
```

---

# 6. Performance & Time Limit Cases

These cases verify that the solver completes within acceptable time limits.

## Verification Notes:

- 8 tetrominoes: ≤ 1 second
- 11 tetrominoes: ≤ 3 seconds
- 12 tetrominoes: ≤ 5 seconds
- Efficient backtracking with pruning
- No unnecessary recursion depth

---

## Summary Table of Performance Cases

| ID  | Description            | Input File              | Expected Behavior / Output |
|-----|------------------------|-------------------------|----------------------------|
| 51  | 8 tetros performance   | `goodexample02.txt`     | Complete in ≤ 1 second. |
| 52  | 11 tetros performance  | `goodexample03.txt`     | Complete in ≤ 3 seconds. |
| 53  | 12 tetros performance  | `hardexam.txt`          | Complete in ≤ 5 seconds. |

### Command Bundle

```bash
time go run . samples/goodexample02.txt
```
```bash
time go run . samples/goodexample03.txt
```
```bash
time go run . samples/hardexam.txt
```

---

# 7. Output Format Validation Cases

These cases verify correct output formatting and letter assignment.

## Verification Notes:

- Each tetromino assigned unique uppercase letter (A-Z)
- Letters assigned in input file order
- Empty spaces represented by `.`
- No trailing spaces on lines
- Proper newline termination
- Square dimensions match expected size

---

## Summary Table of Output Format Cases

| ID  | Description            | Input File              | Expected Behavior / Output |
|-----|------------------------|-------------------------|----------------------------|
| 61  | Letter assignment order| `goodexample01.txt`     | First tetro = A, second = B, etc. |
| 62  | Empty space rendering  | `goodexample01.txt`     | Exactly 9 `.` characters in output. |
| 63  | No trailing spaces     | `goodexample00.txt`     | Lines end immediately after last char. |
| 64  | Square dimensions      | `goodexample02.txt`     | Output is exactly 6x6 grid. |

### Command Bundle

```bash
go run . samples/goodexample01.txt | grep -o '\.' | wc -l
```
```bash
go run . samples/goodexample00.txt | cat -A
```

---

# 8. Integration & Regression Checklist

## Verification Points:

- **Allowed Packages:**
  ```bash
  go list -deps ./... | grep -v "^tetris-optimizer"
  ```
  Should only show standard library packages.

- **File Presence:**
  ```bash
  ls samples/*.txt
  ```
  Verify all test files exist.

- **Coverage / Regression:**
  ```bash
  go test ./... -cover
  ```
  Minimum 80% code coverage recommended.

- **Formatting / Static Checks:**
  ```bash
  gofmt -l .
  go vet ./...
  ```
  No formatting issues or vet warnings.

- **Build Verification:**
  ```bash
  go build -o tetris-optimizer .
  ```
  Clean compilation with no errors.

---

# 9. Comprehensive Error Category Summary

| Category | Specific Check | Sample File | Expected Output |
|----------|---------------|-------------|-----------------|
| **CLI Arguments** | `len(os.Args) != 2` | N/A | ERROR |
| **File Infrastructure** | Empty file (0 bytes) | `empty_file.txt` | ERROR |
| **Invalid Characters** | Contains tabs | `invalid_chars_tabs.txt` | ERROR |
| **Invalid Characters** | Contains spaces | `invalid_chars_spaces.txt` | ERROR |
| **Invalid Characters** | Contains numbers | `invalid_chars_numbers.txt` | ERROR |
| **Line Length** | Line too long (>4) | `incorrect_line_length_long.txt` | ERROR |
| **Line Length** | Line too short (<4) | `incorrect_line_length_short.txt` | ERROR |
| **Block Height** | Too many lines (>4) | `incorrect_block_height_long.txt` | ERROR |
| **Block Height** | Too few lines (<4) | `incorrect_block_height_short.txt` | ERROR |
| **Separators** | Missing separator | `missing_separator.txt` | ERROR |
| **Separators** | Multiple separators | `badformat.txt` | ERROR |
| **Leading/Trailing** | Leading newline | `leading_newline.txt` | ERROR |
| **Leading/Trailing** | Trailing newlines | `trailing_newlines.txt` | ERROR |
| **Block Integrity** | Too many blocks (>4) | `badexample00.txt` | ERROR |
| **Block Integrity** | Too few blocks (<4) | `too_few_blocks.txt` | ERROR |
| **Block Integrity** | Zero blocks | `badexample03.txt` | ERROR |
| **Connectivity** | Diagonal only | `badexample01.txt` | ERROR |
| **Connectivity** | Disconnected pairs | `badexample04.txt` | ERROR |
| **Connectivity** | Vertical gap | `badexample02.txt` | ERROR |
| **Edge Cases** | Incomplete block at EOF | `half_block_eof.txt` | ERROR |
| **Edge Cases** | Trailing junk content | `ghost_block_trailing.txt` | ERROR |
| **Edge Cases** | Non-uniform blocks | `non_uniform_blocks.txt` | ERROR |

---

# Final Integration Checklist

- **Core Functionality (Section 1):**  
  Cases `01-05` confirm correct tetromino placement and minimal square generation.

- **Error & Validation (Section 2):**  
  Cases `11-20` confirm CLI validation, file access checks, and character/dimension validation.

- **Separator & Formatting (Section 3):**  
  Cases `21-27` confirm strict separator rules and file structure enforcement.

- **Tetromino Integrity (Section 4):**  
  Cases `31-37` confirm block count and connectivity validation.

- **Edge Cases (Section 5):**  
  Cases `41-44` confirm handling of valid boundary conditions.

- **Performance (Section 6):**  
  Cases `51-53` confirm acceptable execution time limits.

- **Output Format (Section 7):**  
  Cases `61-64` confirm correct letter assignment and formatting.

- **Integration Checks (Section 8):**  
  Package restrictions, test coverage, formatting, and build verification.

- **Error Summary (Section 9):**  
  Complete mapping of all error categories to test files.

- **Deterministic Behavior:**  
  All test cases produce consistent, repeatable results.

- **No Memory Leaks:**  
  Program terminates cleanly without resource leaks.
