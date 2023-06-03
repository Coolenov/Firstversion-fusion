import requests
from dataclasses import dataclass
import json
import datetime


@dataclass
class Tag:
    text: str

@dataclass
class Content:
    Id:                 int
    Title:              str
    Link:               str
    Description:        str
    Image_url:          str
    Source:             str
    Publishing_time:    int

@dataclass
class Source:
    name:str


class Nax:
    def __init__(self):
        self.session = requests.Session()
        self.baseUrl = "http://127.0.0.1:9000"
    
    def _formPublishingTime(self, timeStamp:int)->datetime:
        return datetime.datetime.fromtimestamp(timeStamp)

    def _createContentFromResp(self,resp:json)->Content:
        resp = resp["content"]
        return Content(resp['Id'], resp['Title'], resp['Link'], resp['Description'], 
                        resp['Image_url'],resp['Source'],self._formPublishingTime(resp['Publishing_time']))

    def GetTextForTelegramMessage(self,content:Content)->str:
        return f'<b>{content.Title}</b>\n\n{content.Description}\n{content.Publishing_time}'

    def GetLastContentBySource(self,source:str)->Content:
        resp = self.session.post(f"{self.baseUrl}/content/source",data = {"source":source})
        resp = resp.json()
        return self._createContentFromResp(resp)

    def GetNextContent(self,post_id:str)->Content:
        resp = self.session.post(f"{self.baseUrl}/next",data = {"id":post_id})
        resp = resp.json()
        return self._createContentFromResp(resp)

    def GetPreviousContent(self,post_id:str)->Content:
        resp = self.session.post(f"{self.baseUrl}/previous",data = {"id":post_id})
        resp = resp.json()
        return self._createContentFromResp(resp)

    def GetSources(self)->list[str,...]:
        resp = self.session.get(f"{self.baseUrl}/sources")
        resp = resp.json()
        sources = resp["content"]
        return sources







