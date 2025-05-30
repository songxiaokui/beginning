from flask import Flask, request, jsonify
from flask_cors import CORS
import sqlite3
from datetime import date

app = Flask(__name__)
CORS(app)

DB_NAME = './db/checkin.db'


def init_db():
    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    c.execute('''
        CREATE TABLE IF NOT EXISTS checkin (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            date TEXT UNIQUE,
            breakfast INTEGER,
            lunch INTEGER,
            dinner INTEGER,
            exercise INTEGER,
            sleep INTEGER
        )
    ''')
    conn.commit()
    conn.close()


@app.route('/api/checkin', methods=['GET'])
def get_checkin():
    q_date = request.args.get('date', date.today().isoformat())
    conn = sqlite3.connect(DB_NAME)
    c = conn.cursor()
    c.execute('SELECT breakfast, lunch, dinner, exercise, sleep FROM checkin WHERE date = ?', (q_date,))
    row = c.fetchone()
    print(row)
    conn.close()
    if row:
        return jsonify({
            'breakfast': bool(row[0]),
            'lunch': bool(row[1]),
            'dinner': bool(row[2]),
            'exercise': bool(row[3]),
            'sleep': bool(row[4])
        })
    return jsonify({
        'breakfast': False,
        'lunch': False,
        'dinner': False,
        'exercise': False,
        'sleep': False
    }), 200


@app.route('/api/checkin', methods=['POST'])
def post_checkin():
    data = request.get_json()
    q_date = data.get('data', date.today().isoformat())
    config = data.get("config")
    record = (
        q_date,
        int(config.get('breakfast', 0)),
        int(config.get('lunch', 0)),
        int(config.get('dinner', 0)),
        int(config.get('exercise', 0)),
        int(config.get('sleep', 0))
    )
    conn = sqlite3.connect(DB_NAME)
    print(conn)
    c = conn.cursor()
    c.execute('''
        INSERT INTO checkin (date, breakfast, lunch, dinner, exercise, sleep)
        VALUES (?, ?, ?, ?, ?, ?)
        ON CONFLICT(date) DO UPDATE SET
            breakfast=excluded.breakfast,
            lunch=excluded.lunch,
            dinner=excluded.dinner,
            exercise=excluded.exercise,
            sleep=excluded.sleep
    ''', record)
    conn.commit()
    conn.close()
    return jsonify({'message': 'Check-in saved'})


if __name__ == '__main__':
    init_db()
    app.run(debug=True, port=5011, host="0.0.0.0")
