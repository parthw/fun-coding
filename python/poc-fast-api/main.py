from fastapi import FastAPI, Request
from fastapi.responses import HTMLResponse
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates

from models import Payload

templates = Jinja2Templates(directory="web/html")
app = FastAPI()
app.mount("/web/css", StaticFiles(directory="web/css"), name="css")


@app.get("/", response_class=HTMLResponse)
async def root(request: Request):
    return templates.TemplateResponse("index.html", {"request": request})


@app.post("/payload")
async def post_payload(payload: Payload):
    return payload
