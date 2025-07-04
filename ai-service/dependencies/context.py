from contextvars import ContextVar
from dataclasses import dataclass
from typing import Optional


@dataclass
class RequestContext:
    request_id: str
    method: str
    path: str


request_context_var: ContextVar[Optional[RequestContext]] = ContextVar("request_context_var", default=None)
