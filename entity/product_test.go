package entity

import (
	"testing"
)

func TestSizes_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		s    Sizes
		want string
	}{
		{
			name: "ok",
			s:    Sizes{{SizeID: 1, SizeName: "Small"}, {SizeID: 2, SizeName: "Medium"}},
			want: `---	-------	
Id 	Ukuran 	
---	-------	
1	Small	
2	Medium	
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ConvertToTable(); got != tt.want {
				t.Errorf("Sizes.ConvertToTable() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestBrands_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		b    Brands
		want string
	}{
		{
			name: "ok",
			b:    Brands{{BrandID: 1, BrandName: "Zara"}, {BrandID: 2, BrandName: "H&M"}},
			want: `---	-------	
Id 	Ukuran 	
---	-------	
1	Zara	
2	H&M		
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ConvertToTable(); got != tt.want {
				t.Errorf("Brands.ConvertToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
