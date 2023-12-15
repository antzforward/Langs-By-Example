# Rust By Example

## 学习Rust的过程
## 学习过程
### 01 HelloWorld
println!是一个宏（macro），可以将文本输出到控制台（console）
1. xxx!()表示宏的调用
2. print+ln，通常表示一行输出，如果混合输出就用print！
3. 其他的宏，format！输出到字符串，eprint！类似print！但是将文本输出到标准错误（io::stderr)

print的格式化说明，此处与C#的格式非常类似
注意这几个例子
1. println!("{} days", 31);//{}表示被替换的区域。
2. println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");//{0}，{1}出现多次的情况，序号代表参数
3. println!("{subject} {verb} {object}", object="the lazy dog",subject="the quick brown fox",verb="jumps over"); 这种情况，比较近似的方式C#中的interpolated string的方式。有一点的区别，println!是宏，里面的临时变量是可以在里面声明的。

### 02 Primitives
## 总结
