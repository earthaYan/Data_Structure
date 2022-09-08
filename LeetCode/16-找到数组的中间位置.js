/**
 * @param {number[]} nums
 * @return {number}
 */
var findMiddleIndex = function (nums) {
  if (nums.length === 1) {
    return 0;
  } else {
    let total = nums.reduce((prev, cur) => {
      return prev + cur;
    }, 0);
    let sum = 0;
    for (let i = 0; i < nums.length; i++) {
      if (sum === total - nums[i] - sum) {
        return i;
      } else {
        sum = sum + nums[i];
      }
    }
  }
  return -1;
};
let arr1 = [1, 3, 2, 4];
findMiddleIndex(arr1);
