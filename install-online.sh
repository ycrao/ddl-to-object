#!/bin/bash

# DDL to Object Online Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/ycrao/ddl-to-object/main/install-online.sh | bash

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# 配置
REPO="ycrao/ddl-to-object"
API_URL="https://api.github.com/repos/${REPO}/releases/latest"
INSTALL_DIR="$HOME/.local/bin"
CONFIG_DIR="$HOME/.dto"

# 打印函数
print_banner() {
    echo -e "${CYAN}"
    cat << 'EOF'
    ____  ____  __       __           ____  __     _           __ 
   / __ \/ __ \/ /      / /_____     / __ \/ /_   (_)__  _____/ /_
  / / / / / / / /      / __/ __ \   / / / / __ \ / / _ \/ ___/ __/
 / /_/ / /_/ / /___   / /_/ /_/ /  / /_/ / /_/ // /  __/ /__/ /_  
/_____/_____/_____/   \__/\____/   \____/_.___//_/\___/\___/\__/  
                                                                  
EOF
    echo -e "${NC}"
    echo -e "${BLUE}DDL to Object - Online Installer${NC}"
    echo -e "${BLUE}=================================${NC}"
    echo
}

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

print_step() {
    echo -e "${PURPLE}[STEP]${NC} $1"
}

# 检测系统
detect_system() {
    print_step "Detecting system..."
    
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    case "$OS" in
        linux*)
            PLATFORM="linux"
            ;;
        darwin*)
            PLATFORM="mac"
            # 检测 ARM64
            if [ "$ARCH" = "arm64" ]; then
                PLATFORM="mac-arm64"
            fi
            ;;
        *)
            print_error "Unsupported operating system: $OS"
            print_info "Supported systems: Linux, macOS"
            exit 1
            ;;
    esac
    
    print_info "Detected system: $OS ($ARCH)"
    print_info "Platform: $PLATFORM"
}

# 检查依赖
check_dependencies() {
    print_step "Checking dependencies..."
    
    local missing_deps=()
    
    # 检查必需的命令
    for cmd in curl tar; do
        if ! command -v "$cmd" &> /dev/null; then
            missing_deps+=("$cmd")
        fi
    done
    
    if [ ${#missing_deps[@]} -ne 0 ]; then
        print_error "Missing required dependencies: ${missing_deps[*]}"
        print_info "Please install them and try again:"
        
        case "$OS" in
            linux*)
                print_info "  Ubuntu/Debian: sudo apt-get install ${missing_deps[*]}"
                print_info "  CentOS/RHEL: sudo yum install ${missing_deps[*]}"
                ;;
            darwin*)
                print_info "  macOS: brew install ${missing_deps[*]}"
                ;;
        esac
        exit 1
    fi
    
    print_success "All dependencies are available"
}

# 获取最新版本信息
get_latest_release() {
    print_step "Fetching latest release information..."
    
    # 获取最新版本信息
    local release_info
    release_info=$(curl -fsSL "$API_URL" 2>/dev/null) || {
        print_error "Failed to fetch release information from GitHub API"
        print_info "Please check your internet connection and try again"
        exit 1
    }
    
    # 解析版本号
    VERSION=$(echo "$release_info" | grep '"tag_name"' | sed -E 's/.*"tag_name": "([^"]+)".*/\1/')
    
    if [ -z "$VERSION" ]; then
        print_error "Failed to parse version information"
        exit 1
    fi
    
    print_info "Latest version: $VERSION"
    
    # 构建下载URL
    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/ddl-to-object-${PLATFORM}-${VERSION}.tar.gz"
    
    print_info "Download URL: $DOWNLOAD_URL"
}

# 创建安装目录
create_directories() {
    print_step "Creating installation directories..."
    
    # 创建安装目录
    mkdir -p "$INSTALL_DIR"
    mkdir -p "$CONFIG_DIR"
    mkdir -p "$CONFIG_DIR/template"
    
    print_success "Directories created"
    print_info "Install directory: $INSTALL_DIR"
    print_info "Config directory: $CONFIG_DIR"
}

# 下载和安装
download_and_install() {
    print_step "Downloading and installing..."
    
    # 创建临时目录
    local temp_dir
    temp_dir=$(mktemp -d)
    
    # 确保清理临时目录
    trap "rm -rf '$temp_dir'" EXIT
    
    print_info "Downloading from: $DOWNLOAD_URL"
    
    # 下载文件
    local archive_file="$temp_dir/ddl-to-object.tar.gz"
    if ! curl -fsSL -o "$archive_file" "$DOWNLOAD_URL"; then
        print_error "Failed to download release archive"
        print_info "Please check the URL and try again"
        exit 1
    fi
    
    print_success "Download completed"
    
    # 解压文件
    print_info "Extracting archive..."
    tar -xzf "$archive_file" -C "$temp_dir"
    
    # 查找解压后的目录
    local extracted_dir="$temp_dir/$PLATFORM"
    if [ ! -d "$extracted_dir" ]; then
        print_error "Extracted directory not found: $extracted_dir"
        exit 1
    fi
    
    # 安装二进制文件
    if [ -f "$extracted_dir/ddl-to-object" ]; then
        cp "$extracted_dir/ddl-to-object" "$INSTALL_DIR/"
        chmod +x "$INSTALL_DIR/ddl-to-object"
        print_success "Binary installed: $INSTALL_DIR/ddl-to-object"
    else
        print_error "Binary file not found in archive"
        exit 1
    fi
    
    # 安装模板文件
    if [ -d "$extracted_dir/template" ]; then
        cp -r "$extracted_dir/template"/* "$CONFIG_DIR/template/"
        print_success "Templates installed: $CONFIG_DIR/template/"
    else
        print_warning "Template directory not found in archive"
    fi
    
    # 安装配置文件
    if [ -f "$extracted_dir/config.json" ]; then
        if [ ! -f "$CONFIG_DIR/config.json" ]; then
            cp "$extracted_dir/config.json" "$CONFIG_DIR/"
            print_success "Config file installed: $CONFIG_DIR/config.json"
        else
            print_info "Config file already exists, skipping"
        fi
    else
        print_warning "Config file not found in archive"
    fi
}

# 更新PATH
update_path() {
    print_step "Updating PATH..."
    
    # 检查PATH是否已包含安装目录
    if [[ ":$PATH:" == *":$INSTALL_DIR:"* ]]; then
        print_info "PATH already contains $INSTALL_DIR"
        return
    fi
    
    # 添加到shell配置文件
    local shell_configs=("$HOME/.bashrc" "$HOME/.zshrc" "$HOME/.profile")
    local updated=false
    
    for config_file in "${shell_configs[@]}"; do
        if [ -f "$config_file" ]; then
            # 检查是否已经添加过
            if ! grep -q "$INSTALL_DIR" "$config_file"; then
                echo "" >> "$config_file"
                echo "# Added by ddl-to-object installer" >> "$config_file"
                echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$config_file"
                print_success "Updated $config_file"
                updated=true
            fi
        fi
    done
    
    if [ "$updated" = true ]; then
        print_info "Please reload your shell or run: source ~/.bashrc or source ~/.zshrc or source ~/.profile"
    else
        print_warning "No shell config files found to update"
        print_info "Please manually add $INSTALL_DIR to your PATH"
    fi
}

# 验证安装
verify_installation() {
    print_step "Verifying installation..."
    
    # 检查二进制文件
    if [ -x "$INSTALL_DIR/ddl-to-object" ]; then
        print_success "Binary is executable"
        
        # 尝试运行版本命令
        if "$INSTALL_DIR/ddl-to-object" version &>/dev/null; then
            local installed_version
            installed_version=$("$INSTALL_DIR/ddl-to-object" version 2>/dev/null | head -1 || echo "unknown")
            print_success "Installation verified: $installed_version"
        else
            print_warning "Binary exists but version check failed"
        fi
    else
        print_error "Binary is not executable"
        exit 1
    fi
    
    # 检查配置目录
    if [ -d "$CONFIG_DIR" ]; then
        print_success "Config directory exists"
    else
        print_warning "Config directory not found"
    fi
}

# 显示完成信息
show_completion() {
    echo
    print_success "Installation completed successfully!"
    echo
    echo -e "${CYAN}Installation Summary:${NC}"
    echo "  Version: $VERSION"
    echo "  Binary: $INSTALL_DIR/ddl-to-object"
    echo "  Config: $CONFIG_DIR/"
    echo "  Templates: $CONFIG_DIR/template/"
    echo
    echo -e "${CYAN}Usage:${NC}"
    echo "  # Show version"
    echo "  ddl-to-object version"
    echo
    echo "  # Show help"
    echo "  ddl-to-object --help"
    echo
    echo "  # Generate Go struct"
    echo "  ddl-to-object go -f your_table.sql -t ./output/"
    echo
    echo -e "${YELLOW}Note: You may need to reload your shell or run 'source ~/.bashrc' or 'source ~/.zshrc' or 'source ~/.profile' ${NC}"
    echo
}

# 主函数
main() {
    print_banner
    
    # 检查是否以root身份运行
    if [ "$EUID" -eq 0 ]; then
        print_warning "Running as root is not recommended"
        print_info "This installer will install to user directories"
    fi
    
    detect_system
    check_dependencies
    get_latest_release
    create_directories
    download_and_install
    update_path
    verify_installation
    show_completion
}

# 错误处理
handle_error() {
    print_error "Installation failed!"
    print_info "Please check the error messages above and try again"
    exit 1
}

# 设置错误处理
trap handle_error ERR

# 运行主函数
main "$@"