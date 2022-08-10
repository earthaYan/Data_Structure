/**
 *
 * @param {character[][]} matrix
 * @returns {number}
 */
var maximalSquare = function (matrix) {
  let maxRow = matrix.length;
  let maxCol = matrix[0].length;
  let maxSide = 0;
  if (maxRow === 0 || maxCol === 0) {
    return 0;
  }
  for (let row in matrix) {
    let currentRow = matrix[row];
    for (let col in currentRow) {
      if (currentRow[col] === '1') {
        // 遇到一个 1 作为正方形的左上角
        maxSide = Math.max(maxSide, 1);
        // 计算可能的最大正方形边长
        let currentMaxSide = Math.min(maxRow - row, maxCol - col);
        for (let k = 0; k < currentMaxSide; k++) {
          // 判断新增的一行一列是否均为 1
          let flag = true;
          if (matrix[row + k][col + k] == '0') {
            break;
          }
          for (let m = 0; m < k; m++) {
            if (
              matrix[row + k][col + m] == '0' ||
              matrix[row + m][col + k] == '0'
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
  let maxSquare = Math.pow(maxSide, 2);
  console.log(maxSide, maxSquare);
  return maxSquare;
};
maximalSquare([
  ['1', '0', '1', '0', '0'],
  ['1', '0', '1', '1', '1'],
  ['1', '1', '1', '1', '1'],
  ['1', '0', '0', '1', '0'],
]);
// 这题采用的暴力解法，官方题解提供了动态规划，之后再看
