from typing import Dict, List, Literal

from pydantic import BaseModel
from pydantic_ai.messages import ModelRequest, ModelResponse


class ChatMessage(BaseModel):
    role: Literal["user", "assistant"]
    content: str


class ChatHistory:
    _instance = None
    _initialized = False

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance

    def __init__(self):
        if not self._initialized:
            self._initialized = True
            self.user_histories: Dict[str, List[ChatMessage]] = {}

    def _get_user_history(self, user_id: str) -> List[ChatMessage]:
        return self.user_histories.setdefault(user_id, [])

    def add_user_message(self, user_id: str, content: str):
        self._get_user_history(user_id).append(ChatMessage(role="user", content=content))
        self._trim_old_messages(user_id)

    def add_assistant_message(self, user_id: str, content: str):
        self._get_user_history(user_id).append(ChatMessage(role="assistant", content=content))
        self._trim_old_messages(user_id)

    def get_history(self, user_id: str) -> List[Dict]:
        """Returns history for a user in dict format (e.g., for model input)."""
        return [msg.model_dump() for msg in self._get_user_history(user_id)]

    def get_pydantic_ai_history(self, user_id: str) -> List[ModelRequest | ModelResponse]:
        """Returns history in PydanticAI Message format."""
        messages = []
        for msg in self._get_user_history(user_id):
            if msg.role == "user":
                messages.append(ModelRequest.user_text_prompt(msg.content))
            else:
                # For assistant messages, we need to create a ModelResponse
                # This is a simplified approach - in practice you might want to store more details
                from pydantic_ai.messages import ModelResponse, TextPart

                messages.append(ModelResponse(parts=[TextPart(content=msg.content)]))
        return messages

    def clear(self, user_id: str):
        """Clears chat history for a specific user."""
        self.user_histories[user_id] = []

    def _trim_old_messages(self, user_id: str, max_messages: int = 25):
        """Trims the oldest messages if history exceeds `max_messages`."""
        history = self._get_user_history(user_id)
        if len(history) > max_messages:
            self.user_histories[user_id] = history[-max_messages:]  # Keep only last N messages
