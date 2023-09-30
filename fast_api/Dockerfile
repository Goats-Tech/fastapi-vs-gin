FROM python:3.11-alpine as build

WORKDIR /app

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

FROM python:3.11-alpine

WORKDIR /app

COPY --from=build /opt/venv /opt/venv
COPY src src/
COPY migrations migrations/
COPY alembic.ini alembic.ini

ENV PYTHONPATH /app
ENV PATH="/opt/venv/bin:$PATH"
EXPOSE 4444

CMD ["python", "src/main.py"]
