package mapindex

import (
	"github.com/fxamacker/cbor"
	"strings"
)

func Load(keyMame string, data []byte) MapIndex {
	res := strings.SplitN(keyMame, "_", 3)
	var mapIndex MapIndex
	switch res[2] {
	case "MapElementID2Refs":
		mapIndex = MapElementID2Refs{}
	case "MapRegion2IDEntry":
		mapIndex = MapRegion2IDEntry{}
	case "MapFeatureID2CurrentObjectIDEntry":
		mapIndex = MapFeatureID2CurrentObjectIDEntry{}
	case "MapID2ItemEntry":
		mapIndex = MapID2ItemEntry{}
	default:
		panic("Unrecognized type")
	}
	err := cbor.Unmarshal(data, &mapIndex)
	if err != nil {
		panic(err)
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
