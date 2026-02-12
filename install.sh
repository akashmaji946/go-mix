#!/bin/bash

# Go-Mix Installation Script
# This script builds go-mix from source and installs it to a standard system path
# making it accessible from anywhere in the system

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper functions
print_header() {
    echo ""
    echo -e "${BLUE}════════════════════════════════════════════════════════${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}════════════════════════════════════════════════════════${NC}"
    echo ""
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_error() {
    echo -e "${RED}✗ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ $1${NC}"
}

# Check if running as root (required for system-wide installation)
check_root() {
    if [[ $EUID -ne 0 ]]; then
        echo ""
        print_warning "This script requires root privileges to install to /usr/local/bin"
        print_info "Re-running with sudo..."
        echo ""
        exec sudo "$0" "$@"
    fi
}

# Check if Go is installed
check_go() {
    print_info "Checking for Go installation..."
    
    # Try multiple common Go paths
    local go_paths=(
        "$(command -v go 2>/dev/null)"
        "/usr/local/go/bin/go"
        "/usr/bin/go"
        "$HOME/go/bin/go"
        "/opt/go/bin/go"
    )
    
    local go_binary=""
    for path in "${go_paths[@]}"; do
        if [ -x "$path" ] 2>/dev/null; then
            go_binary="$path"
            break
        fi
    done
    
    if [ -z "$go_binary" ]; then
        print_error "Go is not installed or not found in PATH"
        echo ""
        echo "Tried looking in:"
        for path in "${go_paths[@]}"; do
            echo "  - $path"
        done
        echo ""
        echo "Please install Go from: https://golang.org/doc/install"
        echo "Minimum required version: Go 1.18"
        exit 1
    fi
    
    # Add Go to PATH if not already there
    if [[ ":$PATH:" != *":$(dirname "$go_binary"):"* ]]; then
        export PATH="$(dirname "$go_binary"):$PATH"
        print_info "Added Go to PATH: $(dirname "$go_binary")"
    fi
    
    local go_version=$("$go_binary" version | awk '{print $3}' | sed 's/go//')
    print_success "Found Go version: $go_version"
    print_info "Go binary: $go_binary"
}

# Build go-mix
build_gomix() {
    print_header "Building Go-Mix"
    
    # Get the directory where this script is located
    local script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    
    if [ ! -f "$script_dir/go.mod" ]; then
        print_error "go.mod not found. Make sure you're running this script from the go-mix root directory"
        exit 1
    fi
    
    print_info "Building from: $script_dir"
    
    cd "$script_dir"
    
    # Build the executable from main package
    print_info "Running: go build -o go-mix ./main"
    if go build -o go-mix ./main; then
        print_success "Build completed successfully"
    else
        print_error "Build failed"
        exit 1
    fi
    
    if [ ! -f "go-mix" ]; then
        print_error "go-mix executable not found after build"
        exit 1
    fi
    
    print_success "Executable created: ./go-mix"
}

# Install go-mix
install_gomix() {
    print_header "Installing Go-Mix"
    
    local script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    local source_binary="$script_dir/go-mix"
    local install_path="/usr/local/bin/go-mix"
    
    if [ ! -f "$source_binary" ]; then
        print_error "Source binary not found at $source_binary"
        exit 1
    fi
    
    print_info "Installing to: $install_path"
    
    # Copy binary to installation path
    if cp "$source_binary" "$install_path"; then
        print_success "Binary copied to $install_path"
    else
        print_error "Failed to copy binary"
        exit 1
    fi
    
    # Make it executable
    if chmod +x "$install_path"; then
        print_success "Made executable: chmod +x $install_path"
    else
        print_error "Failed to make executable"
        exit 1
    fi
}

# Verify installation
verify_installation() {
    print_header "Verifying Installation"
    
    # Check if go-mix is in PATH
    if ! command -v go-mix &> /dev/null; then
        print_warning "go-mix not found in PATH"
        print_info "You may need to restart your terminal or run: source ~/.bashrc"
        return 1
    fi
    
    print_success "go-mix found in PATH"
    
    # Show location
    local location=$(which go-mix)
    print_info "Location: $location"
    
    # Show version/help
    print_info "Testing execution..."
    if go-mix samples/functions/01_basic_functions.gm > /dev/null 2>&1; then
        print_success "go-mix executed successfully"
    else
        print_warning "Could not test with sample file (may not be in expected location)"
    fi
    
    return 0
}

# Cleanup
cleanup() {
    print_header "Cleanup"
    
    local script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    
    # You can optionally remove the built binary from the source directory
    # Uncomment below if you want to clean up
    # if [ -f "$script_dir/go-mix" ]; then
    #     rm "$script_dir/go-mix"
    #     print_success "Removed source binary"
    # fi
    
    print_success "Cleanup complete"
}

# Show final information
show_final_info() {
    print_header "Installation Complete!"
    
    echo "Go-Mix has been successfully installed to your system."
    echo ""
    echo "You can now use go-mix from anywhere:"
    echo ""
    echo -e "${GREEN}  go-mix /path/to/script.gm${NC}        # Run a Go-Mix script"
    echo -e "${GREEN}  go-mix${NC}                             # Start interactive REPL"
    echo ""
    echo "Useful commands:"
    echo -e "  ${BLUE}which go-mix${NC}                   # Show installation path"
    echo -e "  ${BLUE}go-mix -h${NC}                     # Show help (if available)"
    echo ""
    echo "Next steps:"
    echo "  1. Try running a sample: go-mix samples/algo/05_factorial.gm"
    echo "  2. Install VS Code extension: sudo $0 install-extension"
    echo "  3. Or search 'GoMix' in VS Code Extensions (Ctrl+Shift+X)"
    echo "  4. Check documentation: https://github.com/akashmaji946/go-mix"
    echo ""
}

# Install VS Code Extension
install_vscode_extension() {
    print_header "Installing VS Code Extension"
    
    # Check if VS Code is installed
    if ! command -v code &> /dev/null; then
        print_warning "VS Code is not installed or not in PATH"
        echo ""
        echo "To install the GoMix VS Code extension:"
        echo "  1. Install VS Code from: https://code.visualstudio.com/"
        echo "  2. Open Extensions (Ctrl+Shift+X)"
        echo "  3. Search for 'GoMix'"
        echo "  4. Click Install"
        echo ""
        echo "Or use this command after installing VS Code:"
        echo "  code --install-extension vscode-ext/go-mix-*.vsix"
        return 1
    fi
    
    print_success "Found VS Code installation"
    
    # Find the most recent .vsix file in vscode-ext directory
    local script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    local vsix_file=""
    
    if [ -d "$script_dir/vscode-ext" ]; then
        vsix_file=$(find "$script_dir/vscode-ext" -name "*.vsix" -type f | sort -V | tail -n1)
    fi
    
    if [ -z "$vsix_file" ] || [ ! -f "$vsix_file" ]; then
        print_warning "No .vsix file found in vscode-ext directory"
        echo ""
        echo "To build the extension:"
        echo "  cd vscode-ext"
        echo "  npm install"
        echo "  npm run vscode:prepublish"
        echo "  vsce package"
        echo ""
        echo "Then run this script again."
        return 1
    fi
    
    print_info "Found extension package: $vsix_file"
    print_info "Installing with VS Code..."
    
    if code --install-extension "$vsix_file" --force; then
        print_success "VS Code extension installed successfully!"
        print_info "Package: $(basename "$vsix_file")"
        echo ""
        echo "The GoMix extension is now ready to use!"
        echo "Restart VS Code to activate the extension."
    else
        print_error "Failed to install VS Code extension"
        echo "Try installing manually with:"
        echo "  code --install-extension $vsix_file"
        return 1
    fi
}

# Uninstall function
uninstall_gomix() {
    print_header "Uninstalling Go-Mix"
    
    local install_path="/usr/local/bin/go-mix"
    
    if [ -f "$install_path" ]; then
        if rm "$install_path"; then
            print_success "Go-Mix uninstalled successfully"
            print_info "Removed: $install_path"
        else
            print_error "Failed to remove Go-Mix"
            exit 1
        fi
    else
        print_warning "Go-Mix not found at $install_path"
    fi
}

# Main execution
main() {
    echo ""
    print_header "Go-Mix Installation Script"
    
    # Parse command line arguments
    case "${1:-install}" in
        uninstall)
            check_root
            uninstall_gomix
            ;;
        install)
            check_root
            check_go
            build_gomix
            install_gomix
            cleanup
            
            # Small delay to let system update PATH
            sleep 1
            
            if verify_installation; then
                show_final_info
            else
                print_warning "Installation completed but verification had issues"
                echo "Try restarting your terminal and running: go-mix"
            fi
            ;;
        install-extension)
            # Install VS Code extension without requiring root
            install_vscode_extension
            ;;
        rebuild)
            # Rebuild without requiring root (builds to current directory)
            check_go
            build_gomix
            print_info "Binary built to: ./go-mix"
            print_info "To install system-wide, run: sudo $0 install"
            ;;
        clean)
            print_header "Cleaning Build Artifacts"
            local script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
            if [ -f "$script_dir/go-mix" ]; then
                rm "$script_dir/go-mix"
                print_success "Removed: $script_dir/go-mix"
            fi
            ;;
        *)
            echo "Usage: $0 [command]"
            echo ""
            echo "Commands:"
            echo "  install              Build and install go-mix to /usr/local/bin (default)"
            echo "  install-extension    Install GoMix VS Code extension"
            echo "  uninstall            Remove go-mix from /usr/local/bin"
            echo "  rebuild              Build go-mix to current directory (no installation)"
            echo "  clean                Remove build artifacts"
            echo ""
            echo "Examples:"
            echo "  sudo $0                        # Install system-wide (with sudo)"
            echo "  $0 install-extension           # Install VS Code extension"
            echo "  $0 rebuild                     # Just build, don't install"
            echo "  sudo $0 uninstall              # Remove from system"
            exit 1
            ;;
    esac
    
    echo ""
}

# Run main function
main "$@"
