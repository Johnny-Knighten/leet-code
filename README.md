# Leet Code Solutions

This repo contains my leet code submissions for 2024 and onward. I had another repo with mostly Java solutions from back around 2019, but decided to start fresh again. I will be trying to answer all questions in both Python and Go.

## Go

Each problem is labled `problem_*.go` and has an assoicated unit test file called `problem_*_test.go`.

To run all tests, open a terminal and change to the `go` directory then:
```bash
$ go test -v
```

To run a specific test:
```bash
$ go test -v problem_1.go problem_1_test.go
```

## Python

For all Python solutions test cases are included in the script entry point; that is the `if __name__ == "__main__"` section in each script.

To run a specific Python solution, open a terminal and change to the `python` directory then:
```bash
$ python problem_1.py
```

## Dev Container

I have included a devcontainer spec that has all dependencies needed to run the provided solutions. You should be able to open it in VS Code or a JetBrains IDE without any issue.

See [here](https://code.visualstudio.com/docs/devcontainers/containers) for more details.
