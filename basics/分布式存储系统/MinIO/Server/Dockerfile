FROM python:3.12-rc
EXPOSE 9115

WORKDIR /app
COPY . /app

RUN pip install --upgrade pip && pip install -r requirements.txt

CMD ["python", "storage.py"]
