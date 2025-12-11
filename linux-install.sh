#!/bin/bash

# ddl-to-object installation script
# Install ddl-to-object to user home directory

set -e

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print colored messages
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check required commands
check_dependencies() {
    print_info "Checking dependencies..."
    
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed, please install Go first"
        exit 1
    fi
    
    print_success "Dependencies check passed"
}

# Create necessary directories
create_directories() {
    print_info "Creating directory structure..."
    
    # Create .local/bin directory
    mkdir -p "$HOME/.local/bin"
    
    # Create .dto config directory
    mkdir -p "$HOME/.dto"
    
    # Create .dto/template directory
    mkdir -p "$HOME/.dto/template"
    
    print_success "Directory creation completed"
}

# Build binary file
build_binary() {
    print_info "Building ddl-to-object binary..."
    
    # Get version information
    VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
    GIT_COMMIT=$(git rev-parse HEAD 2>/dev/null || echo "unknown")
    BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
    
    # Build flags
    LDFLAGS="-X 'ddl-to-object/cmd.Version=${VERSION}' -X 'ddl-to-object/cmd.GitCommit=${GIT_COMMIT}' -X 'ddl-to-object/cmd.BuildTime=${BUILD_TIME}' -s -w"
    
    # Build binary file
    go build -ldflags "$LDFLAGS" -o "$HOME/.local/bin/ddl-to-object" .
    
    # Set execute permission
    chmod +x "$HOME/.local/bin/ddl-to-object"
    
    print_success "Binary build completed"
}

# Copy configuration files
copy_config() {
    print_info "Copying configuration files..."
    
    # Copy example config file
    if [ -f "config.example.json" ]; then
        cp "config.example.json" "$HOME/.dto/config.json"
        print_success "Config file copied: ~/.dto/config.json"
    else
        print_warning "config.example.json not found, creating default config file"
        cat > "$HOME/.dto/config.json" << 'EOF'
{
  "default_packages": {
    "go": "models",
    "java": "com.yourcompany.domain.entity",
    "php": "App\\Models",
    "python": ""
  },
  "template_dir": "~/.dto/template",
  "log_level": "info",
  "output_settings": {
    "create_directories": true,
    "overwrite_files": true,
    "backup_existing": false
  }
}
EOF
    fi
}

# Copy template files
copy_templates() {
    print_info "Copying template files..."
    
    if [ -d "template" ]; then
        cp -r template/* "$HOME/.dto/template/"
        print_success "Template files copied"
    else
        print_warning "Template directory not found, skipping template copy"
    fi
}

# Update PATH
update_path() {
    print_info "Updating PATH environment variable..."
    
    # Check if .local/bin is already in PATH
    if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
        # Add to .bashrc
        if [ -f "$HOME/.bashrc" ]; then
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.bashrc"
            print_success "Added ~/.local/bin to ~/.bashrc"
        fi
        
        # Add to .zshrc (if exists)
        if [ -f "$HOME/.zshrc" ]; then
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.zshrc"
            print_success "Added ~/.local/bin to ~/.zshrc"
        fi
        
        # Add to .profile
        if [ -f "$HOME/.profile" ]; then
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.profile"
            print_success "Added ~/.local/bin to ~/.profile"
        fi
        
        print_warning "Please reload shell or run: source ~/.bashrc"
    else
        print_info "PATH already contains ~/.local/bin"
    fi
}

# Verify installation
verify_installation() {
    print_info "Verifying installation..."
    
    if [ -f "$HOME/.local/bin/ddl-to-object" ]; then
        print_success "Binary installed successfully: ~/.local/bin/ddl-to-object"
    else
        print_error "Binary installation failed"
        exit 1
    fi
    
    if [ -f "$HOME/.dto/config.json" ]; then
        print_success "Config file installed successfully: ~/.dto/config.json"
    else
        print_error "Config file installation failed"
        exit 1
    fi
    
    if [ -d "$HOME/.dto/template" ]; then
        print_success "Template directory created successfully: ~/.dto/template"
    else
        print_error "Template directory creation failed"
        exit 1
    fi
}

# Show completion information
show_completion() {
    print_success "Installation completed!"
    echo
    echo "Installation locations:"
    echo "  Binary file: ~/.local/bin/ddl-to-object"
    echo "  Config file: ~/.dto/config.json"
    echo "  Template dir: ~/.dto/template/"
    echo
    echo "Usage:"
    echo "  # Reload shell environment"
    echo "  source ~/.bashrc"
    echo
    echo "  # Check version"
    echo "  ddl-to-object version"
    echo
    echo "  # Show help"
    echo "  ddl-to-object --help"
    echo
    echo "  # Generate Go struct"
    echo "  ddl-to-object go -f your_table.sql -t ./output/"
    echo
}

# Main function
main() {
    echo "========================================"
    echo "    ddl-to-object Installation Script"
    echo "========================================"
    echo
    
    check_dependencies
    create_directories
    build_binary
    copy_config
    copy_templates
    update_path
    verify_installation
    show_completion
}

# Run main function
main "$@"