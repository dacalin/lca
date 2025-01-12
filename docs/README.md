# Debian Packaging

This document provides instructions on building and installing the Debian package for **lca**.

## Prerequisites

Ensure you have the necessary tools installed:

- `dpkg-deb`: For building Debian packages.
- `make`: For running the build script.

You can install these tools using the following command:

```bash
sudo apt-get update
sudo apt-get install dpkg-dev make
```

## Building the Debian Package

To build the Debian package, follow these steps:

1. **Navigate to the Debian Packaging Directory:**

    ```bash
    cd linux
    ```

2. **Run the Build Script:**

    ```bash
    ./make_deb.sh
    ```

    This script creates the `lca_debian.deb` package in the `linux` directory.

## Installing the Debian Package

After building the package, install it using `dpkg`:

```bash
sudo dpkg -i lca_debian.deb
```

If there are dependency issues, resolve them with:

```bash
sudo apt-get install -f
```

## Control File

The Debian control file is located at `linux/DEBIAN/control`. It contains metadata about the package, such as:

- **Package**: The name of the package.
- **Version**: The package version.
- **Section**: The package section.
- **Priority**: The package priority.
- **Architecture**: The supported architecture.
- **Maintainer**: The maintainer's contact information.
- **Description**: A brief description of the package.

### Example Control File

```plaintext
Package: lca
Version: 1.0.0
Section: base
Priority: optional
Architecture: all
Maintainer: casale.candoit@gmail.com
Description: LCA: An alternative to ls that computes file hashes, provides color-coded, user-friendly output, and supports configurable recursion depth.
```

Ensure this file is properly configured before building the package.

## Uninstalling the Package

To remove the package, use:

```bash
sudo dpkg -r lca
```

## Troubleshooting

- **Dependency Issues**: If installation fails due to missing dependencies, run `sudo apt-get install -f` to fix them.
- **Build Errors**: Ensure all prerequisites are installed and the `make_deb.sh` script has execute permissions (`chmod +x make_deb.sh`).
- **Script Execution Issues**: If the build script fails to execute, verify that all necessary files are present in the `linux` directory and that the script has the correct permissions.

## Additional Information

For more details on Debian packaging, refer to the [Debian New Maintainers' Guide](https://www.debian.org/doc/manuals/maint-guide/).
