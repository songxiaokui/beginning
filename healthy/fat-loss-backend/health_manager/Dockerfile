FROM python:3.11-slim
EXPOSE 5001

WORKDIR /app
COPY . /app

RUN pip install --upgrade pip && pip install -r requirements.txt
RUN python manage.py collectstatic --noinput

CMD ["gunicorn", "health_manager.wsgi:application", "--bind", "0.0.0.0:5011", "--workers", "3"]
