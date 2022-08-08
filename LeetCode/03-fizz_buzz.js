/**
 * 给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，
 * 并用字符串数组 answer（下标从 1 开始）返回结果，其中：
answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
answer[i] == "Fizz" 如果 i 是 3 的倍数。
answer[i] == "Buzz" 如果 i 是 5 的倍数。
answer[i] == i （以字符串形式）如果上述条件全不满足
 */
/**
 * @param {number} n
 * @return {string[]}
 */
// 自己的算法，通过能否取余判断是否需要替换成对应字符串
var fizzBuzzSelef = function (n) {
  let answer = new Array(n).fill(1);
  for (let i = 1; i <= n; i++) {
    if (i % 3 === 0 && i % 5 === 0) {
      answer[i - 1] = 'FizzBuzz';
    } else if (i % 3 === 0) {
      answer[i - 1] = 'Fizz';
    } else if (i % 5 === 0) {
      answer[i - 1] = 'Buzz';
    } else {
      answer[i - 1] = i;
    }
  }

  return answer;
};
// 官方题解
var fizzBuzz = function (n) {
  const answer = [];
  for (let i = 1; i <= n; i++) {
    const sb = [];
    if (i % 3 === 0) {
      sb.push('Fizz');
    }
    if (i % 5 === 0) {
      sb.push('Buzz');
    }
    if (sb.length === 0) {
      sb.push(i);
    }
    answer.push(sb.join(''));
  }
  console.log(answer);
  return answer;
};
fizzBuzz(15);
