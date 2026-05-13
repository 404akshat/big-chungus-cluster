# Aether OS: Self-Hosted PaaS for Kubernetes

[![Language](https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Infrastructure](https://img.shields.io/badge/Orchestration-Kubernetes-326CE5?style=flat&logo=kubernetes)](https://kubernetes.io/)
[![Status](https://img.shields.io/badge/Status-In_Progress-orange)](https://404akshat.github.io/SkyNet/)

Aether OS is a vision for a self-hosted Platform-as-a-Service (PaaS) that removes the friction of Kubernetes management. The goal is simple: **you provide a Docker image, and Aether OS handles the rest.**

> [!CAUTION]
> **Active Development:** This project is currently in the conceptual and early-build phase. The codebase and live status are updated weekly at the link below.

---

## 💡 The Core Concept

Managing Kubernetes manifests, ingress controllers, and scaling policies is a full-time job. Aether OS aims to be the "Operating System" for your cluster, written entirely in **Go** for maximum performance and native K8s compatibility.

### The Pipeline:
1. **Push:** You provide a container image.
2. **Abstract:** Aether OS generates the necessary K8s CRDs, Services, and Deployments.
3. **Manage:** The system monitors health and scales automatically.
4. **Harden:** Integrated chaos engineering modules intentionally stress the app to ensure reliability.

---

## 🚀 Future Roadmap

* **Go-Based Controller:** A custom Kubernetes operator to manage application lifecycles.
* **Automated Observability:** Built-in metrics and logging without manual sidecar injection.
* **Resilience Testing:** Native chaos engineering tools to simulate pod failures and network latency.
* **Security Layer:** Automated mTLS and secret management out of the box.

---

## 🔧 Current Status

The project is currently being built out. Weekly progress reports and live updates can be tracked here:
* **🔗 Progress Dashboard:** [My LinkedIN](https://www.linkedin.com/in/akshat-sharma-devsecop/)

---

## 📄 License

MIT License

Copyright (c) 2026 404akshat

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---
**Maintained by [404akshat](https://github.com/404akshat)**
