// 此函数取得堆分配的内存的所有权
fn destroy_box(c: Box<i32>) {
    println!("Destroying a box that contains {}", c);

    // `c` 被销毁且内存得到释放
}
//转移一次试试
fn wrap_box(c:Box<i32>)->Box<i32>{
	println!("wrap_box:{:?}",c);
	c
}
fn main() {
    // 栈分配的整型
    let x = 5u32;

    // 将 `x` *复制*到 `y`——不存在资源移动
    let y = x;

    // 两个值各自都可以使用
    println!("x is {}, and y is {}", x, y);

    // `a` 是一个指向堆分配的整数的指针
    let a = Box::new(5i32);

    println!("a contains: {}", a);

    // *移动* `a` 到 `b`
    let b = a;
    // 把 `a` 的指针地址（而非数据）复制到 `b`。现在两者都指向
    // 同一个堆分配的数据，但是现在是 `b` 拥有它。

    // 报错！`a` 不能访问数据，因为它不再拥有那部分堆上的内存。
    //println!("a contains: {}", a);
    // 试一试 ^ 去掉此行注释

    // 此函数从 `b` 中取得堆分配的内存的所有权
    destroy_box(b);

    // 此时堆内存已经被释放，这个操作会导致解引用已释放的内存，而这是编译器禁止的。
    // 报错！和前面出错的原因一样。
    //println!("b contains: {}", b);
    // 试一试 ^ 去掉此行注释
    
    //尝试问题1
    let a = Box::new(5i32);
    println!("a contains: {}", a);
    
    let mut b = wrap_box(a);
    println!("b contains: {}", b);
    //第二次
    b = wrap_box(b);
    println!("b contains: {}", b);
    //转移了，但没接收
    wrap_box(b);
	//println!("b contains: {}", b);//value borrowed here after move
}
//问题1：move 到 destroy_box，能否用返回值move到其他变量？再move一次？如果可以，上一个操作如果没人assign操作会是啥？
