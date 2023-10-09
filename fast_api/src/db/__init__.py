from sqlalchemy import create_engine

from src.config import DatabaseConfig

db_config = DatabaseConfig()
conn_string = (
    f"{db_config.username}:{db_config.pwd}@{db_config.host}:{db_config.port}/{db_config.name}"
)
db_url = f"postgresql+psycopg://{conn_string}"

db_engine = create_engine(db_url)


def db_connection():
    conn = db_engine.engine.connect()
    yield conn
    conn.close()
