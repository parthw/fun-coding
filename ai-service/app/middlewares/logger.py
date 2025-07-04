import time

from fastapi import Request
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.requests import Request

from dependencies import logger


class WebServerLoggingMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        start_time = time.monotonic()
        response = await call_next(request)
        latency = (time.monotonic() - start_time) * 1000  # in milliseconds
        if 200 <= response.status_code < 400:
            logger.info("access_log", extra={"latency": f"{latency:.2f}"})
        else:
            logger.error("access_log", extra={"latency": f"{latency:.2f}"})
        return response
