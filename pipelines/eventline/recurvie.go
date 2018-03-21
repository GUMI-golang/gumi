package eventline

import "github.com/GUMI-golang/gumi"

func recurviePreOccurEvent(node Node, event *gumi.Event){
	*event = node.PreHandle(*event)
	if *event == nil{
		return
	}
	for _, v := range node.GetChildrun(){
		recurviePreOccurEvent(v, event)
	}
}
func recurvieOccurEvent(node Node, event gumi.Event) {
	node.Handle(event)
	for _, v := range node.GetChildrun(){
		recurvieOccurEvent(v, event)
	}
}
