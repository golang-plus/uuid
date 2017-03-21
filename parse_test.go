package uuid

import (
	"fmt"
	"testing"

	testing2 "github.com/golang-plus/testing"
)

func TestParse(t *testing.T) {
	uuids := []string{
		"00000000-0000-0000-0000-000000000000", // nil
		"945f6800-b463-11e4-854a-0002a5d5c51b", // v1: time based
		"000001f5-b465-21e4-8ffe-98fe945016ea", // v2: dce security
		"6fa459ea-ee8a-3ca4-894e-db77e160355e", // v3: name based md5
		"0f8fad5b-d9cb-469f-a165-70867728950e", // v4: random
		"886313e1-3b8a-5372-9b90-0c9aee199e5d", // v5: name based sha1
	}
	for i, v := range uuids {
		u, err := Parse(v)
		if err != nil {
			t.Fatalf("%s: %s", v, err)
		}

		switch i {
		case 0:
			testing2.ExpectEqual(t, u, Nil)
		case 1:
			testing2.ExpectEqual(t, u.Version(), VersionTimeBased)
		case 2:
			testing2.ExpectEqual(t, u.Version(), VersionDCESecurity)
		case 3:
			testing2.ExpectEqual(t, u.Version(), VersionNameBasedMD5)
		case 4:
			testing2.ExpectEqual(t, u.Version(), VersionRandom)
		case 5:
			testing2.ExpectEqual(t, u.Version(), VersionNameBasedSHA1)
		}

		fmt.Printf("%s, %s, %s\n", u.String(), u.Version(), u.Layout())
	}
}

func TestParseErrors(t *testing.T) {
	uuids := map[string]bool{
		"00000000-0000-0000-0000-000000000000":   true,  // "is an correct pattern",
		"00000000-00000000-0000-0000-00000000":   false, // "is an incorrect pattern",
		"945f6800-b463-11e4-854a-0002a5d5c51b2a": false, // "is too long",
		"000001f5-b465-21h4-8ffe-98fe945016ea":   false, // "contains an invalid character",
		"0f8fad5bd9cb469fa1657086772895":         false, // "is too short",
	}

	i := 0
	for k, v := range uuids {
		_, err := Parse(k)
		testing2.ExpectEqualL(t, err == nil, v, i)
		i++
	}
}
