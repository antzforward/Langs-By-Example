// 将 `deeply::nested::function` 路径绑定到 `other_function`。
// 与c#的use 语法相近的，其中一种写法。C#中use 可使用的地方太多了，还有static use等。
use deeply::nested::function as other_function;

fn function() {
    println!("called `function()`");
}

mod deeply {
    pub mod nested {
        pub fn function() {
            println!("called `deeply::nested::function()`")
        }
    }
}

fn main() {
    // 更容易访问 `deeply::nested::funcion`
    other_function();

    println!("Entering block");
    {
        // 这和 `use deeply::nested::function as function` 等价。
        // 此 `function()` 将遮蔽外部的同名函数。
        use deeply::nested::function;
        function();

        // `use` 绑定拥有局部作用域。在这个例子中，`function()`
        // 的遮蔽只存在这个代码块中。
        println!("Leaving block");
    }

    function();
}
