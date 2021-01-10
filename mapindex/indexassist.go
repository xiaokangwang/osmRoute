package mapindex

import (
	"github.com/paulmach/osm"
	"strconv"
	"strings"
)

type MapID2ItemEntry struct {
	ID         osm.ObjectID
	SkipLength int64
}

func (m MapID2ItemEntry) ToIndexKey() string {
	return "Index_MapID2ItemEntry_" + m.ID.String()
}

func (m MapID2ItemEntry) Install(oldVal MapIndex) MapIndex {
	return m
}

type MapElementID2CurrentObjectIDEntry struct {
	ID     osm.ObjectID
	IDElem osm.ElementID
}

func (m MapElementID2CurrentObjectIDEntry) ToIndexKey() string {
	return "Index_MapElementID2CurrentObjectIDEntry_" + m.IDElem.String()
}

func (m MapElementID2CurrentObjectIDEntry) Install(oldVal MapIndex) MapIndex {
	if oldVal.(MapElementID2CurrentObjectIDEntry).ID.Version() > m.ID.Version() {
		return oldVal
	}
	return m
}

type MapRegion2IDEntry struct {
	IDs osm.ObjectIDs
	Lat float64
	Lon float64
}

func (m MapRegion2IDEntry) ToIndexKey() string {
	return "Index_MapRegion2IDEntry_" + InterlacedString(FloatToConstantLengthFloat(m.Lat), FloatToConstantLengthFloat(m.Lon))
}

func (m MapRegion2IDEntry) Install(oldVal MapIndex) MapIndex {
	ety := oldVal.(MapRegion2IDEntry)
	ety.IDs = append(ety.IDs, m.IDs...)
	return ety
}

type MapIndex interface {
	//returned value should be prefixed with type
	ToIndexKey() string
	Install(oldVal MapIndex) MapIndex
}

func InterlacedString(a, b string) string {
	var sb strings.Builder
	if len(a) != len(b) {
		panic("Strings of Different size")
	}
	lenstr := len(a)
	for i := 0; i < lenstr; i++ {
		sb.WriteByte(a[i])
		sb.WriteByte(b[i])
	}

	return sb.String()
}

func FloatToConstantLengthFloat(input float64) string {
	return strconv.FormatFloat(input, 'f', 32, 64)
}
