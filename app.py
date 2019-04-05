from urllib.request import urlopen
from time import strftime, strptime
from flask import Flask, render_template, request
from bs4 import BeautifulSoup

app = Flask(__name__)
url = 'https://snowstar.org/'
pages = []
articles_list = []
while url:
    page = BeautifulSoup(urlopen(url).read().decode('utf-8'), 'html.parser')
    pages.append(page)
    url = page.select('.next')[0]['href'] if page.select('.next') else None
for page in pages:
    articles = []
    for article in page.select('article'):
        articles.append([
            # Link
            article.select('h2 a')[0]['href'],
            # Title
            article.select('h2')[0].get_text(),
            # Date
            strftime('%Y-%m-%d',
                     strptime(
                         article.select('.posted-date')[0].get_text(),
                         '%Y年%m月%d日'))
            if article.select('.posted-date') else None,
            # Excerpt
            article.select('.entry-summary p:first-child')[0].get_text()
        ])
    articles_list.append(articles)


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/ncuos')
def ncuos():
    return render_template('ncuos.html')


@app.route('/crawler')
def crawler():
    if request.args.get('q'):
        search_url = 'https://snowstar.org/?s=' + request.args.get('q')
        search_results = []
        while search_url:
            search_page = BeautifulSoup(
                urlopen(search_url).read().decode('utf-8'), 'html.parser')
            search_url = search_page.select('.next')[0][
                'href'] if search_page.select('.next') else None
            for search_article in search_page.select('article'):
                search_results.append([
                    # Link
                    search_article.select('h2 a')[0]['href'],
                    # Title
                    search_article.select('h2')[0].get_text(),
                    # Excerpt
                    article.select('.entry-summary p:first-child')[0]
                    .get_text()
                ])
        return render_template(
            'crawler.html',
            articles=search_results,
            keyword=request.args.get('q'))
    else:
        return render_template(
            'crawler.html',
            articles=articles_list[int(request.args.get('p')) - 1 if request.args.get('p') else 0],
            page=[
                # Current
                int(request.args.get('p')) if request.args.get('p') else 1,
                # Total
                len(articles_list)
            ])


if __name__ == '__main__':
    app.run()
