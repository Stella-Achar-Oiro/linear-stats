# Linear-Stats

## Project Overview

`linear-stats` is a command-line application written in Go that reads numerical data from a file, calculates the Linear Regression Line and Pearson Correlation Coefficient, and outputs the results. The project is designed to help users learn about statistical and probability calculations.

## Features

- Reads data from a specified file
- Calculates the Linear Regression Line (`y = mx + b`)
- Calculates the Pearson Correlation Coefficient
- Outputs results with specified precision

## Requirements

- Go 1.16 or higher

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Stella-Achar-Oiro/linear-stats.git
   cd linear-stats
   ```

## Usage

To run the program, use the following command:

```bash
./bin/linear-stats 
```
while comparing with your program 
```bash
go run . data.txt
```

### Data File Format

The data file should contain comma-separated numerical values, one line per entry. Example:

```
189, 113, 121, 114, 145, 110
132, 119, 137, 115, 142, 120
```

### Output Format

The program outputs the Linear Regression Line and Pearson Correlation Coefficient in the following format:

```
Linear Regression Line: y = <slope>x + <intercept>
Pearson Correlation Coefficient: <value>
```

### Precision

- Linear Regression Line slope and intercept are printed with 6 decimal places.
- Pearson Correlation Coefficient is printed with 10 decimal places.

## Error Handling

- The program will print an error message and exit if the data file cannot be opened or read.
- If the data file contains invalid numbers or numbers out of the range, an error message will be printed and the program will exit.

### Example Error Message

```
2024/08/06 12:30:32 Usage: go run . <data file>
exit status 1
```

## Testing

The program can be tested by running it with various data files and comparing the output to expected results.

1. **Generate a random data set and run the program**:

   ```bash
   go run . data.txt
   ```

2. **Verify the output format and precision**:

   - Check that the Linear Regression Line is in the format `y = <value>x + <value>`.
   - Ensure the values on the Linear Regression Line contain 6 decimal places.
   - Check that the Pearson Correlation Coefficient contains 10 decimal places.

3. **Compare results**:

   Run the program multiple times with different data sets to ensure consistent and accurate results.

## Example Data

Below is an example of data to use for testing:

```
1267
1329
1354
1303
1299
1358
```

## Contributing

If you wish to contribute to the project, please fork the repository and submit a pull request with your changes. For major changes, please open an issue first to discuss what you would like to change.

## License

The project is licensed under the MIT License.
