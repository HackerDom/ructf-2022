import requests
import websocket
import json
import random
import time
from math import gcd

URL = "0.0.0.0:8080"

def multiply(left_perm, right_perm):
    new_perm = [0 for _ in range(len(left_perm))]
    for i, element in enumerate(right_perm):
        new_perm[i] = left_perm[element]
    
    return new_perm


def invert(perm):
    new_perm = [0 for _ in range(len(perm))]
    for i in range(len(perm)):
        new_perm[perm[i]] =  i
    
    return new_perm


def exponentiation(perm, power):
    result = [i for i in range(len(perm))]

    if power < 0:
        perm = invert(perm)
        power = -power

    while power > 0:
        if power & 1 == 1:
            result = multiply(result, perm)
        
        perm = multiply(perm, perm)
        power  = power >> 1
    
    return result

def create_game():
    init_perm = [157, 79, 170, 8, 108, 234, 163, 16, 251, 181, 23, 148, 55, 162, 211, 186, 194, 222, 152, 207, 57, 97, 87, 45, 245, 141, 142, 40, 13, 92, 89, 64, 191, 102, 247, 178, 28, 138, 118, 68, 226, 24, 151, 103, 15, 139, 154, 244, 180, 83, 82, 196, 171, 167, 31, 155, 63, 246, 38, 200, 228, 120, 218, 204, 10, 238, 47, 56, 146, 185, 172, 158, 133, 53, 117, 42, 193, 241, 206, 86, 161, 0, 77, 243, 149, 239, 121, 129, 2, 85, 159, 59, 96, 164, 81, 220, 114, 18, 214, 65, 60, 125, 188, 201, 104, 174, 153, 75, 240, 223, 126, 35, 189, 113, 27, 236, 122, 143, 124, 73, 227, 43, 49, 67, 187, 48, 99, 250, 39, 20, 165, 115, 1, 177, 93, 232, 202, 249, 116, 54, 6, 242, 252, 69, 255, 22, 176, 197, 110, 5, 61, 169, 254, 183, 19, 229, 109, 150, 111, 131, 156, 253, 208, 145, 58, 179, 76, 7, 91, 78, 37, 233, 212, 9, 215, 192, 62, 209, 33, 32, 198, 168, 17, 195, 136, 166, 98, 130, 71, 248, 90, 217, 25, 30, 112, 34, 231, 3, 237, 21, 80, 224, 100, 66, 52, 84, 106, 4, 101, 205, 26, 105, 128, 225, 210, 135, 137, 175, 95, 70, 132, 203, 182, 29, 219, 190, 199, 44, 235, 140, 147, 74, 144, 46, 123, 216, 221, 14, 94, 127, 119, 36, 184, 88, 107, 12, 41, 134, 213, 72, 173, 160, 50, 51, 11, 230]
    order = 8933296680
    power = random.randint(1, order - 1)
    while gcd(order, power) != 1:
        power = random.randint(1, order - 1)
    body = {
        "secret": str(random.randint(1000, 893329668)),
        "init": exponentiation(init_perm, power),
        "flag": "FLAGGG" + str(random.randint(1, 10000))
    }

    res = requests.post(f"http://{URL}/create", json=body)
    games = res.json()["id"]

def list_games():
    body = {
        "limit": int(10),
        "offset": int(0)
    }
    res = requests.post(f"http://{URL}/gameList", json=body)
    return res.json()["ids"]

game = None
counter = 0
fields_counter = 0

counters_fields = []

def print_map(game_map):
    for el in game_map:
        print(" ".join(el))

def on_message(ws, message):
    global counters_fields
    global fields_counter
    data = json.loads(message)
    print(data)
    if fields_counter == 29:
        send_data = {"direction": "w",
                    "closeGame": True,
                    "newGame": False}
        fields_counter = 0
    elif data.get("gameResult"):
        fields_counter += 1
        counters_fields.append((data["counter"], data["permutation"]))
        send_data = {"direction": "w",
                    "closeGame": False,
                    "newGame": True}
    else:
        direction = "w"
        print_map(data["gameMap"])
        send_data = {"direction": direction,
                    "closeGame": False,
                    "newGame": False}
    ws.send(str(json.dumps(send_data)))


def on_error(ws, error):
    print(ws)
    print("Error!!:", error)


def on_close(ws, close_status_code, close_msg):
    print(ws, close_status_code, close_msg)
    print("### closed ###")


def on_open(ws):
    global games
    global counter
    ws.send(str(json.dumps({"id": games[1]})))
    counter += 1    


def sploit():   
    global counters_fields
    # for i in list_games():
    counters_fields = []
    ws = websocket.WebSocketApp(f"ws://{URL}/play",
                                    on_message=on_message,
                                    on_error=on_error,
                                    on_close=on_close,
                                    on_open=on_open)
    ws.run_forever()
    print(counters_fields)

# websocket.enableTrace(True)
sploit()
