## LFP - License for Project
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Language](https://img.shields.io/badge/Go-1.10-blue.svg)](https://golang.org/)
[![Build Status](https://travis-ci.org/YuriyLisovskiy/lfp.svg?branch=master)](https://travis-ci.org/YuriyLisovskiy/lfp)
[![Project Status](https://img.shields.io/badge/status-development-red.svg)](https://travis-ci.org/YuriyLisovskiy/lfp)

[LFP](https://github.com/YuriyLisovskiy/lfp) - tool which can help you to create the license and license notice for
every file in the project.

Currently it supports only built-in open source licenses, see [licenses](docs/licenses.md)
for details. Using this tool you can create your own license notice template as described [here](docs/custom-notice.md),
and use it in your project.

Read [docs](docs) for more details.

## Download/Update and Install
Download an archive with the latest release manually
[here](https://github.com/YuriyLisovskiy/lfp/releases)   

```bash
sudo tar -xvzf ~/Downloads/<version>.tar.gz -C /usr/local/bin/
```

If you have already installed LFP tool of version at least 1.0.2-stable, run the next
command to get update:
```bash
sudo lfp --update-latest
```
Ensure that LFP have been updated:
```bash
lfp --version
```

## Usage
Add license to your project, see how to create config file [here](docs/create-config.md):
```bash
lfp -c path/to/config/config.yml
```
Get the latest release:
```bash
sudo lfp --update latest
```
Get specific release version, for example 1.0.1-beta:
```bash
sudo lfp --update 1.0.1-beta
```
Read usage:
```bash
lfp --help
```
Check current version:
```bash
lfp --version
```

### Issues
Read [issue template](.github/ISSUE_TEMPLATE.md) before opening a new issue.

### Contributing
Read [contributing](.github/CONTRIBUTING.md) docs before opening new pull request.

### Author
* [Yuriy Lisovskiy](https://github.com/YuriyLisovskiy)

### License
[LFP](https://github.com/YuriyLisovskiy/lfp) is licensed under the terms of the
[MIT](https://opensource.org/licenses/MIT) software license, see the [LICENSE](LICENSE) file for more information.
