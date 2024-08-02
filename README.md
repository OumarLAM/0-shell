# 🐚 0-Shell: Your Simple Unix Shell 🌟

## Overview

Welcome to the **0-Shell Project**! This project aims to create a simple Unix shell, allowing us to explore the core of the Unix system, focusing on process creation and synchronization.

## 🎯 Objectives

The main objectives of this project are:
1. **Create a Simple Unix Shell**: Implement a basic shell capable of executing common commands.
2. **Process Creation and Synchronization**: Learn how commands in a shell create and synchronize processes.
3. **Command Implementation**: Implement commands from scratch without calling external binaries.

## 🚀 Features

Our 0-Shell supports the following commands:
- `echo`
- `cd`
- `ls` (with flags `-l`, `-a`, and `-F`)
- `pwd`
- `cat`
- `cp`
- `rm` (with flag `-r`)
- `mv`
- `mkdir`
- `exit`

## 🔧 Instructions

### Basic Requirements

1. **Command Line Interface**: Our shell displays a simple `$` prompt and wait for a command.
2. **Execution**: The `$` prompt reappears only after the command execution is complete.
3. **Error Handling**: Display appropriate error messages.
4. **Interruption Management**: Handle program interruptions like `Ctrl + D`.
5. **Compiled Language**: In our case, we're using the Go Programming Language.
6. **Good Practices**: Follow good coding practices and guidelines.

### Implementing Commands

We've implemented each command from scratch without using external binaries. Here's a brief overview of what each command does:

- `echo`: Print text to the console.
- `cd`: Change the current directory.
- `ls`: List directory contents with optional flags.
- `pwd`: Print the current working directory.
- `cat`: Concatenate and display files.
- `cp`: Copy files or directories.
- `rm`: Remove files or directories.
- `mv`: Move or rename files or directories.
- `mkdir`: Create a new directory.
- `exit`: Exit the shell.

## 🎁 Bonus Features

Enhance our shell with the following bonuses:
- **Low-Level System Calls**: Implement commands using low-level system calls directly.
- **Ctrl + C Management**: Handle interruption signals like `Ctrl + C`.
- **Autocomplete**: Implement command-line autocompletion.
- **Piping and Redirection**: Add support for command piping and redirection.
- **Path in Prompt**: Display the current path in the shell prompt.
- **Colored Output**: Add colors for directories and errors.
- **Advanced Commands**: Implement additional Unix commands as you see fit.

## 📜 Usage Example

```shell
student$ ./0-shell
$ cd dev
$ pwd
dev
$ ls -l
total 0
crw-------  1 root   root     10,    58 fev  5 09:21 acpi_thermal_rel
crw-r--r--  1 root   root     10,   235 fev  5 09:21 autofs
drwxr-xr-x  2 root   root           540 fev  5 09:21 block
crw-------  1 root   root     10,   234 fev  5 09:21 btrfs-control
drwxr-xr-x  3 root   root            60 fev  5 09:20 bus
drwxr-xr-x  2 root   root          4400 fev  5 09:21 char
crw-------  1 root   root      5,     1 fev  5 09:21 console
lrwxrwxrwx  1 root   root            11 fev  5 09:20 core -> /proc/kcore
drwxr-xr-x  2 root   root            60 fev  5 09:20 cpu
crw-------  1 root   root     10,    59 fev  5 09:21 cpu_dma_latency
$ something
Command 'something' not found
$ echo "Hello There"
Hello There
$ exit
student$
```

## Authors

👤 **Oumar LAM**
- Github: [Oumar LAM](https://github.com/OumarLAM)
- Twitter: [@OumarLAM](https://twitter.com/oumarlam_fcb)

👤 **Abdou TOP**
- Github: [Oumar LAM](https://github.com/abdotop)

## Let's Connect! 🌐

Got questions, ideas, or just want to say hi? Reach out to me on [Twitter](https://twitter.com/OumarLAM) or by sending a mail to oumarlam154@gmail.com . I'd love to hear from you!

Happy coding and enjoy building your 0-Shell! 🌟🔧🐚