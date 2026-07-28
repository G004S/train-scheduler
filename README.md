# Pathfinder

A Go program that finds efficient routes for multiple trains through a railway network.

The program reads a railway map from a `.map` file, validates the input, finds possible routes between two stations, selects the best combination of non-overlapping routes, simulates train movement, and prints every train movement turn by turn.

---

## Features

- Reads railway maps from `.map` files
- Validates map structure and input arguments
- Ignores comments beginning with `#`
- Finds **possible simple routes** using DFS
- Groups routes that do not share intermediate stations
- Chooses the fastest route combination through simulation
- Simulates multiple trains moving simultaneously
- Prevents collisions between trains
- Prints train movements turn by turn
- All errors are printed to **stderr**

---

## Project Structure
```text
july_pathfinder/
├── internal/          # Project packages
├── testmaps/          # Maps used for automated tests
├── london.map         # Sample network
├── paths.map          # Additional example map
├── 10k.map            # Map with 10,000 stations
├── 10001_stations.map # Invalid map (>10,000 stations)
├── main.go
├── pathfinder_test.go
├── go.mod
└── README.md
```

## Usage

```bash
go run . <map_file> <start_station> <end_station> <number_of_trains>
```

Example:

```bash
go run . network.map waterloo st_pancras 4
```

---

## Arguments

| Argument | Description |
|----------|-------------|
| `map_file` | Path to the railway map |
| `start_station` | Starting station |
| `end_station` | Destination station |
| `number_of_trains` | Positive integer |

---

## Map format

A map contains two sections:

```text
stations:
```

and

```text
connections:
```

Example:

```text
stations:
waterloo,0,0
victoria,1,0
bank,2,0
st_pancras,3,0

connections:
waterloo-victoria
victoria-bank
bank-st_pancras
```

Comments beginning with `#` are ignored.

Example:

```text
# London network

stations:
waterloo,0,0

connections:
# main route
waterloo-bank
```

---

## Output

The program prints one line per turn.

Example:

```text
T1-victoria T2-bank
T1-bank T2-st_pancras
T1-st_pancras
```

Each movement has the format

```text
T<train_id>-<station_name>
```

---

# Route selection

The algorithm works in several stages.

## 1. Parse the map

The map file is read and validated.

---

## 2. Route Search

The program uses a combination of **Breadth-First Search (BFS)** and **Depth-First Search (DFS)** to find candidate routes between the start and end stations.

### BFS

BFS is first used to verify that a path exists and to determine the length of the shortest possible route. This value is later used to limit the DFS search space.

### DFS

DFS enumerates all valid routes that satisfy the following constraints:

- The route must not contain repeated stations.
- The route length may not exceed **(shortest path length + 5 stations)**.
- At most **100 routes** are collected.
- The search is limited to **15 seconds**. If this limit is exceeded, the program terminates with an error.

These limits were introduced to prevent excessive execution times on large or highly connected maps. Without them, the number of possible routes can grow exponentially, causing the search to become impractically slow or consume excessive memory.

Because of these optimizations, the program does **not** guarantee that every possible route is explored. Instead, it searches within practical limits and uses the collected routes to determine an efficient train schedule.

---

## 3. Build route groups

Routes are grouped so that no two routes in the same group share intermediate stations.

Start and destination stations may be shared.

---

## 4. Simulate train movement

Each route group is simulated independently.

The group requiring the fewest turns is selected.

---

## 5. Print movements

The movements of every train are printed one turn at a time.

---

# Train movement rules

- Only one train may occupy an intermediate station at a time.
- Multiple trains may occupy the destination station.
- Trains move at most one station per turn.
- Trains cannot overtake each other.
- The simulation finishes when every train reaches the destination.

---

# Validation

The program validates both command-line arguments and the map file.

## Command-line validation

- missing arguments
- too many arguments
- start station equals destination station
- number of trains is not an integer
- number of trains is less than one
- map file cannot be found
- map file cannot be accessed

---

## Map validation

The parser checks for:

- empty file
- missing `stations:` section
- missing `connections:` section
- empty `stations:` section
- empty `connections:` section
- malformed station definition
- malformed connection definition
- duplicate station names
- duplicate coordinates
- coordinates that are not positive integers
- unknown station referenced in a connection
- self-connections
- duplicate connections
- more than 10,000 stations
- more than 3000 trains

---

### Train limit

Our implementation accepts up to 3000 trains.

This limit was introduced as a safeguard to avoid excessively large simulations and impractically large output. It is an implementation decision rather than a requirement of the original project specification.

The value was chosen as a reasonable upper bound for the current implementation. Exactly **3000** trains are accepted, while values above **3000** are rejected.

---

## Route validation

The program also reports errors when

- the start station does not exist
- the destination station does not exist
- no route exists between the stations

---

# Data structures

## Station

Represents one railway station.

```go
type Station struct {
    Name       string
    X          int
    Y          int
    Neighbours []*Station
}
```

---

## Route

Represents one complete route.

```go
type Route struct {
	ID                int
	Stations          []*Station
	ConflictingRoutes map[int]bool
}
```

---

## Train

Represents one train during simulation.

```go
type Train struct {
    ID       int
    Position int
    Route    []*Station
}
```

---

# Algorithm complexity

### DFS

Worst case:

```
O(number of simple paths)
```

Finding every simple path is exponential in the worst case.

---

### Route grouping

Approximately

```
O(R² × L)
```

where

- **R** is the number of routes
- **L** is the average route length

---

### Simulation

Approximately

```
O(T × M)
```

where

- **T** is the number of turns
- **M** is the number of trains

---

# Error handling

All errors are written to **stderr** and the program exits with a non-zero status code.

Example:

```text
Error: station "waterloo" does not exist
```

or

```text
Error: duplicate connection between "bank" and "victoria"
```

---

## Testing

Run all tests:
```bash
go test -v
```

The test suite covers:

- valid routes for the London network, bond_square and beethoven maps;
- invalid train counts (zero and negative);
- invalid start and end stations;
- identical start and end stations;
- duplicate stations;
- duplicate coordinates;
- duplicate connections;
- invalid coordinates;
- missing stations: section;
- missing connections: section;
- empty map file;
- connections to unknown stations;
- self-connections;
- non-existent map file.

Additional test maps are located in the testmaps/ directory.
