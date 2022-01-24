func isValid(s string) bool {
    stack := NewStack()
    for _, v := range s {
        if v == ')' && stack.Peak() == '(' {
            stack.Pop()
        } else if v == '}' && stack.Peak() == '{' {
            stack.Pop()
        } else if v == ']' && stack.Peak() == '[' {
            stack.Pop()
        } else {
            stack.Push(v)
        }
    }
    return stack.Empty()
}

type Stack struct {
    list *list.List 
    lock *sync.RWMutex
}

func NewStack() *Stack {
    list := list.New()
    l := &sync.RWMutex{}
    return &Stack{list, l}
}

func (stack *Stack) Push(val interface{}) {
    stack.lock.Lock()
    defer stack.lock.Unlock()
    stack.list.PushBack(val)
}

func (stack *Stack) Pop() interface{} {
    stack.lock.Lock()
    defer stack.lock.Unlock()
    e := stack.list.Back()
    if e != nil {
        stack.list.Remove(e)
        return e.Value
    }
    return nil
}

func (stack *Stack) Peak() interface{} {
    e := stack.list.Back()
    if e != nil {
        return e.Value
    }
    return nil
}

func (stack *Stack) Len() int {
    return stack.list.Len()
}

func (stack *Stack) Empty() bool {
    return stack.list.Len() == 0
}
