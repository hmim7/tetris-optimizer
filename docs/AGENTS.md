# Project Agenda - Tetris-Optimizer

---

## 1. Project Overview & Mission

The mission is to create a Go CLI program that arranges tetrominoes into the smallest possible square.
The tool reads tetromino definitions from text files and computes optimal placement using backtracking algorithms.

You are a **Senior Go Developer**. Your mission is to implement a robust tetromino solver using a modular, pipeline-based architecture while ensuring strict input validation and deterministic output.

You must strictly follow:
- `docs/PRD.md`
- `docs/error_cases.md`, `docs/golden-tests.md`
- Modular Pipeline Architecture

---

## 2. Technical Strategy & Constraints

- **Language:** Go (Golang)
- **Packages:** Go Standard Library only. No external dependencies.
- **Architecture:** Linear pipeline with fail-fast validation + backtracking solver.
- **Logic:** Prefer deterministic placement order; avoid randomization.
- **Input Format:** Strict 4x4 grid validation with connectivity checking.
- **File I/O:** Read input files once into memory for processing.
- **Error Handling:** Respect exact `ERROR` output behavior from `PRD.md` section 4.6.
- **Comments:** Use valid Go comments (`//` or `/* */`) only when needed.

Core validation rules:
- **Grid Structure:** Each tetromino must be exactly 4x4 characters.
- **Character Set:** Only `#`, `.`, and `\n` characters allowed.
- **Block Count:** Exactly 4 `#` characters per tetromino.
- **Connectivity:** All blocks must be connected edge-to-edge (no diagonal-only).
- **Separation:** Exactly one empty line between tetromino definitions.

Solver rules:
- Use depth-first backtracking for tetromino placement.
- Start with minimal square size based on total block count.
- Increment board size if no solution exists at current size.
- Assign letters A-Z based on input file order.
- Fill empty spaces with `.` character.
- Ensure deterministic placement order for consistent results.

---

## 3. Development Workflow

- Run `gofmt` before every commit.
- Run `go test ./...` after each module integration.
- Keep functions small and single-purpose.
- Validate all error cases produce exactly `ERROR` output.
- Test performance against time limits specified in acceptance criteria.

---

## 4. Execution Protocol

Before final delivery, follow these phases mapped to PRD milestones.

### Phase 1: Foundation (Milestones 1-2)
- Validate CLI argument handling and file reading.
- Implement strict input validation with immediate error detection.
- Confirm all error cases output exactly `ERROR`.

### Phase 2: Core Logic (Milestones 3-4)
- Implement tetromino parsing and validation.
- Add connectivity checking using flood fill or adjacency counting.
- Implement coordinate normalization for tetromino representation.
- Add comprehensive unit tests for validation logic.

### Phase 3: Solver Implementation (Milestones 4-5)
- Implement backtracking algorithm for tetromino placement.
- Add minimal square size calculation.
- Ensure deterministic placement order.
- Optimize for performance requirements (≤1s, ≤3s, ≤5s).

### Phase 4: Validation & Delivery (Milestone 6)
- Validate against all test cases in `docs/golden-tests.md`.
- Confirm acceptance criteria alignment with `docs/PRD.md` section 6.
- Verify performance meets specified time limits.
- Update README with usage examples.

---

## 5. Project Structure

The repository should follow this structure:

```text
.
|-- docs/
|   |-- PRD.md
|   |-- AGENTS.md
|   |-- error_cases.md
|   |-- golden-tests.md
|   `-- tetrominoes.md
|-- main.go
|-- samples/
|   |-- goodexample00.txt
|   |-- goodexample01.txt
|   |-- goodexample02.txt
|   |-- goodexample03.txt
|   |-- hardexam.txt
|   |-- badexample00.txt
|   |-- badexample01.txt
|   |-- badexample02.txt
|   |-- badexample03.txt
|   |-- badexample04.txt
|   `-- [additional error test files]
|-- unit-tests/
|   |-- input_test.go
|   |-- parser_test.go
|   |-- validator_test.go
|   |-- model_test.go
|   |-- solver_test.go
|   `-- output_test.go
|-- modules/
|   |-- input.go
|   |-- parser.go
|   |-- validator.go
|   |-- model.go
|   |-- solver.go
|   `-- output.go
`-- README.md
```

Responsibility map:
- `main.go`: CLI orchestration and pipeline coordination.
- `input.go`: Command-line argument validation and file reading.
- `parser.go`: Raw input splitting into tetromino blocks.
- `validator.go`: Format validation and connectivity checking.
- `model.go`: Tetromino representation and coordinate normalization.
- `solver.go`: Backtracking algorithm and board management.
- `output.go`: Final square rendering and letter assignment.

---

## 6. Task Card Template

All tasks in `./tasks/taskxx_<short-title>.txt` must follow this structure:

```text
TASK ID: TASKXX
TITLE: [Descriptive Title]
PRIORITY: [High/Medium/Low]
STATUS: TODO

DESCRIPTION:
[Detailed explanation of logic and purpose]

REQUIREMENTS:
- [Technical requirement 1]
- [Technical requirement 2]
- [Standard library only]

FILES TO CREATE/MODIFY:
- .modules/[filename].go
- .unit-tests/[filename]_test.go

GOLDEN TESTS / ACCEPTANCE CRITERIA:
- Input: [Example] -> Expected: [Result]

DEFINITION OF DONE:
- Unit tests pass
- gofmt applied
- Performance requirements met
```

---

## 7. Unit Tests

- Every module must have a corresponding `*_test.go`.
- Test all error cases to ensure exactly `ERROR` output.
- Test valid inputs against expected square outputs.
- Include performance benchmarks for solver algorithm.
- Use clear expected vs actual test logs.
- Keep tests in `unit-tests/`.

---

## 8. Acceptance & Audit Alignment

- Core functionality is defined in `docs/PRD.md` section 6.1.
- Error handling is defined in `docs/PRD.md` sections 6.2-6.5.
- Performance requirements are defined in `docs/PRD.md` section 6.6.
- Output format validation is defined in `docs/PRD.md` section 6.7.
- Use PRD checkbox placeholders as the execution checklist.

---

## 9. README.md Requirements

README must include:
- Project description and tetromino solver purpose.
- Usage examples with sample input files.
- Input format specifications (4x4 grids, connectivity rules).
- Error handling behavior and `ERROR` output cases.
- Performance characteristics and time limits.
- Short technical note on backtracking algorithm approach.
- Example commands and expected output snippets.