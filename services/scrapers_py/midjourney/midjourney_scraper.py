import requests
from bs4 import BeautifulSoup
import json
from pprint3x import pprint
from datetime import datetime



class Midjourney:
	def __init__(self):
		self.url = "https://www.midjourney.com/showcase/recent/"

	def _get_timestamp_from_string(self, date: str) -> int:
		date_string = date
		date_format = "%Y-%m-%d %H:%M:%S.%f"
		timestamp = datetime.strptime(date_string, date_format).timestamp()
		return int(timestamp)

	def getContent(self):
		imageUrls = []
		resp = requests.get(self.url)
		soup = BeautifulSoup(resp.text, 'html.parser')
		body = soup.find("script", {"id":"__NEXT_DATA__"})
		bodyText = body.text
		bodyJson = json.loads(bodyText)
		# pprint(bodyJson)

		props = bodyJson["props"]
		pageProps = props["pageProps"]
		jobs = pageProps["jobs"]

		posts = []
		for job in jobs:
			pprint(job)
			event = job["event"]
			pubTime = job["enqueue_time"]
			imageUrl = event["seedImageURL"]
			
			posts.append({'title': None,
				'link': imageUrl,
				'description': None,
				'imageUrl': imageUrl,
				'source': 'Midjourney',
				'tags': [],
				'publishingTime': self._get_timestamp_from_string(pubTime)})
		# pprint(posts)
		return posts
		

		# print(posts)






	

# obj = Midjourney()
# obj.getContent()