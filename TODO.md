# TODO

## Executable Size Optimization (for 64k demos)

### Current Status
- Minimal program (x := 42): 45KB
- Code segment: 36KB
- Data segment: ~1KB
- Removed unused debug strings: ~500 bytes saved

### Size Reduction Tasks
- [x] Add function usage tracking (usedFunctions map)
- [x] Add emission flags for all runtime functions (foundation laid)
- [ ] Make runtime functions conditionally included (only when used)
  - [x] String concatenation (_c67_string_concat) - wrapped in conditional
  - [ ] String print/println (_c67_string_print, _c67_string_println)
  - [ ] String equality (_c67_string_eq)
  - [ ] String slicing (c67_slice_string)
  - [ ] String conversions (c67_string_to_cstr, cstr_to_c67_string)
  - [ ] List functions (_c67_list_cons, _c67_list_head, _c67_list_tail, _c67_list_length, _c67_list_index, _c67_list_update)
  - [ ] List operations (_c67_list_concat, _c67_list_repeat)
  - [ ] Arena allocator functions (c67_arena_create, c67_arena_alloc, c67_arena_destroy, c67_arena_reset, _c67_arena_ensure_capacity)
  - [ ] Cache functions (c67_cache_lookup, c67_cache_insert)
  - [ ] Printf runtime (full format string parser)
  - [ ] Itoa (_c67_itoa - number to string conversion)
  - [ ] Print syscalls (_c67_print_syscall, _c67_println_syscall - Linux only)
- [ ] Track which runtime functions are actually needed per program
- [ ] Wire emission flags to usage tracking
- [ ] Remove or minimize ELF headers overhead
- [ ] Implement dead code elimination pass
- [ ] Strip unnecessary alignment padding
- [ ] Optimize common patterns (e.g., initialization code)
- [ ] Add `-tiny` flag for demo-optimized builds
- [ ] Target: <8KB for minimal "Hello World"

## Language Features from Design Decisions

### Operator Implementation
- [ ] Implement `µ` operator semantics for memory ownership/movement

### Safety Features
- [ ] Implement optional types with None/Some semantics
- [ ] Add compile-time null safety checks
- [ ] Add division by zero checks
- [ ] Implement stack overflow detection
- [ ] Add integer overflow detection options

### Defer Statement Enhancements
- [ ] Define exception propagation semantics for defer
- [ ] Implement defer stack unwinding on error
- [ ] Add exception propagation semantics for defer statements
- [ ] Add defer ordering guarantees in documentation

### Module-level mutable globals in lambdas
- [ ] Fix variable scope tracking in lambda compilation
- [ ] Ensure mutable globals are properly referenced through rbp

### Register Allocation Improvements
- [ ] Implement live range analysis for better register allocation
- [ ] Add register reuse hints based on live ranges
- [ ] Implement linear scan register allocation to reduce spilling

### Import System
- [ ] Add test for cross-module closure initialization
- [ ] Fix closure variable capture in imported modules
- [ ] Verify import system properly initializes closures across modules
- [ ] Add circular dependency detection

## Architecture-Specific

### ARM64 Optimizations
- [ ] Add CSEL instruction support in ARM64 backend
- [ ] Replace conditional branches with CSEL where beneficial
- [ ] Use conditional select (CSEL) instead of branches on ARM64
- [ ] Add NEON instruction wrappers in ARM64 backend
- [ ] Implement NEON SIMD for vector operations
- [ ] Leverage NEON for SIMD operations on ARM64

### RISC-V Optimizations
- [ ] Add compressed instruction support in RISC-V backend
- [ ] Implement 16-bit instruction encoding for common operations
- [ ] Use compressed instructions for smaller code on RISC-V
- [ ] Add branch compression optimization

## Advanced Features (Future)

### Self-hosting Bootstrap
- [ ] Compile basic C67 parser in C67
- [ ] Compile C67 lexer in C67
- [ ] Compile C67 code generator in C67
- [ ] Implement self-hosting bootstrap (compile C67 compiler in C67)

### Advanced Optimizations
- [ ] Add call site profiling infrastructure
- [ ] Implement method lookup cache for polymorphic calls
- [ ] Add cache invalidation on type changes
- [ ] Implement polymorphic inline caching for dynamic dispatch optimization

### Pattern Matching Enhancements
- [ ] Add support for tuple pattern matching: `(x, y) = tuple`
- [ ] Add support for nested pattern matching: `[[a, b], c] = nested_list`
- [ ] Add support for pattern guards in match expressions
- [ ] Extend pattern matching to support tuple destructuring and nested patterns

### Incremental Compilation
- [ ] Add file change detection (hot reload infrastructure exists)
- [ ] Implement function-level compilation cache
- [ ] Add dependency tracking between compilation units
- [ ] Add incremental compilation result caching

## Fix Core Language Issues

### Issue 1: Mixed Statement-Guard Blocks
**Status**: Not yet needed - current design works well
**Note**: Blocks can be either statement blocks OR match blocks, not mixed. This is clean and unambiguous.

### Issue 2: Ackermann Function (Doubly-Nested Recursive Calls)
**Status**: Known bug (see ERRORS.md)
**Root Cause**: Unknown - possibly memoization or arg evaluation order

Tasks:
- [ ] Create minimal test case for nested recursive calls
- [ ] Add debug logging to recursive call codegen
- [ ] Trace register allocation during nested calls
- [ ] Check if memoization interferes with multi-arg recursion
- [ ] Verify argument evaluation order (left-to-right vs right-to-left)
- [ ] Fix and add regression test

### Issue 3: List Building with Recursive Concatenation
**Status**: Known bug (see ERRORS.md)
**Symptom**: Returns 0 instead of built list

Tasks:
- [ ] Create minimal test case for list concatenation
- [ ] Trace list concatenation codegen
- [ ] Check if list + list operation works correctly
- [ ] Verify accumulator handling in recursive contexts
- [ ] Fix and add regression test

## Code Quality Improvements

### Performance
- [ ] Optimize O(n²) string iteration (codegen.go:10624)

### Code Generation
- [ ] Add explicit float precision conversions (codegen.go:5692)
- [ ] Implement length parameter for string operations (codegen.go:5821)
- [ ] Replace malloc with arena allocation for strings (codegen.go:6491, 7400)
- [ ] Add proper map iteration for string extraction (codegen.go:16853)

### Platform Support
- [ ] Implement Windows decompressor stub with VirtualAlloc (decompressor_stub.go:66)
- [ ] Implement ARM64 decompressor stub (compress.go:262)
- [ ] Implement proper import table generation for PE (pe.go:426)
- [ ] Implement RISC-V PLT generation (pltgot_rv64.go:11)

### Feature Completeness
- [ ] Implement function composition operator `<>` (codegen.go:16770)
- [ ] Handle "host:port" format in network operations (codegen.go:16801)
- [ ] Implement proper transformations for match expressions (codegen.go:18012, 18020)
- [ ] Re-enable blocks-as-arguments feature (parser.go:3689, 4030)
- [ ] Re-enable compression with debugged decompressor (default.go:144)
# Known Errors and Limitations

This document tracks known errors, limitations, and edge cases in the C67 compiler.

## Language Limitations

### 1. Doubly-Nested Recursive Calls with Multiple Arguments

**Status**: FIXED ✅
**Date Fixed**: 2025-12-16

**Description**: Functions with multiple arguments that make recursive calls where one argument is itself a recursive call now work correctly.

**Example**:
```c67
// Ackermann function - NOW WORKS CORRECTLY
ack = (m, n) {
    | m == 0 => n + 1
    | n == 0 => ack(m - 1, 1)
    ~> ack(m - 1, ack(m, n - 1))
}

ack(3, 3)  // Returns 61 (correct!)
ack(3, 4)  // Returns 125 (correct!)
```

**Fix**: Register allocation for nested function calls was corrected to properly save/restore registers.

---

### 2. List Display in Strings

**Status**: Partially Implemented
**Date**: 2025-12-16

**Description**: Lists can be created and manipulated, but when converted to strings (e.g., in f-strings or println), they show as a placeholder "[...]" instead of their actual contents.

**Example**:
```c67
x = [1, 2, 3]
println(f"x = {x}")  // Prints: x = [...]
```

**Status**: Lists work internally but need proper string conversion implementation.

**Note**: List concatenation and other list operations work correctly. This only affects display/printing.

---

## Reserved Keywords

The following are reserved keywords and cannot be used as variable names:

- `max` - Reserved keyword
- `min` - Reserved keyword (likely)

**Error Example**:
```c67
// ERROR: max is a reserved keyword
primes_helper = (current, max, acc) {  // Syntax error
    ...
}
```

**Fix**: Use alternative names like `limit`, `maximum`, `upper_bound`, etc.

---

## Notes for Developers

- Memoization is automatically applied to pure single-argument functions
- Multi-argument functions do not currently benefit from memoization
- Tail-call optimization works for properly tail-recursive functions
- Forward references work via automatic function reordering

---

## Testing

The following benchmark programs demonstrate working and problematic patterns:

### Working Benchmarks:
- `factorial.c67` - Recursive and tail-recursive factorial (both work perfectly)
- `fib.c67` / `fib_bench.c67` - Fibonacci with automatic memoization
- `primes.c67` - Prime counting (modified to avoid list building)

### Problematic Examples:
- `ackermann.c67` - Demonstrates the doubly-nested recursive call issue
