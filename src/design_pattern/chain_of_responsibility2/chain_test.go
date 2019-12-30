package chain_of_responsibility

import (
	"testing"
)

func TestChain(t *testing.T) {
	user1 := newDirector("u1")
	meeting := newCongress("meeting")

	user1.SetSuccessor(meeting)

	p1 := PurchaseRequest{45000, 100, "购买4.5万物品"}
	user1.ProcessRequest(p1)

	p2 := PurchaseRequest{60000, 100, "购买6万物品"}
	user1.ProcessRequest(p2)

	p3 := PurchaseRequest{160000, 100, "购买16万物品"}
	user1.ProcessRequest(p3)

	p4 := PurchaseRequest{800000, 100, "购买80万物品"}
	user1.ProcessRequest(p4)

}
