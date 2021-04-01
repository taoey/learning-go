package load_balance

type LoadBalance interface {
	Add(...string) error
	Get(string) (string, error) // string 参数为一致性hash时使用
}
