package main

import (
	"ddl-to-object/lib"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ConvertRequest 转换请求结构
type ConvertRequest struct {
	DDL      string `json:"ddl"`
	Language string `json:"language"`
	Package  string `json:"package,omitempty"`
}

// ConvertResponse 转换响应结构
type ConvertResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Error   string `json:"error,omitempty"`
}

// 处理转换请求
func handleConvert(w http.ResponseWriter, r *http.Request) {
	// 设置安全头
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	// 处理 OPTIONS 请求
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 限制请求体大小 (1MB)
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)

	var req ConvertRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // 不允许未知字段
	if err := decoder.Decode(&req); err != nil {
		sendErrorResponse(w, "Invalid JSON: "+err.Error())
		return
	}

	// 验证输入
	if strings.TrimSpace(req.DDL) == "" {
		sendErrorResponse(w, "DDL content is required")
		return
	}

	// 限制DDL长度
	if len(req.DDL) > 100000 { // 100KB
		sendErrorResponse(w, "DDL content too large")
		return
	}

	// 验证语言
	allowedLanguages := map[string]bool{
		"go":     true,
		"java":   true,
		"php":    true,
		"python": true,
	}
	
	if !allowedLanguages[req.Language] {
		sendErrorResponse(w, "Unsupported language")
		return
	}

	// 验证包名格式
	if req.Package != "" {
		if len(req.Package) > 200 {
			sendErrorResponse(w, "Package name too long")
			return
		}
		// 简单的包名验证
		if strings.Contains(req.Package, "..") || 
		   strings.Contains(req.Package, "/") ||
		   strings.Contains(req.Package, "\\") {
			sendErrorResponse(w, "Invalid package name")
			return
		}
	}

	// 解析 DDL
	result, err := lib.Parse(req.DDL)
	if err != nil {
		sendErrorResponse(w, "Failed to parse DDL: "+err.Error())
		return
	}

	// 设置包名
	switch req.Language {
	case "go":
		if req.Package != "" {
			packageArr := strings.Split(req.Package, ".")
			if len(packageArr) > 0 {
				result.GoPackageName = packageArr[len(packageArr)-1]
			}
		}
	case "java":
		if req.Package != "" {
			result.JavaPackageName = req.Package
		}
	case "php":
		if req.Package != "" {
			result.PhpNamespaceName = req.Package
		}
	}

	// 生成代码
	code, err := generateCode(result, req.Language)
	if err != nil {
		sendErrorResponse(w, "Failed to generate code: "+err.Error())
		return
	}

	// 返回成功响应
	response := ConvertResponse{
		Success: true,
		Code:    code,
	}

	json.NewEncoder(w).Encode(response)
}

// 生成代码
func generateCode(result lib.ParsedResult, language string) (string, error) {
	// 查找模板文件
	templatePath := fmt.Sprintf("../template/%s.template", language)
	
	// 如果本地模板不存在，尝试用户目录
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		homeDir, _ := os.UserHomeDir()
		templatePath = filepath.Join(homeDir, ".dto", "template", language+".template")
	}

	// 解析模板
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to load template: %w", err)
	}

	// 执行模板
	var buf strings.Builder
	if err := tpl.Execute(&buf, result); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// 发送错误响应
func sendErrorResponse(w http.ResponseWriter, message string) {
	response := ConvertResponse{
		Success: false,
		Error:   message,
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}

// 处理静态文件
func handleStatic(w http.ResponseWriter, r *http.Request) {
	// 获取请求路径
	requestPath := r.URL.Path
	
	// 根路径重定向到 index.html
	if requestPath == "/" {
		requestPath = "/index.html"
	}
	
	// 移除开头的斜杠
	requestPath = strings.TrimPrefix(requestPath, "/")
	
	// 严格的安全检查
	if strings.Contains(requestPath, "..") || 
	   strings.Contains(requestPath, "\\") ||
	   strings.HasPrefix(requestPath, "/") ||
	   strings.Contains(requestPath, "~") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	
	// 只允许特定的文件
	allowedFiles := map[string]bool{
		"index.html":  true,
		"app.js":      true,
		"style.css":   true,
		"favicon.ico": true, // 允许网站图标
	}
	
	if !allowedFiles[requestPath] {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	
	// 构建安全的文件路径
	safePath := filepath.Join(".", requestPath)
	
	// 再次验证路径是否在当前目录内
	absPath, err := filepath.Abs(safePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	currentDir, err := filepath.Abs(".")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	if !strings.HasPrefix(absPath, currentDir) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	
	// 检查文件是否存在
	if _, err := os.Stat(safePath); os.IsNotExist(err) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	
	// 设置内容类型
	if strings.HasSuffix(requestPath, ".html") {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	} else if strings.HasSuffix(requestPath, ".js") {
		w.Header().Set("Content-Type", "application/javascript")
	} else if strings.HasSuffix(requestPath, ".css") {
		w.Header().Set("Content-Type", "text/css")
	}
	
	// 设置安全头
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	
	// 提供文件
	http.ServeFile(w, r, safePath)
}

// 健康检查
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"service": "ddl-to-object-web",
	})
}

// 安全中间件
func securityMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 记录请求
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		
		// 设置通用安全头
		w.Header().Set("Server", "ddl-to-object-web")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		
		// 调用下一个处理器
		next(w, r)
	}
}

// 限制中间件
func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 简单的速率限制 - 在生产环境中应该使用更复杂的实现
		// 这里只是示例，实际应该使用 Redis 或内存存储来跟踪请求
		next(w, r)
	}
}

func main() {
	// 设置路由，添加安全中间件
	http.HandleFunc("/api/convert", securityMiddleware(rateLimitMiddleware(handleConvert)))
	http.HandleFunc("/health", securityMiddleware(handleHealth))
	http.HandleFunc("/", securityMiddleware(handleStatic))

	// 获取端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 DDL to Object Web Server starting on port %s\n", port)
	fmt.Printf("📱 Open http://localhost:%s in your browser\n", port)
	fmt.Printf("🔗 API endpoint: http://localhost:%s/api/convert\n", port)
	fmt.Printf("🔒 Security features enabled\n")

	// 启动服务器
	log.Fatal(http.ListenAndServe(":"+port, nil))
}