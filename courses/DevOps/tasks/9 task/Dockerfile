FROM python:3.12-slim AS builder

COPY requirements-main.txt ./

RUN apt update && apt install -y --no-install-recommends \
      build-essential \
    && pip install --upgrade pip \
    && pip wheel --no-cache-dir --wheel-dir /wheels -r requirements-main.txt \
    && apt purge -y build-essential \
    && rm -rf /var/lib/apt/lists/*


FROM python:3.12-slim AS runtime

WORKDIR /opt/app
COPY --from=builder /wheels /wheels
RUN apt update && apt install wget -y --no-install-recommends \
    && pip install --no-cache-dir /wheels/* \
    && rm -rf /wheels /var/lib/apt/lists/*

COPY app ./app

RUN mkdir -p log/main

RUN adduser --disabled-password --gecos '' appuser && \
    chown -R appuser:appuser /opt/app
USER appuser

EXPOSE 8000

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
