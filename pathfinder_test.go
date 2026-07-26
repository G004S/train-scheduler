package main

import (
	"os/exec"
	"strings"
	"testing"
)

type testCase struct {
	name      string
	args      []string
	wantError bool
	expected  string
	lastTrain string
}

func TestProgram(t *testing.T) {
	tests := []testCase{
		{
			name: "London 2 trains",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"2",
			},
			wantError: false,
			expected: `T1-victoria T2-euston
T1-st_pancras T2-st_pancras`,
		},
		{
			name: "London 3 trains",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"3",
			},
			wantError: false,
			expected: `T1-victoria T2-euston
T1-st_pancras T2-st_pancras T3-victoria
T3-st_pancras`,
		},
		{
			name: "London 4 trains",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"4",
			},
			wantError: false,
			expected: `T1-victoria T2-euston
T1-st_pancras T2-st_pancras T3-victoria T4-euston
T3-st_pancras T4-st_pancras`,
		},
		{
			name: "London 100 trains",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"100",
			},
			wantError: false,
			expected: `T1-victoria T2-euston
T1-st_pancras T2-st_pancras T3-victoria T4-euston
T3-st_pancras T4-st_pancras T5-victoria T6-euston
T5-st_pancras T6-st_pancras T7-victoria T8-euston
T7-st_pancras T8-st_pancras T9-victoria T10-euston`,
			lastTrain: "T100-st_pancras",
		},
		{
			name:      "Too few arguments",
			args:      []string{},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Too many arguments",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"2",
				"extra",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Train count is not a number",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"hello",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Zero trains",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"0",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Negative trains",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"-2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Start station does not exist",
			args: []string{
				"london.map",
				"unknown",
				"st_pancras",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "End station does not exist",
			args: []string{
				"london.map",
				"waterloo",
				"unknown",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Start and end are same",
			args: []string{
				"london.map",
				"waterloo",
				"waterloo",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Duplicate station",
			args: []string{
				"duplicate_stations.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Duplicate coordinates",
			args: []string{
				"duplicate_coordinates.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Duplicate connection",
			args: []string{
				"testmaps/duplicate_connections.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Invalid coordinates",
			args: []string{
				"testmaps/invalid_coordinates.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Missing stations section",
			args: []string{
				"testmaps/missing_stations.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Missing connections section",
			args: []string{
				"testmaps/missing_connections.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Empty file",
			args: []string{
				"testmaps/empty.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Unknown station in connection",
			args: []string{
				"testmaps/unknown_station.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Station connected to itself",
			args: []string{
				"testmaps/self_connection.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Map file does not exist",
			args: []string{
				"testmaps/does_not_exist.map",
				"waterloo",
				"st_pancras",
				"2",
			},
			wantError: true,
			expected:  "Error",
		},
		{
			name: "Empty stations section",
			args: []string{
				"testmaps/empty_stations.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "stations section is empty",
		},
		{
			name: "Empty connections section",
			args: []string{
				"testmaps/empty_connections.map",
				"a",
				"b",
				"2",
			},
			wantError: true,
			expected:  "connections section is empty",
		},
		{
			name: "bond_square",
			args: []string{
				"testmaps/bond_square.map",
				"bond_square",
				"space_port",
				"4",
			},
			wantError: false,
			expected: `T1-apple_avenue
T1-orange_junction T2-apple_avenue
T1-space_port T2-orange_junction T3-apple_avenue
T2-space_port T3-orange_junction T4-apple_avenue
T3-space_port T4-orange_junction
T4-space_port`,
		},
		{
			name: "beethoven",
			args: []string{
				"testmaps/beethoven.map",
				"beethoven",
				"part",
				"9",
			},
			wantError: false,
			expected: `T1-verdi T3-handel
T1-part T2-verdi T3-mozart T5-handel
T2-part T3-part T4-verdi T5-mozart T7-handel
T4-part T5-part T6-verdi T7-mozart T9-handel
T6-part T7-part T8-verdi T9-mozart
T8-part T9-part`,
		},
		{
			name: "more than 10k stations",
			args: []string{
				"10001_stations.map",
				"station_0",
				"station_1",
				"1",
			},
			wantError: true,
			expected:  "Error: too many stations: maximum is 10000, provided 10001",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			commandArgs := append([]string{"run", "."}, test.args...)

			cmd := exec.Command("go", commandArgs...)

			output, err := cmd.CombinedOutput()
			result := string(output)

			if test.wantError && err == nil {
				t.Fatal("expected an error, but program succeeded")
			}

			if !test.wantError && err != nil {
				t.Fatalf(
					"program failed: %v\noutput:\n%s", err, output)
			}

			if test.expected != "" && !strings.Contains(result, test.expected) {
				t.Errorf(
					"expected output to contain:\n%s\n\ngot:\n%s", test.expected, output)
			}

			if test.lastTrain != "" && !strings.Contains(result, test.lastTrain) {
				t.Errorf("expected output to contain:\n%s\n\ngot:\n%s", test.lastTrain, output)
			}
		})
	}
}
