package chain_of_responsibility

import "fmt"

// 采购单
type PurchaseRequest struct {
	Amount  float64 // 采购金额
	Number  int     // 采购单编号
	Purpose string  // 采购目的
}

// 抽象审批者动作接口
type ApproverAction interface {
	SetNext(approver ApproverAction)
	ProcessRequest(request *PurchaseRequest)
}

// ----------------抽象审批者----------------------
type Approver struct {
	ApproverAction
	Name string // 审批者姓名
}

// 抽象审批者构造函数
func newApprover(name string) Approver {
	return Approver{
		Name: name,
	}
}

// 设置后继者：所有具体审批者共用此方法
func (this *Approver) SetNext(approver ApproverAction) {
	this.ApproverAction = approver
}

// 抽象请求处理方法：具体审批者需要单独设置审批流程
func (this *Approver) ProcessRequest(request *PurchaseRequest) {
	fmt.Println("默认处理")
}

// ----------------------具体审批者：主任-------------------
type Director struct {
	Approver
}

func newDirector(name string) Director {
	return Director{
		Approver: newApprover(name),
	}
}

// 处理5万以下采购单
func (this *Director) ProcessRequest(request *PurchaseRequest) {
	if request.Amount < 50000 {
		fmt.Println(fmt.Sprintf("主任%s审批采供但单：%+v", this.Name, request))
	} else {
		// this.Approver.ApproverAction.ProcessRequest(request)
		this.Approver.ApproverAction.ProcessRequest(request)
	}
}

// ------------------具体审批者：副董事长----------------------
type VicePresident struct {
	Approver
}

func newVicePresident(name string) VicePresident {
	return VicePresident{
		Approver: newApprover(name),
	}
}

// 处理10万以下采购单
func (this *VicePresident) ProcessRequest(request *PurchaseRequest) {
	if request.Amount < 100000 {
		fmt.Println(fmt.Sprintf("副董事长%s审批采供但单：%+v", this.Name, request))
	} else {
		this.Approver.ApproverAction.ProcessRequest(request)
	}
}

// -------------------具体审批者：董事长-----------------------
type President struct {
	Approver
}

func newPresident(name string) President {
	return President{
		Approver: newApprover(name),
	}
}

// 处理50万以下采购单
func (this *President) ProcessRequest(request *PurchaseRequest) {
	if request.Amount < 500000 {
		fmt.Println(fmt.Sprintf("董事长%s审批采供但单：%+v", this.Name, request))
	} else {
		this.Approver.ApproverAction.ProcessRequest(request)
	}
}

// -------------------具体审批者：董事会----------------------
type Congress struct {
	Approver
}

func newCongress(name string) Congress {
	return Congress{
		Approver: newApprover(name),
	}
}

// 处理50万以上采购单
func (this *Congress) ProcessRequest(request *PurchaseRequest) {
	if request.Amount >= 500000 {
		fmt.Println(fmt.Sprintf("董事会%s审批采供但单：%+v", this.Name, request))
	}
}
