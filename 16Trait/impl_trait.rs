use std::iter;
use std::vec::IntoIter;

// 该函数组合了两个 `Vec <i32>` 并在其上返回一个迭代器。
// 看看它的返回类型多么复杂！
fn combine_vecs_explicit_return_type(
    v: Vec<i32>,
    u: Vec<i32>,
) -> iter::Cycle<iter::Chain<IntoIter<i32>, IntoIter<i32>>> {
    v.into_iter().chain(u.into_iter()).cycle()
}

// 这是完全相同的函数，但其返回类型使用 `impl Trait`。
// 看看它多么简单！
fn combine_vecs(
    v: Vec<i32>,
    u: Vec<i32>,
) -> impl Iterator<Item=i32> {
    v.into_iter().chain(u.into_iter()).cycle()
}

// 返回一个将输入和 `y` 相加的函数
fn make_adder_function(y: i32) -> impl Fn(i32) -> i32 {
    let closure = move |x: i32| { x + y };
    closure
}

fn double_positives<'a>(numbers: &'a Vec<i32>) -> impl Iterator<Item = i32> + 'a {
    numbers
        .iter()
        .filter(|x| x > &&0)
        .map(|x| x * 2)
}


fn main() {
    let v1 = vec![1, 2, 3];
    let v2 = vec![4, 5];
    let mut v3 = combine_vecs(v1, v2);
    assert_eq!(Some(1), v3.next());
    assert_eq!(Some(2), v3.next());
    assert_eq!(Some(3), v3.next());
    assert_eq!(Some(4), v3.next());
    assert_eq!(Some(5), v3.next());
    let plus_one = make_adder_function(1);
    assert_eq!(plus_one(2), 3);
    println!("all done");
    
    let v1 = vec![1, 2, 3];
    let mut v3 = double_positives(&v1);
    println!("Some(1) : {:?}",v3.next());
    println!("Some(2) : {:?}",v3.next());
    println!("Some(3) : {:?}",v3.next());
    println!("Some(4) : {:?}",v3.next());
}
