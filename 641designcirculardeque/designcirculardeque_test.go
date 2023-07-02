package _41designcirculardeque

import "testing"

func TestDesigncirculardeque(t *testing.T) {
   obj := Constructor(3)

	 obj.InsertLast(1)
	 obj.InsertLast(2)
     obj.InsertLast(3)
     obj.InsertFront(4)
     obj.GetRear()
     obj.IsFull()
     obj.DeleteLast()
     obj.InsertFront(4)
     obj.GetFront()
     t.Log(obj)
}
