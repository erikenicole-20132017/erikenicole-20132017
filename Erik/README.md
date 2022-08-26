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
In the previous chapter, we talked about binary numaries, which have controllers. The track has 2, the octal has 8. We're still focused on the binary 🛤️ for a while.
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
