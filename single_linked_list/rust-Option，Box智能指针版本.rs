struct Node {
    val: i32,
    next: Option<Box<Node>>,
}
struct NodeList {
    head: Option<Box<Node>>,
}
impl NodeList {
    fn new() -> NodeList {
        NodeList {
            // 定义一个伪头节点
            head: Some(Box::new(Node {
                val: -1,
                next: None,
            })),
        }
    }
    fn insert(&mut self, val: i32) {
        let mut cur = &mut self.head;
        while let Some(node) = cur {
            if node.next.is_none() || node.next.as_ref().unwrap().val >= val {
                let new_node = Box::new(Node {
                    val,
                    // .take()是Option的方法，会取走Option里面的值，并将原Option设置为None
                    next: node.next.take(),
                });
                node.next = Some(new_node);
                return;
            }
            cur = &mut node.next;
        }
    }
    fn remove(&mut self, val: i32){
        let mut cur = &mut self.head;
        while let Some(node)=cur{
            if let Some(node_next)=&mut node.next{
                if node_next.val == val{
                    node.next = node_next.next.take();
                    return;
                }
            }
            cur = &mut node.next;
        }
    }
    fn search(&mut self,val:i32)->bool{
        let mut cur = &mut self.head;
        while let Some(node) = cur {
            if node.val == val{
                return true;
            }
            cur = &mut node.next;
        }
        false
    }
    fn print(&self) {
        let mut cur = &self.head;
        while let Some(node) = cur {
            println!("{}", node.val);
            cur = &node.next;
        }
    }
}
fn main() {
    let mut list = NodeList::new();
    list.insert(1);
    list.insert(2);
    list.remove(1);
    list.print();
}