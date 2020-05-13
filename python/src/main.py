from flask import Flask, render_template
import os
import random

app = Flask(__name__)

images = [
    "https://media1.giphy.com/media/yFQ0ywscgobJK/giphy.webp",
    "https://media0.giphy.com/media/WYEWpk4lRPDq0/200w.webp",
    "https://media0.giphy.com/media/zWuSfeDJkqj0A/200w.webp",
    "https://media2.giphy.com/media/mJqQXx8vK9kD6/200w.webp",
    "https://media2.giphy.com/media/12mWfQYoxRqslq/200.webp",
    "https://media3.giphy.com/media/Vj5ZgHbXa3kWY/200w.webp",
    "https://media0.giphy.com/media/40Fpxgn6Yq640/giphy.webp",
    "https://media3.giphy.com/media/jpPTyP6YghtiU/200.webp",
    "https://media2.giphy.com/media/3oriO0OEd9QIDdllqo/200.webp",
    "https://media0.giphy.com/media/OmK8lulOMQ9XO/200w.webp",
]

@app.route('/')
@app.route('/index')
def index():
    url = random.choice(images)
    return render_template("index.html", url=url)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=int(os.environ.get("PORT", 5000)))
