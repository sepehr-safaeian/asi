
# AIS - Auto Installer Software

AutoInstaller is a command-line interface (CLI) application written in Go (Golang) that facilitates the installation of various software packages. It automates the download and installation process based on a predefined list of software.


## Features

- **Search:** Users can search for specific software packages within the authorized installer list by providing keywords.
- **Install:** ASI allows users to install desired software packages from the predefined list of allowed installers.
- **List:** Users can view a comprehensive list of all available software packages supported by ASI.

## Why ASI?
**Simplify Installation:** ASI simplifies the complex process of software installation by providing a unified interface and automated installation procedure.

**Standardization:** By using a predefined list of software packages and configurations, ASI ensures consistency and standardization across different systems.

**Reduce Errors:** Automation reduces the likelihood of human errors during the installation process, enhancing system reliability.

## How It Works
ASI (Auto Software Installer) operates by leveraging a predefined list of software packages and their download URLs to automate the process of software installation. This system is designed to eliminate the need for manual intervention and tedious search processes by autonomously fetching the latest versions of designated software from their official websites.

## Key Features and Mechanisms
**Dynamic Update:** ASI regularly updates its internal database of software packages and their download URLs to accommodate changes and ensure the retrieval of the latest versions.

**Customization:** Users can customize ASI's behavior by modifying the list of software packages and their download URLs in the configuration file. This flexibility allows for tailored installations based on specific requirements or preferences.


## Get Started

Clone the repository:

```bash
  git clone https://github.com/sepehr-safaeian/ASI
```
Navigate to the project directory:
```bash
  cd ASI
```
Build the application:
```bash
  go build
```
## Usage/Examples
**Help**
Display detailed information about all available commands:

```bash
./asi help
```
**Search**
for software within the authorized installer list:
```bash
./asi search <software>
```
**Install**
the specified software from the allowed installer list:
```bash
./asi install <software>
```
**List**
List all available software packages:
```bash
./asi list
```
## Configuration

The list of software packages and their download URLs are configured in the **config/softwares.yaml** file. You can add, remove, or update software packages in this file as needed.
