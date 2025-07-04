import uuid

from fastapi import Request
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.requests import Request

from dependencies import context


class WebServerReqContextMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        ctx = context.RequestContext(
            request_id=request.headers.get("x-request-id", str(uuid.uuid4())),
            method=request.method,
            path=request.url.path,
        )
        context.request_context_var.set(ctx)

        response = await call_next(request)
        response.headers["x-request-id"] = ctx.request_id
        return response
