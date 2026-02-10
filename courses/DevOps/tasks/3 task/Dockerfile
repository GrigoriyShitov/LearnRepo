FROM python:3.12-slim AS builder

COPY app/requirements.txt ./

RUN apt-get update && apt-get install -y --no-install-recommends \
      build-essential \
    && pip install --upgrade pip \
    && pip wheel --no-cache-dir --wheel-dir /wheels -r requirements.txt \
    && apt-get purge -y build-essential \
    && rm -rf /var/lib/apt/lists/*


FROM python:3.12-slim AS runtime

WORKDIR /opt/app
COPY --from=builder /wheels /wheels
RUN pip install --no-cache-dir /wheels/* && rm -rf /wheels

COPY app ./app

RUN useradd -r -u 10001 appuser
USER appuser

EXPOSE 8000

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
