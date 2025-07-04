from fastapi import APIRouter, HTTPException
from fastapi.responses import JSONResponse
from pydantic import BaseModel

from dependencies import logger
from services.chat_history import ChatHistory

apiChatHistoryRouter = APIRouter(
    prefix="/api/chat_history",
    tags=["api", "chat_history"],
    responses={404: {"description": "Not found"}},
)


class UserHistoryRequest(BaseModel):
    user_id: str


@apiChatHistoryRouter.post("/get")
async def get_chat_history(payload: UserHistoryRequest):
    try:
        history = ChatHistory().get_history(payload.user_id)
        return JSONResponse(content={"history": history})
    except Exception as e:
        logger.error(f"Error getting chat history: {e}")
        raise HTTPException(status_code=500, detail="Internal server error")


@apiChatHistoryRouter.post("/clear")
async def clear_chat_history(payload: UserHistoryRequest):
    try:
        ChatHistory().clear(payload.user_id)
        return JSONResponse(content={"message": f"History cleared for user {payload.user_id}"})
    except Exception as e:
        logger.error(f"Error clearing chat history: {e}")
        raise HTTPException(status_code=500, detail="Internal server error")
