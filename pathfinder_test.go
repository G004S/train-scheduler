package main

import (
	"bytes"
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
			name: "London 1 trains",
			args: []string{
				"london.map",
				"waterloo",
				"st_pancras",
				"1",
			},
			wantError: false,
			expected: `T1-victoria
T1-st_pancras`,
		},
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
			name: "jungle_desert",
			args: []string{
				"testmaps/jungle_desert.map",
				"jungle",
				"desert",
				"10",
			},
			wantError: false,
			expected: `T1-farms T2-grasslands T3-green_belt
T1-downtown T2-suburbs T3-village T4-farms T5-grasslands T6-green_belt
T1-metropolis T2-clouds T3-mountain T4-downtown T5-suburbs T6-village T7-farms T8-grasslands T9-green_belt
T1-industrial T2-wetlands T3-treetop T4-metropolis T5-clouds T6-mountain T7-downtown T8-suburbs T9-village T10-farms
T1-desert T2-desert T3-desert T4-industrial T5-wetlands T6-treetop T7-metropolis T8-clouds T9-mountain T10-downtown
T4-desert T5-desert T6-desert T7-industrial T8-wetlands T9-treetop T10-metropolis
T7-desert T8-desert T9-desert T10-industrial
T10-desert`,
		},
		{
			name: "reverse duplicate connections",
			args: []string{
				"testmaps/jungle_desert2.map",
				"jungle",
				"desert",
				"10",
			},
			wantError: true,
			expected:  "Error: line 35: dublicate connection between \"mountain\" and \"farms\"",
		},
		{
			name: "invalid station names",
			args: []string{
				"testmaps/jungle_desert3.map",
				"jungle",
				"desert",
				"10",
			},
			wantError: true,
			expected:  "Error: line 2: station name \"Jungle\" must be in lower case",
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
		{
			name: "no path",
			args: []string{
				"testmaps/no_path.map",
				"start",
				"end",
				"1",
			},
			wantError: true,
			expected:  "Error: no path exists between \"start\" and \"end\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			commandArgs := append([]string{"run", "."}, test.args...)

			cmd := exec.Command("go", commandArgs...)

			var stdout bytes.Buffer
			var stderr bytes.Buffer

			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()

			stdoutResult := stdout.String()
			stderrResult := stderr.String()

			if test.wantError {
				if err == nil {
					t.Fatal("expected an error, but program succeeded")
				}

				if stderrResult == "" {
					t.Fatal("expected error message in stderr, but stderr was empty")
				}

				if test.expected != "" &&
					!strings.Contains(stderrResult, test.expected) {

					t.Errorf(
						"expected stderr to contain:\n%s\n\ngot stderr:\n%s",
						test.expected,
						stderrResult,
					)
				}

				if stdoutResult != "" {
					t.Errorf(
						"expected stdout to be empty for error case, got:\n%s",
						stdoutResult,
					)
				}

				return
			}

			if err != nil {
				t.Fatalf(
					"program failed: %v\nstderr:\n%s",
					err,
					stderrResult,
				)
			}

			if stderrResult != "" {
				t.Errorf(
					"expected stderr to be empty, got:\n%s",
					stderrResult,
				)
			}

			if test.expected != "" &&
				!strings.Contains(stdoutResult, test.expected) {

				t.Errorf(
					"expected stdout to contain:\n%s\n\ngot stdout:\n%s",
					test.expected,
					stdoutResult,
				)
			}

			if test.lastTrain != "" &&
				!strings.Contains(stdoutResult, test.lastTrain) {

				t.Errorf(
					"expected stdout to contain:\n%s\n\ngot stdout:\n%s",
					test.lastTrain,
					stdoutResult,
				)
			}
		})
	}
}
