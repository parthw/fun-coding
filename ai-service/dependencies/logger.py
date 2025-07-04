import json
import logging
from datetime import datetime
from typing import Optional

from dependencies import context
from dependencies.config import config


class Color:
    Black = 30
    Red = 31
    Green = 32
    Yellow = 33
    Blue = 34
    Magenta = 35
    Cyan = 36
    White = 37

    @staticmethod
    def apply(color_code, text):
        """Apply the given color code to the text."""
        return f"\x1b[{color_code}m{text}\x1b[0m"


def extract_values_as_string(data):
    result = []
    stack = [("", data)]

    while stack:
        prefix, current = stack.pop()

        if isinstance(current, dict):
            for key, value in current.items():
                full_key = f"{prefix}.{key}" if prefix else key
                stack.append((full_key, value))
        elif isinstance(current, list):
            for i, item in enumerate(current):
                full_key = f"{prefix}[{i}]" if prefix else f"[{i}]"
                stack.append((full_key, item))
        elif isinstance(current, (str, int, float, bool)) or current is None:
            result.append(f"{prefix}={current}")

    return " ".join(result[::-1])


class _LoggerRequestContextFilter(logging.Filter):
    def filter(self, record: logging.LogRecord) -> bool:
        try:

            ctx: Optional[context.RequestContext] = context.request_context_var.get()
            if ctx is not None:
                record.request_id = ctx.request_id
                record.request_method = ctx.method
                record.request_path = ctx.path
        except LookupError:
            record.request_id = "N/A"
            record.request_method = "N/A"
            record.request_path = "N/A"

        return True


class _LoggerJSONFormatter(logging.Formatter):
    def format(self, record: logging.LogRecord) -> str:
        log_record = {
            "timestamp": datetime.now().isoformat(),
            "level": record.levelname,
            "message": record.getMessage(),
        }
        if getattr(record, "request_id", None) is not None:
            log_record["request_id"] = getattr(record, "request_id")

        if record.levelname != "INFO":
            log_record["file_name"] = record.pathname

        standard_attrs = logging.LogRecord("", 1, "", 0, "", (), None).__dict__.keys()
        for attr, value in record.__dict__.items():
            if attr not in standard_attrs and attr not in log_record:
                log_record[attr] = value
        return json.dumps(log_record)


class _LoggerJSONFormatterForDevMode(logging.Formatter):
    def format(self, record: logging.LogRecord) -> str:
        log_record = {
            "level": record.levelname,
            "message": record.getMessage(),
        }
        if getattr(record, "request_id", None) is not None:
            log_record["request_id"] = getattr(record, "request_id")

        if record.levelname != "INFO":
            log_record["file_name"] = record.pathname

        standard_attrs = logging.LogRecord("", 1, "", 0, "", (), None).__dict__.keys()
        for attr, value in record.__dict__.items():
            if attr not in standard_attrs and attr not in log_record:
                log_record[attr] = value

        # Apply colors based on log level using ANSI escape codes
        if record.levelname == "INFO":
            log_record["message"] = Color.apply(Color.Green, log_record["message"])  # Green for INFO
            log_record["level"] = Color.apply(Color.Green, log_record["level"])
        elif record.levelname == "ERROR":
            log_record["message"] = Color.apply(Color.Red, log_record["message"])  # Red for ERROR
            log_record["level"] = Color.apply(Color.Red, log_record["level"])
        elif record.levelname == "WARNING":
            log_record["message"] = Color.apply(Color.Yellow, log_record["message"])  # Yellow for WARNING
            log_record["level"] = Color.apply(Color.Yellow, log_record["level"])
        elif record.levelname == "DEBUG":
            log_record["message"] = Color.apply(Color.Cyan, log_record["message"])  # Cyan for DEBUG
            log_record["level"] = Color.apply(Color.Cyan, log_record["level"])
        else:
            log_record["message"] = Color.apply(Color.White, log_record["message"])  # White for other levels
            log_record["level"] = Color.apply(Color.White, log_record["level"])
        return extract_values_as_string(log_record)


# Set up logger
app_logger = logging.getLogger("app_logger")
app_logger.setLevel(logging.INFO)

_handler = logging.StreamHandler()
if config.dev_mode.Enabled:
    _handler.setFormatter(_LoggerJSONFormatterForDevMode())
else:
    _handler.setFormatter(_LoggerJSONFormatter())
_handler.addFilter(_LoggerRequestContextFilter())

app_logger.addHandler(_handler)


def info(msg: str, *args, **kwargs):
    app_logger.info(msg, *args, **kwargs)


def error(msg: str, *args, **kwargs):
    app_logger.error(msg, *args, **kwargs)


def debug(msg: str, *args, **kwargs):
    app_logger.debug(msg, *args, **kwargs)


def warning(msg: str, *args, **kwargs):
    app_logger.warning(msg, *args, **kwargs)
