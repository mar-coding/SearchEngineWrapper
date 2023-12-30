# 'https://um.ac.ir/members/professors/index.html',
# 'https://um.ac.ir/members/staffs/index.html',
# 'https://um.ac.ir/members/students/index.html'


import scrapy
import re
from scrapy.linkextractors import LinkExtractor
from elasticsearch import Elasticsearch

es = Elasticsearch(['https://localhost:9200'], basic_auth=('elastic', '6qYp22cj9aO4XM7fAK'), verify_certs=False)
class UMScraperSpider(scrapy.Spider):
    name = 'umScraper'
    start_urls = [
        'https://um.ac.ir/members/professors/index.html',
        'https://um.ac.ir/members/staffs/index.html',
        'https://um.ac.ir/members/students/index.html'
    ]
    def parse(self, response):
        all_text = response.xpath('//text()').getall()
        website_title = response.xpath('//title/text()').get()
        doc = {'title': website_title, 'text': all_text, 'url': response.url}
        es.index(index='ferdowsi_university', body=doc)

        current_domain = response.url.split('/')[2]
        linkExtractor = LinkExtractor()
        links = linkExtractor.extract_links(response)
        for link in links:
            if link.url.startswith('/') or current_domain in link.url:
                yield scrapy.Request(link.url, callback=self.parse)
