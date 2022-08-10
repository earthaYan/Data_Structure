/**
 *
 * @param {character[][]} matrix
 * @returns {number}
 */
var maximalSquare = function (matrix) {
  let maxRow = matrix.length;
  let maxCol = matrix[0].length;
  let maxSide = 0;
  if (!matrix || maxRow === 0 || maxCol === 0) {
    return 0;
  }
  for (let row = 0; row < maxRow; row++) {
    for (let col = 0; col < maxCol; col++) {
      // 找到正方形左上角
      if (matrix[row][col] === '1') {
        maxSide = Math.max(maxSide, 1);
        // 计算可能的最大正方形边长
        let currentMaxSide = Math.min(maxRow - row, maxCol - col);
        // 判断新增的一行一列是否均为 1
        for (let k = 0; k < currentMaxSide; k++) {
          let flag = true;
          if (matrix[row + k][col + k] === '0') {
            break;
          }
          for (let m = 0; m < k; m++) {
            // 判断当前是否是正方形边界线
            if (
              matrix[row + k][col + m] === '0' ||
              matrix[row + m][col + k] === '0'
            ) {
              flag = false;
              break;
            }
          }
          if (flag) {
            maxSide = Math.max(maxSide, k + 1);
          } else {
            break;
          }
        }
      }
    }
  }
  return Math.pow(maxSide, 2);
};
maximalSquare([
  ['1', '0', '1', '0', '0'],
  ['1', '0', '1', '1', '1'],
  ['1', '1', '1', '1', '1'],
  ['1', '0', '0', '1', '0'],
]);
// 这题采用的暴力解法，官方题解提供了动态规划，之后再看
