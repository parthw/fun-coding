from fastapi import APIRouter, Request
from fastapi.responses import HTMLResponse, JSONResponse
from fastapi.templating import Jinja2Templates
from pydantic import BaseModel

from dependencies import logger
from services import default_agent
from services.chat_history import ChatHistory

uiRouter = APIRouter(
    prefix="/ui",
    tags=["ui"],
    responses={404: {"description": "Not found"}},
)

templates = Jinja2Templates(directory="templates")


@uiRouter.get("/", response_class=HTMLResponse)
async def get_ui(request: Request):
    return templates.TemplateResponse("index.html", {"request": request})


class ChatRequest(BaseModel):
    message: str
    userId: str | None = None


@uiRouter.post("/chat")
async def chat(request: ChatRequest):
    try:
        logger.info(f"Processing chat request: {request.message[:50]}")
        user_id = request.userId or "anonymous"
        agent = default_agent.Agent().get_agent()
        history = ChatHistory()

        chat_messages = history.get_pydantic_ai_history(user_id)

        result = await agent.run(request.message, message_history=chat_messages)
        response_data = result.output.format()

        # Add both user message and assistant response to history
        history.add_user_message(user_id, request.message)
        history.add_assistant_message(user_id, response_data)

        # Return comprehensive response
        return JSONResponse(content={"response": response_data, "userId": user_id})

    except Exception as e:
        logger.error(f"Error in chat endpoint: {str(e)}")
        return JSONResponse(
            status_code=500,
            content={
                "response": "I apologize, but I'm experiencing technical difficulties. Please try again in a moment.",
            },
        )
