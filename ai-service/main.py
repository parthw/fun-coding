import argparse

import uvicorn
from fastapi import FastAPI

from dependencies import logger
from dependencies.config import config
from app.web_server import WebServer


def start_server():
    server = WebServer()
    uvicorn.run(server.app, host="0.0.0.0", port=8000, access_log=False)


def server_debug_mode() -> FastAPI:
    server = WebServer()
    return server.app


if __name__ == "__main__":
    if config.dev_mode.Enabled:
        logger.info("Starting the application in development mode")

    parser = argparse.ArgumentParser(description="Command line interface for the application.")
    parser.add_argument("command", choices=["server", "script"], help="Choose one of: server | script")
    args = parser.parse_args()

    if args.command == "server":
        start_server()
    elif args.command == "script":
        # Placeholder for script execution
        logger.info("Script execution is not implemented yet.")
    else:
        logger.error("Invalid command. Use 'server' or 'script'.")
        exit(1)
