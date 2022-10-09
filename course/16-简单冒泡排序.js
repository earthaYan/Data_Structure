function bubbleSort(arr) {
  // 上升法:自下而上扫描，最小元素上升到数组顶部
  for (let i = 0; i < arr.length - 1; i++) {
    for (let j = arr.length - 2; j >= i; j--) {
      if (arr[j] > arr[j + 1]) {
        let temp = arr[j];
        arr[j] = arr[j + 1];
        arr[j + 1] = temp;
      }
    }
  }
  console.log('上升法', arr);
  return arr;
}

function bubbleSortPro(arr) {
  // 上升法改进：当没有发生逆序交换即flag=0的时候停止扫描
  let flag = 1;
  while (flag) {
    flag = 0;
    for (let i = arr.length - 2; i >= 0; i--) {
      if (arr[i] > arr[i + 1]) {
        let temp = arr[i];
        arr[i] = arr[i + 1];
        arr[i + 1] = temp;
        flag = 1; //发现交换说明无序
      }
    }
  }
  console.log('上升法升级版：', arr);
  return arr;
}
function bubbleSortDown(arr) {
  // 下降法：自上而下扫描，最大元素下降到数组底部
  for (let i = 0; i < arr.length - 1; i++) {
    for (let j = 0; j <= arr.length - 1; j++) {
      if (arr[j] > arr[j + 1]) {
        let temp = arr[j];
        arr[j] = arr[j + 1];
        arr[j + 1] = temp;
      }
    }
  }
  console.log(arr);
  return arr;
}
let arr1 = [3, 32, 45, 41, 12, 1, 9, 1];
// bubbleSort(arr1);
// bubbleSortPro(arr1);
bubbleSortDown(arr1);
