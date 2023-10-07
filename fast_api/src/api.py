from fastapi import FastAPI, Depends
from pydantic import BaseModel
from sqlalchemy.ext.asyncio import AsyncSession

from src.db import db_connection, db_engine
from src.db.helpers import db_read, db_write

app = FastAPI()


@app.on_event("shutdown")
def shutdown():
    """
    Closes the database connection when the FastAPI application is shut down.
    """
    db_engine.dispose()


class Product(BaseModel):
    name: str
    value: float


@app.get("/read", response_model=list[Product])
async def read(db: AsyncSession = Depends(db_connection)):
    return await db_read(db)


@app.post("/write")
async def write(db: AsyncSession = Depends(db_connection)):
    return await db_write(db)
