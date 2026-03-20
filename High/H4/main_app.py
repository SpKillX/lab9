import requests
import rust_crypto

def process_orchestrator_data():
    response = requests.get("http://localhost:8082/next-task")
    task = response.json()
    
    result_hash = rust_crypto.secure_hash(task['payload'])
    
    return {"id": task['id'], "hash": result_hash}

if __name__ == "__main__":
    print(process_orchestrator_data())