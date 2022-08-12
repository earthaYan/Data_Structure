/**
 * 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，
 * 使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
 * 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
 */
/**
 * @param {number[]} nums
 * @return {number}
 */
// 双指针解法（非独立解出）
// 想到了维护统计每个数字出现的次数和可能应该使用双指针
// 但是没想到把这两者结合起来
var removeDuplicates = function (nums) {
  let slow = 0,
    fast = 0,
    count = 1,
    len = nums.length;
  if (len <= 2) return len;
  while (fast < len - 1) {
    fast++;
    if (nums[slow] === nums[fast]) {
      count++;
      if (count < 3) {
        slow++;
        nums[slow] = nums[fast];
      }
    } else {
      slow++;
      nums[slow] = nums[fast];
      count = 1;
    }
  }
  nums.length = slow + 1;
  return nums.length;
};
removeDuplicates([1, 1, 2, 2, 2, 3, 3, 4, 6, 6, 7]);
