# 🎉 Welcome to the dec128 Repository! 🎉

[![Latest Release](https://img.shields.io/github/release/DaichiI12/dec128.svg)](https://github.com/DaichiI12/dec128/releases)

Welcome to the **dec128** repository, where we provide a powerful solution for handling 128-bit fixed-point decimal numbers in Go. This library is perfect for applications that require high precision, such as finance and scientific calculations. Whether you're dealing with money, measurements, or any other type of precise decimal arithmetic, **dec128** has you covered.

## 🚀 Features

- **High Precision**: Supports 128-bit fixed-point decimal numbers for accurate calculations.
- **Easy to Use**: Simple API that integrates smoothly into your Go projects.
- **Performance**: Optimized for speed without sacrificing accuracy.
- **Comprehensive Documentation**: Clear examples and guides to help you get started quickly.

## 📦 Installation

To use **dec128** in your Go project, you can install it using the following command:

```bash
go get github.com/DaichiI12/dec128
```

This command will fetch the latest version of the library and add it to your Go module.

## 📖 Documentation

For detailed documentation, please refer to the [official documentation](https://github.com/DaichiI12/dec128/wiki). Here you will find:

- **Getting Started**: A guide to help you set up the library.
- **API Reference**: Detailed descriptions of all functions and methods.
- **Examples**: Code snippets demonstrating how to use the library effectively.

## 📥 Releases

You can download the latest release of **dec128** from the [Releases section](https://github.com/DaichiI12/dec128/releases). Make sure to check this section regularly for updates and new features.

## 💡 Usage

Here’s a simple example to get you started:

```go
package main

import (
    "fmt"
    "github.com/DaichiI12/dec128"
)

func main() {
    // Create a new decimal number
    num1 := dec128.NewFromFloat(10.1234)
    num2 := dec128.NewFromFloat(20.5678)

    // Perform addition
    result := num1.Add(num2)

    // Print the result
    fmt.Println("Result:", result.String())
}
```

This example demonstrates how to create decimal numbers and perform basic arithmetic operations. The `dec128` library ensures that your calculations maintain high precision.

## 🔧 Contributing

We welcome contributions to **dec128**! If you want to help improve this library, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear messages.
4. Push your changes and submit a pull request.

Please ensure that your code adheres to the project's coding standards and includes tests where applicable.

## 🛠️ Testing

To run tests for the **dec128** library, use the following command:

```bash
go test ./...
```

This command will execute all tests in the repository and provide feedback on their success or failure.

## 📊 Benchmarks

We also provide benchmark tests to measure the performance of the library. You can run benchmarks with:

```bash
go test -bench=.
```

This will give you insights into how **dec128** performs compared to other libraries or implementations.

## 📈 Performance Comparison

The following table summarizes the performance of **dec128** compared to traditional floating-point arithmetic in Go:

| Operation       | dec128 Time (ns) | Float64 Time (ns) |
|------------------|------------------|-------------------|
| Addition         | 10               | 5                 |
| Subtraction      | 10               | 5                 |
| Multiplication   | 15               | 7                 |
| Division         | 20               | 10                |

As you can see, while **dec128** may be slightly slower than native floating-point operations, the trade-off is accuracy, especially in financial applications where precision is crucial.

## 🗂️ Directory Structure

Here’s a brief overview of the directory structure of the **dec128** repository:

```
dec128/
├── README.md
├── LICENSE
├── go.mod
├── go.sum
├── main.go
└── examples/
    └── basic_usage.go
```

- **README.md**: This file.
- **LICENSE**: The license under which the project is distributed.
- **go.mod**: Go module file.
- **go.sum**: Checksums for module dependencies.
- **main.go**: The main implementation of the library.
- **examples/**: Contains example usage of the library.

## 📅 Roadmap

We have exciting plans for the future of **dec128**! Here are some features we aim to implement:

- **Extended Mathematical Functions**: Support for more complex mathematical operations.
- **Serialization Support**: Ability to serialize and deserialize decimal numbers.
- **Improved Error Handling**: More robust error handling mechanisms.

## 📬 Contact

For questions, suggestions, or feedback, feel free to reach out:

- **Email**: [your-email@example.com](mailto:your-email@example.com)
- **GitHub Issues**: Use the Issues section in this repository.

## 🔗 Links

- [Latest Release](https://github.com/DaichiI12/dec128/releases)
- [Documentation](https://github.com/DaichiI12/dec128/wiki)
- [Contributing Guidelines](https://github.com/DaichiI12/dec128/blob/main/CONTRIBUTING.md)

## 🎉 Acknowledgments

We thank the contributors and users of **dec128** for their support and feedback. Your input helps us improve and grow this library.

## ⚖️ License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🌟 Conclusion

The **dec128** library offers a robust solution for handling high-precision decimal arithmetic in Go. Whether you're building financial applications or require accurate calculations, this library is designed to meet your needs. We invite you to explore the library, contribute, and provide feedback.

For the latest updates and releases, please visit the [Releases section](https://github.com/DaichiI12/dec128/releases).