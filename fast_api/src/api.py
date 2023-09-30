from fastapi import FastAPI, Depends
from pydantic import BaseModel
from sqlalchemy.engine import Connection

from src.db import db_connection
from src.db.helpers import db_read, db_write

app = FastAPI()


class Product(BaseModel):
    name: str
    value: float


@app.get("/read", response_model=list[Product])
def read(db: Connection = Depends(db_connection)):
    return db_read(db)


@app.post("/write")
def write(db: Connection = Depends(db_connection)):
    return db_write(db)
