/**
 * 给定一个头结点为 head 的非空单链表，返回链表的中间结点。
 * 如果有两个中间结点，则返回第二个中间结点。
 */
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * head:{val:1,next:{val:2,next:{val:3,next:{val:4,next:{val:5}}}}}
 * @return {ListNode}
 */
// 自己的解法:两次遍历链表
var middleNodeSelf = function (head) {
  let currentNode = head;
  let listLen = 1;
  while (currentNode.next) {
    currentNode.depth = listLen;
    listLen++;
    currentNode = currentNode.next;
  }
  // 相对官方题解的同样思路,实现上出现了差错,第二次遍历的时候,while的条件应该是k<targetIndex
  let targetIndex =
    listLen % 2 === 0 ? listLen / 2 + 1 : Math.ceil(listLen / 2);
  currentNode = head;
  if (listLen === 1) {
    return currentNode;
  } else if (listLen === 2) {
    return currentNode.next;
  }
  while (currentNode.next) {
    if (currentNode.depth === targetIndex) {
      return currentNode;
    }
    currentNode = currentNode.next;
  }
};
// 官方解法1:
/**
 * 
链表的缺点在于不能通过下标访问对应的元素。
因此我们可以考虑对链表进行遍历，同时将遍历到的元素依次放入数组 A 中。
如果我们遍历到了 N 个元素，那么链表以及数组的长度也为 N，对应的中间节点即为 A[N/2]
 */
var middleNode = function (head) {
  let A = [head];
  while (A[A.length - 1].next != null) A.push(A[A.length - 1].next);
  return A[Math.trunc(A.length / 2)];
};
/**
 * 官方题解2:单指针法
 * 我们可以对方法一进行空间优化，省去数组 A。
我们可以对链表进行两次遍历。
第一次遍历时，我们统计链表中的元素个数 N；
第二次遍历时，我们遍历到第 N/2 个元素（链表的首节点为第 0 个元素）时，将该元素返回即可。
 */
var middleNode = function (head) {
  n = 0;
  cur = head;
  while (cur != null) {
    ++n;
    cur = cur.next;
  }
  k = 0;
  cur = head;
  while (k < Math.trunc(n / 2)) {
    ++k;
    cur = cur.next;
  }
  return cur;
};
