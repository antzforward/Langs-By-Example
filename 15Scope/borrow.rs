/*Rust 使用了借用（borrowing）机制。对象可以通过引用（&T）来传递，从而取代通过值（T）来传递。
编译器（通过借用检查）静态地保证了引用总是指向有效的对象。也就是说，当存在引用指向一个对象时，该对象不能被销毁。
*/
// 此函数取得一个 box 的所有权并销毁它
fn eat_box_i32(boxed_i32: Box<i32>) {
    println!("Destroying box that contains {}", boxed_i32);
}

// 此函数借用了一个 i32 类型
fn borrow_i32(borrowed_i32: &i32) {
    println!("This int is: {}", borrowed_i32);
}

fn main() {
    // 创建一个装箱的 i32 类型，以及一个存在栈中的 i32 类型。
    let boxed_i32 = Box::new(5_i32);
    let stacked_i32 = 6_i32;

    // 借用了 box 的内容，但没有取得所有权，所以 box 的内容之后可以再次借用。
    // 译注：请注意函数自身就是一个作用域，因此下面两个函数运行完成以后，
    // 在函数中临时创建的引用也就不复存在了。
    borrow_i32(&boxed_i32);
    borrow_i32(&stacked_i32);

    {
        // 取得一个对 box 中数据的引用
        let _ref_to_i32: &i32 = &boxed_i32;
        // 报错！
        // 当 `boxed_i32` 里面的值之后在作用域中被借用时，不能将其销毁。
        //eat_box_i32(boxed_i32);
        // 改正 ^ 注释掉此行

        // 在 `_ref_to_i32` 里面的值被销毁后，尝试借用 `_ref_to_i32`
        //（译注：如果此处不借用，则在上一行的代码中，eat_box_i32(boxed_i32)可以将 `boxed_i32` 销毁。）
        borrow_i32(_ref_to_i32);
        // `_ref_to_i32` 离开作用域且不再被借用。
    }

    // `boxed_i32` 现在可以将所有权交给 `eat_i32` 并被销毁。
    //（译注：能够销毁是因为已经不存在对 `boxed_i32` 的引用）
    eat_box_i32(boxed_i32);
}
//问题1：这里的borrow实际上跟接近ref，也就是它的引用的生命周期只要小于等于本体就行。引用要出现在本体之后，也要遵循RAII 资源获取即初始化。
//问题2：borrow和move都与mut与否无关，含义不同。只有const是imut的其他的只要你接用，转移给可修改的变量或者引用，就可以修改了。
//问题3：borrow也只能一次被借入给一个吧。错，borrow可以借给多个。那能否修改啊？应该可以吧。