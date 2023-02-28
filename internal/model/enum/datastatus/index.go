package datastatus

type Value int

const (
	// Normal 正常
	Normal Value = iota
	// Invalid 失效
	Invalid
	// Deleted 删除
	Deleted
)
