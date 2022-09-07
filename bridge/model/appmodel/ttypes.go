package appmodel

import "bridge/bridge/pbs/apppb"

type RuntimeType string

var (
	Python RuntimeType = "Python"
	Golang RuntimeType = "Golang"
)

func (r RuntimeType) ToPBEnum() apppb.App_RuntimeType {
	// TODO: RuntimeType 和 apppb.App_RuntimeType 是一一对应的，
	// 所以这里不做检查，如果对应不上，程序会直接 Panic 退出
	runtime := apppb.App_RuntimeType_value[string(r)]
	return apppb.App_RuntimeType(runtime)
}
