from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.asyncio import create_async_engine
from sqlalchemy.ext.asyncio import async_sessionmaker

from src.config import DatabaseConfig

db_config = DatabaseConfig()
conn_string = (
    f"{db_config.username}:{db_config.pwd}@{db_config.host}:{db_config.port}/{db_config.name}"
)
db_url = f"postgresql+asyncpg://{conn_string}"

db_engine = create_async_engine(db_url)

DBSession = async_sessionmaker(db_engine)

async def db_connection():
    session = DBSession()
    try:
        yield session
    finally:
        session.close()
