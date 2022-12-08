import os
from fastapi import FastAPI, Request

deployment_type = os.getenv("DEPLOYMENT_TYPE", "api")
app = FastAPI()


@app.get("/")
async def root_handler():
    return f"deployment type = {deployment_type}"


@app.get("/health")
async def health_handler():
    return "Healthy"
