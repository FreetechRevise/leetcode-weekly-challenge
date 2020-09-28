package week208

type ThroneNode struct{
    IsDead bool 
    Name string 
    childs []*ThroneNode
}

type ThroneInheritance struct {
    Tree *ThroneNode
    FamilyMap map[string]*ThroneNode
}


func Constructor(kingName string) ThroneInheritance {
    t := ThroneInheritance{}
    t.Tree = new(ThroneNode)
    t.Tree.Name = kingName
    t.Tree.childs = make([]*ThroneNode, 0)
    t.FamilyMap = make(map[string]*ThroneNode)
    t.FamilyMap[kingName] = t.Tree
    return t 
}

func (this *ThroneInheritance) Birth(parentName string, childName string)  {
    if address, ok := this.FamilyMap[parentName]; ok {
        newNode := new(ThroneNode)
        newNode.Name = childName
        newNode.childs = make([]*ThroneNode, 0)
        this.FamilyMap[childName] = newNode
        address.childs = append(address.childs, newNode)
    }
}


func (this *ThroneInheritance) Death(name string)  {
    if node, ok := this.FamilyMap[name]; ok {
        node.IsDead = true 
    }
}

func (node *ThroneNode) getInheritanceOrder(result *[]string){
    if !node.IsDead{
        *result = append(*result, node.Name)
    }
    if len(node.childs) == 0 {
        return 
    }
    for _, child := range node.childs{
        child.getInheritanceOrder(result)
    }
}

func (this *ThroneInheritance) getInheritanceOrder(result *[]string) {
    this.Tree.getInheritanceOrder(result)
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
    result := make([]string, 0)
    this.getInheritanceOrder(&result)
    return result 
}



/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */