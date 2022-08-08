/**
 * 给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。
 * 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。
 */
/**
 * @param {number} num
 * @return {number}
 */
// 自己的解法:使用while循环,每操作一次step+1,num===0的时候终止循环
var numberOfStepsSelf = function (num) {
  let step = 0;
  while (num !== 0) {
    if (num % 2 === 0) {
      num = num / 2;
    } else {
      num = num - 1;
    }
    step++;
  }
  return step;
};
// 官方题解:将num 与 1进行位运算来判断num 的奇偶性
var numberOfSteps = function (num) {
  let ret = 0;
  while (num > 0) {
    ret += (num > 1 ? 1 : 0) + (num & 0x01);
    num >>= 1;
  }
  return ret;
};
