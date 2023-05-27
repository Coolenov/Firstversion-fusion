from flask import Flask, request
from online_ua_news import UaNews
import json


app = Flask(__name__)


@app.route('/get/news', methods=['GET', "POST"])
def get_news():

    ua = UaNews()
    ready_news = ua.get_news()

    return json.dumps(ready_news)


if __name__ == "__main__":
    app.run(host='', port=40390, debug=True)