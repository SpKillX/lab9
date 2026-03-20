import requests

response = requests.get("http://localhost:8081/fib", params={"n": 30})
print(f"Фибоначчи (Go): {response.json()['result']}")