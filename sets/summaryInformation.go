package sets

import "github.com/richardlehane/msoleps/types"

var SummaryInformation = PropertySetDef{
	types.MustGuidFromString("{F29F85E0-4FF9-1068-AB91-08002B27B3D9}"),
	[]IDName{
		IDName{0x00000002, "Title"},
		IDName{0x00000003, "Subject"},
		IDName{0x00000004, "Author"},
		IDName{0x00000005, "Keywords"},
		IDName{0x00000006, "Comments"},
		IDName{0x00000007, "Template"},
		IDName{0x00000008, "LastAuthor"},
		IDName{0x00000009, "RevNumber"},
		IDName{0x0000000A, "EditTime"},
		IDName{0x0000000B, "LastPrinted"},
		IDName{0x0000000C, "CreateTime"},
		IDName{0x0000000D, "LastSaveTime"},
		IDName{0x0000000E, "PageCount"},
		IDName{0x0000000F, "WordCount"},
		IDName{0x00000010, "CharCount"},
		IDName{0x00000011, "Thumbnail"},
		IDName{0x00000012, "AppName"},
		IDName{0x00000011, "DocSecurity"},
	},
}
