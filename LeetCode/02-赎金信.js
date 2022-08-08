/**
 * 给你两个字符串：ransomNote 和 magazine ，
 * 判断 ransomNote 能不能由 magazine 里面的字符构成。
 * 如果可以，返回 true ；否则返回 false 。
 * magazine 中的每个字符只能在 ransomNote 中使用一次。
 */
/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 */
// 自己的算法:比较每个字符在目标串和源串的次数
var canConstructSelf = function (ransomNote, magazine) {
  // ransomNote:目标字符串
  // magazine:源字符串
  const targetObj = {};
  const sourceObj = {};
  ransomNote.split('').forEach((item) => {
    if (Object.keys(targetObj).includes(item)) {
      targetObj[item] = targetObj[item] + 1;
    } else {
      targetObj[item] = 1;
    }
  });
  magazine.split('').forEach((item) => {
    if (Object.keys(sourceObj).includes(item)) {
      sourceObj[item] = sourceObj[item] + 1;
    } else {
      sourceObj[item] = 1;
    }
  });
  return Object.keys(targetObj).every((key) => {
    return targetObj[key] <= sourceObj[key];
  });
};
// 官方题解
/**
 * 
只需要满足字符串 magazine 中的每个英文字母'a'-'z'的统计次数都大于等于 ransomNote 中相同字母的统计次数即可。
1.如果字符串magazine 的长度小于字符串 ransomNote 的长度，
则可以肯定magazine 无法构成ransomNote，此时直接返回false。
2.首先统计magazine 中每个英文字母 aa 的次数cnt[a]，
再遍历统计ransomNote 中每个英文字母的次数，
如果发现 ransomNote 中存在某个英文字母 cc 的统计次数大于 magazine 中该字母统计次数 cnt[c]
则直接返回false。
 */
var canConstruct = function (ransomNote, magazine) {
  if (ransomNote.length > magazine.length) {
    return false;
  }
  const cnt = new Array(26).fill(0);
  for (const c of magazine) {
    cnt[c.charCodeAt() - 'a'.charCodeAt()]++;
  }
  for (const c of ransomNote) {
    cnt[c.charCodeAt() - 'a'.charCodeAt()]--;
    if (cnt[c.charCodeAt() - 'a'.charCodeAt()] < 0) {
      return false;
    }
  }
  return true;
};
