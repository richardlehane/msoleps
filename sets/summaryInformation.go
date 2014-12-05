package sets

import "github.com/richardlehane/msoleps/types"

var SummaryInformation = PropertySetDef{
	types.MustGuidFromString("{F29F85E0-4FF9-1068-AB91-08002B27B3D9}"),
	map[uint32]string{
		0x00000002: "Title",
		0x00000003: "Subject",
		0x00000004: "Author",
		0x00000005: "Keywords",
		0x00000006: "Comments",
		0x00000007: "Template",
		0x00000008: "LastAuthor",
		0x00000009: "RevNumber",
		0x0000000A: "EditTime",
		0x0000000B: "LastPrinted",
		0x0000000C: "CreateTime",
		0x0000000D: "LastSaveTime",
		0x0000000E: "PageCount",
		0x0000000F: "WordCount",
		0x00000010: "CharCount",
		0x00000011: "Thumbnail",
		0x00000012: "AppName",
		0x00000013: "DocSecurity",
	},
}
