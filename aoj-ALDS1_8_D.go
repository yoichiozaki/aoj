
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var MIN int = 0

var sc = bufio.NewScanner(os.Stdin)

type Node struct {
    key, pri    int
    left, rigth *Node
}

func next() int {
    sc.Scan()
    i, _ := strconv.Atoi(sc.Text())
    return i

}

func nextString() string {
    sc.Scan()
    return sc.Text()

}

func rightRotate(t *Node) *Node {
    s := t.left
    t.left = s.rigth
    s.rigth = t
    return s
}

func leftRotate(t *Node) *Node {
    s := t.rigth
    t.rigth = s.left
    s.left = t
    return s
}

func insert(t *Node, key int, pri int) *Node {
    if t == nil {
        return &Node{key: key, pri: pri}
    }
    if key == t.key {
        return t
    }
    if key < t.key {

        t.left = insert(t.left, key, pri)
        if t.pri < t.left.pri {
            t = rightRotate(t)
        }
    } else {
        t.rigth = insert(t.rigth, key, pri)
        if t.pri < t.rigth.pri {
            t = leftRotate(t)
        }

    }
    return t
}

func find(t *Node, key int) bool {
    if t == nil {
        return false
    }

    if t.key == key {
        return true
    }

    if t.key > key {
        return find(t.left, key)

    } else if t.key < key {
        return find(t.rigth, key)

    }
    panic("error")
}

func erase(t *Node, key int) *Node {
    if t == nil {
        return nil
    }
    if key == t.key {
        if t.left == nil && t.rigth == nil {
            return nil

        } else if t.left == nil {
            t = leftRotate(t)
        } else if t.rigth == nil {
            t = rightRotate(t)

        } else {
            if t.left.pri > t.rigth.pri {
                t = rightRotate(t)

            } else {

                t = leftRotate(t)

            }

        }
        return erase(t, key)

    }
    if key < t.key {
        t.left = erase(t.left, key)

    } else {
        t.rigth = erase(t.rigth, key)

    }
    return t
}

func printInOrder(t *Node) {
    if t == nil {
        return
    }

    printInOrder(t.left)
    fmt.Printf(" %d", t.key)
    printInOrder(t.rigth)

}

func printPreOrder(t *Node) {
    if t == nil {
        return
    }

    fmt.Printf(" %d", t.key)
    printPreOrder(t.left)
    printPreOrder(t.rigth)

}

func main() {
    sc.Split(bufio.ScanWords)

    var head *Node

    m := next()
    for i := 0; i < m; i++ {

        switch nextString() {

        case "insert":
            k, p := next(), next()
            head = insert(head, k, p)

        case "print":
            printInOrder(head)
            fmt.Println()
            printPreOrder(head)
            fmt.Println()
        case "find":
            k := next()
            if find(head, k) {
                fmt.Println("yes")
            } else {
                fmt.Println("no")
            }
        case "delete":
            k := next()
            head = erase(head, k)

        }

    }
}
