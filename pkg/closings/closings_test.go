package closings

import "testing"

func TestGet(t *testing.T) {
	tests := []struct {
		name     string
		iterator int
		want     string
	}{
		{
			name:     "1",
			iterator: 1,
			want:     "May the God of hope fill us with all joy and peace in believing through the power of the Holy Spirit.",
		},
		{
			name:     "4",
			iterator: 4,
			want:     "May the God of hope fill us with all joy and peace in believing through the power of the Holy Spirit.",
		},
		{
			name:     "3",
			iterator: 3,
			want:     "The grace of our Lord Jesus Christ, and the love of God, and the fellowship of the Holy Spirit, be with us all evermore.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.iterator); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
