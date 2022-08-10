/**
 * 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
 * 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
 * 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素
 */
/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
// 自己的解法：同12类似，只是把0换成了常数
var removeElementSelf = function (nums, val) {
  const len = nums.length;
  let slow = 0,
    fast = 0;
  while (fast < len) {
    if (nums[fast] !== val) {
      nums[slow] = nums[fast];
      slow++;
    }
    fast++;
  }
  nums.length = nums.length - (fast - slow);
  return nums.length;
};
removeElementSelf([3, 2, 2, 3], 3);
