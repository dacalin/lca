# lca

## Description

**lca** is an alternative to `ls` that computes file hashes, provides color-coded, user-friendly output, and supports configurable recursion depth. It leverages both Go and Ruby to deliver robust performance and flexibility.

## Installation

### Using Debian Package

To install the `lca` Debian package, follow these steps:

1. Download the Debian package:

    ```bash
    wget https://github.com/dacalin/lca/releases/download/v1.0.0/lca_debian.deb
    ```

2. Install the package using `dpkg`:

    ```bash
    sudo dpkg -i lca_debian.deb
    ```

3. Resolve any dependencies:

    ```bash
    sudo apt-get install -f
    ```

### From Source

If you prefer to build `lca` from source:

1. Clone the repository:

    ```bash
    git clone https://github.com/dacalin/lca.git
    ```

2. Navigate to the project directory:

    ```bash
    cd lca
    ```

3. Build the project:

    ```bash
    make
    ```

4. Install the binary:

    ```bash
    sudo make install
    ```

## Usage

After installation, you can use `lca` via the command line:

```bash
lca [options]
```

### Examples

- **Display Help Information:**
    ```bash
    lca -h
    ```

- **Compute MD5 File Hashes:**
    ```bash
    lca --hash md5
    ```

- **Compute SHA1 File Hashes (Using Shorthand):**
    ```bash
    lca -H sha1
    ```

- **Show Permissions and Owners:**
    ```bash
    lca -p
    ```

- **Set Recursion Depth to 2:**
    ```bash
    lca -r 2
    ```

- **Combine Multiple Options:**
    ```bash
    lca --hash sha256 -p -r 3
    ```

## Documentation

Comprehensive documentation is available [here](docs/README.md).

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch:
    ```bash
    git checkout -b feature/YourFeature
    ```
3. Commit your changes:
    ```bash
    git commit -m "Add your message"
    ```
4. Push to the branch:
    ```bash
    git push origin feature/YourFeature
    ```
5. Open a pull request.

Please ensure all tests pass and adhere to the code style guidelines.

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

For any inquiries or support, please contact [casale.candoit@gmail.com](mailto:casale.candoit@gmail.com).
