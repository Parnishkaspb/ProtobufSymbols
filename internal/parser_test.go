package internal

import (
	"testing"
)

func TestParseProto(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Pattern
	}{
		{
			name: "simple import",
			input: `syntax = "proto3";
			import "google/protobuf/timestamp.proto";`,
			expected: []Pattern{
				{
					Name:     "google/protobuf/timestamp.proto",
					Type:     "import",
					Line:     2,
					StartPos: 8,
					EndPos:   41,
				},
			},
		},
		{
			name: "service with rpc",
			input: `service Greeter {
			  rpc SayHello(HelloRequest) returns (HelloReply);
			}`,
			expected: []Pattern{
				{
					Name:     "Greeter",
					Type:     "service",
					Line:     1,
					StartPos: 9,
					EndPos:   16,
				},
				{
					Name:     "SayHello",
					Type:     "method",
					Line:     2,
					StartPos: 7,
					EndPos:   15,
				},
			},
		},
		{
			name: "nested message should be ignored",
			input: `message Outer {
			  message Inner {
				string x = 1;
			  }
			}`,
			expected: []Pattern{
				{
					Name:     "Outer",
					Type:     "message",
					Line:     1,
					StartPos: 9,
					EndPos:   14,
				},
			},
		},
		{
			name: "train test",
			input: `syntax = "proto3";

				package example;
				
				import "google/protobuf/timestamp.proto";
				
				option go_package = "gitlab.ozon.ru/example/api/example;example";
				
				service Example {
				  rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
				}
				
				enum ExampleEnum {
				  ONE = 0;
				  TWO = 1;
				  THREE = 2;
				}
				
				message ExampleRPCRequest {
				  message Emb { string field11 = 1; }
				  ExampleEnum field1 = 1;
				  Emb filed2 = 2;
				  google.protobuf.Timestamp filed3 = 3;
				}
				
				message ExampleRPCResponse {}`,
			expected: []Pattern{
				{
					Name:     "google/protobuf/timestamp.proto",
					Type:     "import",
					Line:     5,
					StartPos: 8,
					EndPos:   41,
				},
				{
					Name:     "Example",
					Type:     "service",
					Line:     9,
					StartPos: 9,
					EndPos:   16,
				},
				{
					Name:     "ExampleRPC",
					Type:     "method",
					Line:     10,
					StartPos: 7,
					EndPos:   16,
				},
				{
					Name:     "ExampleEnum",
					Type:     "enum",
					Line:     13,
					StartPos: 6,
					EndPos:   17,
				},
				{
					Name:     "ExampleRPCRequest",
					Type:     "message",
					Line:     19,
					StartPos: 9,
					EndPos:   26,
				},
				{
					Name:     "ExampleRPCResponse",
					Type:     "message",
					Line:     26,
					StartPos: 9,
					EndPos:   27,
				},
			},
		},
	}

	parser := &Parser{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.ParseProto(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("ожидалось %d символов, получено %d", len(tt.expected), len(result))
			}

			for i, pat := range tt.expected {
				if i >= len(result) {
					break
				}
				if result[i].Name != pat.Name || result[i].Type != pat.Type || result[i].Line != pat.Line {
					t.Errorf("ожидалось %+v, получено %+v", pat, result[i])
				}
			}
		})
	}
}
