# Notes
1. `for...of` works on `iterables` object. E.g. Array, String, TypedArray. Map, Set, NodeList, arguments, generators, user-defined ones
2. `for await...of` works on async iterables as well as sync
3. `for...in` works on `enumerables string properties` of an object.[ Symbol props ignored ]. It includes inherited enumerable props as well
4. `arguments` object in function is both array-like and iterable, but is not an array. Due to prototypal inheritance. So no arr.methods allowed
5. `arguments` object is not there in case of `arrow` functions.
6. `spread` syntax uses iterators to gather elements
7. `LexicalEnvironment` - every function is associated with the one, the whole script too. It has 2 parts.
   1. `EnvironmentRecord` - object to store all local vars & its props.
   2. `Reference` -  to outer LexicalEnvironment.
8. Initially when the script starts, JS engine knows about variables but they are in uninitialized state.
9. Function definitions are instantly fully initialized. Thats why we can use them even before declaration.
10. All functions remember the lexical environments in which they were made by using a hidden prop [[Environment]]
11. A Variable is updated in the `LE` where it lives.
12. `LE` is `GarbageCollected` after function call finishes. But not if a nested function reference its LE
13. Make a cloure null if not being used anymore to get `GarbageCollected`
14. `var`
    1.  No block scope
    2.  if var is declared inside a function, it becomes a function level var due to old practices (Js had no blocks `LE` previously)
    3.  It tolerates re-declaration.
    4.  Can be declared below their use becuase of `hoisting (raised)`.
    5.  Declarations are hoisted but assignments are not.
15. Hack: Use IIFE to declare and call the function in the same line. E.g. (function(){})(), [!, + - see if u remember]
16. var variables becomes global property & can be accessed using `window.[var]` or `global.[var]`
17. `NFE` - Named Function Expression. ``` let f = function funcName(params) {}```
18. `Functions are objects` too. So we can do add/delete etc
19. To get a function name use `functionName.name`
20. Functions have a `length` prop - it returns a function param count. *[rest params are not counted].
21. `new Function(arg1, arg2, ...., functionBody)` - every argument is string, even a function body.
22. `new Function` - its `[[Environment]]` points to the global one not `LE`.
23. `func.call(context, ...args)` - args: list of arguments.
24. `func.apply(context, args)` - args: array-like objects.
25. `func.bind(context, [arg1], ...)` - context: bind an object to the function call.
26. Arrow functions
    1.  No this (means they borrow this using LE approach)
    2.  can't be called with new
    3.  No arguments
27. ```Object.defineProperty({}, propName, { enumerable: true, writable: true, readable: true, value: val })```
    1.  non-enumerable: not listed in loops
    2.  non-writable: value can be changed
    3.  non-configurable: property can be deleted & these attributes can be modified, otherwise not
28. ```Object.getOwnPropertyDescriptor({}, propName)```
29. Object
    1.  preventExtension - Forbids the addition of new properties to the object. 
    2.  seal - Forbids adding/removing of properties. Sets configurable: false for all existing properties.
    3.  freeze - Forbids adding/removing/changing of properties. 
    4.  isExtensible
    5.  isSealed
    6.  isFrozen
30. 