import string
import random

from sqlalchemy.engine import Connection

from src.db.schemas import products


def random_string(length: int = 10):
    letters = string.ascii_lowercase
    return "".join(random.choice(letters) for i in range(length))


def random_float(_min: int = 0, _max: int = 9):
    return random.uniform(_min, _max)


def db_read(db: Connection):
    q = products.select().limit(100)
    return db.execute(q).all()


def db_write(db: Connection):
    data = products.insert().values(
        name=random_string(),
        value=random_float(),
    )
    db.execute(data)
    db.commit()
