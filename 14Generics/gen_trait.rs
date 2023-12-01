// 不可复制的类型。
struct Empty;
struct Null;

// `T` 的泛型 trait。
trait DoubleDrop<T> {
    // 定义一个调用者的方法，接受一个额外的参数 `T`，但不对它做任何事。
    fn double_drop(self, _: T);
}

// 对泛型的调用者类型 `U` 和任何泛型类型 `T` 实现 `DoubleDrop<T>` 。
impl<T, U> DoubleDrop<T> for U {
    // 此方法获得两个传入参数的所有权，并释放它们。
    fn double_drop(self, _: T) {}
}

fn main() {
    let empty = Empty;
    let null  = Null;

    // 释放 `empty` 和 `null`。
    empty.double_drop(null);

    //empty
    //null;
    // ^ 试一试：去掉这两行的注释。
}

//问题1：C#中 多个类型模板定义，也是这样的写法吗？ YES
//问题2：DoubleDrop代表啥意思，为什么能有释放的意思？Because self is not  &mut self or &self
//问题3：显然empty所有权丢失了，那么null，也是类似self的写法，接受的参数类型为 T,调用和定义都改成&T，null那行就可以了。