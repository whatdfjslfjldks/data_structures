use std::ptr::null_mut;
use std::vec;
use rand::random;

const MAX_LEVEL: i32 = 16; // 允许的最大层级
const P: f64 = 0.5; // 上升一个层级的概率

struct Node {
    val: i32,
    next: *mut Node,
    down: *mut Node, // 指向下一层级的指针
}

struct SkipList {
    head: *mut Node,
    level: i32, // 当前最大层级
}

impl SkipList {
    fn new() -> SkipList {
        SkipList {
            head: Box::into_raw(Box::new(Node {
                val: -1,
                next: null_mut(),
                down: null_mut(),
            })), // 定义头节点为伪节点
            level: 1, // 初始层级
        }
    }

    /// 随机生成元素最大层级
    fn random_level() -> i32 {
        let mut level = 1;
        while random::<f64>() < P && level < MAX_LEVEL {
            level += 1;
        }
        level
    }

    fn insert(&mut self, val: i32) {
        // 生成当前元素的随机最大层级
        let level = Self::random_level();

        // 如果随机层级大于当前最大层级，增加新层级并更新头节点
        if level > self.level {
            for _ in self.level..level {
                let new_head = Box::into_raw(Box::new(Node {
                    val: -1,
                    next: null_mut(),
                    down: self.head,
                }));
                self.head = new_head;
            }
            self.level = level;
        }

        // 维护前驱节点数组
        let mut update = vec![null_mut(); level as usize];
        let mut cur = self.head;

        // 遍历所有层级，找到每一层的前驱节点
        for i in (0..self.level).rev() {
            unsafe {
                while !(*cur).next.is_null() && (*(*cur).next).val < val {
                    cur = (*cur).next;
                }
                // 如果是当前所在最大层级内
                if i < level {
                    update[i as usize] = cur; // 记录前驱节点
                }
                if !(*cur).down.is_null() {
                    cur = (*cur).down;
                }
            }
        }

        let mut down: *mut Node = null_mut();
        // 从底层开始插入新节点
        for i in 0..level {
            let new_node = Box::into_raw(Box::new(Node {
                val,
                next: null_mut(),
                down,
            }));
            unsafe {
                (*new_node).next = (*update[i as usize]).next;
                (*update[i as usize]).next = new_node;
            }
            down = new_node;
        }
    }

    // 查找某个元素是否存在
    fn search(&mut self,val:i32)->bool{
        let mut cur = self.head;
        while !cur.is_null(){
            unsafe {
                while !(*cur).next.is_null() && (*(*cur).next).val < val {
                    cur = (*cur).next;
                }
                if (*(*cur).next).val == val {
                    return true;
                }
                // 到下一个层级
                if !(*cur).down.is_null(){
                    cur=(*cur).down
                }else{
                    break
                }
            }
        }
        false
    }

    fn remove(&mut self, val: i32){
        let mut remove=vec::Vec::new();
        let mut cur = self.head;

        while !cur.is_null(){
            unsafe {
                // 在当前层级搜索
                while !(*cur).next.is_null()&& (*(*cur).next).val<val{
                    cur=(*cur).next
                }
                if !(*cur).next.is_null() && (*(*cur).next).val==val{
                    // 找到要删除的节点
                    remove.push(cur);
                }
                // 向下一个层级搜索
                if !(*cur).down.is_null(){
                    cur=(*cur).down
                }else{
                    break
                }
            }
        }

        // 删除节点
        for i in 0..remove.len(){
            unsafe {
                (*remove[i]).next=(*(*remove[i]).next).next
            }
        }

        // 更新当前最大层级
        unsafe {
            while !self.head.is_null() && (*self.head).next.is_null(){
                self.head=(*self.head).down;
                self.level-=1
            }
        }
    }

    fn print(&self) {
        let mut cur_level = self.head;
        for i in (1..=self.level).rev() {
            print!("第 {} 层: ", i);
            let mut cur = cur_level;
            unsafe {
                while !(*cur).next.is_null() {
                    print!("{} -> ", (*(*cur).next).val);
                    cur = (*cur).next;
                }
                println!("结束");
                cur_level = (*cur_level).down; // 进入下一层
            }
        }
    }
}

fn main() {
    let mut list = SkipList::new();
    list.insert(1);
    list.insert(0);
    list.insert(2);
    list.insert(3);

    list.print();
    list.remove(3);
    list.print()
}
