# TODO

## High Priority - Executable Size Optimization (for 64k demos)

### Current Status
- Minimal program (x := 42): 45KB
- Code segment: 36KB
- Data segment: ~1KB
- Function naming harmonized: all runtime functions now use `_c67_` prefix
- Removed duplicate emit* flags: using only usedFunctions map for tracking

### Size Reduction Tasks
- [ ] Make runtime functions conditionally included (only when used)
  - [ ] String operations (_c67_string_print, _c67_string_println, _c67_string_eq, _c67_slice_string)
  - [ ] List functions (_c67_list_cons, _c67_list_head, _c67_list_tail, _c67_list_length, _c67_list_index, _c67_list_update, _c67_list_concat, _c67_list_repeat)
  - [ ] Arena allocator functions (_c67_arena_create, _c67_arena_alloc, _c67_arena_destroy, _c67_arena_reset, _c67_arena_ensure_capacity)
  - [ ] Cache functions (_c67_cache_lookup, _c67_cache_insert)
  - [ ] Printf runtime (full format string parser)
  - [ ] Itoa (_c67_itoa - number to string conversion)
  - [ ] Print syscalls (_c67_print_syscall, _c67_println_syscall - Linux only)
- [ ] Remove or minimize ELF headers overhead
- [ ] Implement dead code elimination pass
- [ ] Strip unnecessary alignment padding
- [ ] Optimize common patterns (e.g., initialization code)
- [ ] Add `-tiny` flag for demo-optimized builds
- [ ] Target: <8KB for minimal "Hello World"

## High Priority - Language Features from Design Decisions

### Operator Implementation
- [ ] Implement `µ` operator semantics for memory ownership/movement
- [ ] Add `?` suffix for optional types (e.g., `x?: int`)
- [ ] Implement `.?` safe navigation operator
- [ ] Add `??` null coalescing operator

### Safety Features
- [ ] Implement optional types with None/Some semantics
- [ ] Add compile-time null safety checks
- [ ] Add division by zero checks
- [ ] Implement stack overflow detection
- [ ] Add integer overflow detection options

### Defer Statement Enhancements
- [ ] Define exception propagation semantics for defer
- [ ] Implement defer stack unwinding on error
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

### x86-64 Optimizations
- (All basic BMI1/BMI2 optimizations are implemented)

### ARM64 Optimizations
- [ ] Add CSEL instruction support in ARM64 backend
- [ ] Replace conditional branches with CSEL where beneficial
- [ ] Add NEON instruction wrappers in ARM64 backend
- [ ] Implement NEON SIMD for vector operations

### RISC-V Optimizations
- [ ] Add compressed instruction support in RISC-V backend
- [ ] Implement 16-bit instruction encoding for common operations
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
