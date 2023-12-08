trait Person {
    fn name(&self) -> String;
}

// Person 是 Student 的父 trait。
// 实现 Student 需要你也 impl 了 Person。
trait Student: Person {
    fn university(&self) -> String;
}

trait Programmer {
    fn fav_language(&self) -> String;
}

// CompSciStudent (computer science student，计算机科学的学生) 是 Programmer 和 Student 两者的子类。
// 实现 CompSciStudent 需要你同时 impl 了两个父 trait。
trait CompSciStudent: Programmer + Student {
    fn git_username(&self) -> String;
}

fn comp_sci_student_greeting(student: &dyn CompSciStudent) -> String {
    format!(
        "My name is {} and I attend {}. My favorite language is {}. My Git username is {}",
        student.name(),
        student.university(),
        student.fav_language(),
        student.git_username()
    )
}

struct TheOne{}

impl Person for TheOne {
	fn name(&self) -> String{
		"Peter Liu".to_string()
	}
}
impl Student for TheOne {
	fn university(&self) -> String{
		"UESTC".to_string()
	}
}
impl Programmer for TheOne {
	fn fav_language(&self) -> String{
		"C++".to_string()
	}
}
impl CompSciStudent for TheOne {
	fn git_username(&self) -> String{
		"antzforward".to_string()
	}
}
fn main() {
	let one = TheOne{};
	println!("greeting by {} ",comp_sci_student_greeting(&one));
}

//问题1：实现过程缺不能只实现继承的trait，只能分段实现。