use pyo3::prelude::*;

#[pyfunction]
fn encrypt_data(data: String) -> PyResult<String> {
    let reversed: String = data.chars().rev().collect();
    Ok(reversed)
}

#[pymodule]
fn rust_crypto(m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(encrypt_data, m)?)?;
    Ok(())
}