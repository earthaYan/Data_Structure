/**
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 * 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
 */
// 自己的解法:
var findKthLargestSelf = function (nums, k) {
  // 注意：此题只需要降序排序而不需要去重
  let sortNums = nums.sort((a, b) => b - a);
  return sortNums[k - 1];
};
//官方题解无javascript写法，题解表明主要考察的是快速排序
findKthLargestSelf([3, 2, 3, 1, 2, 4, 5, 5, 6], 4);
