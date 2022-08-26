# Welcome in the Erik Operating System
Here there will be only the source folder of the operating system and some information about it.

## Numerics
### Binary utility file

Now look, there will be information about the hard drive where the operating system will be set. Here are 0 bytes:
```
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
[contain 0 bytes]
```
Of course it is easy for 0 bytes to all contain zeros in the binary.

The file will be created as a `utility.bin` that will show the direct reports (the direct reports modify the attributes for once, the indirect ones instead keep them by sorting it by date)

Example for indirect file used from `24 Jan 2001 - 25 Jan 2001`:
```
24/01/2001
00 00 00 01 00 00 00 00 00 00 00 00 00 01 10 
00 10 01 10 00 10 00 10 10 00 10 10 10 00 01 
00 00 00 10 01 11 11 01 10 10 00 01 01 00 10 
25/01/2001
00 00 00 01 00 00 00 00 00 00 00 00 00 01 10 
00 10 01 10 00 10 00 10 10 00 10 10 10 00 01 
00 00 00 10 01 10 00 10 10 00 10 01 00 00 10
```

Example for direct file used from `24 Jan 2001 - 25 Jan 2001`:

Modified by 24 Jan 2001: 
`utility.bin`
```
00 00 00 01 00 00 00 00 00 00 00 00 00 01 10 
00 10 01 10 00 10 00 10 10 00 10 10 10 00 01 
00 00 00 10 01 11 11 01 10 10 00 01 01 00 10 
```

Modified by 25 Jan 2001: 
`utility.bin`
```
00 00 00 01 00 00 00 00 00 00 00 00 00 01 10 
00 10 01 10 00 10 00 10 10 00 10 10 10 00 01 
00 00 00 10 01 10 00 10 10 00 10 01 00 00 10
```

Code to edit file:
```
when i receive utilfile {
   show.plot [exporting]
   show.plot [edit]
   key press holderplace = "any key"
   const variable;
   edit variable = switch degress to +1
}
```
We talked about the work on the hard drive, there are binary numbers. Called "controllers", but strange, the previous example was too small for a large storage space. The controllers, when they are too small, the disk will be reported in partition as invalid. depending on the binaries, there are also controllers that support multiple capacities: octal. Everything comes from the track, I say well? In fact, the track has only 2 controllers, but the hexadecimal has as many as 8. 

### Octal controllers
In the previous chapter, we talked about binary numaries, which have controllers. The binary has 2, the octal has 8. We're still focused on the binary 🛤️ for a while.
The track has 2 controllers: 0-1. Instead the octal has a scale from 0 to 7: 0-1-2-3-4-5-6-7. The octal only supports these numbers, but not as many as 8 and 9. Otherwise it would be the decimal scale, which will reign over the next chapter. Now let's move on to the octal. An octal sequence would look like this:
```
0 1 2 3 4 5 6 7
1 2 3 4 5 6 7 0
2 3 4 5 6 7 0 1
```
The example shows that, when the next number of 7, is 0, not 8. Reviewing the octal number scale ...

` 0 1 2 3 4 5 6 7 ` `0 1 2 3 4 5 6 7 `
... you see that when the 7 touches the 0, a new scale begins. Let's bring 7 and 0 closer.

![Octal](https://i.postimg.cc/RCXndBbq/octal.jpg)

Here: given the interruption, we can move to decimals.

### Decimal controllers
Everyone knows the decimal with a scale of 0-1-2-3-4-5-6-7-8-9. So I will not dare to talk to you about decimals. And don't try to ask me what decimals are. Here is an example of numerical allocation:
```
0 1 2 3 4 5 6 7 8 9
1 2 3 4 5 6 7 8 9 0
2 3 4 5 6 7 8 9 0 1
```

And now the numerical allocation method that has the most capacity:

### Here is the hexadecimal controller
Here is the hexadecimal: allocate multiple files in the internal controller. In fact it has a number scale. It goes like this: 0-1-2-3-4-5-6-7-8-9-A-B-C-D-E-F. But it has a different allocation method, because it has a complex list that goes on indefinitely. Here is a small part of the list:

Here are the pieces that compose it:
```
00000000   00000070   000000E0
00000010   00000080   000000F0
00000020   00000090
00000030   000000A0
00000040   000000B0
00000050   000000C0
00000060   000000D0
```
![Hexadecimal](https://i.postimg.cc/52yX1NqC/hexyy-2.png)

Screenshot of a hexadecimal editor:
![Hexadecimal Editor](https://i.postimg.cc/vHHKQp1n/hexeditor-hexyy.png)

Example:

![Hexadecimal Code](https://i.postimg.cc/br3pZkH9/hex-example.png)


## The length in the codes

In computing, a word is the natural unit of data used by a particular processor design. A word is a fixed-sized datum handled as a unit by the instruction set or the hardware of the processor. The number of bits or digits[a] in a word (the word size, word width, or word length) is an important characteristic of any specific processor design or computer architecture.

The size of a word is reflected in many aspects of a computer's structure and operation; the majority of the registers in a processor are usually word sized and the largest datum that can be transferred to and from the working memory in a single operation is a word in many (not all) architectures. The largest possible address size, used to designate a location in memory, is typically a hardware word (here, "hardware word" means the full-sized natural word of the processor, as opposed to any other definition used).

Documentation for older computers with fixed word size commonly states memory sizes in words rather than bytes or characters. The documentation sometimes uses metric prefixes correctly, sometimes with rounding, e.g., 65 kilowords (KW) meaning for 65536 words, and sometimes uses them incorrectly, with kilowords (KW) meaning 1024 words (210) and megawords (MW) meaning 1,048,576 words (220). With standardization on 8-bit bytes and byte addressability, stating memory sizes in bytes, kilobytes, and megabytes with powers of 1024 rather than 1000 has become the norm, although there is some use of the IEC binary prefixes.

Computer architecture bit widths: Bit
```
1 2 4 8 12 16 18 20 24 26 28 30 31 32 36 40 45 48 60 64 128 256 512
```

Application
```
8 16 32 64
```

Binary floating-point precision
```
16 (×½)2432 (×1)4064 (×2)80128 (×4)256 (×8)
```

Decimal floating-point precision
```
32 64 128
```

Name      | Capacity
--------- | ------
Byte      | x8
Word      | x16
Dword     | x32
Qword     | x64
KByte     | x1024
MByte     | x1000000
GByte     | x1000000000
TByte     | x1000000000000
PByte     | x1000000000000000
EByte     | x1000000000000000000
ZByte     | x1000000000000000000000
YByte     | x1000000000000000000000000


Here: we concluded with the numbers.
