from flask import Flask, request
from kommersant import kommersantNews
import json


app = Flask(__name__)


@app.route('/get/news', methods=['GET', "POST"])
def get_news():

    k = kommersantNews()
    ready_news = k.get_news()

    return json.dumps(ready_news)


if __name__ == "__main__":
    app.run(host='', port=9000, debug=True)
    