# Traefik Plugin Development Environment

This project provides a preconfigured development environment for building and testing [Traefik](https://traefik.io/) plugins in **Local Plugins mode**.  
It is based on a VS Code Dev Container and a `docker-compose` setup that runs Traefik, a test service (`whoami`), and your local plugin together.

## Purpose

The goal of this project is to offer a consistent, reproducible, and convenient environment for Traefik plugin development, supporting both local unit tests and integration tests with a running Traefik instance.

This configuration enables:
- **Development of Traefik plugins in Go** with full IDE support (IntelliSense, linting, debugging).
- **Automated provisioning** of a Traefik test setup with the `whoami` service via `docker compose`.
- **Fast development feedback loop** through tasks for starting, restarting, and viewing logs.

## Architecture Overview

- **Dev Container (VS Code)**  
  Contains the Go toolchain, Delve debugger, and additional tools for plugin development.
- **Traefik Container**  
  Loads the plugin in Local Plugins mode and exposes it via a configured router.
- **Whoami Container**  
  Minimal HTTP test service to validate plugin behavior.

## Go Toolchain in the Dev Container – Purpose and Benefits

Traefik plugins in Local Plugins mode are not shipped as precompiled Go binaries.  
Instead, Traefik loads the plugin source code at runtime and interprets it using [Yaegi](https://github.com/traefik/yaegi).  
Therefore, a Go installation inside the Traefik container is **not** required for execution.

However, having a Go toolchain in the development environment is recommended for:

1. **Running Unit Tests**  
   Test plugin logic in isolation from Traefik (`go test ./...`) to significantly reduce the feedback loop.

2. **Modern Development Tooling**  
   Code completion, navigation, static analysis (`go vet`, `staticcheck`), and refactoring tools are readily available.

3. **Dependency Management**  
   Commands like `go mod tidy` keep `go.mod` and `go.sum` consistent, preventing runtime errors when Traefik loads the plugin.

4. **Quality Assurance**  
   Local tests and static analysis help catch issues early and improve long-term maintainability.

**Conclusion:**  
The Go installation in the Dev Container exists solely for development and quality assurance.  
Traefik itself does not require Go to execute Local Plugins.

## Start & Test

1. **VS Code:** Open the project in VS Code and select **“Reopen in Container”**.  
2. **Start Traefik:** Run the VS Code task `compose up`.  
3. **Verify functionality:**  
   ```bash
   curl -i -H "Host: demo.localhost" http://localhost:8080/
   ```
   The response should contain the HTTP header `X-Plugin: Hello`.  
4. **Change code & retest:**  
   After modifying the plugin code, run the VS Code task `compose restart traefik` and repeat the test command.

## Requirements

- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- [Visual Studio Code](https://code.visualstudio.com/) with the [Dev Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
- Shared project directories in Docker Desktop (on Windows/macOS)

## License

[MIT](LICENSE)