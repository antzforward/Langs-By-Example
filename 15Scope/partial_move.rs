/*在单个变量的解构内，可以同时使用 by-move 和 by-reference 模式绑定。这样做将导致变量的部分移动（partial move），
这意味着变量的某些部分将被移动，而其他部分将保留。在这种情况下，后面不能整体使用父级变量，但是仍然可以使用只引用（而不移动）的部分。
*/
fn main() {
    #[derive(Debug)]
    struct Person {
        name: String,
        age: u8,
    }

    let person = Person {
        name: String::from("Alice"),
        age: 20,
    };

    // `name` 从 person 中移走，但 `age` 只是引用
    let Person { name, ref age } = person;

    println!("The person's age is {}", age);

    println!("The person's name is {}", name);

    // 报错！部分移动值的借用：`person` 部分借用产生
    //println!("The person struct is {:?}", person);

    // `person` 不能使用，但 `person.age` 因为没有被移动而可以继续使用
    println!("The person's age from person struct is {}", person.age);
    //println!("The person's name from person struct is {}", person.name);//person.name丧失了所有权。
    
    //如果name是区域的应该是怎么样的？
    let person = Person {
        name: String::from("Alice"),
        age: 20,
    };
    {
		// `name` 从 person 中移走，但 `age` 只是引用
		let Person { name, ref age } = person;

		println!("The person's age is {}", age);

		println!("The person's name is {}", name);
    }
    println!("The person's age from person struct is {}", person.age);
    //println!("The person's name from person struct is {}", person.name);//还是不能使用，在内部就丢失了所有权，出了之后就永久失去了。
}
//问题1：这确实是个非常强大的功能，强大到我还想不到为什么要部分移动。其中name还不能把所有权转回去。