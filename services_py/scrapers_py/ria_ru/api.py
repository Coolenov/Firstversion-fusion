from flask import Flask, request
from ria_news import RiaNews
import json


app = Flask(__name__)


@app.route('/get/news', methods=['GET', "POST"])
def get_news():

    ria = RiaNews()
    ready_news = ria.get_news()

    return json.dumps(ready_news)


if __name__ == "__main__":
    app.run(host='', port=9040, debug=True)