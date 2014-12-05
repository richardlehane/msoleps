package sets

import "github.com/richardlehane/msoleps/types"

type PropertySetDef struct {
	FMTID types.Guid
	Dict  map[uint32]string
}
