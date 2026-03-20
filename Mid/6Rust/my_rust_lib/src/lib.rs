pub fn add_numbers(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_positive_numbers() {
        assert_eq!(add_numbers(10, 20), 30);
    }

    #[test]
    fn test_negative_numbers() {
        assert_eq!(add_numbers(-5, -5), -10);
    }

    #[test]
    fn test_zero() {
        assert_eq!(add_numbers(0, 100), 100);
    }
}