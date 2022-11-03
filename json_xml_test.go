package perf

import (
	"reflect"
	"testing"
)

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
		{name: "Test1", userString: "<User><ID>1</ID><Name>SomeName</Name></User>", want: &User{ID: 1, Name: "SomeName"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := XMLToStruct(tt.userString)
			if (err != nil) != tt.wantErr {
				t.Errorf("XMLToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		{name: "Test1", want: "<User><ID>1</ID><Name>SomeName</Name></User>", user: &User{ID: 1, Name: "SomeName"}, wantErr: false},
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

func BenchmarkToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := StructToJSON(&User{ID: 1, Name: "BenchName"})
		if err != nil {
			b.Fail()
		}
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

func BenchmarkFromJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := JSONToStruct(`{"userid":1,"username":"SoemName"}`)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkFromXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := XMLToStruct("<User><ID>1</ID><Name>Some Name</Name></User>")
		if err != nil {
			b.Fail()
		}
	}
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

func BenchmarkParFromXML(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_, err := XMLToStruct("<User><ID>1</ID><Name>Some Name</Name></User>")
			if err != nil {
				b.Fail()
			}
		}
	})
}
