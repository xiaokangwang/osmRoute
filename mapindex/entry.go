package mapindex

import (
	"github.com/fxamacker/cbor"
	"strings"
)

func Load(keyMame string, data []byte) MapIndex {
	res := strings.SplitN(keyMame, "_", 3)
	var mapIndex MapIndex
	switch res[1] {
	case "D":
		var unm MapFeatureID2Refs
		err := cbor.Unmarshal(data, &unm)
		if err != nil {
			panic(err)
		}
		mapIndex = unm
	case "C":
		var unm MapRegion2IDEntry
		err := cbor.Unmarshal(data, &unm)
		if err != nil {
			panic(err)
		}
		mapIndex = unm
	case "B":
		var unm MapFeatureID2CurrentObjectIDEntry
		err := cbor.Unmarshal(data, &unm)
		if err != nil {
			panic(err)
		}
		mapIndex = unm
	case "A":
		var unm MapID2ItemEntry
		err := cbor.Unmarshal(data, &unm)
		if err != nil {
			panic(err)
		}
		mapIndex = unm
	case "E":
		var unm MapFTS
		err := cbor.Unmarshal(data, &unm)
		if err != nil {
			panic(err)
		}
		mapIndex = unm
	default:
		panic("Unrecognized type")
	}

	return mapIndex
}

func Store(index MapIndex) (string, []byte) {
	encoded, err := cbor.Marshal(index, cbor.CanonicalEncOptions())
	if err != nil {
		panic(err)
	}
	return index.ToIndexKey(), encoded
}
