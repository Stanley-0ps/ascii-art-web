# ASCII Art Web

ASCII Art Web is a Go web application that converts plain text into stylized ASCII art. It provides a browser interface for choosing a banner font, entering text, generating the result, previewing available fonts, and downloading the generated output as a text file.

## Features

- Generate ASCII art from browser input.
- Choose from embedded banner fonts stored in `asciiart/banners`.
- Preview the same input across available fonts.
- Download generated art as `ascii.txt`.
- Serve HTML templates and CSS from Go embedded files.
- Optional dark mode and alternate UI style.
- Custom runtime fonts can be loaded from a local `banners/` directory.

## Requirements

- Go 1.22.2 or newer.
- A web browser.

## Run Locally

Clone the project:

```bash
git clone https://github.com/Stanley-0ps/ascii-art-web.git
cd ascii-art-web
```

Start the server:

```bash
go run .
```

Open the app in your browser:

```text
http://localhost:8080
```

To build and run a binary:

```bash
go build -o ascii-art-web .
./ascii-art-web
```

## Configuration

The server reads these environment variables:

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `8080` | HTTP port used by the web server. |
| `STYLE` | `A` | UI style. Set to `K` to use the alternate template and stylesheet. |

Examples:

```bash
PORT=3000 go run .
```

```bash
STYLE=K go run .
```

## Docker

The repository includes a Dockerfile and helper scripts.

Build an image:

```bash
docker build -t ascii-art-web .
```

Run a container:

```bash
docker run --rm -p 8080:8080 -e PORT=8080 ascii-art-web
```

Or use the interactive helper script:

```bash
./run_Docker.sh
```

## How It Works

The application exposes three main routes:

| Route | Method | Purpose |
| --- | --- | --- |
| `/` | `GET` | Render the home page and font selector. |
| `/ascii-art` | `POST` | Generate, preview, or download ASCII art from form input. |
| `/css/` | `GET` | Serve embedded CSS files. |

The ASCII generation logic is in `asciiart/asciiArt.go`. The web server is in the `server` package, and `main.go` embeds the frontend files before starting the server.

## Project Structure

```text
.
|-- asciiart/              # ASCII art generation logic and banner fonts
|-- frontend/
|   |-- css/               # Stylesheets
|   `-- templates/         # HTML templates
|-- server/                # HTTP handlers, validation, and errors
|-- main.go                # Application entry point and embedded files
|-- Dockerfile             # Container build file
`-- README.md
```

## Input Notes

Only printable ASCII characters are supported, plus newlines. Non-ASCII characters are rejected by the generator.

Runtime font files placed in `./banners` must follow the supported banner format used by the project.

## Author

[Stanley-0ps](https://github.com/Stanley-0ps)
