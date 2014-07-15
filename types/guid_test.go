package types

import "testing"

func TestGuid(t *testing.T) {
	g, err := GuidFromString("{F29F85E0-4FF9-1068-AB91-08002B27B3D9}")
	if err != nil {
		t.Error(err)
	}
	r := g.String()
	if r != "{F29F85E0-4FF9-1068-AB91-08002B27B3D9}" {
		t.Errorf("GUID round trip failed, expecting {F29F85E0-4FF9-1068-AB91-08002B27B3D9}, got %v", r)
	}
}
