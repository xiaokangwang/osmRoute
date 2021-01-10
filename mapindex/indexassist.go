package mapindex

import (
	"github.com/paulmach/osm"
	"modernc.org/sortutil"
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

type MapFeatureID2CurrentObjectIDEntry struct {
	ID     osm.ObjectID
	IDFeat osm.FeatureID
}

func (m MapFeatureID2CurrentObjectIDEntry) ToIndexKey() string {
	return "Index_MapElementID2CurrentObjectIDEntry_" + m.IDFeat.String()
}

func (m MapFeatureID2CurrentObjectIDEntry) Install(oldVal MapIndex) MapIndex {
	if oldVal.(MapFeatureID2CurrentObjectIDEntry).ID.Version() > m.ID.Version() {
		return oldVal
	}
	return m
}

type MapRegion2IDEntry struct {
	IDs  osm.ObjectIDs
	FIDs osm.FeatureIDs
	Lat  float64
	Lon  float64
}

func (m MapRegion2IDEntry) ToIndexKey() string {
	return "Index_MapRegion2IDEntry_" + InterlacedString(FloatToConstantLengthFloat(m.Lat), FloatToConstantLengthFloat(m.Lon))
}

func (m MapRegion2IDEntry) Install(oldVal MapIndex) MapIndex {
	ety := m
	if oldVal != nil {
		ety = oldVal.(MapRegion2IDEntry)
	}
	ety.IDs = append(ety.IDs, m.IDs...)
	sortutil.Dedupe(ObjectIDSlice(ety.IDs))
	ety.FIDs = append(ety.FIDs, m.FIDs...)
	sortutil.Dedupe(FeatureIDSlice(ety.FIDs))
	return ety
}

type MapElementID2Refs struct {
	IDElem osm.FeatureID

	Refs osm.FeatureIDs
}

func (m MapElementID2Refs) ToIndexKey() string {
	return "Index_MapElementID2Refs_" + m.IDElem.String()
}

func (m MapElementID2Refs) Install(oldVal MapIndex) MapIndex {
	if oldVal != nil {
		m.Refs = append(m.Refs, oldVal.(MapElementID2Refs).Refs...)
	}
	sortutil.Dedupe(FeatureIDSlice(m.Refs))
	return m
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

type ObjectIDSlice []osm.ObjectID

func (p ObjectIDSlice) Len() int { return len(p) }

func (p ObjectIDSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p ObjectIDSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type FeatureIDSlice []osm.FeatureID

func (p FeatureIDSlice) Len() int { return len(p) }

func (p FeatureIDSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p FeatureIDSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
