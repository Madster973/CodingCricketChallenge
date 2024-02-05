# Text Analyzer

This Go program analyzes text input, either from a file or standard input (stdin), to count bytes, lines, or words. It's a handy tool for basic text analysis tasks, similar in spirit to Unix utilities like `wc`.

## Features

- **Byte Count**: Counts the number of bytes in the input.
- **Line Count**: Counts the number of lines in the input.
- **Word Count**: Counts the number of words in the input.
- **Combined Metrics**: Provides all the above counts in one go.

## Requirements

To run this program, you need to have Go installed on your system. You can download it from [the official Go website](https://golang.org/dl/).

## Building the Program

1. Clone the repository to your local machine (assuming you have git installed):
   ```sh
   git clone https://yourrepositoryurl.com/path/to/repo.git
   ```
   
2. Navigate into the project directory:
   ```sh
   cd path/to/your/cloned/repo
   ```

3. Build the program using the Go compiler:
   ```sh
   go build -o textanalyzer
   ```

## Usage

After building the program, you can run it directly from the command line.

### Counting Bytes
To count the number of bytes:
```sh
./textanalyzer -c [filename]
```
Or, for stdin:
```sh
cat [filename] | ./textanalyzer -c
```

### Counting Lines
To count the number of lines:
```sh
./textanalyzer -l [filename]
```
Or, for stdin:
```sh
cat [filename] | ./textanalyzer -l
```

### Counting Words
To count the number of words:
```sh
./textanalyzer -w [filename]
```
Or, for stdin:
```sh
cat [filename] | ./textanalyzer -w
```

### Combined Metrics
To get bytes, lines, and words count all together:
```sh
./textanalyzer -m [filename]
```
Or, for stdin:
```sh
cat [filename] | ./textanalyzer -m
```

## Contributing
This is a solution to one of the Coding Challanges by John Cricket.
Please refer to this link for [more challenges](https://codingchallenges.fyi/challenges/challenge-wc/). 
Contributions are welcome! Please feel free to submit a pull request.

## License

This project is open-sourced under the MIT License. See the LICENSE file for more details.