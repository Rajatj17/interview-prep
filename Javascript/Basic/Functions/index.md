Summary:

1. **Function as Objects:**
   - In JavaScript, functions are objects, and they can be treated as callable "action objects."
   - Functions can be called and manipulated like other objects in JavaScript.

2. **"name" Property:**
   - Functions have a "name" property that represents the name of the function.
   - The name can be accessed using the `functionName.name` syntax.

3. **"length" Property:**
   - Functions have a built-in "length" property that returns the number of parameters the function expects.
   - Rest parameters are not counted in the "length" property.

4. **Custom Properties:**
   - Custom properties can be added to functions. These properties can be used to store additional information or track the state of the function.

5. **Function Properties vs. Variables:**
   - Properties assigned to a function do not define local variables inside it. Function properties and variables are distinct.

6. **Named Function Expression (NFE):**
   - Named Function Expressions allow functions to reference themselves internally.
   - The name is not visible outside the function, providing a reliable way for the function to call itself.

7. **Tasks:**
   - Modify the `makeCounter` function to allow setting, decreasing, and getting the counter value.
   - Implement a `sum` function that can handle an arbitrary number of brackets for adding numbers.

In summary, functions in JavaScript are versatile objects with properties like "name" and "length." They can have custom properties and can be used to create powerful patterns like closures and named function expressions. Understanding these concepts is crucial for advanced work with functions in JavaScript.