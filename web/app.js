// DDL to Object Web App JavaScript

// 国际化文本
const i18nTexts = {
    zh: {
        title: 'DDL to Object',
        subtitle: '将 MySQL DDL 转换为各种编程语言的对象结构',
        'ddl-label': 'MySQL DDL 语句:',
        'ddl-placeholder': '请粘贴您的 MySQL CREATE TABLE 语句...\n\n例如:\nCREATE TABLE `users` (\n  `id` bigint unsigned NOT NULL AUTO_INCREMENT,\n  `name` varchar(255) NOT NULL,\n  `email` varchar(255) NOT NULL,\n  `created_at` timestamp NULL DEFAULT NULL,\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;',
        'load-example': '加载示例 DDL',
        'select-language': '选择目标语言:',
        'go-desc': '生成 struct 结构体',
        'java-desc': '生成 Entity 类',
        'php-desc': '生成 Model 类',
        'python-desc': '生成 Class 类',
        'generate-btn': '🚀 生成代码',
        'generating': '正在生成代码...',
        'output-label': '生成的代码:',
        'copy-btn': '📋 复制代码',
        'copy-success': '✅ 已复制!',
        'error-empty-ddl': '请输入 DDL 语句',
        'error-generation': '生成代码时发生错误: ',
        'error-copy': '复制失败，请手动复制',
        'page-title': 'DDL to Object - 在线转换工具',
        'syntax-highlight': '语法高亮'
    },
    en: {
        title: 'DDL to Object',
        subtitle: 'Convert MySQL DDL to object structures in various programming languages',
        'ddl-label': 'MySQL DDL Statement:',
        'ddl-placeholder': 'Please paste your MySQL CREATE TABLE statement...\n\nExample:\nCREATE TABLE `users` (\n  `id` bigint unsigned NOT NULL AUTO_INCREMENT,\n  `name` varchar(255) NOT NULL,\n  `email` varchar(255) NOT NULL,\n  `created_at` timestamp NULL DEFAULT NULL,\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;',
        'load-example': 'Load Example DDL',
        'select-language': 'Select Target Language:',
        'go-desc': 'Generate struct',
        'java-desc': 'Generate Entity class',
        'php-desc': 'Generate Model class',
        'python-desc': 'Generate Class',
        'generate-btn': '🚀 Generate Code',
        'generating': 'Generating code...',
        'output-label': 'Generated Code:',
        'copy-btn': '📋 Copy Code',
        'copy-success': '✅ Copied!',
        'error-empty-ddl': 'Please enter DDL statement',
        'error-generation': 'Error occurred while generating code: ',
        'error-copy': 'Copy failed, please copy manually',
        'page-title': 'DDL to Object - Online Conversion Tool',
        'syntax-highlight': 'Syntax Highlight'
    }
};

// 语法高亮状态
let ddlHighlightEnabled = false;
let outputHighlightEnabled = true;
let currentOutputLanguage = 'go';

// 当前语言
let currentLanguage = 'zh';

// 示例 DDL
const exampleDDL = {
    zh: `CREATE TABLE \`users\` (
  \`id\` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  \`username\` varchar(50) NOT NULL COMMENT '用户名',
  \`email\` varchar(100) NOT NULL COMMENT '邮箱地址',
  \`password\` varchar(255) NOT NULL COMMENT '密码',
  \`phone\` varchar(20) DEFAULT NULL COMMENT '手机号',
  \`avatar\` varchar(255) DEFAULT NULL COMMENT '头像URL',
  \`status\` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 1-正常, 0-禁用',
  \`created_at\` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  \`updated_at\` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (\`id\`),
  UNIQUE KEY \`uk_username\` (\`username\`),
  UNIQUE KEY \`uk_email\` (\`email\`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';`,
    en: `CREATE TABLE \`users\` (
  \`id\` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  \`username\` varchar(50) NOT NULL COMMENT 'Username',
  \`email\` varchar(100) NOT NULL COMMENT 'Email address',
  \`password\` varchar(255) NOT NULL COMMENT 'Password',
  \`phone\` varchar(20) DEFAULT NULL COMMENT 'Phone number',
  \`avatar\` varchar(255) DEFAULT NULL COMMENT 'Avatar URL',
  \`status\` tinyint NOT NULL DEFAULT '1' COMMENT 'Status: 1-Active, 0-Disabled',
  \`created_at\` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  \`updated_at\` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (\`id\`),
  UNIQUE KEY \`uk_username\` (\`username\`),
  UNIQUE KEY \`uk_email\` (\`email\`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Users table';`
};

// 切换语言
function switchLanguage(lang) {
    currentLanguage = lang;
    
    // 更新语言按钮状态
    document.querySelectorAll('.lang-btn').forEach(btn => {
        btn.classList.remove('active');
        if (btn.getAttribute('data-lang') === lang) {
            btn.classList.add('active');
        }
    });
    
    // 更新页面文本
    updatePageTexts();
    
    // 更新HTML lang属性
    document.documentElement.lang = lang === 'zh' ? 'zh-CN' : 'en';
}

// 更新页面文本
function updatePageTexts() {
    const texts = i18nTexts[currentLanguage];
    
    // 更新所有带有data-i18n属性的元素
    document.querySelectorAll('[data-i18n]').forEach(element => {
        const key = element.getAttribute('data-i18n');
        if (texts[key]) {
            if (element.tagName === 'TITLE') {
                element.textContent = texts[key];
            } else {
                element.textContent = texts[key];
            }
        }
    });
    
    // 更新placeholder
    document.querySelectorAll('[data-i18n-placeholder]').forEach(element => {
        const key = element.getAttribute('data-i18n-placeholder');
        if (texts[key]) {
            element.placeholder = texts[key];
        }
    });
}

// 加载示例 DDL
function loadExample() {
    document.getElementById('ddl-input').value = exampleDDL[currentLanguage];
    updateDDLHighlight();
}

// 切换DDL语法高亮
function toggleDDLHighlight() {
    ddlHighlightEnabled = document.getElementById('ddl-highlight-toggle').checked;
    const textarea = document.getElementById('ddl-input');
    const display = document.getElementById('ddl-display');
    
    if (ddlHighlightEnabled) {
        textarea.style.display = 'none';
        display.classList.add('active');
        updateDDLHighlight();
    } else {
        textarea.style.display = 'block';
        display.classList.remove('active');
    }
}

// 更新DDL语法高亮
function updateDDLHighlight() {
    if (!ddlHighlightEnabled || !window.hljs) return;
    
    const code = document.getElementById('ddl-input').value;
    const display = document.getElementById('ddl-display');
    
    if (code.trim()) {
        const highlighted = hljs.highlight(code, { language: 'sql' });
        display.innerHTML = highlighted.value;
    } else {
        display.innerHTML = '';
    }
}

// 切换输出语法高亮
function toggleOutputHighlight() {
    outputHighlightEnabled = document.getElementById('output-highlight-toggle').checked;
    const textarea = document.getElementById('output-code');
    const display = document.getElementById('output-display');
    
    if (outputHighlightEnabled) {
        textarea.style.display = 'none';
        display.style.display = 'block';
        updateOutputHighlight();
    } else {
        textarea.style.display = 'block';
        display.style.display = 'none';
    }
}

// 更新输出语法高亮
function updateOutputHighlight() {
    if (!outputHighlightEnabled || !window.hljs) return;
    
    const code = document.getElementById('output-code').value;
    const display = document.getElementById('output-display');
    
    if (code.trim()) {
        // 映射语言名称
        const languageMap = {
            'go': 'go',
            'java': 'java',
            'php': 'php',
            'python': 'python'
        };
        
        const language = languageMap[currentOutputLanguage] || currentOutputLanguage;
        const highlighted = hljs.highlight(code, { language: language });
        display.innerHTML = highlighted.value;
    } else {
        display.innerHTML = '';
    }
}



// 生成代码
async function generateCode() {
    const ddlInput = document.getElementById('ddl-input').value.trim();
    const selectedLanguage = document.querySelector('input[name="language"]:checked').value;
    
    if (!ddlInput) {
        showError(i18nTexts[currentLanguage]['error-empty-ddl']);
        return;
    }
    
    // 更新当前输出语言
    currentOutputLanguage = selectedLanguage;
    
    // 显示加载状态
    showLoading(true);
    hideError();
    hideOutput();
    
    try {
        // 调用后端 API 或者使用 WebAssembly
        const result = await callDDLToObjectAPI(ddlInput, selectedLanguage);
        
        if (result.success) {
            showOutput(result.code);
        } else {
            showError(result.error || '生成代码失败');
        }
    } catch (error) {
        console.error('Error:', error);
        showError(i18nTexts[currentLanguage]['error-generation'] + error.message);
    } finally {
        showLoading(false);
    }
}

// 调用 DDL to Object API
async function callDDLToObjectAPI(ddl, language) {
    try {
        const response = await fetch('/api/convert', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                ddl: ddl,
                language: language
            })
        });
        
        const result = await response.json();
        return result;
    } catch (error) {
        // 如果 API 不可用，回退到模拟转换
        console.warn('API not available, using simulation:', error);
        return new Promise((resolve) => {
            setTimeout(() => {
                try {
                    const result = simulateConversion(ddl, language);
                    resolve({ success: true, code: result });
                } catch (error) {
                    resolve({ success: false, error: error.message });
                }
            }, 500);
        });
    }
}

// 模拟转换 (简化版本)
function simulateConversion(ddl, language) {
    // 简单的 DDL 解析 (这只是一个演示，实际应该使用后端)
    const tableMatch = ddl.match(/CREATE\s+TABLE\s+`?(\w+)`?\s*\(/i);
    if (!tableMatch) {
        throw new Error('无法解析表名');
    }
    
    const tableName = tableMatch[1];
    const className = toPascalCase(tableName);
    
    // 提取字段
    const fields = extractFields(ddl);
    
    switch (language) {
        case 'go':
            return generateGoStruct(className, fields);
        case 'java':
            return generateJavaClass(className, fields);
        case 'php':
            return generatePHPClass(className, fields);
        case 'python':
            return generatePythonClass(className, fields);
        default:
            throw new Error('不支持的语言');
    }
}

// 提取字段信息
function extractFields(ddl) {
    const fields = [];
    const fieldRegex = /`(\w+)`\s+(\w+)(?:\([\d,]+\))?\s*([^,\n]*)/gi;
    let match;
    
    while ((match = fieldRegex.exec(ddl)) !== null) {
        const [, name, type, attributes] = match;
        
        // 跳过主键和索引定义
        if (name.toLowerCase() === 'primary' || name.toLowerCase() === 'key' || 
            name.toLowerCase() === 'unique' || name.toLowerCase() === 'index') {
            continue;
        }
        
        const comment = extractComment(attributes);
        const nullable = !attributes.toLowerCase().includes('not null');
        
        fields.push({
            name: name,
            type: type.toLowerCase(),
            comment: comment,
            nullable: nullable,
            camelName: toCamelCase(name),
            pascalName: toPascalCase(name)
        });
    }
    
    return fields;
}

// 提取注释
function extractComment(attributes) {
    const commentMatch = attributes.match(/COMMENT\s+['"](.*?)['"]/i);
    return commentMatch ? commentMatch[1] : '';
}

// 转换为驼峰命名
function toCamelCase(str) {
    return str.replace(/_([a-z])/g, (match, letter) => letter.toUpperCase());
}

// 转换为帕斯卡命名
function toPascalCase(str) {
    return str.charAt(0).toUpperCase() + toCamelCase(str).slice(1);
}

// 生成 Go 结构体
function generateGoStruct(className, fields) {
    let code = `package models\n\nimport (\n\t"time"\n)\n\n`;
    code += `// ${className} 结构体\n`;
    code += `type ${className} struct {\n`;
    
    fields.forEach(field => {
        const goType = mapToGoType(field.type, field.nullable);
        const jsonTag = field.name;
        const dbTag = field.name;
        
        code += `\t${field.pascalName} ${goType} \`json:"${jsonTag}" db:"${dbTag}"\``;
        if (field.comment) {
            code += ` // ${field.comment}`;
        }
        code += '\n';
    });
    
    code += '}\n';
    return code;
}

// 生成 Java 类
function generateJavaClass(className, fields) {
    let code = `package com.example.entity;\n\n`;
    code += `import lombok.Data;\n`;
    code += `import java.time.LocalDateTime;\n\n`;
    code += `/**\n * ${className} 实体类\n */\n`;
    code += `@Data\n`;
    code += `public class ${className} {\n\n`;
    
    fields.forEach(field => {
        const javaType = mapToJavaType(field.type);
        if (field.comment) {
            code += `    /** ${field.comment} */\n`;
        }
        code += `    private ${javaType} ${field.camelName};\n\n`;
    });
    
    code += '}\n';
    return code;
}

// 生成 PHP 类
function generatePHPClass(className, fields) {
    let code = `<?php\n\nnamespace App\\Models;\n\n`;
    code += `/**\n * ${className} 模型类\n */\n`;
    code += `class ${className}\n{\n`;
    
    fields.forEach(field => {
        const phpType = mapToPHPType(field.type);
        if (field.comment) {
            code += `    /** @var ${phpType} ${field.comment} */\n`;
        }
        code += `    public $${field.camelName};\n\n`;
    });
    
    code += '}\n';
    return code;
}

// 生成 Python 类
function generatePythonClass(className, fields) {
    let code = `from typing import Optional\nfrom datetime import datetime\n\n`;
    code += `class ${className}:\n`;
    code += `    """${className} 数据类"""\n\n`;
    code += `    def __init__(self):\n`;
    
    fields.forEach(field => {
        const pythonType = mapToPythonType(field.type, field.nullable);
        code += `        self.${field.name}: ${pythonType} = None`;
        if (field.comment) {
            code += `  # ${field.comment}`;
        }
        code += '\n';
    });
    
    return code;
}

// 类型映射函数
function mapToGoType(mysqlType, nullable) {
    const typeMap = {
        'bigint': 'int64',
        'int': 'int32',
        'tinyint': 'int8',
        'varchar': 'string',
        'text': 'string',
        'timestamp': 'time.Time',
        'datetime': 'time.Time',
        'date': 'time.Time'
    };
    
    let goType = typeMap[mysqlType] || 'interface{}';
    
    if (nullable && goType !== 'interface{}') {
        if (goType === 'string') {
            goType = 'sql.NullString';
        } else if (goType.includes('int')) {
            goType = 'sql.NullInt64';
        } else if (goType === 'time.Time') {
            goType = 'sql.NullTime';
        }
    }
    
    return goType;
}

function mapToJavaType(mysqlType) {
    const typeMap = {
        'bigint': 'Long',
        'int': 'Integer',
        'tinyint': 'Integer',
        'varchar': 'String',
        'text': 'String',
        'timestamp': 'LocalDateTime',
        'datetime': 'LocalDateTime',
        'date': 'LocalDateTime'
    };
    
    return typeMap[mysqlType] || 'Object';
}

function mapToPHPType(mysqlType) {
    const typeMap = {
        'bigint': 'int',
        'int': 'int',
        'tinyint': 'int',
        'varchar': 'string',
        'text': 'string',
        'timestamp': 'string',
        'datetime': 'string',
        'date': 'string'
    };
    
    return typeMap[mysqlType] || 'mixed';
}

function mapToPythonType(mysqlType, nullable) {
    const typeMap = {
        'bigint': 'int',
        'int': 'int',
        'tinyint': 'int',
        'varchar': 'str',
        'text': 'str',
        'timestamp': 'datetime',
        'datetime': 'datetime',
        'date': 'datetime'
    };
    
    let pythonType = typeMap[mysqlType] || 'Any';
    
    if (nullable) {
        pythonType = `Optional[${pythonType}]`;
    }
    
    return pythonType;
}

// UI 控制函数
function showLoading(show) {
    document.getElementById('loading').style.display = show ? 'block' : 'none';
    document.querySelector('.generate-btn').disabled = show;
}

function showOutput(code) {
    document.getElementById('output-code').value = code;
    document.getElementById('output-section').style.display = 'block';
    
    // 更新语法高亮
    updateOutputHighlight();
}

function hideOutput() {
    document.getElementById('output-section').style.display = 'none';
}

function showError(message) {
    const errorDiv = document.getElementById('error-message');
    errorDiv.textContent = message;
    errorDiv.style.display = 'block';
}

function hideError() {
    document.getElementById('error-message').style.display = 'none';
}

// 复制到剪贴板
function copyToClipboard() {
    const outputCode = document.getElementById('output-code');
    
    // 如果使用现代API
    if (navigator.clipboard && window.isSecureContext) {
        navigator.clipboard.writeText(outputCode.value).then(() => {
            showCopySuccess();
        }).catch(err => {
            console.error('Copy failed:', err);
            fallbackCopy();
        });
    } else {
        fallbackCopy();
    }
}

// 备用复制方法
function fallbackCopy() {
    const outputCode = document.getElementById('output-code');
    outputCode.select();
    outputCode.setSelectionRange(0, 99999); // 移动端兼容
    
    try {
        document.execCommand('copy');
        showCopySuccess();
    } catch (err) {
        console.error('Copy failed:', err);
        alert(i18nTexts[currentLanguage]['error-copy']);
    }
}

// 显示复制成功
function showCopySuccess() {
    const copyBtn = document.querySelector('.copy-btn');
    const originalText = copyBtn.textContent;
    copyBtn.textContent = i18nTexts[currentLanguage]['copy-success'];
    copyBtn.style.background = '#28a745';
    
    setTimeout(() => {
        copyBtn.textContent = originalText;
        copyBtn.style.background = '#28a745';
    }, 2000);
}

// 页面加载完成后的初始化
document.addEventListener('DOMContentLoaded', function() {
    // 初始化页面文本
    updatePageTexts();
    
    // 检测浏览器语言
    const browserLang = navigator.language || navigator.userLanguage;
    if (browserLang.startsWith('en')) {
        switchLanguage('en');
    }
    
    // 等待highlight.js加载完成
    if (window.hljs) {
        initializeHighlighting();
    } else {
        // 如果highlight.js还没加载完成，等待一下
        setTimeout(() => {
            if (window.hljs) {
                initializeHighlighting();
            }
        }, 100);
    }
    
    console.log('DDL to Object Web App loaded with highlight.js');
});

// 初始化语法高亮
function initializeHighlighting() {
    // 设置默认状态 - 输出默认启用语法高亮
    document.getElementById('output-highlight-toggle').checked = true;
    outputHighlightEnabled = true;
    
    // 确保输出区域默认显示高亮版本
    const textarea = document.getElementById('output-code');
    const display = document.getElementById('output-display');
    textarea.style.display = 'none';
    display.style.display = 'block';
    
    // 配置highlight.js
    hljs.configure({
        ignoreUnescapedHTML: true,
        throwUnescapedHTML: false
    });
    
    console.log('Highlight.js initialized with default syntax highlighting');
}