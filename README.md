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



