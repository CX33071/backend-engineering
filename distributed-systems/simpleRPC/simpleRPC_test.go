package main
import(
	"fmt"
	"testing"
	"6.5840/labrpc"
)
type Args struct {
	A int
	B int
}
type Reply struct {
	Result int
}
type Calculator struct{
}
func (c*Calculator)Add(args*Args,reply *Reply){
	reply.Result=args.A+args.B
}
func TestRPC(t*testing.T){
	server := labrpc.MakeServer()
	ca := new(Calculator)
	server.AddService(labrpc.MakeService(ca))
	network := labrpc.MakeNetwork()
	client:=network.MakeEnd("client")
	network.AddServer("server",server)
	network.Connect("client","server")
	args := Args{
		A:10,
		B:20,
	}
	var reply Reply
	ok:=client.Call("Calculator.Add",&args,&reply)
	if !ok {
		t.Fatal("RPC failed")
	}
	fmt.Println(reply.Result)
}