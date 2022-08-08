/**
 * 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
 * 返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
 */
/**
 * @param {number[]} nums
 * @return {number}
 */
// 自己的解法:遍历数组,如果有相同的,删除数据，然后递归
var removeDuplicatesSelf = function (nums) {
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] === nums[i + 1]) {
      nums.splice(i, 1);
      removeDuplicates(nums);
    }
  }
  return nums.length;
};
var removeDuplicatesSelf2 = function (nums) {
  let slow,
    fast = 1;
  if (nums.length === 0) {
    return 0;
  }
  while (fast < nums.length) {
    if (nums[fast] !== nums[fast - 1]) {
      nums[slow] = nums[fast];
      ++slow;
    }
    ++fast;
  }
};
// 官方题解：双指针
var removeDuplicates = function (nums) {
  const n = nums.length;
  if (n === 0) {
    return 0;
  }
  let fast = 1,
    slow = 1;
  while (fast < n) {
    if (nums[fast] !== nums[fast - 1]) {
      nums[slow] = nums[fast];
      ++slow;
    }
    ++fast;
  }
  return slow;
};
