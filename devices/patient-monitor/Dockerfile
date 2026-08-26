FROM python:3.12-alpine

WORKDIR /app

RUN pip install --no-cache-dir requests

COPY simulator.py .

CMD ["python", "simulator.py"]