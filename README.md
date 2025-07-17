# 🌑 Dark Web Website (Vanity Onion Service)

A minimalist and secure `.onion` website served over the Tor network. Built with Go, containerized using Docker, and powered by a vanity Tor hidden service.

---

## 📁 Project Structure

```
.
├── ansible/                 # Ansible playbook for server setup
│   └── playbook.yml
├── app.go                   # Main app entrypoint
├── config/                  # Docker and Tor configuration
│   ├── compose.yml          # Docker Compose file
│   ├── config.go            # App config
│   ├── Dockerfile           # Dockerfile for the Go app
│   └── tor/
│       ├── Dockerfile       # Dockerfile for Tor hidden service
│       ├── torrc            # Tor config
│       └── vanity/          # Vanity onion keys
│           ├── hostname
│           ├── hs_ed25519_public_key
│           └── hs_ed25519_secret_key
├── go.mod                   # Go module file
├── LICENSE
├── Makefile                 # CLI commands
├── README.md                # This file
└── server/
    ├── assets/              # Static assets
    │   └── tor.png
    ├── pages/               # HTML templates
    │   ├── about.html
    │   ├── error.html
    │   ├── index.html
    │   └── layout.html
    └── server.go            # HTTP handlers and routes
```

---

## ⚙️ Getting Started

### 🐳 Run with Docker

Run the app and Tor hidden service in detached mode:

```bash
make run
```

Stop everything and clean up:

```bash
make stop
```

This will:

* Build and start the Go app
* Start a Tor container exposing the site as a hidden service
* Output the `.onion` address from `config/tor/vanity/hostname`

To see the address:

```bash
cat config/tor/vanity/hostname
```

---

### 💻 Local Development

Run the Go app on port `8080` directly without Tor:

```bash
make dev
```

It will serve the site at: [http://localhost:8080](http://localhost:8080)

---

## 🧅 Vanity Onion Service

To use your own `.onion` address:

1. Generate a vanity address using [`mkp224o`](https://github.com/cathugger/mkp224o).
2. Place the keys inside: `config/tor/vanity/`
3. Tor will serve your site at that address.

Make sure your `torrc` has:

```conf
HiddenServiceDir /var/lib/tor/hidden_service/
HiddenServicePort 80 127.0.0.1:8080
```

---

## 🧪 Test with Tor Locally

If Tor is running on your system:

```bash
curl --socks5-hostname 127.0.0.1:9050 http://yourvanityaddress.onion
```

---

## 🛡️ Onion Location Header (Optional)

If you want clearnet visitors to see your `.onion` option in Brave or Tor Browser, add this header in your Go server:

```http
Onion-Location: http://yourvanityaddress.onion
```

---

## 🧰 Makefile Commands

| Command     | Description                         |
| ----------- | ----------------------------------- |
| `make run`  | Start the Tor + Go app stack        |
| `make stop` | Stop and remove all containers      |
| `make dev`  | Run the Go app locally on port 8080 |

---

## 📦 Deployment with Ansible

Use the provided `ansible/playbook.yml` to deploy to a VPS (Docker and SSH required). Example:

```bash
ansible-playbook ansible/playbook.yml -i your_vps_ip,
```

---

## 🔐 Security Notes

* Do **not** commit `hs_ed25519_secret_key` to public repos.
* Use firewall rules or reverse proxies as needed.
* Keep your base image and Go binary minimal and hardened.

---

## 📄 License

MIT © 2025 Yasasvi Gumma
Use it freely, modify it responsibly.
