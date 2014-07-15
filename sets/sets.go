package sets

import "github.com/richardlehane/msoleps/types"

type IDName struct {
	ID   uint32
	Name string
}

type PropertySetDef struct {
	FMTID   types.Guid
	IDNames []IDName
}
