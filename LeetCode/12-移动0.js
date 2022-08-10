// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
// 自己的解法：灵感来源于08-删除有序数组中的重复项的官方题解：双指针
// 通过双指针先排查出非0的数组元素，索引到slow为止，循环fast，将多出来的元素位置置为0
var moveZeroesSelf = function (nums) {
  const len = nums.length;
  let slow = 0,
    fast = 0;
  while (fast < len) {
    if (nums[fast] !== 0) {
      nums[slow] = nums[fast];
      ++slow;
    }
    ++fast;
  }
  for (let i = slow; i < fast; i++) {
    nums[i] = 0;
  }
};
moveZeroesSelf([0, 1, 0]);
// 官方题解：类似，只给出了python的实现方式
