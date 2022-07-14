   /**
   * Definition for singly-linked list.
     * type ListNode struct {
     *     Val int
     *     Next *ListNode
     * }
     */
   func removeElements(head *ListNode, val int) *ListNode {
      dummy := head
      prev , curr := dummy , head
      for curr != nil {
          if dummy.Val == val {
              dummy = dummy.Next
          }
          nxt := curr.Next
          if curr.Val == val {
              prev.Next = nxt
          } else {
              prev = curr
          }
          curr = nxt
      }
      return dummy
   }
