from flask import Flask, request
from lenta_ru import lentaRu
import json


app = Flask(__name__)


@app.route('/get/news', methods=['GET', "POST"])
def get_news():

    lenta_ru = lentaRu()
    ready_news = lenta_ru.get_news()

    return json.dumps(ready_news)


if __name__ == "__main__":
    app.run(host='', port=9010, debug=True)
    