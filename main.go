package main

import (
	"fmt"
	"fundamentals-of-go-/learning"
)

func main() {

	/*	learning.Const()
		learning.Variable()
		learning.Values()
		learning.Forloop()
		learning.Ifelse()
		learning.Switch()
		learning.Array()
		learning.Slice()
		learning.Map()
	*/
	/*
		fmt.Println(learning.Add(2, 4))
		fmt.Println(learning.Addm(2, 4, 9))
		a, b := learning.Vals()
		fmt.Println(a, b)
		_, c := learning.Vals()
		fmt.Println(c)
	*/

	/*
		learning.Varfunc(2, 4)
		learning.Varfunc(2, 3, 4, 5)

		nums := []int{1, 2, 3, 4}
		learning.Varfunc(nums...)
	*/

	/*
		va := learning.Clouser()
		fmt.Println(va())
		fmt.Println(va())
		fmt.Println(va())
		va1 := learning.Clouser()
		fmt.Println(va1())
	*/
	/*
		fmt.Println(learning.Recur(7))
		var fib func(n int) int

		fib = func(n int) int {
			if n < 2 {
				return n
			}
			return fib(n-1) + fib(n-2)

		}
		fmt.Println(fib(7))
	*/

	//learning.Rangeovertypes()

	/*i := 1
	learning.Zeroval(i)
	fmt.Println(i)
	learning.Zeropoint(&i)
	fmt.Println(i)
	fmt.Println(&i)*/

	//	learning.Runesbytesstrings()

	//learning.Struc()
	//learning.Result()

	/* CODE FOR INTERFACES
	rectan := learning.Rectangle{Width: 6, Height: 7}
	circ := learning.Circle{Radius: 5}
	shapes := []learning.Shape{rectan, circ}

	for _, Shape := range shapes {
		// by creating a slice we can call it like this
		fmt.Println(Shape.Area())
		// by craeting a func in interface class we can call like this
		fmt.Println(learning.Calculate(Shape))
	}
	Periofrectan := learning.Rectangle{Width: 6, Height: 7}
	periofcirc := learning.Circle{Radius: 5}
	shapes1 := []learning.Shape2{Periofrectan, periofcirc}

	for _, Shape2 := range shapes1 {
		// by creating a slice we can call it like this

		fmt.Println((Shape2.Perim()))
		// by craeting a func in interface class we can call like this

		fmt.Println(learning.Calculate1(Shape2))
	}*/

	/* Emuns
	Analyzer := learning.PacketAnalyzer{}

	for i := 0; i < 10; i++ {
		packet := learning.PacketGenerator()
		Status := Analyzer.Analyze(packet)
		fmt.Printf("Packet #%d, Protocol %s, SourceIP %s, DestinationIP %s, Size %d bytes, Status %s\n  ",
			packet.ID, packet.Protocol, packet.SourceIP, packet.Destinationip, packet.Size, Status)
	}

	fmt.Println("Total Packets\n", Analyzer.TotalPackets)
	fmt.Println("Accepted Packets\n", Analyzer.AcceptedPackets)
	fmt.Println("Rejected Packets\n", Analyzer.RejectedPackets)
	fmt.Println("Suspicious Packets\n", Analyzer.Suspiciouspackets)*/
	/*
		player := learning.Newplayer()
		player.Move(6.8, 9.9)
		fmt.Println(player.Postion)
		player.Teleport(6.8, 9.9)
		fmt.Println(player.Postion)
		player.MoveSpecial(6.8, 9.9)
		fmt.Println(player.Postion)

		zombie := learning.NewEnemy()
		zombie.Move(9, 7)
		fmt.Println(zombie.Postion)
		zombie.Teleport(9, 7)
		fmt.Println(zombie.Postion)
		zombie.MoveSpecial(9, 7)
		fmt.Println(zombie.Postion) */
	//Generics
	//simple FUnctio

	fmt.Println(learning.HAs([]string{"a", "b", "c"}, "c"))
	// now we use generic function when we call it their we specify the type of list
	fmt.Println(learning.HAS([]string{"a", "b", "c"}, "c"))
	fmt.Println(learning.HAS([]string{"a", "b", "c"}, "d"))
	fmt.Println(learning.HAS([]int{1, 2, 3}, 3))
	fmt.Println(learning.HAS([]int{1, 2, 3}, 4))

	learning.Typeperimeters(2, 2, 2, 3)
	learning.Typeperimeters("2", "2", 2, 3)
	learning.Typeperimeters(2, 2, "b", 3)
	// method call can be use for strings
	intcheck := learning.Sets(1, 2, 3, 4)
	fmt.Println(intcheck.Check(7))

}
