# Go-Mix Installation Guide

This guide explains how to install Go-Mix on your system so it's accessible from anywhere.

## Quick Start (Recommended)

```bash
cd /path/to/go-mix
sudo ./install.sh
```

That's it! Go-Mix will be installed to `/usr/local/bin/go-mix` and accessible from anywhere.

## Installation Script Usage

The `install.sh` script in the root directory handles all installation tasks.

### Prerequisites

- **Go 1.18 or higher** - Download from https://golang.org/
- **sudo access** (for system-wide installation)

### Commands

#### Install to System (Default)
```bash
sudo ./install.sh
# or
sudo ./install.sh install
```

Installs go-mix to `/usr/local/bin/go-mix` making it available system-wide.

#### Rebuild Only (No Installation)
```bash
./install.sh rebuild
```

Just builds the executable to `./go-mix` in the current directory without installing system-wide.

#### Uninstall
```bash
sudo ./install.sh uninstall
```

Removes go-mix from `/usr/local/bin/go-mix`.

#### Clean Build Artifacts
```bash
./install.sh clean
```

Removes the compiled `go-mix` binary from the project directory.

#### Install VS Code Extension
```bash
./install.sh install-extension
```

Automatically finds and installs the GoMix VS Code extension (.vsix file) from the vscode-ext directory. 
Requires VS Code to be installed and accessible from the command line.

Alternatively, install manually from VS Code Marketplace:
1. Open VS Code
2. Go to Extensions (Ctrl+Shift+X)
3. Search for "GoMix"
4. Click Install

Or use VS Code command directly:
```bash
code --install-extension vscode-ext/*.vsix --force
```

## What Gets Installed

| Item | Location | Description |
|------|----------|-------------|
| Executable | `/usr/local/bin/go-mix` | The Go-Mix interpreter |
| VS Code Extension | VS Code Extensions | GoMix language support (optional) |
| PATH | Automatic | go-mix added to system PATH on installation |

## VS Code Extension (Optional)

After installing the go-mix interpreter, you can optionally install the VS Code extension:

```bash
./install.sh install-extension
```

The extension provides:
- 🎨 Syntax highlighting for `.gm` files
- ⚡ Run code with `Ctrl+Shift+G`
- 📝 15+ code snippets
- 🔌 REPL integration
- 💡 Keyword hover information
- 🔄 Smart indentation & formatting

## Verifying Installation

After installation, verify it worked:

```bash
# Check if go-mix is in PATH
which go-mix
# Output: /usr/local/bin/go-mix

# Check version/help
go-mix --help

# Run a sample program
go-mix samples/algo/05_factorial.gm
```

## Usage After Installation

Once installed, you can run Go-Mix from anywhere:

```bash
# Run a script
go-mix /path/to/script.gm

# Start interactive REPL
go-mix

# From any directory
cd ~
go-mix path/to/your/script.gm
```

## Manual Installation

If you prefer to install manually:

```bash
# Build
cd /path/to/go-mix
go build -o go-mix

# Install to /usr/local/bin
sudo cp go-mix /usr/local/bin/
sudo chmod +x /usr/local/bin/go-mix

# Verify
go-mix --help
```

## System Paths Explained

### `/usr/local/bin`

This is the standard location for user-installed programs on Unix-like systems:

- **Why this location?**
  - Part of PATH by default on most systems
  - Accessible from any directory
  - Designated for manually installed executables
  - Survives system package manager updates

- **Alternative locations:**
  - `$HOME/.local/bin` - User-specific (not system-wide)
  - `/usr/bin` - Reserved for package manager
  - `/opt` - Larger applications

### PATH

The PATH environment variable tells your shell where to look for executable programs.

Check your PATH:
```bash
echo $PATH
```

Output example:
```
/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
```

Note that `/usr/local/bin` is checked first, so programs there take priority.

## Troubleshooting

### "Permission denied" Error
```bash
# Make sure to use sudo
sudo ./install.sh install

# Or check file permissions
ls -l /usr/local/bin/go-mix
# Should show: -rwxr-xr-x (755 permissions)
```

### "go-mix: command not found" After Installation

Try one of these solutions:

```bash
# 1. Restart terminal (simplest)
exit
# Open new terminal window

# 2. Reload shell configuration
source ~/.bashrc
# or for zsh
source ~/.zshrc

# 3. Verify installation
which go-mix
/usr/local/bin/go-mix --help
```

### "Go is not installed" Error

Install Go first:

```bash
# macOS with Homebrew
brew install go

# Ubuntu/Debian
sudo apt-get install golang-go

# Or download from https://golang.org/dl/
```

### Script "Permission denied"

Make the script executable:

```bash
chmod +x /path/to/go-mix/install.sh
./install.sh
```

## Environment Variables

After installation, you can optionally set these variables:

```bash
# Add to ~/.bashrc or ~/.zshrc:

# Set GOMIX_HOME (optional)
export GOMIX_HOME=/path/to/go-mix

# GOMIX_HOME is useful if you want to refer to go-mix installation directory
```

## Uninstallation

To remove Go-Mix from your system:

```bash
# Using the script
sudo ./install.sh uninstall

# Or manually
sudo rm /usr/local/bin/go-mix
```

## Next Steps

After installation:

1. **Test it:** `go-mix samples/algo/05_factorial.gm`
2. **Install VS Code Extension:** Search "Go-Mix" in VS Code Extensions
3. **Read Documentation:** https://github.com/akashmaji946/go-mix
4. **Write Code:** Create `.gm` files and run them!

## Support

For issues with installation:

1. Verify Go is installed: `go version`
2. Check permissions: `ls -l install.sh`
3. Review error messages carefully
4. Check GitHub: https://github.com/akashmaji946/go-mix/issues

## Advanced: Building from Source

If you want to customize the build:

```bash
# Clean build
go clean
go build -o go-mix

# Build with debug info
go build -ldflags="-s=false" -o go-mix

# Build for specific OS/Architecture
GOOS=linux GOARCH=amd64 go build -o go-mix

# Then install
sudo cp go-mix /usr/local/bin/
sudo chmod +x /usr/local/bin/go-mix
```

## Build Information

The installation script automatically:

1. ✅ Checks for Go installation
2. ✅ Builds from source
3. ✅ Installs to `/usr/local/bin`
4. ✅ Sets executable permissions
5. ✅ Verifies installation
6. ✅ Provides helpful output

No additional configuration needed!

---

**Enjoy using Go-Mix!** 🚀
