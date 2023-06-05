FROM python:alpine

WORKDIR /app

RUN pip install Flask && pip install aiogram && pip install requests && pip install python-dotenv


COPY . .

CMD ["python", "bot.py"]