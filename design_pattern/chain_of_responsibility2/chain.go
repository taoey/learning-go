package chain_of_responsibility

import (
	"fmt"
)

// 抽象审批者动作接口
type IApprover interface {
	SetSuccessor(approver IApprover)
	ProcessRequest(request PurchaseRequest)
}

// 采购单
type PurchaseRequest struct {
	Amount  float64 // 采购金额
	Number  int     // 采购单编号
	Purpose string  // 采购目的
}

// ----------------------具体审批者：主任-------------------
type Director struct {
	IApprover
	Name string
}

func newDirector(name string) *Director {
	return &Director{
		Name: name,
	}
}
func (this *Director) SetSuccessor(approver IApprover) {
	this.IApprover = approver
}

// 处理5万以下采购单
func (this *Director) ProcessRequest(request PurchaseRequest) {
	if request.Amount < 50000 {
		fmt.Println(fmt.Sprintf("主任%s审批采供但单：%+v", this.Name, request))
	} else {
		this.IApprover.ProcessRequest(request)
	}
}

// -------------------具体审批者：董事会----------------------
type Congress struct {
	IApprover
	Name string
}

func newCongress(name string) *Congress {
	return &Congress{
		Name: name,
	}
}

func (this *Congress) SetSuccessor(approver IApprover) {
	this.IApprover = approver
}

// 处理50万以上采购单
func (this *Congress) ProcessRequest(request PurchaseRequest) {
	if request.Amount >= 500000 {
		fmt.Println(fmt.Sprintf("董事会%s审批采供但单：%+v", this.Name, request))
	}
}
