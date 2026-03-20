use pyo3::prelude::*;

#[pyfunction]
fn secure_hash(data: String) -> PyResult<String> {
    let hashed = format!("{:x}", md5::compute(data.as_bytes())); 
    Ok(hashed)
}

#[pymodule]
fn rust_crypto(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(secure_hash, m)?)?;
    Ok(())
}