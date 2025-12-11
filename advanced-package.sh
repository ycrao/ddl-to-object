#!/bin/bash

# 高级打包脚本 - 支持多种选项

set -e

# 默认配置
DEFAULT_VERSION=$(date +%Y%m%d)
RELEASE_DIR="release"
OUTPUT_DIR="dist"

# 帮助信息
show_help() {
    cat << EOF
Usage: $0 [OPTIONS]

Options:
    -v, --version VERSION    Set version (default: $DEFAULT_VERSION)
    -o, --output DIR         Output directory (default: $OUTPUT_DIR)
    -r, --release DIR        Release directory (default: $RELEASE_DIR)
    -c, --clean             Clean output directory before packaging
    -h, --help              Show this help message

Examples:
    $0                      # Package with default settings
    $0 -v 1.2.3            # Package with version 1.2.3
    $0 -v 1.2.3 -c         # Clean and package with version 1.2.3
    $0 --output packages    # Output to 'packages' directory

Platforms:
    - Linux: .tar.gz
    - Mac: .tar.gz
    - Mac ARM64: .tar.gz
    - Windows: .zip
EOF
}

# 解析命令行参数
parse_args() {
    VERSION="$DEFAULT_VERSION"
    CLEAN=false
    
    while [[ $# -gt 0 ]]; do
        case $1 in
            -v|--version)
                VERSION="$2"
                shift 2
                ;;
            -o|--output)
                OUTPUT_DIR="$2"
                shift 2
                ;;
            -r|--release)
                RELEASE_DIR="$2"
                shift 2
                ;;
            -c|--clean)
                CLEAN=true
                shift
                ;;
            -h|--help)
                show_help
                exit 0
                ;;
            *)
                echo "Unknown option: $1"
                show_help
                exit 1
                ;;
        esac
    done
}

# 创建输出目录
setup_output_dir() {
    if [ "$CLEAN" = true ] && [ -d "$OUTPUT_DIR" ]; then
        echo "🧹 Cleaning output directory: $OUTPUT_DIR"
        rm -rf "$OUTPUT_DIR"
    fi
    
    mkdir -p "$OUTPUT_DIR"
    echo "📁 Output directory: $OUTPUT_DIR"
}

# 打包函数
package_platform() {
    local platform=$1
    local archive_type=$2
    local source_dir="$RELEASE_DIR/$platform"
    
    if [ ! -d "$source_dir" ]; then
        echo "⚠️  Platform directory not found: $source_dir"
        return 1
    fi
    
    local filename="ddl-to-object-$platform-$VERSION"
    
    echo "📦 Packaging $platform..."
    
    case $archive_type in
        "tar.gz")
            tar -czf "$OUTPUT_DIR/$filename.tar.gz" -C "$RELEASE_DIR" "$platform/"
            ;;
        "zip")
            (cd "$RELEASE_DIR" && zip -r "../$OUTPUT_DIR/$filename.zip" "$platform/")
            ;;
        *)
            echo "❌ Unknown archive type: $archive_type"
            return 1
            ;;
    esac
    
    echo "✅ Created: $OUTPUT_DIR/$filename.$archive_type"
}

# 生成校验和
generate_checksums() {
    echo "🔐 Generating checksums..."
    
    cd "$OUTPUT_DIR"
    
    # 生成SHA256校验和
    if command -v sha256sum &> /dev/null; then
        sha256sum *.tar.gz *.zip > checksums.sha256 2>/dev/null || true
    elif command -v shasum &> /dev/null; then
        shasum -a 256 *.tar.gz *.zip > checksums.sha256 2>/dev/null || true
    fi
    
    # 生成MD5校验和
    if command -v md5sum &> /dev/null; then
        md5sum *.tar.gz *.zip > checksums.md5 2>/dev/null || true
    elif command -v md5 &> /dev/null; then
        md5 *.tar.gz *.zip > checksums.md5 2>/dev/null || true
    fi
    
    cd - > /dev/null
    
    if [ -f "$OUTPUT_DIR/checksums.sha256" ]; then
        echo "✅ Created: $OUTPUT_DIR/checksums.sha256"
    fi
    
    if [ -f "$OUTPUT_DIR/checksums.md5" ]; then
        echo "✅ Created: $OUTPUT_DIR/checksums.md5"
    fi
}

# 显示结果
show_results() {
    echo
    echo "📊 Package Summary:"
    echo "=================="
    echo "Version: $VERSION"
    echo "Output: $OUTPUT_DIR"
    echo
    
    if [ -d "$OUTPUT_DIR" ]; then
        echo "Generated files:"
        ls -lh "$OUTPUT_DIR"/ | grep -E '\.(tar\.gz|zip|sha256|md5)$' || echo "No packages found"
        
        echo
        echo "Total size:"
        du -sh "$OUTPUT_DIR" 2>/dev/null || echo "Unable to calculate size"
    fi
}

# 主函数
main() {
    echo "🚀 DDL to Object Advanced Packager"
    echo "=================================="
    
    parse_args "$@"
    
    echo "Version: $VERSION"
    echo "Release directory: $RELEASE_DIR"
    
    # 检查release目录
    if [ ! -d "$RELEASE_DIR" ]; then
        echo "❌ Release directory not found: $RELEASE_DIR"
        echo "Run 'make build-all' first to create release files"
        exit 1
    fi
    
    setup_output_dir
    
    echo
    echo "📦 Starting packaging process..."
    
    # 打包各平台
    package_platform "linux" "tar.gz"
    package_platform "mac" "tar.gz"
    package_platform "mac-arm64" "tar.gz"
    package_platform "win" "zip"
    
    # 生成校验和
    generate_checksums
    
    show_results
    
    echo
    echo "🎉 Packaging completed successfully!"
}

# 运行主函数
main "$@"