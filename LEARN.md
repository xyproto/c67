# C67 Language Learning Guide

This document contains everything an AI needs to know to understand, write, and analyze C67 code.

## 1. Core Philosophy: The Universal Map
**CRITICAL**: C67 has exactly **ONE** runtime type: `map[uint64]float64`.

*   **Everything is a map**:
    *   **Numbers**: `{0: value}` (e.g., `42` is `{0: 42.0}`)
    *   **Strings**: `{0: char0, 1: char1, ...}` (e.g., `"Hi"` is `{0: 72.0, 1: 105.0}`)
    *   **Lists**: `{0: item0, 1: item1, ...}` (e.g., `[1, 2]` is `{0: 1.0, 1: 2.0}`)
    *   **Objects/Structs**: `{hash("key"): value, ...}` (e.g., `{x: 1}` is `{hash("x"): 1.0}`)
    *   **Functions**: `{0: code_ptr, 1: closure_data, ...}`
    *   **Pointers**: C pointers are stored as `float64` values in the map.

*   **Implications**:
    *   No primitives. No "int" vs "float" vs "string" at runtime.
    *   All values are IEEE 754 float64s internally.
    *   Type annotations (`x: num`, `s: str`) are **metadata only** for intent and FFI; they do not change runtime representation.

## 2. Syntax & Variables

### Assignment
*   **Immutable (`=`)**: `x = 42` (Cannot be reassigned or modified). **Default for functions.**
*   **Mutable (`:=`)**: `y := 42` (Can be reassigned).
*   **Update (`<-`)**: `y <- 100` (Update value of mutable variable).
*   **Shadowing**: `shadow x = 10` is **required** to shadow a variable from an outer scope.

### Functions & Lambdas
*   **Definition**: `add = (x, y) -> x + y` (Use `=` for function definitions).
*   **Blocks**: `start = -> { ... }`.
*   **Implicit Lambda**: `run = { println("Hi") }` desugars to `run = -> { println("Hi") }` in assignment context.
*   **Return**: Last expression is returned implicitly. Explicit `ret value`.
*   **Variadic**: `(args...) -> ...` (Last param only).

### Types (Annotations)
*   **Native**: `num`, `str`, `list`, `map`.
*   **Foreign (FFI)**: `cint`, `cfloat`, `cstring` (`char*`), `cptr` (`void*`).
*   **Casting**: `val as int32`, `ptr as cstr`.

## 3. Control Flow

### Loops (`@`)
*   **Range**: `@ i in 0..<10 { ... }` (0 to 9). `0..10` is inclusive.
*   **For-each**: `@ item in list { ... }`.
*   **While**: `@ x < 10 { ... }`.
*   **Infinite**: `@ { ... }`.
*   **Labels & Returns**:
    *   `ret @`: Break innermost loop.
    *   `ret @1`: Break outermost loop.
    *   `ret @ 42`: Break loop returning value 42.
    *   `ret`: Return from function.
*   **Max Iterations**: `@ ... max 1000 { ... }` (Safety limit).

### Pattern Matching (`{ ... }`)
*   **Value Match** (Expression before `{`):
    ```c67
    x {
        0 => "zero"
        _ => "other" // or ~>
    }
    ```
*   **Guard Match** (No expression, `|` at start of line):
    ```c67
    {
        | x > 0 => "positive"
        | x < 0 => "negative"
        ~> "zero"
    }
    ```
*   **Mixed Blocks**: Can contain statements before guards.

### Error Handling
*   **Result Type**: Errors are encoded in the map (Type byte `0xE0`).
*   **Check**: `.error` accessor returns error string (e.g., "dv0" for div by zero) or "" on success.
*   **Or Operator (`or!`)**:
    *   `val = risky() or! 0` (Default value).
    *   `val = risky() or! { exit(1) }` (Block execution).
    *   Checks for **NaN** (Error) or **0.0** (Null pointer).

## 4. System Programming & FFI

### C Interop
*   **Import**: `import sdl3 as sdl`. Parses C headers (via DWARF/pkg-config).
*   **Call**: `sdl.SDL_Init(...)`.
*   **Null Pointers**: `0`, `[]`, `{}` are all null pointer (0.0).
*   **Memory**:
    *   `c.malloc(size)`, `c.free(ptr)`.
    *   **Arena**: `arena { ... }`. Allocations inside are freed at end of block.

### CStruct
*   Define C-layout structs for direct memory access.
    ```c67
    cstruct Point { x as float64, y as float64 }
    p = Point(1, 2)
    ```

### Unsafe
*   `unsafe` blocks for raw pointer manipulation (often with `cstruct` or assembly generation).

## 5. Concurrency

### Parallelism
*   **Parallel Loop**: `|| i in 0..10 { ... }`. Spawns processes (fork).
*   **Parallel Map**: `list || x -> heavy_work(x)`.

### ENet Channels (Message Passing)
*   **Address**: `&8080`, `&"localhost:9000"`.
*   **Send**: `&8080 <- "msg"`.
*   **Receive**: `msg <= &8080`.

## 6. Object-Oriented Programming (OOP)
*   **Classes** are syntactic sugar for maps + closures.
*   **`class Point { ... }`**: Creates a constructor function returning a map.
*   **`init = ...`**: Constructor method.
*   **`.field`**: Access instance field (`this.field`).
*   **`. `**: Returns `this` (e.g., for chaining).
*   **Composition**: `class Robot <> Walker <> Talker { ... }`. Mixes in behavior maps. No inheritance.

## 7. Operators & Gotchas

### Operators
*   **Bitwise**: Must use `b` suffix: `&b`, `|b`, `<<b`.
*   **Pipe**: `data | transform`.
*   **Composition**: `f <> g` (Do `g`, then `f`).
*   **Ownership**: `µ` (prefix) for move semantics/ownership transfer.
*   **Random**: `??` (CSPRNG).

### Gotchas
*   **Shadowing**: You MUST use `shadow x = ...` to shadow. `x = ...` will error if `x` exists in outer scope.
*   **Loop Labels**: `@` is innermost, `@1` is outermost.
*   **Types**: Don't rely on type checks. Check structure or values.
*   **Semicolons**: Only needed for multiple statements on one line.
*   **Comments**: `//` only.

## 8. Memory Management
*   **No Garbage Collection**.
*   **Defer**: `defer cleanup()` (LIFO execution).
*   **Arenas**: Preferred for high performance.
*   **Manual**: `c.free` if using `c.malloc`.

## 9. Comparison Guide

| Feature | C67 | Go | Rust | C |
| :--- | :--- | :--- | :--- | :--- |
| **Type System** | Universal Map | Statically Typed | Statically Typed | Statically Typed |
| **GC** | No (Arenas/Manual) | Yes | No (Ownership) | No (Manual) |
| **Generics** | Implicit (Duck Typing) | Explicit | Explicit | N/A (Void*) |
| **Concurrency** | ENet / Fork | Goroutines | Threads/Async | Threads |
| **Error Handling** | `or!`, Result map | `if err != nil` | `Result<T, E>` | Return codes |
| **OOP** | Composition (`<>`) | Composition | Traits | Structs |

## 10. Example Code Snippet

```c67
import sdl3 as sdl

class Game {
    init = {
        .window = sdl.SDL_CreateWindow("Game", 800, 600, 0) or! {
            exitf("Error: %s", sdl.SDL_GetError())
        }
    }

    run = {
        defer sdl.SDL_DestroyWindow(.window)
        @ {
            // Game Loop
            | should_quit() => ret @
        }
    }
}

main = {
    sdl.SDL_Init(sdl.SDL_INIT_VIDEO)
    defer sdl.SDL_Quit()

    game := Game()
    game.run()
}
```
