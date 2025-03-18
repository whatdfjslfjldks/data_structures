
#[derive(PartialEq,Clone)]
struct Node{
    val:i32,
    next:Option<Box<Node>>
}
struct NodeList{
    head:Option<Box<Node>>
}
// 循环链表，定义头节点为哨兵节点
impl NodeList{
    fn new()->NodeList{
        let mut sentinel=Box::new(Node{
            val:-1,
            next:None
        });
        sentinel.next=Some(sentinel.clone());
        NodeList{
            head:Some(sentinel)
        }
    }

    fn insert(&mut self, val: i32) {
        let mut cur = &mut self.head;
        // 遍历链表，找到合适的插入位置
        while let Some(node) = cur {
            if let Some(next_node) = &node.next {
                // 如果下一个节点的值小于要插入的值且不是哨兵节点，继续遍历
                if next_node.val < val && next_node.val != -1 {
                    cur = &mut node.next;
                    continue;
                }
            }
            // 找到合适的插入位置，创建新节点
            let new_node = Box::new(Node {
                val,
                next: node.next.take(),
            });
            // 将新节点插入到当前节点之后
            node.next = Some(new_node);
            return;
        }
    }
    fn print(&self){
        let mut cur=&self.head;
        while let Some(node)=cur{
            if node.next==self.head{
                break;
            }
            println!("{}",node.val);
            cur=&node.next;
        }
    }

}


fn main(){
    let mut list=NodeList::new();
    list.insert(1);
    list.insert(2);
    list.insert(3);
    list.print();

}