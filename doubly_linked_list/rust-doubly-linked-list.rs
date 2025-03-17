use std::ptr::null_mut;

struct Node{
    val:i32,
    prev:*mut Node,
    next:*mut Node
}
struct NodeList{
    head:*mut Node
}

impl NodeList{
    fn new()->NodeList{
        NodeList{
            // 创建一个伪节点作为头节点
            head:Box::into_raw(Box::new(Node{
                val:-1,
                prev:null_mut(),
                next:null_mut()
            }))
        }
    }

    fn insert(&mut self,val:i32){
        let mut cur = self.head;
        unsafe {
            // 遍历找到插入位置的前驱节点
            while !(*cur).next.is_null() && (*(*cur).next).val < val{
                cur = (*cur).next;
            }
            let new_node = Box::into_raw(Box::new(Node{
                val,
                prev:cur,
                next:(*cur).next
            }));
            if !(*cur).next.is_null(){
                (*(*cur).next).prev=new_node;
            }
            (*cur).next=new_node;
        }
    }

    fn search(&mut self,val:i32)->bool{
        unsafe {
            let mut cur = self.head;
            while !(*cur).next.is_null() && (*(*cur).next).val < val{
                cur = (*cur).next;
            }
            if (*(*cur).next).val == val{
                return true;
            }
            false
        }
    }

    fn remove(&mut self,val:i32){
        unsafe {
            let mut cur = self.head;
            while !(*cur).next.is_null() && (*(*cur).next).val < val{
                cur = (*cur).next;
            }
            if !(*cur).next.is_null() && (*(*cur).next).val == val{
                let next = (*cur).next;
                (*cur).next = (*next).next;
                if !(*next).next.is_null(){
                    (*(*next).next).prev = cur;
                }
                drop(Box::from_raw(next));
            }
        }
    }

    fn print(&mut self){
        let mut cur = self.head;
        unsafe {
            while !cur.is_null(){
                print!("{}->",(*cur).val);
                cur = (*cur).next;
            }
            println!("结束");
        }
    }
}

fn main() {
    let mut list = NodeList::new();
    list.insert(1);
    list.insert(2);
    list.insert(3);
    list.insert(5);
    list.insert(4);
    list.remove(5);
    list.print();
}