from pydantic import BaseModel


class Payload(BaseModel):
    payload: str
