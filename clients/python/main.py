import socket

HOST = "localhost"
PORT = 8080

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    print(f"Connecting to -> {HOST}:{PORT}")

    print("write somethin, or 'EXIT' to close: ")
    while True:
        msg = input("> ")

        if msg == "EXIT":
            break

        s.sendall((msg + "\n").encode("utf-8"))
