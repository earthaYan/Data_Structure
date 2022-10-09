// 递归方法
function FibD(n) {
  let f;
  if (n < 2) {
    f = 1;
  } else {
    f = FibD(n - 2) + FibD(n - 1);
  }
  console.log(f);
  return f;
}
const res = FibD(5);

