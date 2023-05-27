from flask import Flask, request
from banki_ru import bankiRu
import json


app = Flask(__name__)


@app.route('/get/news', methods=['GET', "POST"])
def get_news():

    banki_ru = bankiRu()
    ready_news = banki_ru.get_news()

    return json.dumps(ready_news)


if __name__ == "__main__":
    app.run(host='', port=40394, debug=True)
    