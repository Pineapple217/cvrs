# /// script
# requires-python = ">=3.12"
# dependencies = [
#     "bs4",
#     "requests",
# ]
# ///
import os
import requests
from bs4 import BeautifulSoup
from concurrent.futures import ThreadPoolExecutor

TAGS = ["hip-hop", "indie+rock", "drum+and+bass", "pop+punk", "electronic"]
BASE_URL = "https://www.last.fm/"
MAX_PAGE = 10
TOKEN = os.getenv("CVRS_TOKEN")
BACKEND_URL = "http://localhost:3000/api"
ARTISTS_ADD_URL = BACKEND_URL + "/artists/add"

session = requests.Session()

def handle_artist(link):
    name = link.text
    print("processing:", name)
    img = get_artists_picture(link.attrs["href"])
    post_artist(name, img)

def main():
    with ThreadPoolExecutor(max_workers=16) as executor:
        for tag in TAGS:
            url = f"{BASE_URL}/tag/{tag}/artists"
            params = {"page": 1}
            while params["page"] <= MAX_PAGE:
                resp = session.get(url, params=params)
                page = BeautifulSoup(resp.content, "html.parser")
                for a in page.find_all("h3", class_="big-artist-list-title"):
                    link = a.find("a")
                    executor.submit(handle_artist, link)
                params["page"] += 1

def get_artists_picture(artist_url: str):
    url = f"{BASE_URL}/{artist_url}/+images"
    resp = session.get(url)
    page = BeautifulSoup(resp.content, "html.parser")
    for img in page.find_all("a", class_="image-list-item"):
        img_page = session.get(BASE_URL + img["href"])
        soup2 = BeautifulSoup(img_page.content, "html.parser")
        img_url = soup2.find("img", class_="js-gallery-image")["src"]
        resp2 = session.get(img_url)
        mime = resp2.headers["Content-Type"]
        img_name = img_url.split("/")[-1].split("#")[0]
        if mime == "image/jpeg":
            print(mime, img_name)
            return (img_name, resp2.content, mime)

def post_artist(name: str, img):
    resp = session.post(
        ARTISTS_ADD_URL,
        headers={"Authorization": f"Bearer {TOKEN}"},
        files={"img": img},
        data={"json": f'{{"name": "{name}"}}'},
    )
    if resp.status_code != 200:
        print(resp.status_code, resp.content)

if __name__ == "__main__":
    main()
