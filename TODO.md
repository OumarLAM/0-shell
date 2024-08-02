# Unix Shell in Go - TODO List

## 1. Project Setup
- [x] Initialize a new Go project
- [x] Set up the project structure
- [x] Create a main.go file

## 2. Basic Shell Loop
- [x] Implement a basic loop that:
  - [x] Displays the prompt ($)
  - [x] Reads user input
  - [x] Processes the input
  - [x] Executes commands
  - [x] Repeats

## 3. Command Parsing
- [x] Implement a function to parse user input into command and arguments
- [x] Handle quoting and escaping in command arguments

## 4. Built-in Commands Implementation
- [x] echo
- [x] cd
- [ ] ls (with -l, -a, -F flags)
- [x] pwd
- [ ] cat
- [ ] cp
- [ ] rm (with -r flag)
- [ ] mv
- [ ] mkdir

## 5. Error Handling
- [ ] Implement error handling for each command
- [ ] Display appropriate error messages

## 6. Signal Handling
- [ ] Implement Ctrl+D (EOF) handling to exit the shell

## 7. Testing
- [ ] Write unit tests for each command
- [ ] Perform integration testing

## 8. Code Optimization and Refactoring
- [ ] Review and optimize the code
- [ ] Ensure good coding practices are followed

## 9. Documentation
- [ ] Write comments and documentation for the code

## Next Steps
1. Add error handling for each command
2. Implement the flags for ls (-l, -a, -F) and rm (-r)