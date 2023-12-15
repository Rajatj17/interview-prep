/**
 * Write a Closure function
 * @returns 
 */
function Closure() {
  let count = 0;
  return function counter() {
    return count++
  }
}

// Both the syntaxes are same
const counter = new Closure()
const counter2 = Closure()
console.log(counter())
console.log(counter())

console.log(counter2())
console.log(counter2())

/**
 * Make a set of “ready to use” filters:

  inBetween(a, b) – between a and b or equal to them (inclusively).
  inArray([...]) – in the given array.
  The usage must be like this:

  arr.filter(inBetween(3,6)) – selects only values between 3 and 6.
  arr.filter(inArray([1,2,3])) – selects only elements matching with one of the members of [1,2,3]. 
*/

let arr = [1, 2, 3, 4, 5, 6, 7];

function inBetween(a, b) {
  return (element) => element >= a && element <= b
}

function inArray(list) {
  return (element) => list.includes(element)
}

console.log( arr.filter(inBetween(3, 6)) ); // 3,4,5,6

console.log( arr.filter(inArray([1, 2, 10])) ); // 1,2


/**
 * We’ve got an array of objects to sort:
  let users = [
    { name: "John", age: 20, surname: "Johnson" },
    { name: "Pete", age: 18, surname: "Peterson" },
    { name: "Ann", age: 19, surname: "Hathaway" }
  ];
  The usual way to do that would be:

  // by name (Ann, John, Pete)
  users.sort((a, b) => a.name > b.name ? 1 : -1);

  // by age (Pete, Ann, John)
  users.sort((a, b) => a.age > b.age ? 1 : -1);
  Can we make it even less verbose, like this?

  users.sort(byField('name'));
  users.sort(byField('age'));
 * 
 */

let users = [
  { name: "John", age: 20, surname: "Johnson" },
  { name: "Pete", age: 18, surname: "Peterson" },
  { name: "Ann", age: 19, surname: "Hathaway" }
];

function byField(field) {
  return (a, b) => a[field] > b[field] ? 1 : -1
}

users.sort(byField('name'));
console.log('Sort By Name:\n', users)
users.sort(byField('age'));
console.log('Sort By Age:\n', users)



