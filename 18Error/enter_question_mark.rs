use std::num::ParseIntError;

fn multiply(first_number_str: &str, second_number_str: &str) -> Result<i32, ParseIntError> {
    //let first_number = first_number_str.parse::<i32>()?;
    //let second_number = second_number_str.parse::<i32>()?;
    let first_number = try!(first_number_str.parse::<i32>());
    let second_number = try!(second_number_str.parse::<i32>());

    Ok(first_number * second_number)
}

fn print(result: Result<i32, ParseIntError>) {
    match result {
        Ok(n)  => println!("n is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

fn main() {
    print(multiply("10", "2"));
    print(multiply("t", "2"));
}
//问题1：C#上用？。或者？？更多，需要仔细知道为什么啊。
//问题2：带？的含义是啥？与early return相反。
