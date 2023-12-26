# Event Loop

## Phases
> timers
> pending callbacks
> idle/prepare
> poll
> check
> close callback

`process.nextTick` and `setImmediate` are both mechanisms in Node.js for scheduling the execution of a function in the next iteration of the event loop, but they have some key differences in terms of when they execute and their priority in the event loop.

### `process.nextTick`:

1. **Timing:**
   - `process.nextTick` callbacks are executed immediately after the current operation completes, before the event loop moves to the next phase.

2. **Event Loop Phase:**
   - Callbacks scheduled using `process.nextTick` are placed at the front of the event loop queue. They are executed before I/O events, timers, or other types of events.

3. **Priority:**
   - `process.nextTick` callbacks have a higher priority than other types of callbacks scheduled by `setTimeout` or `setImmediate`. This means that if there are both `process.nextTick` and `setImmediate` callbacks, the `process.nextTick` callbacks will be executed first.

4. **Recursive Calls:**
   - If a `process.nextTick` callback itself calls `process.nextTick`, it forms a recursive chain, and all callbacks in the chain are executed before moving on to the next phase of the event loop.

5. **Example:**
   ```javascript
   process.nextTick(() => {
     console.log('process.nextTick callback');
   });

   console.log('Main script execution');
   ```

### `setImmediate`:

1. **Timing:**
   - `setImmediate` callbacks are executed in the next iteration of the event loop, but after I/O events.

2. **Event Loop Phase:**
   - `setImmediate` callbacks are placed in a separate queue and are executed after I/O events but before timers.

3. **Priority:**
   - `setImmediate` callbacks have a lower priority compared to `process.nextTick` callbacks. If both are present, `process.nextTick` callbacks are executed first.

4. **Recursive Calls:**
   - `setImmediate` callbacks do not form a recursive chain like `process.nextTick`. If a `setImmediate` callback itself schedules another `setImmediate`, the second one is placed in the next iteration of the event loop.

5. **Example:**
   ```javascript
   setImmediate(() => {
     console.log('setImmediate callback');
   });

   console.log('Main script execution');
   ```

### Choosing Between `process.nextTick` and `setImmediate`:

- **Use `process.nextTick` when:**
  - You want to ensure that a callback is executed immediately after the current operation.
  - You want the callback to have a higher priority than `setImmediate` callbacks.
  - You want to create a recursive chain of callbacks.

- **Use `setImmediate` when:**
  - You want to ensure that a callback is executed in the next iteration of the event loop, after I/O events.
  - You want the callback to have a lower priority than `process.nextTick` callbacks.
  - You want to avoid creating a recursive chain of callbacks.

It's important to note that while `process.nextTick` and `setImmediate` serve similar purposes, they are not interchangeable, and the choice between them depends on the specific requirements of your code and when you want the callback to be executed in the event loop.