package catalogue

import (
	"encoding/json"
	"testing"
	"time"
)

func TestFlexibleFloat_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{
			name:    "unmarshal from float",
			input:   `7.5`,
			want:    7.5,
			wantErr: false,
		},
		{
			name:    "unmarshal from string",
			input:   `"7.5"`,
			want:    7.5,
			wantErr: false,
		},
		{
			name:    "unmarshal from integer",
			input:   `10`,
			want:    10.0,
			wantErr: false,
		},
		{
			name:    "unmarshal from string integer",
			input:   `"10"`,
			want:    10.0,
			wantErr: false,
		},
		{
			name:    "unmarshal zero",
			input:   `0`,
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "unmarshal negative",
			input:   `"-3.14"`,
			want:    -3.14,
			wantErr: false,
		},
		{
			name:    "invalid string",
			input:   `"not a number"`,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got FlexibleFloat
			err := json.Unmarshal([]byte(tt.input), &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("FlexibleFloat.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got.Float64() != tt.want {
				t.Errorf("FlexibleFloat.UnmarshalJSON() = %v, want %v", got.Float64(), tt.want)
			}
		})
	}
}

func TestFlexibleFloat_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		value   FlexibleFloat
		want    string
		wantErr bool
	}{
		{
			name:    "marshal float",
			value:   FlexibleFloat(7.5),
			want:    `7.5`,
			wantErr: false,
		},
		{
			name:    "marshal zero",
			value:   FlexibleFloat(0),
			want:    `0`,
			wantErr: false,
		},
		{
			name:    "marshal negative",
			value:   FlexibleFloat(-3.14),
			want:    `-3.14`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("FlexibleFloat.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && string(got) != tt.want {
				t.Errorf("FlexibleFloat.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestFlexibleDate_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "unmarshal YYYY-MM-DD",
			input:   `"1983-08-06"`,
			want:    time.Date(1983, 8, 6, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "unmarshal RFC3339",
			input:   `"2023-01-15T10:30:00Z"`,
			want:    time.Date(2023, 1, 15, 10, 30, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "unmarshal RFC3339 with timezone",
			input:   `"2023-01-15T10:30:00-05:00"`,
			want:    time.Date(2023, 1, 15, 15, 30, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "unmarshal datetime without timezone",
			input:   `"2023-01-15T10:30:00"`,
			want:    time.Date(2023, 1, 15, 10, 30, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "unmarshal datetime with space",
			input:   `"2023-01-15 10:30:00"`,
			want:    time.Date(2023, 1, 15, 10, 30, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "unmarshal empty string",
			input:   `""`,
			want:    time.Time{},
			wantErr: false,
		},
		{
			name:    "unmarshal null",
			input:   `null`,
			want:    time.Time{},
			wantErr: false,
		},
		{
			name:    "invalid date format",
			input:   `"not a date"`,
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got FlexibleDate
			err := json.Unmarshal([]byte(tt.input), &got)

			if (err != nil) != tt.wantErr {
				t.Errorf("FlexibleDate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !got.Time.Equal(tt.want) {
				t.Errorf("FlexibleDate.UnmarshalJSON() = %v, want %v", got.Time, tt.want)
			}
		})
	}
}

func TestFlexibleDate_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		value   FlexibleDate
		want    string
		wantErr bool
	}{
		{
			name:    "marshal date",
			value:   FlexibleDate{Time: time.Date(1983, 8, 6, 0, 0, 0, 0, time.UTC)},
			want:    `"1983-08-06T00:00:00Z"`,
			wantErr: false,
		},
		{
			name:    "marshal datetime",
			value:   FlexibleDate{Time: time.Date(2023, 1, 15, 10, 30, 0, 0, time.UTC)},
			want:    `"2023-01-15T10:30:00Z"`,
			wantErr: false,
		},
		{
			name:    "marshal zero time",
			value:   FlexibleDate{Time: time.Time{}},
			want:    `null`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("FlexibleDate.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && string(got) != tt.want {
				t.Errorf("FlexibleDate.MarshalJSON() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestWork_UnmarshalWithFlexibleTypes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		check   func(*testing.T, *Work)
	}{
		{
			name: "work with string userRating",
			input: `{
				"id": "work123",
				"workType": "movie",
				"slug": "test-movie",
				"title": "Test Movie",
				"originalTitle": "Test Movie",
				"userRating": "7.5",
				"spokenLanguages": [],
				"languages": [],
				"isStreamable": false,
				"isInTheatre": false,
				"featured": false,
				"genres": [],
				"createdAt": "2023-01-01T00:00:00Z",
				"updatedAt": "2023-01-01T00:00:00Z"
			}`,
			wantErr: false,
			check: func(t *testing.T, w *Work) {
				if w.UserRating == nil {
					t.Error("UserRating should not be nil")
					return
				}
				if w.UserRating.Float64() != 7.5 {
					t.Errorf("UserRating = %v, want 7.5", w.UserRating.Float64())
				}
			},
		},
		{
			name: "work with float userRating",
			input: `{
				"id": "work123",
				"workType": "movie",
				"slug": "test-movie",
				"title": "Test Movie",
				"originalTitle": "Test Movie",
				"userRating": 8.2,
				"spokenLanguages": [],
				"languages": [],
				"isStreamable": false,
				"isInTheatre": false,
				"featured": false,
				"genres": [],
				"createdAt": "2023-01-01T00:00:00Z",
				"updatedAt": "2023-01-01T00:00:00Z"
			}`,
			wantErr: false,
			check: func(t *testing.T, w *Work) {
				if w.UserRating == nil {
					t.Error("UserRating should not be nil")
					return
				}
				if w.UserRating.Float64() != 8.2 {
					t.Errorf("UserRating = %v, want 8.2", w.UserRating.Float64())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var work Work
			err := json.Unmarshal([]byte(tt.input), &work)

			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.check != nil {
				tt.check(t, &work)
			}
		})
	}
}

func TestPerson_UnmarshalWithFlexibleTypes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		check   func(*testing.T, *Person)
	}{
		{
			name: "person with YYYY-MM-DD birthDate",
			input: `{
				"id": "person123",
				"name": "Test Person",
				"slug": "test-person",
				"birthDate": "1983-08-06",
				"deceased": false,
				"aliases": [],
				"nationality": [],
				"externalLinks": [],
				"featured": false,
				"createdAt": "2023-01-01T00:00:00Z",
				"updatedAt": "2023-01-01T00:00:00Z"
			}`,
			wantErr: false,
			check: func(t *testing.T, p *Person) {
				if p.BirthDate == nil {
					t.Error("BirthDate should not be nil")
					return
				}
				expected := time.Date(1983, 8, 6, 0, 0, 0, 0, time.UTC)
				if !p.BirthDate.Time.Equal(expected) {
					t.Errorf("BirthDate = %v, want %v", p.BirthDate.Time, expected)
				}
			},
		},
		{
			name: "person with RFC3339 birthDate",
			input: `{
				"id": "person123",
				"name": "Test Person",
				"slug": "test-person",
				"birthDate": "1983-08-06T00:00:00Z",
				"deceased": false,
				"aliases": [],
				"nationality": [],
				"externalLinks": [],
				"featured": false,
				"createdAt": "2023-01-01T00:00:00Z",
				"updatedAt": "2023-01-01T00:00:00Z"
			}`,
			wantErr: false,
			check: func(t *testing.T, p *Person) {
				if p.BirthDate == nil {
					t.Error("BirthDate should not be nil")
					return
				}
				expected := time.Date(1983, 8, 6, 0, 0, 0, 0, time.UTC)
				if !p.BirthDate.Time.Equal(expected) {
					t.Errorf("BirthDate = %v, want %v", p.BirthDate.Time, expected)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var person Person
			err := json.Unmarshal([]byte(tt.input), &person)

			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.check != nil {
				tt.check(t, &person)
			}
		})
	}
}
