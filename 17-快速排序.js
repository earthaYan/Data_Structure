function quickSort(arr) {
  var len = arr.length,
    index,
    pivot,
    left = [],
    right = [];
  if (len <= 1) return arr;
  index = Math.floor(len / 2);
  pivot = arr[index];
  len--;
  for (var i = 0; i < len; i++) {
    if (arr[i] < pivot) {
      left.push(arr[i]);
    } else {
      right.push(arr[i]);
    }
  }
  const res = quickSort(left).concat(pivot, quickSort(right));
  return res;
}

function quickSort2(arr, left, right) {
  var len = arr.length,
    partitionIndex,
    left = typeof left !== 'number' ? 0 : left,
    right = typeof right !== 'number' ? len - 1 : right;
  // 递归的终止条件；
  if (left < right) {
    partitionIndex = partition(arr, left, right);
    quickSort2(arr, left, partitionIndex - 1);
    quickSort2(arr, partitionIndex + 1, right);
  }
  return arr;
}

function partition(arr, left, right) {
  var pivot = left,
    index = pivot + 1;
  for (var i = index; i <= right; i++) {
    if (arr[i] < arr[pivot]) {
      swap(arr, i, index);
      index++;
    }
  }
  swap(arr, pivot, index - 1);
  return index - 1;
}

function swap(arr, i, j) {
  var temp = arr[i];
  arr[i] = arr[j];
  arr[j] = temp;
}
let arr1 = [3, 32, 45, 41, 12, 1, 9, 1];
let res = quickSort2(arr1);
console.log(arr1, res);
