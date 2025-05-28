# fmse_exam
repo for FMSE ITMO course exam tasks

There are two tasks: 
- library for DFA and NFA modeling
- simple Simplex algorithm realisation
  
Both are written on pure go, to run tests for DFA\NFA use:
```bash
go test -v ./...
```

For Simplex, both example and implementation are in the same file, to run it, use:
```bash
go run simplex/simplex.go
```

***
*Note: to run it, you should have golang installed on your pc locally, otherwise there should be a test and run results in .pdf file for the task*