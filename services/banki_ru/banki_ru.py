import requests
import feedparser
from datetime import datetime
import re


class bankiRu:
	def __init__(self):
		self.base_url = "https://www.banki.ru/xml/news.rss"

	def _get_timestamp_from_string(self, date: str) -> int:
		date_format = '%a, %d %b %Y %H:%M:%S %z'
		date_object = datetime.strptime(date, date_format)
		timestamp = date_object.timestamp()

		return int(timestamp)

	def _clean_text_from_tags(self, text: str) -> str:
		text_without_tags = re.sub('<.*?>', '', text)

		return text_without_tags

	def get_news(self) -> list:
		all_news = feedparser.parse(self.base_url)
		print(all_news)
		ready_news_list = []

		for news in all_news['entries']:
			ready_news_list.append({'title': news['title'],
				'link': news['links'][0]['href'],
				'description': self._clean_text_from_tags(news['summary']),
				'image_url': news['links'][1]['href'] if len(news['links']) > 1 else None,
				'source': 'banki.ru',
				'tags': [tag['term'] for tag in news['tags']],
				'publishing_time': self._get_timestamp_from_string(news['published'])})
		
		return ready_news_list

b = bankiRu()

news = b.get_news()

print(news)