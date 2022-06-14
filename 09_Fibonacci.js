console.time('common');
let rabbits = new Array(40);
rabbits[0] = 0;
rabbits[1] = 1;
for (let i = 2; i < 40; i++) {
  rabbits[i] = rabbits[i - 1] + rabbits[i - 2];
}
console.log(rabbits);
console.timeEnd('common');
// ==========================================================
console.time('fib');
function Fbi(num) {
  if (num < 2) {
    return num;
  }
  return Fbi(num - 1) + Fbi(num - 2);
}
const rabbitsArr = new Array(40);
for (let i = 0; i < 40; i++) {
  rabbitsArr[i] = Fbi(i);
}
console.log(rabbitsArr);
console.timeEnd('fib');
