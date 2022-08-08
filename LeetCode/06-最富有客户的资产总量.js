/**
 * @param {number[][]} accounts
 * @return {number}
 */
// 自己的解法:遍历二维数组获取每一个的资产总量,然后返回最大值
var maximumWealthSelf = function (accounts) {
  let result = [];
  for (let accountIndex in accounts) {
    const res = accounts[accountIndex].reduce((previous, current) => {
      return previous + current;
    }, 0);
    result.push(res);
  }
  return Math.max(...result);
};
// 官方题解
var maximumWealth = function (accounts) {
  let maxWealth = -Number.MAX_VALUE;
  for (const account of accounts) {
    maxWealth = Math.max(
      maxWealth,
      account.reduce((a, b) => a + b)
    );
  }
  return maxWealth;
};
maximumWealth([
  [1, 2, 3],
  [3, 2, 1],
]);
