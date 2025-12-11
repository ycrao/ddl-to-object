# DDL to Object Web Interface

A bilingual web interface for converting MySQL DDL statements to object structures in various programming languages.

## Features

- 🎨 **Syntax Highlighting**: Full syntax highlighting for SQL input and generated code using highlight.js
- � **Bhilingual Support**: Chinese (中文) and English
- 🚀 **Real-time Conversion**: Paste DDL, select language, generate code instantly
- 📋 **One-click Copy**: Copy generated code to clipboard
- 🎯 **Multi-language Support**: Go, Java, PHP, Python
- 📱 **Responsive Design**: Works on desktop and mobile devices
- 🔒 **Security**: Secure file serving and input validation

## Language Support

The interface automatically detects browser language and defaults to:

- Chinese for Chinese browsers
- English for other browsers

Users can manually switch between languages using the toggle buttons in the header.

## Supported Programming Languages

| Language | Output | Description |
|----------|--------|-------------|
| Go | struct | Generate Go struct with tags |
| Java | Entity class | Generate Java entity with Lombok annotations |
| PHP | Model class | Generate PHP model with namespace |
| Python | Class | Generate Python class with type hints |

## Usage

1. **Start the server**:

   ```bash
   cd web
   go run server.go
   ```

2. **Open in browser**: http://localhost:8080

3. **Use the interface**:
   - Paste your MySQL DDL statement
   - Toggle syntax highlighting for input (optional)
   - Select target programming language
   - Click "Generate Code" / "生成代码"
   - View syntax-highlighted output (enabled by default)
   - Copy the generated code

## API Usage

The web interface also provides a REST API:

```bash
curl -X POST http://localhost:8080/api/convert \
  -H "Content-Type: application/json" \
  -d '{
    "ddl": "CREATE TABLE users (id INT PRIMARY KEY, name VARCHAR(255));",
    "language": "go",
    "package": "models"
  }'
```

## Security Features

- Whitelist-based file serving
- Input validation and sanitization
- Request size limits
- Security headers
- Path traversal protection

## Internationalization

The interface uses a simple i18n system with:

- Language detection based on browser settings
- Manual language switching
- Localized error messages
- Localized example DDL statements

### Adding New Languages

To add a new language:

1. Add translations to `i18nTexts` object in `app.js`
2. Add example DDL in the new language
3. Add language button in HTML
4. Update the `switchLanguage` function

## Files

- `index.html` - Main HTML interface with i18n attributes
- `app.js` - JavaScript with i18n support and API calls
- `style.css` - Additional styles
- `server.go` - Secure Go web server
- `README.md` - This documentation

## Development

The interface is built with vanilla HTML, CSS, and JavaScript for simplicity and security. It uses highlight.js from CDN for syntax highlighting functionality.

## Browser Support

- Modern browsers with ES6+ support
- Chrome, Firefox, Safari, Edge
- Mobile browsers (iOS Safari, Chrome Mobile)

## Contributing

When contributing to the web interface:

1. Maintain bilingual support for all user-facing text
2. Test both language modes
3. Ensure responsive design works
4. Follow security best practices
5. Update documentation for new features