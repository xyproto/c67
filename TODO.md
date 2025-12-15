# TODO

- [ ] Let `println` handle multiple arguments.
- [ ] Fix module-level mutable globals in lambdas (closure capture issue at codegen.go:6644).
- [ ] Implement peephole optimization patterns (infrastructure exists in optimizer.go).
- [ ] Implement live range analysis for better register allocation.
- [ ] Implement linear scan register allocation to reduce spilling.
- [ ] Auto-vectorize simple parallel loops using existing SIMD infrastructure.
- [ ] Use BMI1/BMI2 instructions when available on x86-64 (POPCNT, TZCNT, LZCNT).
- [ ] Use conditional select (CSEL) instead of branches on ARM64.
- [ ] Leverage NEON for SIMD operations on ARM64.
- [ ] Use compressed instructions for smaller code on RISC-V.
- [ ] Implement self-hosting bootstrap (compile C67 compiler in C67).
- [ ] Add incremental compilation result caching (hot reload infrastructure exists).
- [ ] Implement polymorphic inline caching for dynamic dispatch optimization.
- [ ] Extend pattern matching to support tuple destructuring and nested patterns.
- [ ] Add exception propagation semantics for defer statements.
- [ ] Verify import system properly initializes closures across modules.
