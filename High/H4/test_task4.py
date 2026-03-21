import pytest
import requests
from main_app import process_orchestrator_data

def test_full_integration_flow():
    try:
        result = process_orchestrator_data()
        
        assert result['id'] == 101
        assert len(result['hash']) > 0
        assert isinstance(result['hash'], str)
        
    except requests.exceptions.ConnectionError:
        pytest.fail("Оркестратор Go не запущен на порту 8082")

def test_rust_module_directly():
    import rust_crypto
    data = "test_data"
    hash1 = rust_crypto.encrypt_data(data)
    hash2 = rust_crypto.encrypt_data(data)
    
    assert hash1 == hash2
    assert hash1 != data