/*
比较 trait: Eq, PartialEq, Ord, PartialOrd
Clone, 用来从 &T 创建副本 T。
Copy，使类型具有 “复制语义”（copy semantics）而非 “移动语义”（move semantics）。
Hash，从 &T 计算哈希值（hash）。
Default, 创建数据类型的一个空实例。
Debug，使用 {:?} formatter 来格式化一个值。
*/
// `Centimeters`，可以比较的元组结构体
#[derive(PartialEq, PartialOrd)]
struct Centimeters(f64);

// `Inches`，可以打印的元组结构体
#[derive(Debug)]
struct Inches(i32);

impl Inches {
    fn to_centimeters(&self) -> Centimeters {
        let &Inches(inches) = self;

        Centimeters(inches as f64 * 2.54)
    }
}

// `Seconds`，不带附加属性的元组结构体
struct Seconds(i32);

fn main() {
    let _one_second = Seconds(1);

    // 报错：`Seconds` 不能打印；它没有实现 `Debug` trait
    //println!("One second looks like: {:?}", _one_second);
    // 试一试 ^ 取消此行注释

    // 报错：`Seconds`不能比较；它没有实现 `PartialEq` trait
    //let _this_is_true = (_one_second == _one_second);
    // 试一试 ^ 取消此行注释

    let foot = Inches(12);

    println!("One foot equals {:?}", foot);

    let meter = Centimeters(100.0);

    let cmp =
        if foot.to_centimeters() < meter {
            "smaller"
        } else {
            "bigger"
        };

    println!("One foot is {} than one meter.", cmp);
}
//问题，也就是常规能加上attribute的dervie 包括了以上的几种，Debug我用的多一点，其他不多。