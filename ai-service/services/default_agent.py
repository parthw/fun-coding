from pydantic import BaseModel, Field
from pydantic_ai import Agent as PydanticAIAgent
from pydantic_ai.models.openai import OpenAIModel
from pydantic_ai.providers.openai import OpenAIProvider

from dependencies.config import config

SYSTEM_PROMPT = """
You are a helpful AI assistant.
Your primary role is to assist users with their queries and provide accurate, helpful responses.
"""


class AgentOutput(BaseModel):
    message: str = Field(description="The output message from the AI agent")

    def format(self) -> str:
        return self.message.strip()


# Initialize the AI agent (singleton pattern)
class Agent:
    _instance = None
    _agent = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance

    def get_agent(self):
        if self._agent is None:
            self._agent = self._create_agent()
        return self._agent

    def _create_agent(self):
        model = OpenAIModel(
            model_name=config.default_agent.model,
            provider=OpenAIProvider(base_url=config.default_agent.provider_base_url),
        )
        return PydanticAIAgent[None, AgentOutput](
            model=model,
            result_type=AgentOutput,
            system_prompt=SYSTEM_PROMPT,
            result_retries=2,
            end_strategy="early",
        )
