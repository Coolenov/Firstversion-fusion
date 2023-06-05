FROM python:alpine

WORKDIR /app

RUN pip install Flask
RUN pip install feedparser
RUN pip install requests



COPY . .

CMD ["python", "api.py"]