from pathlib import Path

import yaml
from pydantic import BaseModel, Field
from pydantic_settings import BaseSettings


class DevMode(BaseModel):
    Enabled: bool = Field(default=False, alias="enabled")

    class Config:
        case_sensitive = False
        validate_assignment = True
        frozen = True
        strict = True


class DefaultAgent(BaseModel):
    model: str = Field(alias="model")
    provider_base_url: str = Field(alias="provider_base_url")

    class Config:
        case_sensitive = False
        validate_assignment = True
        frozen = True
        strict = True


class Config(BaseSettings):
    dev_mode: DevMode = Field(alias="dev_mode")
    default_agent: DefaultAgent = Field(alias="default_agent")

    class Config:
        case_sensitive = False
        validate_assignment = True
        frozen = True
        strict = True


CONFIG_FILE_PATH = "./config/config.yaml"
_config_data = yaml.safe_load(Path(CONFIG_FILE_PATH).read_text())
config = Config(**_config_data)
