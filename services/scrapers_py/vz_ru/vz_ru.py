import requests
import feedparser
from datetime import datetime


class VZ:
	def __init__(self):
		self.base_url = "https://vz.ru/export/yandex.xml"

	def _get_timestamp_from_string(self, date: str) -> int:
		date_format = '%a, %d %b %Y %H:%M:%S %z'
		date_object = datetime.strptime(date, date_format)
		timestamp = date_object.timestamp()

		return int(timestamp)

	def get_news(self) -> list:
		all_news = feedparser.parse(self.base_url)
		ready_news_list = []

		for news in all_news['entries']:
			ready_news_list.append({'title': news['title'],
				'link': news['links'][0]['href'],
				'description': news['summary'],
				'imageUrl': None,
				'source': 'VzRu',
				'tags': [tag['term'] for tag in news['tags']],
				'publishingTime': self._get_timestamp_from_string(news['published'])})
		
		return ready_news_list
