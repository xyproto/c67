# TODO

## Priority 0: CRITICAL - Windows PE Arena Initialization Crash

**CURRENT STATUS:** Significant progress made, but arena programs still crash.

### What Works ✅
1. Minimal programs (`main = { 42 }`) - EXIT CODE 42 ✅
2. Variable arithmetic (`x = 10; y = 32; x + y`) - EXIT CODE 42 ✅
3. Function calls (`add(20, 22)`) - EXIT CODE 42 ✅
4. PC relocations for arena metadata - NO MORE 0xDEADBEEF ✅
5. Arena init function generation - CALL patched correctly ✅
6. No compiler errors or warnings ✅

### What Doesn't Work ❌
- Programs with arenas crash with ACCESS_VIOLATION (0xC0000005)
- Arena init function appears to be called but crashes during execution

### Investigation
**Root cause identified:** Arena metadata symbols (_vibe67_arena_meta, etc.) were not defined when `usesArenas` was checked too early.

**Fix applied:**
1. Moved symbol definitions to `generateRuntimeHelpers()` (after code gen)
2. Created `generateArenaInitFunction()` - standalone function for init
3. Patch 5 NOPs at offset with CALL to `_vibe67_init_arenas`
4. All relocations now patched correctly

**Still failing:**
- Runtime crash in arena init function execution
- Likely issues: shadow space, calling conventions, HeapAlloc usage
- Need WinDbg debugging to see exact crash location

### Next Steps
1. Debug with WinDbg to find exact crash instruction
2. Verify shadow space allocation in arena init function
3. Check Windows calling convention compliance (rcx, rdx, r8, r9)
4. Test simpler arena init without HeapAlloc (use malloc only)
5. Compare working Linux mmap version with Windows HeapAlloc version

---

## Priority 1: Fixes and verifications

- [ ] Fix variable scope tracking in lambda compilation (module-level mutable globals).
- [ ] Verify import system properly initializes closures across modules.

## Priority 2: Demoscene & Size Optimization (<8KB Goal)

The goal is to enable the creation of competitive 64k intros (Linux/x86_64) using SDL3/RayLib. "Hello World" should be <1KB, not 21KB.

- [ ] **Tiny ELF Writer ("-tiny" flag)**
    - [ ] Implement custom ELF header generation (overlapping headers/segments).
    - [ ] Remove page alignment padding (align=1, disable standard 0x1000 alignment).
    - [ ] Merge `.text`, `.data`, and `.rodata` into a single `RX` or `RWX` segment.
    - [ ] Implement `DT_HASH` usage for symbol resolution (smaller than `DT_GNU_HASH`).
- [ ] **Dead Code Elimination (DCE)**
    - [ ] Implement function-level reachability analysis.
    - [ ] Strip unused global variables and constants.
    - [ ] Aggressively remove unused runtime helper functions (e.g., FMA checks if FMA unused).
- [ ] **Asset Compression**
    - [ ] Finish the built-in decompressor stub (LZ4 or custom simple algorithm).
    - [ ] Allow embedding compressed resources directly into the `.text` segment.
- [ ] **Shader Minification**
    - [ ] Add support for embedding and minifying GLSL strings at compile time.

## Priority 2: Language quality and tooling

- [ ] **Developer Tooling (Critical)**
    - [ ] **Language Server Protocol (LSP)**: Implement a basic LSP for VS Code/Neovim (Go-to-definition, simple completions).
    - [ ] **Debug Info**: Generate DWARF v5 debug information for GDB/LLDB support.
    - [ ] **Formatter**: Implement `c67 fmt` for canonical code style.
- [ ] **Compiler Correctness & Robustness**
    - [ ] **Fix Unsafe Bug**: Fix register assignment limitation (`rax <- ptr`) to allow raw memory iteration.
    - [ ] **Register Allocation**: Upgrade from simple allocator to Linear Scan or Graph Coloring for denser code.
    - [ ] **Fuzzing**: Set up fuzz testing for the parser to prevent crashes on invalid input.
- [ ] **Performance Proof**
    - [ ] Create a benchmark suite comparing C67 vs C (gcc -O2/-O3) vs Go.
    - [ ] Optimize the `match` compiler to generate jump tables for density/speed.

## Priority 3: Language Features

Refining the "Vibe" into a rigorous specification.

- [ ] **Safety & Types**
    - [ ] Implement `µ` operator semantics for explicit memory ownership/movement.
    - [ ] Add `??` null coalescing operator and `?` optional type suffix.
    - [ ] Implement compile-time division-by-zero checks.
- [ ] **Metaprogramming**
    - [ ] "Comptime" evaluation: Execute pure C67 functions at compile time to generate constants (tables, sin/cos LUTS).
- [ ] **Advanced Pattern Matching**
    - [ ] Tuple destructuring: `(x, y) = point`.
    - [ ] Nested patterns: `[[a, b], c] = list`.

## Priority 3: Platform & Architecture

- [ ] **Linux/ARM64**: Polish the ARM64 backend to parity with x86_64.
- [ ] **Windows/x86_64**: Fix code generation gap causing crashes (see Priority 0)
- [ ] **macOS**: Finish Mach-O support.

## Priority 5: Self-hosting

The ultimate proof of language quality.

- [ ] Write a basic C67 parser in C67.
- [ ] Write the Code Generator in C67.
- [ ] Bootstrap: Use the Go compiler to compile the C67 compiler, then use that to compile itself.

## Priority 6: Nice to have and optimizations

- [ ] Implement Windows decompressor stub with VirtualAlloc (LOW PRIORITY - compression is optional).
- [x] Implement print/println for Windows using printf. ✅ COMPLETE
- [ ] Implement proper import table generation for PE files.
- [ ] Optimize O(n²) string iteration in codegen.
