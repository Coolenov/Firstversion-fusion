FROM python:alpine

WORKDIR /app

RUN pip install Flask && pip install requests && pip install beautifulsoup4 && pip install pprint3x

COPY . .

CMD ["python", "api.py"]