import requests
import pytest
import time

BASE_URL = "http://localhost:8081" 

def test_go_microservice_heavy_calc():
    params = {"n": 10}
    try:
        response = requests.get(f"{BASE_URL}/fib", params=params, timeout=10)
        
        assert response.status_code == 200
        
        data = response.json()
        assert "result" in data
        
        assert data["result"] == 55
    except requests.exceptions.ConnectionError:
        pytest.fail("Go-микросервис не запущен на порту 8081")

def test_go_microservice_performance():
    start_time = time.time()
    requests.get(f"{BASE_URL}/fib", params={"n": 20})
    duration = time.time() - start_time
    
    assert duration < 2.0