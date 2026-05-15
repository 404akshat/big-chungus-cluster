# Aether OS — Big Chungus Cluster
### A Self-Hosted PaaS for Kubernetes, Built From Scratch

[![Language](https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Orchestration](https://img.shields.io/badge/Orchestration-Kubernetes-326CE5?style=flat&logo=kubernetes)](https://kubernetes.io/)
[![Containerization](https://img.shields.io/badge/Containerization-Docker-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![Status](https://img.shields.io/badge/Status-Week_1_Complete-brightgreen)](https://github.com/404akshat/big-chungus-cluster)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Aether OS is a self-hosted Platform-as-a-Service built entirely from scratch. The goal is simple: **you provide a Docker image, Aether OS handles the rest** — deployments, scaling, monitoring, security, and self-healing. All of it. No managed cloud. No abstractions you don't own.

> [!CAUTION]
> **Active Development:** Built and documented weekly in public. This is not a tutorial project, every component is production-grade or working toward it. Progress is tracked weekly on LinkedIn.

---

## 💡 The Core Concept

Managing Kubernetes manifests, ingress controllers, autoscalers, and security policies is a full-time job. Aether OS is the operating system for your cluster — a platform that sits on top of Kubernetes and makes it feel like you have a DevOps team running behind the scenes.

### The Pipeline:
1. **Deploy:** Send a request to `bcc-api` or use `bcc-ctl` from your terminal
2. **Abstract:** Aether OS generates all K8s Deployments, Services, and Ingress resources automatically
3. **Observe:** Built-in metrics, logs, and traces with no manual configuration
4. **Harden:** Chaos engineering modules stress-test the platform continuously
5. **Recover:** Self-healing watcher detects failures and rolls back automatically

---

## 🏗️ Architecture — 21 Components

| # | Component | What It Does | Status |
|---|-----------|-------------|--------|
| 1 | `bcc-api` | The platform brain — Go HTTP server that drives everything | ✅ Week 1 |
| 2 | `bcc-ctl` | CLI tool to control the platform from your terminal | 🔧 Week 2 |
| 3 | PostgreSQL | Persistent memory — stores all app metadata and state | 🔧 Week 2 |
| 4 | Terraform | Provisions the entire cluster infrastructure as code | 📋 Planned |
| 5 | Helm Charts | Packages every component for repeatable deployments | 📋 Planned |
| 6 | GitHub Actions | CI/CD pipeline — test, build, scan, push, deploy | 📋 Planned |
| 7 | ArgoCD | GitOps engine — Git is the source of truth | 📋 Planned |
| 8 | Prometheus | Scrapes and stores all cluster and app metrics | 📋 Planned |
| 9 | Grafana | Turns raw metrics into dashboards you actually want to look at | 📋 Planned |
| 10 | Loki | Centralizes logs from every container in the cluster | 📋 Planned |
| 11 | OpenTelemetry + Jaeger | Distributed tracing — finds exactly where things slow down | 📋 Planned |
| 12 | Custom K8s Controller | A real Kubernetes Operator written in Go | 📋 Planned |
| 13 | LitmusChaos | Deliberately breaks things to prove the platform survives | 📋 Planned |
| 14 | Auto-Recovery Watcher | Self-healing Go service that detects and rolls back failures | 📋 Planned |
| 15 | Trivy | Scans every Docker image for vulnerabilities before deployment | 📋 Planned |
| 16 | OPA Gatekeeper | Policy enforcer — the bouncer at the cluster door | 📋 Planned |
| 17 | Falco | Runtime threat detection at the kernel level | 📋 Planned |
| 18 | HashiCorp Vault | Secrets management — nothing sensitive ever touches Git | 📋 Planned |
| 19 | Cilium | Zero-trust networking with mTLS between every pod | 📋 Planned |
| 20 | KEDA | Event-driven autoscaling based on real platform metrics | 📋 Planned |
| 21 | React Dashboard | Visual control plane — the face of the whole platform | 📋 Planned |

---

## ✅ Week 1 — What's Built

### `bcc-api`
The core Go HTTP server. The first and most critical component — everything else in Aether OS talks to or builds on top of this.

**Endpoints:**
| Method | Route | Description |
|--------|-------|-------------|
| GET | `/health` | Returns platform health status |
| GET | `/apps` | Lists all registered applications |
| POST | `/apps` | Registers a new application |
| GET | `/apps/:name` | Fetches a specific application |
| DELETE | `/apps/:name` | Removes an application |

**Running locally:**
```bash
cd bcc-api
go run main.go
# Server starts on :8080
```

**Running in Docker:**
```bash
docker build -t bcc-api .
docker run -p 8080:8080 bcc-api
```

**Deploying to Kubernetes:**
```bash
eval $(minikube docker-env)
docker build -t bcc-api .
kubectl apply -f k8s/bcc-api/deployment.yaml
kubectl apply -f k8s/bcc-api/service.yaml
minikube service bcc-api-service
```

---

## 📁 Repository Structure

```
big-chungus-cluster/
├── bcc-api/                  # Go HTTP server — the platform brain
│   ├── main.go
│   ├── handlers/
│   │   └── apps.go
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── k8s/                      # Kubernetes manifests
│   ├── deployment.yaml
│   └── service.yaml
└── README.md
```

*Structure grows every week as new components are added.*

---

## 🗓️ Build Log

| Week | Focus | Deliverable |
|------|-------|------------|
| **Week 1** | Go API + Docker + Kubernetes | `bcc-api` running on local K8s cluster |
| **Week 2** | PostgreSQL + `bcc-ctl` CLI | Database-backed API + terminal control tool |
| **Week 3** | Terraform + Helm + ArgoCD | Infrastructure as code + GitOps pipeline |
| **Week 4** | GitHub Actions + CI/CD | Fully automated build and deploy pipeline |
| **Week 5+** | Observability, Security, Chaos | The platform gets serious |

---

## 🚀 Roadmap

- [x] REST API in Go with JSON endpoints
- [x] Docker containerization
- [x] Kubernetes deployment on local cluster
- [ ] PostgreSQL integration
- [ ] CLI tool (`bcc-ctl`)
- [ ] Terraform infrastructure
- [ ] Full CI/CD pipeline
- [ ] Observability stack
- [ ] Custom Kubernetes Operator
- [ ] Chaos engineering + self-healing
- [ ] Security layer
- [ ] React control plane dashboard

---

## 🔗 Follow the Build

Every component gets a LinkedIn post when it's done. Every week is a new post.

**LinkedIn:** [Akshat Sharma](https://www.linkedin.com/in/akshat-sharma-devsecop/)  
**GitHub:** [404akshat/big-chungus-cluster](https://github.com/404akshat/big-chungus-cluster)

---

## 📄 License

MIT License — Copyright (c) 2026 404akshat

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---

**Maintained by [404akshat](https://github.com/404akshat) — Built in public, one week at a time.**
