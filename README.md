# ASI - Auto Software Installer

AutoInstaller is a command-line interface (CLI) application written in Go (Golang) that facilitates the installation of various software packages. It automates the download and installation process based on a predefined list of software.

## Features

- **Search**: Search for software within the authorized installer list based on the provided keyword.
- **Install**: Installs the specified software from the allowed installer list.
- **List**: Lists all available software packages.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/sepehr-safaeian/ASI
   ```

2. Navigate to the project directory:

   ```bash
   cd autoinstaller
   ```

3. Build the application:

   ```bash
   go build
   ```

## Usage

### Help

Display detailed information about all available commands:

```bash
./autoinstaller help
```

### Search

Search for software within the authorized installer list:

```bash
./autoinstaller search <software>
```

### Install

Install the specified software from the allowed installer list:

```bash
./autoinstaller install <software>
```

### List

List all available software packages:

```bash
./autoinstaller list
```

## Configuration

The list of software packages and their download URLs are configured in the `config/softwares.yaml` file. You can add, remove, or update software packages in this file as needed.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
