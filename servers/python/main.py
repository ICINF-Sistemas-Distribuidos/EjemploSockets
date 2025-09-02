import socket

HOST = "0.0.0.0"
PORT = 8081

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.bind((HOST, PORT))
    s.listen()
    print(f"Listening on: {HOST}:{PORT}...")

    conn, addr = s.accept()
    with conn:
        print(f"Connected: {addr}")
        while True:
            data = conn.recv(1024)
            if not data:
                break
            print(f"Received: {data.decode()}")
            text = ">> " + data.decode()
            conn.sendall(text.encode("utf-8"))
