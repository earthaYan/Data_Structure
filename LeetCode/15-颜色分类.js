/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var sortColors = function (nums) {
  for (let i = 0; i < nums.length - 1; i++) {
    for (let j = nums.length - 2; j >= i; j--) {
      if (nums[j] > nums[j + 1]) {
        swap(nums, j, j + 1);
      }
    }
  }
  console.log(nums);
  return nums;
};
function swap(arr, i, j) {
  let temp = arr[i];
  arr[i] = arr[j];
  arr[j] = temp;
}
let arr1 = [3, 32, 45, 41, 12, 1, 9, 1];
sortColors(arr1);
