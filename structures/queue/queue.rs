use std::vec;

struct Stack{
    elements:Vec<i32>
}
impl Stack{
    fn new()->Stack{
        Stack{
            elements:vec![]
        }
    }
    // 向栈中推入一个元素
    fn push(&mut self,element:i32){
        self.elements.push(element);
    }
    // 从栈中弹出一个元素
    fn pop(&mut self)->Option<i32>{
        self.elements.pop()
    }
    // 栈是否为空
    fn is_empty(&self)->bool{
        self.elements.is_empty()
    }
}

// 用两个栈模拟队列
struct Queue{
    in_stack:Stack, // 用来模拟入队
    out_stack:Stack // 用来模拟出队
}
impl Queue{
    fn new()->Queue{
        Queue{
            in_stack:Stack::new(),
            out_stack:Stack::new()
        }
    }
    // 入队
    fn enqueue(&mut self,val:i32){
        self.in_stack.push(val)
    }
    // 出队
    fn dequeue(&mut self)->Option<i32>{
        if self.out_stack.is_empty(){
            while !self.in_stack.is_empty(){
                let val = self.in_stack.pop().unwrap(); // 因为rust的pop()返回Option<T>类型，所以需要unwrap()
                self.out_stack.push(val);
            }
        }
        self.out_stack.pop()
    }

}

fn main(){
    let mut queue = Queue::new();
    queue.enqueue(1);
    queue.enqueue(2);
    queue.enqueue(3);
    let a=queue.dequeue();
    println!("{:?}",a)

}