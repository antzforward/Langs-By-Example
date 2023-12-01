#[derive(Debug)]
struct A;          // 具体类型 `A`。
#[derive(Debug)]
struct S(A);       // 具体类型 `S`。
#[derive(Debug)]
struct SGen<T>(T); // 泛型类型 `SGen`。

// 下面全部函数都得到了变量的所有权，并立即使之离开作用域，将变量释放。

// 定义一个函数 `reg_fn`，接受一个 `S` 类型的参数 `_s`。
// 因为没有 `<T>` 这样的泛型类型参数，所以这不是泛型函数。
fn reg_fn(_s: S) {
	println!("reg_fn : {:#?}",_s)
}

// 定义一个函数 `gen_spec_t`，接受一个 `SGen<A>` 类型的参数 `_s`。
// `SGen<>` 显式地接受了类型参数 `A`，且在 `gen_spec_t` 中，`A` 没有被用作
// 泛型类型参数，所以函数不是泛型的。
fn gen_spec_t(_s: SGen<A>) {
	println!("gen_spec_t : {:#?}",_s)
}

// 定义一个函数 `gen_spec_i32`，接受一个 `SGen<i32>` 类型的参数 `_s`。
// `SGen<>` 显式地接受了类型参量 `i32`，而 `i32` 是一个具体类型。
// 由于 `i32` 不是一个泛型类型，所以这个函数也不是泛型的。
fn gen_spec_i32(_s: SGen<i32>) {
	println!("gen_spec_i32 : {:#?}",_s)
}

// 定义一个函数 `generic`，接受一个 `SGen<T>` 类型的参数 `_s`。
// 因为 `SGen<T>` 之前有 `<T>`，所以这个函数是关于 `T` 的泛型函数。
fn generic<T:std::fmt::Debug>(_s: SGen<T>) {
	println!("generic<T:std::fmt::Debug> : {:#?}",_s)
}

fn main() {
    // 使用非泛型函数
    reg_fn(S(A));          // 具体类型。
    gen_spec_t(SGen(A));   // 隐式地指定类型参数 `A`。
    gen_spec_i32(SGen(6)); // 隐式地指定类型参数 `i32`。

    // 为 `generic()` 显式地指定类型参数 `char`。
    generic::<char>(SGen('a'));

    // 为 `generic()` 隐式地指定类型参数 `char`。
    generic(SGen('c'));
    //c# 貌似没有这么简单的写法，还是要用T声明类型。
}

//问题1：fn中对T的类型声明 在定义函数时就要确认了。并不是根据使用时再推定。
//问题2：还是适配更方便。相对C#来说。