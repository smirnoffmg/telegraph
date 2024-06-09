package telegraph_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/smirnoffmg/telegraph"
)

func TestNodeElement_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		nodeElem *telegraph.NodeElement
		wantJSON string
		wantErr  bool
	}{
		{
			name: "Marshal NodeElement with all fields",
			nodeElem: &telegraph.NodeElement{
				Tag:   "div",
				Attrs: map[string]string{"class": "container"},
				Children: []telegraph.Node{
					"Hello, World!",
				},
			},
			wantJSON: `{"tag":"div","attrs":{"class":"container"},"children":["Hello, World!"]}`,
			wantErr:  false,
		},
		// Add more test cases for different scenarios
		{
			name: "Marshal NodeElement with only tag",
			nodeElem: &telegraph.NodeElement{
				Tag: "div",
			},
			wantJSON: `{"tag":"div"}`,
			wantErr:  false,
		},
		// Add more test cases for different scenarios
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotJSON, err := json.Marshal(tt.nodeElem)
			if (err != nil) != tt.wantErr {
				t.Errorf("telegraph.NodeElement.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(gotJSON), tt.wantJSON) {
				t.Errorf("telegraph.NodeElement.MarshalJSON() = %v, want %v", string(gotJSON), tt.wantJSON)
			}
		})
	}
}

func TestNodeElement_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantNode *telegraph.NodeElement
		wantErr  bool
	}{
		{
			name:     "Unmarshal valid JSON with all fields",
			jsonData: `{"tag":"div","attrs":{"class":"container"},"children":["Hello, World!"]}`,
			wantNode: &telegraph.NodeElement{
				Tag:   "div",
				Attrs: map[string]string{"class": "container"},
				Children: []telegraph.Node{
					"Hello, World!",
				},
			},
			wantErr: false,
		},
		// Add more test cases for different scenarios
		{
			name:     "Unmarshal valid JSON with only tag",
			jsonData: `{"tag":"div"}`,
			wantNode: &telegraph.NodeElement{
				Tag: "div",
			},
			wantErr: false,
		},
		// Add more test cases for different scenarios
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var node telegraph.NodeElement
			err := json.Unmarshal([]byte(tt.jsonData), &node)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(&node, tt.wantNode) {
				t.Errorf("json.Unmarshal() = %+v, want %+v", &node, tt.wantNode)
			}
		})
	}
}
