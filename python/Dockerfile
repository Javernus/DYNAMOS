# ---- Base python ----
# FROM python:3.9-slim-buster AS base
FROM  python:3.12-slim

ARG NAME
ARG WHEEL_NAME
ENV NAME=$NAME

WORKDIR /app
RUN mkdir $NAME
RUN mkdir datasets
RUN mkdir protofiles

COPY requirements.txt .
COPY *_requirements.txt .
COPY ./*.py /app/$NAME
COPY ./protofiles/*.py /app/protofiles

COPY ./dataset[s]/ /app/datasets
COPY $WHEEL_NAME /tmp/

RUN pip install /tmp/$WHEEL_NAME
RUN pip install --no-cache-dir -r requirements.txt
RUN pip install --upgrade numpy pandas

RUN if [ -f "torch_requirements.txt" ]; then \
  pip install -r torch_requirements.txt; \
fi

ENV  ENV="PROD"
ENV PYTHONPATH="/app/protofiles"

CMD ["sh", "-c", "python /app/${NAME}/main.py"]
