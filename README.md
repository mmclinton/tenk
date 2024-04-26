# TENK

## Overview
This tool was designed to increase the efficiency of retrieving a company's 10-k annual report. TENK offers a simple CLI alternative to the SEC's EDGAR tool. TENK only retrieves annual reports filed by a company for a specific year. The tool will retrieve any annual reports, if filed, for years dating back to 1996. You can search over 70,000 symbols.

## Installation

1. **Clone the Repository**:
   Clone the repository to your local machine using the following command:
   ```shell
   git clone https://github.com/mmclinton/tenk.git

2. **Navigate to the Repository**:
   Change into the directory where you downloaded the repository:
   ```shell
   cd tenk

3. **Set Up Your Api Key:**:
   You need to obtain an API key from google.com to use this tool. Once you have obtained your API key, insert it into the following command and run it:
   ```shell
   make api_key="<YOUR API KEY HERE>"
   ```
   This command creates a configuration file (config.json) in the directory ~/.config/tenk with your API key. If something goes wrong during the setup process or if your API key changes later, you can edit the API key directly in the config.json file located in ~/.config/tenk.

## Project Directory Structure
Here's the directory structure of the project:

```mermaid
graph LR;
    root[root]
    root --> cmd
    root --> config
    root --> internal
    root --> go.mod
    root --> Makefile
    root --> README.md
    cmd --> tenk
    tenk --> main.go
    config --> config.go
    internal --> utils
    utils --> fmpAPI.go
    utils --> openBrowser.go
    utils --> outputFormatter.go
    utils --> urlBuilder.go
