function sortByMovingAndFindData(arr) {
  for (let i = 0; i < arr.length; i++) {
    let x = arr[i],
      j;
    for (j = i - 1; j > 0 && x < arr[j]; j--) {
      if (x < arr[j]) {
        arr[j + 1] = arr[j];
      }
    }

    arr[j + 1] = x;
  }
  console.log(arr);
  return arr;
}
let arr1 = [3, 5, 10, 16, 7, 32, 83, 23];
sortByMovingAndFindData(arr1);
