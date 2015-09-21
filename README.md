THE SB6502 VIRTUAL CHIP
=============================

create a new cpu with `new cpu`  
create a new program with `new prog "<instructions>"` - see below for instruction set

See whats loaded with `ps`  
Load Program 0 into CPU 0 with `cpu 0 load 0`  
Run CPU with `cpu 0 run`  

Instruction Set - 16 of them:::  
First 8 are 1 byte instructions:  
0 = Halt  
1 = Add (R0 = R0 + R1)  
2 = Subtract (R0 = R0 – R1)  
3 = Increment R0 (R0 = R0 + 1)  
4 = Increment R1 (R1 = R1 + 1)  
5 = Decrement R0 (R0 = R0 – 1)  
6 = Decrement R1 (R1 = R1 – 1)  
7 = Print Value in R0  
Last 7 are 2 byte instructions (second byte is \<data\>):  
8 = Load \<data\> into R0  
9 = Load \<data\> into R1  
10 = Store R0 into address \<data\>  
11 = Store R1 into address \<data\>  
12 = Store val from address \<data\> in R0  
13 = Jump to address \<data\>  
14 = Jump to address \<data\> if R0 != 0  
15 = Jump to address \<data\> if R0 == 0


Example programs:  
`new prog "8 10 7 5 14 2"`  

which is read as:  
8 10    // load val '10' into R0  
7       // print R0  
5       // decrement R0  
14 2    // go to instruction 2 if R0 != 0  

i.e. it's a loop from 10..1 

more advanced: to go from 1..10, you need to store values in memory too:  
`new prog "8 10 10 31 8 1 10 30 12 30 7 3 10 30 12 31 5 10 31 14 8"`



