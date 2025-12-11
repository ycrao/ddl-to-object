# Security Guidelines

## Web Server Security

The DDL to Object web server has been designed with security in mind. Here are the implemented security measures:

### 1. File System Protection

- **Whitelist-based file serving**: Only specific files (`index.html`, `app.js`, `style.css`) are served
- **Path traversal prevention**: Multiple layers of protection against `../` attacks
- **Absolute path validation**: Ensures served files are within the web directory
- **No directory listing**: Prevents browsing of server directories

### 2. Input Validation

- **Request size limits**: Maximum 1MB request body size
- **DDL content limits**: Maximum 100KB DDL content
- **Language validation**: Only allows supported languages (go, java, php, python)
- **Package name validation**: Prevents malicious package names
- **JSON schema validation**: Disallows unknown fields in requests

### 3. HTTP Security Headers

- `X-Content-Type-Options: nosniff` - Prevents MIME type sniffing
- `X-Frame-Options: DENY` - Prevents clickjacking attacks
- `X-XSS-Protection: 1; mode=block` - Enables XSS protection
- `Referrer-Policy: strict-origin-when-cross-origin` - Controls referrer information
- `Server: ddl-to-object-web` - Custom server header

### 4. CORS Configuration

- Restricted to `http://localhost:8080` in production
- Limited to specific HTTP methods (POST, OPTIONS)
- Controlled allowed headers

### 5. Request Logging

- All requests are logged with method, path, and client IP
- Helps with monitoring and debugging

## Deployment Security

### Production Recommendations

1. **Use HTTPS**: Always deploy with TLS/SSL certificates
2. **Reverse Proxy**: Use nginx or Apache as a reverse proxy
3. **Firewall**: Restrict access to necessary ports only
4. **Rate Limiting**: Implement proper rate limiting (Redis-based)
5. **Authentication**: Add authentication for sensitive deployments

### Environment Variables

- `PORT`: Server port (default: 8080)

### File Permissions

Ensure proper file permissions:

```bash
chmod 644 web/*.html web/*.js web/*.css
chmod 755 web/
chmod 600 config files (if any contain secrets)
```

## Security Checklist

- [ ] Files are served from a restricted whitelist
- [ ] Path traversal attacks are prevented
- [ ] Input validation is in place
- [ ] Security headers are set
- [ ] Request logging is enabled
- [ ] CORS is properly configured
- [ ] File permissions are correct
- [ ] HTTPS is enabled (production)
- [ ] Rate limiting is implemented (production)

## Reporting Security Issues

If you discover a security vulnerability, please:

1. **Do not** create a public GitHub issue
2. Email the maintainers directly
3. Provide detailed information about the vulnerability
4. Allow time for the issue to be addressed before public disclosure

## Security Updates

- Regularly update Go and dependencies
- Monitor security advisories
- Review and update security configurations
- Test security measures regularly

## Known Limitations

1. **Rate Limiting**: Current implementation is basic - use Redis/Memcached for production
2. **Authentication**: No built-in authentication - add if needed for your use case
3. **Input Sanitization**: Basic validation - consider additional sanitization for specific use cases

## Security Testing

Run security tests:

```bash
# Test path traversal
curl "http://localhost:8080/../../../etc/passwd"

# Test large payload
curl -X POST -H "Content-Type: application/json" \
  -d '{"ddl":"'$(head -c 2000000 /dev/zero | tr '\0' 'a')'","language":"go"}' \
  http://localhost:8080/api/convert

# Test invalid language
curl -X POST -H "Content-Type: application/json" \
  -d '{"ddl":"CREATE TABLE test (id INT);","language":"malicious"}' \
  http://localhost:8080/api/convert
```

All these tests should be properly handled and rejected by the server.
