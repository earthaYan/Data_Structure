function sortByMiddle(arr) {
  for (let i = 0; i < arr.length; i++) {
    // 准备好待插入元素
    let x = arr[i];
    let left = 0,
      right = i - 1, //设置查找段的起点和终点
      j;
    while (left <= right) {
      //二分定位
      let middle = Math.floor((left + right) / 2);
      if (x < arr[middle]) {
        right = middle - 1;
      } else {
        left = middle + 1;
      }
    }
    // 元素left之后的元素后移
    for (j = i - 1; j >= left; j--) {
      arr[j + 1] = arr[j];
    }
    arr[left] = x;
  }
  console.log(arr);
  return arr;
}
let arr1 = [3, 5, 10, 16, 7, 32, 83, 23];
sortByMiddle(arr1);
