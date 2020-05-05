# Pandora Security Driver for Go

[![GitHub Release](https://img.shields.io/github/release/pandora-security/pandora-go.svg)](https://github.com/pandora-security/pandora-go/releases)

An official Go module which will help any Go application to communicate with Pandora Security Box APIs.

## List of Contents

* [Pandora Security Driver for Go](#pandora-security-driver-for-go)
  * [List of Contents](#list-of-contents)
  * [Getting Started](#getting-started)
    * [Requirements](#requirements)
    * [Installation](#installation)
  * [Documentation](#documentation)
  * [Built With](#built-with)
  * [Contribution](#contribution)
  * [Version Management](#version-management)
  * [Authors](#authors)
  * [License](#license)

## Getting Started

### Requirements

For this library to be functional, you shall:

 1. have your application registered with Pandora's Author and get your application GUID; and
 2. have Pandora installed on your system.
 
For more information about registering an application and installing Pandora, please visit [Pandora official page](https://pandora-security.github.io/Pandora).

### Installation

To get Pandora included on your project, get this module either using `go get` or `dep`.

```bash
go get -u github.com/pandora-security/pandora-go
```

## Documentation

`Pandora` struct provide representation of Pandora APIs that bound specifically to your application by the provided application GUID.

Documentation about the `Pandora` struct will be available soon. For now, please look into the definition provided within the module.

## Built With

Written in [Go](https://golang.org/) v1.14.

## Contribution

To contribute, simply fork this project, and issue a pull request.

## Version Management

We use [Semantic Versioning](http://semver.org/) for version management. For the versions available, see the [tags on this repository](https://github.com/pandora-security/pandora-go/tags).

## Authors

* **Danang Galuh Tegar Prasetyo** - _Initial work_ - [danang-id](https://github.com/danang-id)

## License

Pandora Security Driver for Go is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details