# nba-lottery

Simple CLI tool built to simulate the NBA draft lottery and view current standings.

## Installation

To download the binary, go to the [Releases](https://github.com/Mahcks/nba-lottery/releases/latest) tab on this GitHub page. The examples in the documentation use the binary in its own folder since it generates its own JSON file.

Note: If you are using MacOS and downloading the binary manually, you may need to adjust its permissions to allow for execution. To do this, run the following command:

```bash
chmod 755 <filename>
```

*Replace `<filename>` with the actual name of the downloaded binary.*

## Commands

### `simulate`

Simulate the NBA draft lottery.

#### Aliases

- `sim`

#### Options

* `-t, --times INT`: The number of times to run the simulation.
* `--help`: Show the help message and exit.

#### Example

```bash
./nba-lottery.exe sim --times 100
```

### `standings`

View the locally stored NBA standings, do `-r` to reload the locally stored data.

#### Aliases

*N/A*

#### Options

- `-r, --reload BOOL`: Reload the standings to get more up-to-date information.
- `--help`: Show the help message and exit.

#### Example

```bash
./nba-lottery.exe standings --reload
```