# Feature Flag Operator

The Feature Flag Operator is a Kubernetes operator written in Go that integrates with various feature flag management services. It automatically fetches feature flags from the configured services via HTTP calls and injects them into Kubernetes resources.

## Supported Feature Flag Solutions

The Feature Flag Operator currently supports the following feature flag management services:

- [Flipt](https://flipt.io/)
- ...

## Features

- Automatically fetches feature flags from configured feature flag management services.
- Injects fetched feature flags into Kubernetes resources.

## Installation

### Prerequisites

- Go (version 1.21+)
- Kubernetes cluster (minikube, kind, or any other Kubernetes cluster)
- Access to the configured feature flag management service(s)

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/usysrc/feature-flag-operator
   ```

2. Navigate to the repository directory:

   ```bash
   cd feature-flag-operator
   ```

3. Build the operator:

   ```bash
   go build -o ffo .
   ```

4. Run the operator:

   ```bash
   ./ffo
   ```

## Usage

1. Ensure that the operator is running in your Kubernetes cluster.

2. Create or update your Kubernetes resources (e.g., Deployments, Pods) with annotations to specify the feature flags you want to enable/disable.

3. The operator will automatically fetch feature flags from the configured feature flag management services and inject them into the annotations of your Kubernetes resources.

## Configuration

Not yet implemented: You can configure the operator by providing environment variables or configuration files. Below are the supported configuration options:

- `FLIPT_BASE_URL`: Base URL of the Flipt service.

## Contributing

Contributions to the Feature Flag Operator are welcome! If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
