// 题目描述
/**
 * 给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
 * 请返回 nums 的动态和。
 **/
/**
 * @param {number[]} nums
 * @return {number[]}
 */
// 自己的算法
var runningSumSelf = function (nums) {
  let sum: number = 0;
  let arr: number[] = [];
  for (let i = 0; i < nums.length; i++) {
    sum = sum + nums[i];
    arr.push(sum);
  }
  return arr;
};
// 官方题解
var runningSum = function (nums) {
  for (let i = 1; i < nums.length; i++) {
    nums[i] = nums[i] + nums[i - 1];
  }
  return nums;
};
