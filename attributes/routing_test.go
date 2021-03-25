package attributes

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRoutingInputAttribute(t *testing.T) {
	type args struct {
		AdditionalInfo map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "empty should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				return m
			}()},
		},
		{
			name:    "objective_time is set should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_time"] = "0.5"
				return m
			}()},
		},
		{
			name:    "objective_cost is set should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_cost"] = "0.5"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "0.5"
				return m
			}()},
		},
		{
			name:    "objective_sustainable + objective_time is set should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_cost"] = "0.5"
				m["objective_sustainable"] = "0.5"
				return m
			}()},
		},
		{
			name:    "objective_sustainable + objective_cost is set should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_cost"] = "0.5"
				m["objective_sustainable"] = "0.5"
				return m
			}()},
		},
		{
			name:    "objective_sustainable + objective_cost is set should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_cost"] = "0.5"
				m["objective_sustainable"] = "0.5"
				return m
			}()},
		},
		{
			name:    "objective_sustainable + objective_cost + objective_time is set should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_cost"] = "0.5"
				m["objective_sustainable"] = "0.5"
				m["objective_time"] = "0.5"
				return m
			}()},
		},
		{
			name:    "empty should be accepted(again)",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				return m
			}()},
		},
		{
			name:    "objective_time is set to 2 should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_time"] = "2"
				return m
			}()},
		},
		{
			name:    "objective_cost is set to 2 should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_cost"] = "2"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set to 2 should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "2"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set to T_T should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "T_T"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set to NaN should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "NaN"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set to null should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "null"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set to {0.2} should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "{0.2}"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set to +0.2 should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "+0.2"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set to +++++++++++++++0.2 should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "+++++++++++++++0.2"
				return m
			}()},
		},
		{
			name:    "objective_sustainable is set to +-+-+-+-0.2 should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["objective_sustainable"] = "+-+-+-+-0.2"
				return m
			}()},
		},
		{
			name:    "unrecognized key should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["0f2b51f9-bb45-4488-af51-15f34c145652"] = "b68f7f24-b853-4a7f-8d6a-772a2160f1dc"
				return m
			}()},
		},
		{
			name:    "can_walkLong set to random should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_walkLong"] = "b68f7f24-b853-4a7f-8d6a-772a2160f1dc"
				return m
			}()},
		},
		{
			name:    "can_walkLong set to TrUe should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_walkLong"] = "TrUe"
				return m
			}()},
		},
		{
			name:    "can_walkLong set to FaLse should be rejected",
			wantErr: true,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_walkLong"] = "FaLse"
				return m
			}()},
		},
		{
			name:    "can_walkLong set to true should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_walkLong"] = "true"
				return m
			}()},
		},
		{
			name:    "can_bike set to false should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_bike"] = "false"
				return m
			}()},
		},
		{
			name:    "can_bike set to true should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_bike"] = "true"
				return m
			}()},
		},
		{
			name:    "can_walkLong set to false should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_walkLong"] = "false"
				return m
			}()},
		},
		{
			name:    "can_drive set to true should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_drive"] = "true"
				return m
			}()},
		},
		{
			name:    "can_drive set to false should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_drive"] = "false"
				return m
			}()},
		},
		{
			name:    "can_drive set to false + objective_time = 0.5 should be accepted",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["can_drive"] = "false"
				m["objective_time"] = "0.5"
				return m
			}()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseRoutingInputAttribute(tt.args.AdditionalInfo); (err != nil) != tt.wantErr {
				t.Errorf("ParseRoutingInputAttribute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAreaToAvoid(t *testing.T) {
	item, err := NewAreaToAvoid(`{"Items":[{"T":1, "B":0, "L":0, "R":1}]}`)
	_ = item
	assert.Nil(t, err)

	t.Run("Check 0.5, 0.5", func(t *testing.T) {
		ret := item.CheckPointInclusion(0.5, 0.5)
		assert.True(t, ret)
	})

	t.Run("Check 1.5, 0.5", func(t *testing.T) {
		ret := item.CheckPointInclusion(1.5, 0.5)
		assert.False(t, ret)
	})
	t.Run("Check -1.5, 0.5", func(t *testing.T) {
		ret := item.CheckPointInclusion(-1.5, 0.5)
		assert.False(t, ret)
	})
	t.Run("Check -1.5, -0.5", func(t *testing.T) {
		ret := item.CheckPointInclusion(-1.5, 0.5)
		assert.False(t, ret)
	})
	t.Run("Check -1.5, -1.5", func(t *testing.T) {
		ret := item.CheckPointInclusion(-1.5, 0.5)
		assert.False(t, ret)
	})
}

func TestBusInfo(t *testing.T) {
	item, err := NewBusInfo(`{"Info":{"node/1234567":{"RemainingTime":2}}}`)
	_ = item
	assert.Nil(t, err)
}

func TestParseRoutingInputAttribute2(t *testing.T) {
	type args struct {
		AdditionalInfo map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "area to avoid",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["area_to_avoid"] = `{"Items":[{"T":1, "B":0, "L":0, "R":1}]}`
				return m
			}()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseRoutingInputAttribute(tt.args.AdditionalInfo); (err != nil) != tt.wantErr {
				t.Errorf("ParseRoutingInputAttribute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBikeStation_ListAllRoutes(t *testing.T) {
	d, err := os.ReadFile("../testing/bike.json")
	assert.Nil(t, err)
	s, err := NewBikeStationData(string(d))
	assert.Nil(t, err)
	w := s.ListAllStations()
	assert.Len(t, w, 105)
}

func TestParseRoutingInputAttribute3(t *testing.T) {
	type args struct {
		AdditionalInfo map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "bike information",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				d, _ := os.ReadFile("../testing/bike.json")
				m["bike"] = string(d)
				return m
			}()},
		}, {
			name:    "bus information",
			wantErr: false,
			args: struct{ AdditionalInfo map[string]string }{AdditionalInfo: func() map[string]string {
				m := make(map[string]string)
				m["bus"] = `{"Info":{"node/1234567":{"RemainingTime":2}}}`
				return m
			}()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseRoutingInputAttribute(tt.args.AdditionalInfo); (err != nil) != tt.wantErr {
				t.Errorf("ParseRoutingInputAttribute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
