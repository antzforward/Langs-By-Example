// raii.rs
fn create_box() {
    // 在堆上分配一个整型数据
    let _box1 = Box::new(3i32);
	println!("create_box ");
    // `_box1` 在这里被销毁，内存得到释放
}

fn main() {
    // 在堆上分配一个整型数据
    let _box2 = Box::new(5i32);

    // 嵌套作用域：
    {
        // 在堆上分配一个整型数据
        let _box3 = Box::new(4i32);

        // `_box3` 在这里被销毁，内存得到释放
    }

    // 创建一大堆 box（只是因为好玩）。
    // 完全不需要手动释放内存！
    for _ in 0u32..1_000 {
        create_box();
    }

    // `_box2` 在这里被销毁，内存得到释放
}
//问题1：这里的Box::new 是不是生成的是局部的指针？类似我们在CPP中将指针用struct 包裹的形式。这个我之前见过，但是不记得写法了。或者C#的IDispose的写法？