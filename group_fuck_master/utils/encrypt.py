from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.backends import default_backend
from base64 import b64encode
import os


def encrypt_data(plain_text: str, aes_key: bytes) -> str:
    nonce = os.urandom(12)
    cipher = Cipher(algorithms.AES(aes_key), modes.GCM(nonce), backend=default_backend())
    encryptor = cipher.encryptor()
    cipher_text = encryptor.update(plain_text.encode()) + encryptor.finalize()
    tag = encryptor.tag

    # 将 nonce、tag 和密文组合并进行 base64 编码
    encrypted_data = nonce + cipher_text + tag
    inner_base64_data = b64encode(encrypted_data).decode()
    outer_base64_data = b64encode(inner_base64_data.encode()).decode()
    print(outer_base64_data)
    return outer_base64_data
