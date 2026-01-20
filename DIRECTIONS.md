# Vibe67 Development Directions

## Core Philosophy

Vibe67 aims to be a **minimal, orthogonal, and beautiful** systems programming language for game development and demoscene productions. Every feature should serve these goals.

## Memory Management

**Prefer arena allocators over malloc/free:**
- Use `arena { ... }` blocks with `alloc()` calls for temporary allocations
- Automatic cleanup on scope exit provides memory safety
- Only use `c.malloc`/`c.free` when interfacing with C libraries that require it
- Arena allocators call kernel directly (mmap/VirtualAlloc), no libc dependency

**Memory Safety:**
- Bounds checking where practical
- Safe buffer abstractions when needed
- Prefer immutable data structures

## Data Access

**Prefer high-level abstractions:**
- Use `list[offset]` for array/list access
- Use map syntax `map[key]` for map access
- Only use low-level operations (`peek32`, `poke32`, etc.) when absolutely necessary for performance or FFI

## Syntax Design

**Minimal and orthogonal:**
- Every feature should have one obvious way to do it
- Avoid redundant syntax
- Keywords should be intuitive and serve clear purposes
- Prefer expressions over statements

**Beautiful code:**
- Clean, readable syntax
- Consistent patterns across the language
- Minimal punctuation noise
- Self-documenting when possible

## Target Use Cases

1. **Game Development** - Steam-ready executables, SDL3/RayLib5 integration
2. **Demoscene** - Minimal binary sizes (< 10KB for realistic programs)
3. **Arch Linux Utilities** - System tools and utilities (bonus goal)

## Performance Goals

- Direct machine code generation (no IR)
- Small binaries through aggressive DCE
- Automatic optimizations (FMA, pure function memoization)
- No GC pauses (manual memory management)
- Competitive with C/C++ performance

## C FFI Integration

- Use `cstruct` for C-compatible structures
- Use `as` casting at Vibe67/C boundaries
- Import C libraries with `import` syntax
- Raw pointers by default in C FFI (no wrapper overhead)
- Link libc only when C FFI is used

## Type System

**Universal type:** Everything is `map[uint64]float64` at runtime
- Compile-time type annotations for safety (`: num`, `: str`, `: cptr`, etc.)
- No runtime type overhead
- Use `as` casting only at C FFI boundaries

## Immutability

**Immutable by default:**
- Use `=` for immutable variables
- Use `:=` for mutable variables
- Update with `<-`, `++`, `--` operators
- Encourages functional programming patterns

## Error Handling

**Explicit and minimal:**
- Use `or!` operator for fallback values
- Error blocks for complex error handling
- No exceptions (systems programming)

## Standard Library

**Minimal but complete:**
- Math functions (via C FFI when needed)
- String utilities
- Collection operations
- File I/O (when needed)
- Network primitives (future)

## Development Priorities

1. **Stability** - Fix all known bugs before adding features
2. **Cross-platform** - Windows, Linux, macOS support
3. **Documentation** - Keep GRAMMAR.md and LANGUAGESPEC.md up to date
4. **Examples** - Real-world examples for all features
5. **Performance** - Profile and optimize hot paths
6. **Binary size** - Keep executables minimal

## What NOT to Do

- Don't add features without clear use cases
- Don't compromise binary size for convenience
- Don't add libc dependencies without good reason
- Don't break orthogonality for syntax sugar
- Don't optimize prematurely (profile first)
- Don't add runtime overhead for safety features (compile-time checks instead)

## Future Directions

- Hot reload for game development
- SIMD vectorization improvements
- Better debugging support (dwarf info)
- Package manager (maybe)
- Build system integration
- Editor tooling (LSP)
