# File Mover

File Mover is a script written in Go that automates the task of organizing downloaded files into their appropriate destination folders based on their file types. This script continuously monitors a specified source folder, such as the Downloads folder, for new files. Upon detection, it moves the files to their corresponding destination folders and maintains a log of the file movements.

## Features

- Automatically organizes downloaded files into destination folders based on file types.
- Monitors the specified source folder continuously for new files.
- Updates a log file with information about each file movement.
- Customizable source folder, destination folders, and log file path.

## Prerequisites

- Go 1.16 or above

## Usage

1. Clone the repository:

https://github.com/aniketbadole/file_mover.git

2. Modify the source folder, destination folders, and log file path in the code to match your system configuration and requirements.

3. Build and run the program:

go run file_mover.go

4. The script will start monitoring the source folder and automatically move files to their appropriate destination folders based on their file types.

## Customization

- **Source Folder**: Modify the `sourceFolder` variable in the code to specify the folder to monitor for new files.

- **Destination Folders**: Update the `destinationFolder` map in the code to define the destination folders for different file types. Add or remove file types and their corresponding folders as needed.

- **Log File**: Adjust the `logFile` variable in the code to set the path for the log file that records file movements.

## License

This project is licensed under the [MIT License](LICENSE).

