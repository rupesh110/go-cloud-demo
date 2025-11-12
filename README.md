# 🐳 Go Cloud Demo — Cloud-Agnostic CI/CD with Docker, GitHub Actions & Artifact Registry

A minimal Go backend demonstrating how to **build once, deploy anywhere** — across **GCP, Azure, and AWS** — using **Docker**, **GitHub Actions**, and **Artifact Registry**.

This project powers my article:  
**[Cloud-Agnostic Docker CI/CD with GitHub Actions & Artifact Registry: How to Switch Between GCP, Azure & AWS Without Rebuilding](#)**  
*(link your Medium article once it’s live!)*

---

## Overview

This demo app shows how a single Docker image can:
- Run locally or on any cloud
- Be built and pushed once
- Be deployed automatically via GitHub Actions

One pipeline. Any cloud. No rebuilds.

---

## Tech Stack

| Component | Purpose |
|------------|----------|
| **Go (Golang)** | Lightweight backend service |
| **Docker** | Containerization for portability |
| **GitHub Actions** | Continuous Integration & Deployment |
| **Google Artifact Registry** | Central image repository (multi-cloud bridge) |
| **GCP Cloud Run / Azure Container Apps / AWS ECS** | Target deployment platforms |

---

## Local Development

### 1️ Clone the repo
```bash
git clone https://github.com/<your-username>/go-cloud-demo.git
cd go-cloud-demo
```

### 2 Create a `.env` file
```bash
PORT=8080
CLOUD_PROVIDER=local
```

### 3 Run locally
```bash
go run main.go
```

Visit [http://localhost:8080](http://localhost:8080)

You’ll see:
```
Hello from Go! Cloud-agnostic demo running on local
```

---

### Testing

The project includes a simple unit test:

```bash
go test ./... -v
```

Verifies HTTP response and environment variable handling.

---

## Docker Usage

### Build the image
```bash
docker build -t go-cloud-demo .
```

### Run the container
```bash
docker run -p 8080:8080 -e CLOUD_PROVIDER=local go-cloud-demo
```

Visit [http://localhost:8080](http://localhost:8080)

### Stop the container
```bash
docker ps
docker stop <container_id>
```

---

## Cloud Deployment

### 1 Push to Google Artifact Registry
```bash
gcloud auth configure-docker REGION-docker.pkg.dev
docker tag go-cloud-demo REGION-docker.pkg.dev/PROJECT_ID/demo-repo/go-cloud-demo:latest
docker push REGION-docker.pkg.dev/PROJECT_ID/demo-repo/go-cloud-demo:latest
```

### 2 Deploy to Google Cloud Run
```bash
gcloud run deploy go-cloud-demo   --image REGION-docker.pkg.dev/PROJECT_ID/demo-repo/go-cloud-demo:latest   --region REGION   --platform managed   --allow-unauthenticated   --set-env-vars "CLOUD_PROVIDER=GCP"
```

### 3 (Optional) Mirror to Azure or AWS

#### Azure:
```bash
az acr import   --name myregistry   --source REGION-docker.pkg.dev/PROJECT_ID/demo-repo/go-cloud-demo:latest   --image go-cloud-demo:latest
```

#### AWS:
```bash
aws ecr get-login-password --region REGION | docker login --username AWS --password-stdin <AWS_ACCOUNT_ID>.dkr.ecr.<region>.amazonaws.com
docker tag go-cloud-demo <AWS_ACCOUNT_ID>.dkr.ecr.<region>.amazonaws.com/go-cloud-demo:latest
docker push <AWS_ACCOUNT_ID>.dkr.ecr.<region>.amazonaws.com/go-cloud-demo:latest
```

---

## CI/CD with GitHub Actions

This repo includes `.github/workflows/ci-cd.yml`, which:

- Runs Go tests  
- Builds and tags the Docker image  
- Pushes it to Google Artifact Registry  
- Deploys automatically to Google Cloud Run

Add the following secrets under  
**Settings → Secrets and Variables → Actions:**

| Secret | Description |
|---------|--------------|
| `GCP_PROJECT_ID` | Your GCP project ID |
| `GCP_REGION` | Example: `australia-southeast1` |
| `GCP_SA_KEY` | JSON key of your service account |

---

## Project Structure

```
go-cloud-demo/
├── .env                # Local environment variables
├── .gitignore
├── .dockerignore
├── Dockerfile
├── go.mod
├── go.sum
├── main.go             # App entry point
├── main_test.go        # Unit test
└── .github/
    └── workflows/
        └── ci-cd.yml   # GitHub Actions pipeline
```

---

## Lessons Learned

- **Containers** are the ultimate portability layer.  
- **GitHub Actions** can automate multi-cloud workflows.  
- **Artifact Registry** bridges multiple clouds seamlessly.  
- True multi-cloud doesn’t mean constant migration — it means **freedom of choice**.

---

## Disclaimer

This is a **personal demo project** for learning and educational purposes.  
All credentials and environment variables shown are examples only.  
Never commit secrets to GitHub — use **GitHub Secrets** or each cloud’s Secret Manager.

---

## Author

**Rupesh Shrestha**  
Developer | Cloud Enthusiast | CI/CD Learner  
[LinkedIn](https://www.linkedin.com/in/rupeshshresthas)