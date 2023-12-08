use std::num::ParseIntError;

// 修改了上一节中的返回类型，现在使用模式匹配而不是 `unwrap()`。
fn multiply(first_number_str: &str, second_number_str: &str) -> Result<i32, ParseIntError> {
    match first_number_str.parse::<i32>() {
        Ok(first_number)  => {
            match second_number_str.parse::<i32>() {
                Ok(second_number)  => {
                    Ok(first_number * second_number)
                },
                Err(e) => Err(e),
            }
        },
        Err(e) => Err(e),
    }
}

// 就像 `Option` 那样，我们可以使用 `map()` 之类的组合算子。
// 除去写法外，这个函数与上面那个完全一致，它的作用是：
// 如果值是合法的，计算其乘积，否则返回错误。
fn multiply2(first_number_str: &str, second_number_str: &str) -> Result<i32, ParseIntError> {
    first_number_str.parse::<i32>().and_then(|first_number| {
        second_number_str.parse::<i32>().map(|second_number| first_number * second_number)
    })
}

fn print(result: Result<i32, ParseIntError>) {
    match result {
        Ok(n)  => println!("n is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

fn main() {
    // 这种情形下仍然会给出正确的答案。
    let twenty = multiply("10", "2");
    print(twenty);

    // 这种情况下就会提供一条更有用的错误信息。
    let tt = multiply("t", "2");
    print(tt);
    
     // 这种情形下仍然会给出正确的答案。
    let twenty = multiply2("10", "2");
    print(twenty);

    // 这种情况下就会提供一条更有用的错误信息。
    let tt = multiply2("t", "2");
    print(tt);
}
