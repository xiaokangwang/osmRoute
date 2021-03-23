package attributes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckRoutingInputAttribute(t *testing.T) {
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
			if err := CheckRoutingInputAttribute(tt.args.AdditionalInfo); (err != nil) != tt.wantErr {
				t.Errorf("CheckRoutingInputAttribute() error = %v, wantErr %v", err, tt.wantErr)
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
	item, err := NewBusInfo(`{"Info":["node/1234567":{"RemainingTime":"2"}]}`)
	_ = item
	assert.Nil(t, err)
}