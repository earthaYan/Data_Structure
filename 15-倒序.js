function reverse(n, str) {
  const strArr = str.split(' ');
  for (let i = 0; i < n; i++) {
    if (i < n / 2) {
      let temp = strArr[i];
      strArr[i] = strArr[n - 1 - i];
      strArr[n - 1 - i] = temp;
    }
  }
  console.log(strArr.join(','));
  return strArr.join(',');
}
reverse(5, '1 2 3 4 5');
