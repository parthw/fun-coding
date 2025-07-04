from contextlib import asynccontextmanager

from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles

from app.middlewares import context, logger
from dependencies import logger as app_logger
from routers import api_chat_history, ui
from services.chat_history import ChatHistory
from services.default_agent import Agent


@asynccontextmanager
async def lifespan(app: FastAPI):
    # STARTUP: Code before yield runs when app starts
    app_logger.info("App starting up...")
    _ = Agent().get_agent()
    _ = ChatHistory()

    yield  # App runs here, serving requests

    # SHUTDOWN: Code after yield runs when app stops
    app_logger.info("App shutting down...")
    # await ChatHistory().clear()


class WebServer:
    def __init__(self):
        self.app = FastAPI(
            lifespan=lifespan,
            title="AI Web Server",
        )
        self.app.mount("/static", StaticFiles(directory="templates"), name="static")
        self.app.add_middleware(logger.WebServerLoggingMiddleware)
        self.app.add_middleware(context.WebServerReqContextMiddleware)

        self.app.include_router(ui.uiRouter)
        self.app.include_router(api_chat_history.apiChatHistoryRouter)
