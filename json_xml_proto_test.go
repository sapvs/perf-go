package perf

import (
	"reflect"
	"testing"
)

var userBytes []byte

func init() {
	userBytes, _ = StructToProto(&UserP{ID: 1, Name: "SomeName"})
}

func TestJSONToStruct(t *testing.T) {
	tests := []struct {
		name       string
		userString string
		want       *User
		wantErr    bool
	}{
		{name: "Test1", userString: `{"userid":1,"username":"SomeName"}`, want: &User{ID: 1, Name: "SomeName"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONToStruct(tt.userString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructToJSON(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		want    string
		wantErr bool
	}{
		{name: "Test1", want: `{"userid":1,"username":"SomeName"}`, user: &User{ID: 1, Name: "SomeName"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToJSON(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StructToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMLToStruct(t *testing.T) {
	tests := []struct {
		name       string
		userString string
		want       *User
		wantErr    bool
	}{
		{name: "Test1", userString: "<user><userid>1</userid><username>SomeName</username></user>", want: &User{ID: 1, Name: "SomeName"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := XMLToStruct(tt.userString)
			if (err != nil) != tt.wantErr {
				t.Errorf("XMLToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != tt.want.ID && got.Name != tt.want.Name {
				t.Errorf("XMLToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructToXML(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		want    string
		wantErr bool
	}{
		{name: "Test1", want: "<user><userid>1</userid><username>SomeName</username></user>", user: &User{ID: 1, Name: "SomeName"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToXML(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StructToXML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := StructToXML(&User{ID: 1, Name: "BenchName"})
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkFromXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := XMLToStruct("<user><ID>1</ID><Name>Some Name</Name></user>")
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkFromJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := JSONToStruct(`{"userid":1,"username":"SoemName"}`)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := StructToJSON(&User{ID: 1, Name: "BenchName"})
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkParToJSON(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_, err := StructToJSON(&User{ID: 1, Name: "BenchName"})
			if err != nil {
				b.Fail()
			}
		}
	})
}

func BenchmarkParFromJSON(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_, err := JSONToStruct(`{"userid":1,"username":"SoemName"}`)
			if err != nil {
				b.Fail()
			}
		}
	})
}

func BenchmarkParToXML(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_, err := StructToXML(&User{ID: 1, Name: "BenchName"})
			if err != nil {
				b.Fail()
			}
		}
	})
}

func BenchmarkParFromXML(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_, err := XMLToStruct("<user><ID>1</ID><Name>Some Name</Name></user>")
			if err != nil {
				b.Fail()
			}
		}
	})
}

func TestStructToProtoToStruct(t *testing.T) {
	tests := []struct {
		name    string
		user    UserP
		wantErr bool
	}{
		{name: "1", user: UserP{ID: 1, Name: "SomeName"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// First conversion
			b, err := StructToProto(&tt.user)
			if err != nil {
				t.Errorf("StructToProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// convertback
			up, err := ProtoToStruct(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if up.ID != tt.user.ID || up.Name != tt.user.Name {
				t.Errorf("TestStructToProtoToStruct() = %v, want %v", up, &tt.user)
			}
		})
	}
}

func BenchmarkToProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := StructToProto(&UserP{ID: 1, Name: "BenchName"})
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkFromProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ProtoToStruct(userBytes)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkParToProto(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_, err := StructToProto(&UserP{ID: 1, Name: "BenchName"})
			if err != nil {
				b.Fail()
			}
		}
	})
}

func BenchmarkParFromProto(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_, err := ProtoToStruct(userBytes)
			if err != nil {
				b.Fail()
			}
		}
	})
}
