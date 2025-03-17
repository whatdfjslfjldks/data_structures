use std::ptr:: null_mut;

struct Node {
    val:i32,
    next:*mut Node
}
struct NodeList{
    head:*mut Node
}

impl NodeList{
    fn new()->NodeList{
        NodeList{
            head:null_mut()
        }
    }

    // insert 插入功能
    fn insert(&mut self,val:i32){
        let node=Box::into_raw(Box::new(Node{
            val,
            next:null_mut()
        }));
        // 判断头节点是否为空，如果为空，新节点就是头节点
        if self.head.is_null(){
            self.head=node;
            return;
        }
        let mut cur=self.head;
        // unsafe -> trust me
        unsafe {
            while !(*cur).next.is_null(){
                cur=(*cur).next;
            }
            (*cur).next=node;
        }
    }

    // remove 删除链表中的一个元素
    fn remove(&mut self,val:i32)->bool{
        if self.head.is_null(){
            return false;
        };
        unsafe {
            if (*self.head).val==val{
                let drop_node=self.head;
                self.head=(*self.head).next;
                drop(Box::from_raw(drop_node));
                return true;
            };
        };
        let mut cur=self.head;
        unsafe {
        while !(*cur).next.is_null(){
                if (*(*cur).next).val==val{
                    let drop_node=(*cur).next;
                    (*cur).next=(*(*cur).next).next;
                    drop(Box::from_raw(drop_node));
                    return true
                }
                cur=(*cur).next;
            }
        }
        false
    }

    // search 查找链表中的某个元素是否存在
    fn search(&self,val:i32)->bool{
        let mut cur=self.head;
        while !cur.is_null(){
            unsafe {
                if (*cur).val==val{
                    return true;
                }
                cur=(*cur).next;
            }
        }
        false
    }

    // print 打印链表
    fn print(&self){
        let mut cur=self.head;
        while !cur.is_null(){
            unsafe {
                println!("{}",(*cur).val);
                cur=(*cur).next;
            }
        }
    }
}

impl Drop for NodeList {
    fn drop(&mut self) {
        while self.remove_first().is_some() {}
    }
}

fn main() {
    let mut list = NodeList::new();
    list.insert(1);
    list.insert(2);
    list.insert(3);

    list.remove(1);

    assert_eq!(list.search(2), true);

    list.print();
}