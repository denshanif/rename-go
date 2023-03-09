# rename-go

# Rename Go is a tool to rename files in a windows directory to the new names in a csv file

# CSV File Preparation

csv file format: <new-name.csv>
name in csv may contain spaces
csv file must contain only one name per line
name in csv must be sorted before running program

# Directory Preparation

directory format: <directory>\pdfs
directory must contain only .pdf files
file must be sorted before running program

# Example

example csv file:
new name 1
new name 2
new name 3

example directory:
c:\temp\pdfs

example directory contents:
old name 1.pdf
old name 2.pdf
old name 3.pdf

example directory contents after running program:
new name 1.pdf
new name 2.pdf
new name 3.pdf

first name in csv file will be renamed to first file in directory

# How to compile

```go build rename.go``` or ```go build -o rename.exe rename.go```

# How to run

```rename.exe <complete path to csv file> <complete path to the directory containing the pdfs>```

example usage: rename.exe c:\temp\newnames.csv c:\temp\pdfs

# Created by Hanif Al Fathoni