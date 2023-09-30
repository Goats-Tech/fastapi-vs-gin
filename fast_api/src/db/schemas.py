from sqlalchemy import Column, Integer, String, Float, MetaData, Table

metadata = MetaData()

products = Table(
    "products",
    metadata,
    Column("id", Integer, primary_key=True, autoincrement=True),
    Column("name", String, nullable=False),
    Column("value", Float, nullable=False),
)
