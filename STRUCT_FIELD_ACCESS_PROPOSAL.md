# Language Improvement Proposal: Type-Aware Struct Field Access

## Problem

Currently, `event.type` fails when `event` is a raw `c.malloc()` pointer because:

1. Compiler doesn't track that `event` points to `SDL_Event`
2. FieldAccessExpr falls back to C67 map lookup
3. Raw memory bytes ≠ C67 map → segfault

## Current Workarounds

### 1. Manual Memory Reading (peek32/peek8)
```c67
event = c.malloc(192)
event_type = peek32(event, 0)  // Read uint32 at offset 0
```
**Status**: Implemented but has bugs in codegen

### 2. CStruct Declaration
```c67
cstruct SDL_Event {
    type as uint32
}
event = SDL_Event()  // Allocates typed struct
event_type = event.type  // Works! Compiler knows layout
```
**Status**: Partially works, but SDL_Event() constructor not implemented

## Proposed Solution: Type Annotations for Pointers

Add syntax to annotate pointer types:

```c67
// Syntax 1: Type annotation on assignment
event: SDL_Event = c.malloc(192)
event_type = event.type  // Compiler knows event is SDL_Event*, can read at offset 0

// Syntax 2: Cast with type
event = c.malloc(192) as SDL_Event
event_type = event.type

// Syntax 3: Declare with type before assignment
event: SDL_Event
event = c.malloc(192)
event_type = event.type
```

### Implementation Plan

1. **Parser**: Accept type annotations on variable declarations
   ```c67
   name: TypeName = expression
   ```

2. **Type Tracking**: Store variable→type mapping in compiler
   ```go
   fc.varTypes map[string]string  // "event" → "SDL_Event"
   ```

3. **FieldAccessExpr Codegen**: Check if object has known struct type
   ```go
   if varType, ok := fc.varTypes[varName]; ok {
       if cstruct, exists := fc.cstructs[varType]; exists {
           // Use direct memory access at known offset
       }
   }
   ```

4. **CStruct Lookup**: Use existing `fc.cstructs` map for field offsets

### Benefits

- **Zero-cost**: No runtime overhead
- **Type-safe**: Compiler validates field names
- **Natural syntax**: `event.type` just works
- **Backward compatible**: Untyped pointers still work
- **Reuses existing infra**: cstruct, FieldAccessExpr already implemented

### Example Usage

```c67
import sdl3 as sdl
import libc as c

cstruct SDL_Event {
    type as uint32,
    timestamp as uint64
}

// With type annotation - enables field access
event: SDL_Event = c.malloc(192)
sdl.SDL_PollEvent(event)

// Now this works!
event_type = event.type  // Read at offset 0
timestamp = event.timestamp  // Read at offset 4

c.free(event)
```

### Alternative: Infer Type from cstruct Constructor

```c67
cstruct Point { x as int32, y as int32 }

// Auto-generate constructor that returns typed pointer
p = Point()  // Returns typed pointer, size from Point.size
p.x = 10     // Works! Compiler knows p is Point*
```

## Recommendation

Implement **Type Annotations** (Syntax 1) because:
- Most explicit and clear
- Works with any allocation method (malloc, mmap, etc.)
- Aligns with existing `: type` syntax
- Can be optional (graceful degradation)

This would make C67's C FFI integration seamless while maintaining zero-cost abstractions.
