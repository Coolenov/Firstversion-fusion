from flask import Flask, request
from vz_ru import VZ
import json


app = Flask(__name__)


@app.route('/get/news', methods=['GET', "POST"])
def get_news():

    vz = VZ()
    ready_news = vz.get_news()

    return json.dumps(ready_news)


if __name__ == "__main__":
    app.run(host='', port=9050, debug=True)
    