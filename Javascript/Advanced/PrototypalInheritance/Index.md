Summary:

1. **Prototypal Inheritance:**
   - Objects in JavaScript have a hidden [[Prototype]] property, which can be either an object or null.
   - The [[Prototype]] property is accessed using `__proto__`.
   - When a property is not found in an object, JavaScript automatically looks for it in its prototype, creating a chain known as the prototype chain.

2. **Setting [[Prototype]]:**
   - The `__proto__` property is used to set the prototype of an object explicitly.
   - Circular references are not allowed, and the value of `__proto__` can be an object or null.

3. **Inheritance in Action:**
   - Properties and methods of the prototype become automatically available in the object.
   - Methods in the prototype can be called on the inheriting object.

4. **Writing and Deleting:**
   - Write and delete operations work directly with the object, not involving the prototype.
   - Accessor properties may behave differently.

5. **__proto__ vs. Prototype:**
   - `__proto__` is a getter/setter for [[Prototype]] and is a bit outdated.
   - Modern JavaScript recommends using `Object.getPrototypeOf` and `Object.setPrototypeOf` instead.

6. **Value of "this":**
   - The value of "this" inside a method is not affected by prototypes. It always refers to the object before the dot in a method call.

7. **for...in Loop:**
   - The `for...in` loop iterates over both own and inherited properties.
   - `hasOwnProperty` can be used to filter out inherited properties.

8. **Key/Value Methods:**
   - Methods like `Object.keys` and `Object.values` operate only on the object itself, ignoring inherited properties.

9. **Tasks:**
   - Tasks involve understanding property lookup in prototype chains, assigning prototypes in a specific order, and fixing issues related to prototype inheritance.

10. **F.prototype and [[Prototype]]:**
   - When new objects are created with a constructor function (e.g., `new F()`), the `F.prototype` property is used to set the `[[Prototype]]` of the new object.
   - The value of `F.prototype` should be an object or null; other values won't work.

11. **Setting Prototype with Example:**
   - Example: `Rabbit.prototype = animal;`
   - This sets the `[[Prototype]]` of the new Rabbit objects to the `animal` object.

12. **F.prototype Property Usage:**
   - `F.prototype` is only used when `new F` is called to assign `[[Prototype]]` to the new object.
   - Changing `F.prototype` after object creation won't affect existing objects.

13. **Default F.prototype and Constructor Property:**
   - Every function has a default `prototype` property, even if not supplied.
   - The default `prototype` is an object with a `constructor` property pointing back to the function itself.

14. **constructor Property Usage:**
   - The `constructor` property in the prototype can be used to create a new object using the same constructor as an existing one.

15. **Caution with constructor Property:**
   - JavaScript does not ensure the correct value of the `constructor` property.
   - Replacing the default prototype entirely removes the `constructor` property.

16. **Recommendations for Preserving constructor:**
   - To preserve the `constructor` property, add/remove properties to the default prototype instead of overwriting it entirely.
   - Alternatively, recreate the `constructor` property manually when overwriting the prototype.

17. **Native Prototypes:**
   - The "prototype" property is utilized by built-in constructor functions in JavaScript.
   - Object.prototype is at the top of the prototype chain for most built-in objects.

18. **Object.prototype:**
   - `Object.prototype` is the prototype for objects created using `{}` or `new Object()`.
   - The `toString` method, for example, is available through the prototype chain.

19. **Built-in Object Prototypes:**
   - Arrays, Dates, Functions, and other built-in objects also have prototypes.
   - The prototype chain for built-ins includes `Object.prototype`.

20. **Inheritance Chain Example:**
   - Example with an array: `arr.__proto__ === Array.prototype`.

21. **Method Overlapping:**
   - Prototypes may have overlapping methods, and the closest one in the chain is used.

22. **Functions as Objects:**
   - Functions are also objects with their methods taken from `Function.prototype`.

23. **Primitives and Wrappers:**
   - String, Number, and Boolean have temporary wrapper objects with methods like `String.prototype`.
   - `null` and `undefined` have no object wrappers or prototypes.

24. **Changing Native Prototypes:**
   - Native prototypes can be modified, but it's generally discouraged due to global conflicts.
   - Modifying native prototypes is acceptable in polyfilling, where a method is manually added for compatibility.

25. **Polyfilling:**
   - Polyfilling involves adding substitute methods to native prototypes for unsupported functionalities.

26. **Borrowing from Prototypes:**
    - Methods from native prototypes can be borrowed and added to other objects.
    - This is useful for adding functionalities to non-array objects.