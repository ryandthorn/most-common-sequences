# most-common-sequences
Outputs a list of the 100 most common three-word sequences in a given text, along with a count of how many times each occurred in the text

## Usage

Text input can be fed to the program in two ways:

stdin: `cat testfiles/simple.txt | go run .`

filepath args: `go run . testfiles/darwin-full.txt`

You also can combine these methods:

`cat testfiles/simple.txt | go run . testfiles/darwin-full.txt



