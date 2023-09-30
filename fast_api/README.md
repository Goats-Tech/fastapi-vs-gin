

## Setup

```bash
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

## Docker image

Change directory to `fast_api`

```shell
cd fast_api
```
```shell
docker build -t fastapi:latest . --no-cache
```
