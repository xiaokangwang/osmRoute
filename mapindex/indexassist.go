package mapindex

import (
	"fmt"
	"github.com/paulmach/osm"
	"modernc.org/sortutil"
	"strings"
)

type MapID2ItemEntry struct {
	ID         osm.ObjectID
	SkipLength int64
}

func (m MapID2ItemEntry) ToIndexKey() string {
	return "I_A_" + m.ID.String()
}

func (m MapID2ItemEntry) Install(oldVal MapIndex) MapIndex {
	return m
}

type MapFeatureID2CurrentObjectIDEntry struct {
	ID         osm.ObjectID
	SkipLength int64

	CachedData []byte

	IDFeat osm.FeatureID
}

func (m MapFeatureID2CurrentObjectIDEntry) ToIndexKey() string {
	return "I_B_" + m.IDFeat.String()
}

func (m MapFeatureID2CurrentObjectIDEntry) Install(oldVal MapIndex) MapIndex {
	if oldVal == nil {
		return m
	}
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

	Significant osm.FeatureIDs
}

func (m MapRegion2IDEntry) ToIndexKey() string {
	return "I_C_" + InterlacedString(FloatToConstantLengthFloat(m.Lat), FloatToConstantLengthFloat(m.Lon))
}

func (m MapRegion2IDEntry) Install(oldVal MapIndex) MapIndex {
	ety := m
	if oldVal != nil {
		ety = oldVal.(MapRegion2IDEntry)
	}
	ety.IDs = append(ety.IDs, m.IDs...)
	n := sortutil.Dedupe(ObjectIDSlice(ety.IDs))
	ety.IDs = ety.IDs[:n]

	ety.FIDs = append(ety.FIDs, m.FIDs...)
	n2 := sortutil.Dedupe(FeatureIDSlice(ety.FIDs))
	ety.FIDs = ety.FIDs[:n2]

	ety.Significant = append(ety.Significant, m.Significant...)
	n3 := sortutil.Dedupe(FeatureIDSlice(ety.Significant))
	ety.Significant = ety.Significant[:n3]
	return ety
}

type MapFeatureID2Refs struct {
	IDElem osm.FeatureID

	Refs osm.FeatureIDs
}

func (m MapFeatureID2Refs) ToIndexKey() string {
	return "I_D_" + m.IDElem.String()
}

func (m MapFeatureID2Refs) Install(oldVal MapIndex) MapIndex {
	if oldVal != nil {
		m.Refs = append(m.Refs, oldVal.(MapFeatureID2Refs).Refs...)
	}
	sortutil.Dedupe(FeatureIDSlice(m.Refs))
	return m
}

type MapFTS struct {
	Name string

	Refs  osm.FeatureIDs
	Count uint32
}

func (m MapFTS) ToIndexKey() string {
	return "I_E_" + m.Name
}

func (m MapFTS) Install(oldVal MapIndex) MapIndex {
	if oldVal == nil {
		return m
	}
	if len(m.Refs) <= 20 {
		m.Refs = append(m.Refs, oldVal.(MapFTS).Refs...)
	}
	m.Count += oldVal.(MapFTS).Count
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
func DeInterlacedString(in string) (string, string) {
	var sba strings.Builder
	var sbb strings.Builder

	lenstr := len(in)
	for i := 0; i < lenstr; i += 2 {
		sba.WriteByte(in[i])
		sbb.WriteByte(in[i+1])
	}

	return sba.String(), sbb.String()
}
func FloatToConstantLengthFloat(input float64) string {
	strv := fmt.Sprintf("%0+16.6f", input)

	return strv
}

type ObjectIDSlice []osm.ObjectID

func (p ObjectIDSlice) Len() int { return len(p) }

func (p ObjectIDSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p ObjectIDSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type FeatureIDSlice []osm.FeatureID

func (p FeatureIDSlice) Len() int { return len(p) }

func (p FeatureIDSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p FeatureIDSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
