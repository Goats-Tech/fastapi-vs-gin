from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from src.config import DatabaseConfig

db_config = DatabaseConfig()
conn_string = (
    f"{db_config.username}:{db_config.pwd}@{db_config.host}:{db_config.port}/{db_config.name}"
)
db_url = f"postgresql+psycopg://{conn_string}"

db_engine = create_engine(db_url)

DBSession = sessionmaker(autocommit=False, autoflush=False, bind=db_engine)

def db_connection():
    session = DBSession()
    try:
        yield session
    finally:
        session.close()
