import os
import sys
from dataclasses import dataclass

from dotenv import load_dotenv

load_dotenv()


@dataclass(frozen=True)
class DatabaseConfig:
    host = os.getenv("DB_HOST") or sys.exit("Environment variable DB_HOST is required")
    port = os.getenv("DB_PORT", 5432)
    username = os.getenv("DB_USERNAME") or sys.exit(
        "Environment variable DB_USERNAME is required"
    )
    pwd = os.getenv("DB_PASSWORD") or sys.exit(
        "Environment variable DB_PASSWORD is required"
    )
    name = os.getenv("DB_NAME", "fastapi")
