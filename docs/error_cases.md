# Audit Cases: Tetris-Optimizer Error Handling

## Audit Case 1: Missing Command Line Argument
**Description:** Tests that the program detects when no input file is provided.

**Input Argument**
```bash
go run .
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 2: Too Many Command Line Arguments
**Description:** Tests that the program rejects multiple file arguments.

**Input Argument**
```bash
go run . file1.txt file2.txt
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 3: Empty File
**Description:** Tests that the program detects and rejects files with 0 bytes.

**Input Argument**
```bash
go run . samples/empty_file.txt
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 4: Invalid Characters - Tabs
**Description:** Tests that the program rejects files containing tab characters.

**Input Argument**
```bash
go run . samples/invalid_chars_tabs.txt
```
**File Content:**
```
...#
	...#
...#
...#
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 5: Invalid Characters - Spaces
**Description:** Tests that the program rejects files containing space characters (leading whitespace).

**Input Argument**
```bash
go run . samples/invalid_chars_spaces.txt
```
**File Content:**
```
 #..
...#
...#
...#
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 6: Invalid Characters - Numbers
**Description:** Tests that the program rejects files containing numeric characters.

**Input Argument**
```bash
go run . samples/invalid_chars_numbers.txt
```
**File Content:**
```
...#
...#
...#
...1
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 7: Incorrect Line Length - Too Long
**Description:** Tests that the program rejects lines longer than 4 characters.

**Input Argument**
```bash
go run . samples/incorrect_line_length_long.txt
```
**File Content:**
```
...#.
...#
...#
...#
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 8: Incorrect Line Length - Too Short
**Description:** Tests that the program rejects lines shorter than 4 characters.

**Input Argument**
```bash
go run . samples/incorrect_line_length_short.txt
```
**File Content:**
```
...#
...#
...#
...
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 9: Incorrect Block Height - Too Many Lines
**Description:** Tests that the program rejects tetromino blocks with more than 4 lines.

**Input Argument**
```bash
go run . samples/incorrect_block_height_long.txt
```
**File Content:**
```
...#
...#
...#
...#
....
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 10: Incorrect Block Height - Too Few Lines
**Description:** Tests that the program rejects tetromino blocks with fewer than 4 lines.

**Input Argument**
```bash
go run . samples/incorrect_block_height_short.txt
```
**File Content:**
```
...#
...#
...#
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 11: Missing Separator Between Tetrominoes
**Description:** Tests that the program detects when two tetromino blocks appear consecutively without an empty line separator.

**Input Argument**
```bash
go run . samples/missing_separator.txt
```
**File Content:**
```
...#
...#
...#
...#
....
....
....
####
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 12: Multiple Separators Between Tetrominoes
**Description:** Tests that the program rejects files with more than one empty line between tetromino definitions.

**Input Argument**
```bash
go run . samples/badformat.txt
```
**File Content:**
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
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 13: Leading Newline
**Description:** Tests that the program rejects files that begin with an empty line.

**Input Argument**
```bash
go run . samples/leading_newline.txt
```
**File Content:**
```

...#
...#
...#
...#
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 14: Trailing Newlines
**Description:** Tests that the program rejects files ending with multiple empty lines.

**Input Argument**
```bash
go run . samples/trailing_newlines.txt
```
**File Content:**
```
...#
...#
...#
...#


```
**Expected Result:**
```
ERROR
```

---

## Audit Case 15: Non-Uniform Block Sizes
**Description:** Tests that the program rejects files with mixed block sizes (some 4 lines, some 5 lines).

**Input Argument**
```bash
go run . samples/non_uniform_blocks.txt
```
**File Content:**
```
...#
...#
...#
...#

....
....
....
####
....
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 16: Half-Block at End of File
**Description:** Tests that the program rejects files with incomplete tetromino definitions at EOF.

**Input Argument**
```bash
go run . samples/half_block_eof.txt
```
**File Content:**
```
...#
...#
...#
...#

....
....
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 17: Ghost Block - Trailing Content
**Description:** Tests that the program rejects files with extra lines (junk) at the end that don't form a complete 4x4 block.

**Input Argument**
```bash
go run . samples/ghost_block_trailing.txt
```
**File Content:**
```
...#
...#
...#
...#

....
.##.
.##.
....

....
....
....
....
....
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 18: Too Many Blocks (More than 4)
**Description:** Tests that the program rejects tetromino blocks containing more than 4 hash symbols.

**Input Argument**
```bash
go run . samples/badexample00.txt
```
**File Content:**
```
####
...#
....
....
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 19: Too Few Blocks (Less than 4)
**Description:** Tests that the program rejects tetromino blocks containing fewer than 4 hash symbols.

**Input Argument**
```bash
go run . samples/too_few_blocks.txt
```
**File Content:**
```
...#
...#
...#
....
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 20: Zero Blocks (All Dots)
**Description:** Tests that the program rejects tetromino blocks with no hash symbols.

**Input Argument**
```bash
go run . samples/badexample03.txt
```
**File Content:**
```
....
....
....
....
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 21: Disconnected Shape - Diagonal Only
**Description:** Tests that the program rejects shapes where blocks only touch at corners (diagonal connections).

**Input Argument**
```bash
go run . samples/badexample01.txt
```
**File Content:**
```
...#
..#.
.#..
#...
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 22: Disconnected Shape - Separate Pairs
**Description:** Tests that the program rejects shapes with blocks that don't form a single contiguous piece.

**Input Argument**
```bash
go run . samples/badexample04.txt
```
**File Content:**
```
..##
....
....
##..
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 23: Disconnected Shape - Vertical Gap
**Description:** Tests that the program rejects shapes with vertical gaps between blocks.

**Input Argument**
```bash
go run . samples/badexample02.txt
```
**File Content:**
```
...#
...#
#...
#...
```
**Expected Result:**
```
ERROR
```

---

## Audit Case 24: Valid Non-Standard Placement
**Description:** Tests that the program correctly handles valid tetrominoes with blocks positioned at the right edge of the 4x4 grid.

**Input Argument**
```bash
go run . samples/non_standard_placement.txt
```
**File Content:**
```
...#
..##
...#
....
```
**Expected Result:**
A valid 2x2 square output:
```
AA
AA
```

---

## Audit Case 25: Zero Tetrominoes (Only Dots)
**Description:** Tests that the program rejects files containing only dots with no valid tetrominoes.

**Input Argument**
```bash
go run . samples/only_dots.txt
```
**File Content:**
```
....
....
....
....
```
**Expected Result:**
```
ERROR
```

---

## Audit Checklist for Error Handling Verification

| Verification Point | Requirement |
|-------------------|-------------|
| **Standard Packages** | Only standard Go packages are used. |
| **CLI Validation** | Exactly one argument required (`len(os.Args) == 2`). |
| **File Access** | File must exist, be readable, and non-empty. |
| **Character Set** | Only `#`, `.`, and `\n` characters allowed. |
| **Grid Dimensions** | Each tetromino must be exactly 4 lines of 4 characters. |
| **Block Count** | Each tetromino must contain exactly 4 `#` blocks. |
| **Connectivity** | All 4 blocks must be connected edge-to-edge (minimum 6 adjacencies). |
| **Separator Rules** | Exactly one empty line between tetrominoes. |
| **No Leading/Trailing** | No empty lines at start or multiple at end of file. |
| **Complete Blocks** | No partial or incomplete tetromino definitions. |
| **Error Output** | All validation failures output exactly `ERROR`. |
| **Exit Behavior** | Program terminates immediately on first error. |
| **Performance** | Validation runs quickly without unnecessary processing. |
| **Code Quality** | Clear error checking logic with proper test coverage. |

---

## Implementation Validation Order

The program should validate in this sequence to catch errors efficiently:

1. **CLI Check:** Verify `len(os.Args) == 2`
2. **File Access:** Confirm file exists, is readable, and non-empty
3. **Line-by-Line Scan:**
   - Exactly 4 lines per block
   - Exactly 4 characters per line
   - Only `#`, `.`, and `\n` characters
   - Exactly one empty line between blocks
4. **Tetromino Validation:**
   - Exactly 4 `#` characters per block
   - All blocks connected edge-to-edge (flood fill or adjacency count ≥ 6)
5. **Global Counter:** Total tetrominoes ≤ 26

**If any check fails at any point, immediately print `ERROR` and exit.**

---

## Summary Table of Error Categories

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
