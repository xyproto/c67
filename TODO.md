# TODO

## Completed
- [x] Let `println` handle multiple arguments.
- [x] Implement peephole optimization patterns (infrastructure exists in optimizer.go).
  - Added double negation optimization: not(not(x)) -> x != 0
  - Added De Morgan's law: (not x) and (not y) -> not(x or y)
  - Added absorption laws: x and (x or y) -> x, x or (x and y) -> x
  - Added boolean algebra simplifications (and/or with true/false)
- [x] Add register pressure tracking to identify spill-heavy code
- [x] Add vector width detection for target platform
- [x] Add test cases to reproduce the closure capture issue
- [x] Add CPUID detection for BMI1/BMI2 support
- [x] Implement POPCNT instruction for bit counting
- [x] Implement TZCNT instruction for trailing zeros
- [x] Implement LZCNT instruction for leading zeros
- [x] Use BMI1/BMI2 instructions when available on x86-64
- [x] Add SIMD intrinsics for common operations
- [x] Implement simple register use/def analysis
- [x] Create data structures for live ranges
- [x] Add loop analysis to detect vectorization candidates
- [x] Implement simple loop dependency analysis
- [x] Auto-vectorize simple parallel loops using existing SIMD infrastructure
  - Supports +, -, * operations on arrays
  - AVX2 support (256-bit, 4 doubles per vector)
  - AVX-512 support (512-bit, 8 doubles per vector) via EnableAVX512 flag
  - Automatic scalar cleanup loop for non-aligned sizes
  - Pattern matching for: result[i] = a[i] OP b[i]

## High Priority - Multi-File Compilation

### Overview
Currently, `c67 file1.c67 file2.c67 -o output` doesn't work as expected. The goal is to enable explicit multi-file compilation where files are concatenated and compiled together, similar to how `gcc file1.c file2.c -o output` works.

### Root Cause Analysis
- [ ] Investigate why sibling file loading mechanism conflicts with explicit multi-file args
- [ ] Determine if parser state is properly reset between files when concatenating
- [ ] Check if lambda parameters are lost during multi-file source combination
- [ ] Verify that the `NewParser()` constructor properly handles concatenated source

### Implementation Steps
1. **File Argument Parsing (cli.go)**
   - [x] Update `cmdBuild()` to collect multiple input files from args
   - [x] Filter out flags (-o, etc.) to get clean file list
   - [x] Add multi-file detection logic (len(inputFiles) > 1)

2. **Multi-File Compiler Function (codegen.go)**
   - [ ] Create `CompileC67MultipleFiles()` that takes []string of file paths
   - [ ] Read each file with `os.ReadFile()` and check for errors
   - [ ] Concatenate sources with proper separators (newlines, not comments)
   - [ ] Use single `NewParser()` call on combined source
   - [ ] Disable sibling file auto-loading when explicit files are provided
   - [ ] Use `NewC67Compiler()` for proper initialization (not manual struct)

3. **Parser State Management**
   - [ ] Verify parser handles multi-file source correctly
   - [ ] Check that lambda parameter scoping works across file boundaries
   - [ ] Ensure line number tracking remains accurate for error reporting
   - [ ] Test that imports/uses in different files don't conflict

4. **Testing & Validation**
   - [ ] Create test case: `add.c67` (defines function) + `main.c67` (calls function)
   - [ ] Test with verbose mode to see file loading sequence
   - [ ] Verify lambda parameters compile correctly in multi-file mode
   - [ ] Test edge cases: circular dependencies, duplicate definitions
   - [ ] Add integration tests for multi-file compilation

5. **Advanced Features**
   - [ ] Add `--no-siblings` flag to disable automatic sibling loading
   - [ ] Support glob patterns: `c67 *.c67 -o output`
   - [ ] Add dependency ordering (topological sort) if needed
   - [ ] Implement parallel file reading for large projects

### Debugging Tips
- Use `-verbose` to see which files are loaded and in what order
- Check combined source with debug print before parsing
- Verify that `fc.variables` map contains expected function names
- Test single-file versions of each component file first
- Use `strace` to see file I/O system calls

### Known Issues
- Lambda parameters appear as "undefined variable" in multi-file mode
- Sibling loading may interfere with explicit file arguments
- Comment separators (// ---- filename ----) might break parser (though C67 supports // comments)

## High Priority - Executable Size Optimization (for 64k demos)

### Current Status
- Minimal program (x := 42): 45KB
- Code segment: 36KB
- Data segment: ~1KB  
- Removed unused debug strings: ~500 bytes saved

### Size Reduction Tasks
- [ ] Make runtime functions conditionally included (only when used)
  - [ ] Arena allocator code (if no `alloc` used)
  - [ ] Bounds checking code (if no array access)
  - [ ] Recursion depth tracking (if no recursion)
  - [ ] Loop iteration limiting (if no loops)
- [ ] Remove or minimize ELF headers overhead
- [ ] Implement dead code elimination pass
- [ ] Strip unnecessary alignment padding
- [ ] Optimize common patterns (e.g., initialization code)
- [ ] Add `-tiny` flag for demo-optimized builds
- [ ] Target: <8KB for minimal "Hello World"

## High Priority - Language Features from Design Decisions

### Operator Implementation
- [x] Implement `~` as bitwise NOT operator (in addition to `!`)
- [x] Add `µ` token for memory ownership/movement
- [ ] Implement `µ` operator semantics for memory ownership/movement
- [ ] Add `?` suffix for optional types (e.g., `x?: int`)
- [ ] Implement `.?` safe navigation operator
- [ ] Add `??` null coalescing operator

### Safety Features
- [ ] Implement optional types with None/Some semantics
- [ ] Add compile-time null safety checks
- [x] Implement bounds checking for array access
- [ ] Add division by zero checks
- [ ] Implement stack overflow detection
- [ ] Add integer overflow detection options

### Defer Statement Enhancements
- [ ] Define exception propagation semantics for defer
- [ ] Implement defer stack unwinding on error
- [ ] Add exception propagation semantics for defer statements
- [ ] Add defer ordering guarantees in documentation

### SIMD and Vectorization
- [x] Add loop analysis to detect vectorization candidates
- [x] Implement simple loop dependency analysis
- [x] Add vector width detection for target platform
- [x] Auto-vectorize simple parallel loops using existing SIMD infrastructure
- [x] Add SIMD intrinsics for common operations

### Module-level mutable globals in lambdas
- [x] Add test cases to reproduce the closure capture issue
- [ ] Fix variable scope tracking in lambda compilation
- [ ] Ensure mutable globals are properly referenced through rbp

### Register Allocation Improvements
- [x] Add register pressure tracking to identify spill-heavy code
- [x] Implement simple register use/def analysis
- [x] Create data structures for live ranges
- [ ] Implement live range analysis for better register allocation
- [ ] Add register reuse hints based on live ranges
- [ ] Implement linear scan register allocation to reduce spilling

### Import System
- [ ] Add test for cross-module closure initialization
- [ ] Fix closure variable capture in imported modules
- [ ] Verify import system properly initializes closures across modules
- [ ] Add circular dependency detection

## Architecture-Specific

### x86-64 Optimizations
- [x] Add CPUID detection for BMI1/BMI2 support
- [x] Implement POPCNT instruction for bit counting
- [x] Implement TZCNT instruction for trailing zeros
- [x] Implement LZCNT instruction for leading zeros
- [x] Use BMI1/BMI2 instructions when available on x86-64

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
