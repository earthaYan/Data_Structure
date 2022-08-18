function shellSort(arr, increaseArr) {
  for (let h = 0; h < increaseArr.length; h++) {
    let k = increaseArr[h];
    for (let i = k; i < arr.length; i++) {
      let j;
      let x = arr[i];
      for (j = i - k; j >= 0 && x < arr[j]; j -= k) {
        arr[j + k] = arr[j];
      }
      arr[j + k] = x;
    }
  }
  console.log(arr);
  return arr;
}
let arr1 = [32, 45, 41, 12, 1, 9];
shellSort(arr1, [3, 2, 1]);
