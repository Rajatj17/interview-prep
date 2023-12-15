let animal = {
  eats: true,
  walk() {
    console.log('Animal Walk')
  }
}

let rabbit = {
  __proto__: animal,
  jumps: true
}
