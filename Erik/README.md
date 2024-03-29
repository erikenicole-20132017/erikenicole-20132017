[Go to back](https://github.com/erikenicole-20132017/Erik-Docs/blob/main/Documents)

# Welcome in the Erik Operating System
Here there will be only the source folder of the operating system and some information about it.

## Numerics
Our system uses 10 symbols or digits, written like this: 0 1 2 3 4 5 6 7 8 9. 0 is the neutral number, that is, it does not change the number. For example: 3 + 0 = 3. In addition, there is a rule based on the sequence of integers, which is the sum of the second factor and the sum of the previous number. For example: 1+1=2. 1+2=3. 2+3=5. 3+5=8. 5+8=13. 8+13=21... And so it continues. 
### Binary
#### Binary Utility File

Now look, there will be information about the hard drive where the operating system will be set. Here are 0 bytes:
```
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 
[contain 0 bytes]
```
Of course it is easy for 0 bytes to all contain zeros in the binary.

The file will be created as a `utility.bin` that will show the direct reports (the direct reports modify the attributes for once, the indirect ones instead keep them by sorting it by date)

Example for indirect file used from `Jan 24 2001 - Jan 25 2001`:
```
2001/01/24
00 00 00 01 00 00 00 00 00 00 00 00 00 01 10 
00 10 01 10 00 10 00 10 10 00 10 10 10 00 01 
00 00 00 10 01 11 11 01 10 10 00 01 01 00 10 
2001/01/25
00 00 00 01 00 00 00 00 00 00 00 00 00 01 10 
00 10 01 10 00 10 00 10 10 00 10 10 10 00 01 
00 00 00 10 01 10 00 10 10 00 10 01 00 00 10
```

Example for direct file used from `24 Jan 2001 - 25 Jan 2001`:

Modified by Jan 24 2001: 
`utility.bin`
```
00 00 00 01 00 00 00 00 00 00 00 00 00 01 10 
00 10 01 10 00 10 00 10 10 00 10 10 10 00 01 
00 00 00 10 01 11 11 01 10 10 00 01 01 00 10 
```

Modified by Jan 25 2001: 
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
We talked about the work on the hard drive, there are binary numbers. Called "controllers", but strange, the previous example was too small for a large storage space. The controllers, when they are too small, the disk will be reported in partition as invalid. depending on the binaries, there are also controllers that support multiple capacities: octal. Everything comes from the binary, I say well? In fact, the track has only 2 controllers, but the octal has as many as 8. 

#### Does location matter?

It is often said that where you are for example in a place at the cinema you should not give a damn but when it comes to binary numbers then you can also change your value if you do not even know where you are. To better understand...



### Octal controllers
In the previous chapter, we talked about binary numaries, which have controllers. The binary has 2, the octal has 8. We're still focused on the binary 🛤️ for a while.
The binary has 2 controllers: 0-1. Instead the octal has a scale from 0 to 7: 0-1-2-3-4-5-6-7. The octal only supports these numbers, but not as many as 8 and 9. Otherwise it would be the decimal scale, which will reign over the next chapter. Now let's move on to the octal. An octal sequence would look like this:
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

Table of the hex codes:
```
0	=	0dec	=	0oct		0	0	0	0	
1	=	1dec	=	1oct		0	0	0	1	
2	=	2dec	=	2oct		0	0	1	0	
3	=	3dec	=	3oct		0	0	1	1	
4	=	4dec	=	4oct		0	1	0	0	
5	=	5dec	=	5oct		0	1	0	1	
6	=	6dec	=	6oct		0	1	1	0	
7	=	7dec	=	7oct		0	1	1	1	
8	=	8dec	=	10oct		1	0	0	0	
9	=	9dec	=	11oct		1	0	0	1	
A	=	10dec	=	12oct		1	0	1	0	
B	=	11dec	=	13oct		1	0	1	1	
C	=	12dec	=	14oct		1	1	0	0	
D	=	13dec	=	15oct		1	1	0	1	
E	=	14dec	=	16oct		1	1	1	0	
F	=	15dec	=	17oct		1	1	1	1	
```

![hexadecimal_number_value](https://user-images.githubusercontent.com/108028311/209473766-e888e07e-43d7-4460-aa47-0b83b159f07a.png)


### omg

It's a
```
🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥
🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥
🟧⬜️⬜️⬜️⬜️🟧⬜️⬜️⬜️⬜️🟧🟧🟧⬜️⬜️🟧🟧🟧⬜️⬜️⬜️⬜️🟧⬜️⬜️⬜️⬜️⬜️🟧
🟧⬜️🟧🟧⬜️🟧⬜️🟧🟧🟧🟧🟧⬜️🟧🟧⬜️🟧🟧⬜️🟧🟧🟧🟧⬜️🟧🟧🟧🟧🟧
🟨⬜️🟨🟨⬜️🟨⬜️🟨🟨🟨🟨⬜️🟨🟨🟨🟨⬜️🟨⬜️🟨🟨🟨🟨⬜️🟨🟨🟨🟨🟨
🟨⬜️⬜️⬜️⬜️🟨⬜️⬜️⬜️⬜️🟨⬜️⬜️⬜️⬜️⬜️⬜️🟨⬜️🟨🟨🟨🟨⬜️⬜️⬜️⬜️⬜️🟨
🟩⬜️🟩🟩🟩🟩⬜️🟩🟩🟩🟩⬜️🟩🟩🟩🟩⬜️🟩⬜️🟩🟩🟩🟩⬜️🟩🟩🟩🟩🟩
🟩⬜️🟩🟩🟩🟩⬜️🟩🟩🟩🟩⬜️🟩🟩🟩🟩⬜️🟩⬜️🟩🟩🟩🟩⬜️🟩🟩🟩🟩🟩
🟦⬜️🟦🟦🟦🟦⬜⬜️⬜️⬜️🟦⬜️🟦🟦🟦🟦⬜️🟦⬜️⬜️⬜️⬜️🟦⬜️⬜️⬜️⬜️⬜️🟦
🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪
🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪🟪
     ⬛️⬛️🟫🟫          ⬛️           ⬛️🟫   ⬛️🟫                     
     ⬛️🟫 ⬛️🟫       ⬛️🟫⬛️🟫       ⬛️  ⬛️🟫                       
     ⬛️🟫  ⬛️🟫      ⬛️⬛️⬛️🟫         ⬛️🟫                         
     ⬛️🟫 ⬛️🟫      ⬛️🟫  ⬛️🟫        ⬛️🟫                         
     ⬛️🟫⬛️🟫       ⬛️🟫  ⬛️🟫        ⬛️🟫                         
     ⬛️⬛️🟫🟫       ⬛️🟫  ⬛️🟫        ⬛️🟫                         
```

## Machine language
Here is the following machine language that the system must perform:
```
   6      5     5     5     5      6     bits
[  op  |  rs |  rt |  rd |shamt| funct]  R-type
[  op  |  rs |  rt | address/immediate]  I-type
[  op  |        target address        ]  J-type
[  op  |  rs |  rt |  rd |shamt| funct]
[  op  |  rs |  rt | address/immediate]
[  op  |        target address        ]
```
```
action 1 ("color")
action 2 ("get machine")
action 3 ("cmd")
   holderplace="write" set blockquote li4 char/write
   run program debug
   error print
   program open
```
Here is the decimal binary conversion:
```
000000 00001 00010 00110 00000 100000   binary
100011 00011 01000 00000 00001 000100   binary
100011 00011 01000 00000 00001 000100   binary
000010 00000 00000 00000 10000 000000   binary
   0     1     2     6     0     32     decimal
  35     3     8           68           decimal
   2                 1024               decimal
```

Let's still stay on the machine code, and try to take a look at the [binary numbers]([link](https://github.com/erikenicole-20132017/erikenicole-20132017/blob/main/Erik/README.md#binary-utility-file)) to clarify ourselves.

Because now we are going to focus on the binary number, so ... fasten your seatbelts because we are about to walk the track! Safe travel!

Let's clarify ourselves better by seeing a binary code.
```
0010 0101 0010 0101
0100 1010 0000 1010
0011 1100 1111 0000
0011 1100 0101 1101
```
Well, that's just the least of professional things I can do like that. In fact, did you know that you can also draw? Stay tuned to find out in the next chapter of this WONDERFUL guide!

## ASCII Art
### Binary ASCII Art
Are you thinking that drawing is very simple, really? Ah, let's see! First of all we have to understand what the binary is and I dare not tell you because I have already said it thousands of times for half a century! Take the spare clothes because you will have to sweat drawing with a keyboard.
> **_NOTE:_** The basic keyboard doesn't have any special symbols, but try the Unicode keyboard. Remember that, many symbols are not colored. The source code does not support formatting, that is, you will not be able to format in bold, italic, underlined, strikethrough, font, size, color, pictures, and maybe even much more. Image and formatting examples have a Markdown format, but in a simple console it will not support formatting. Image and formatting examples have a Markdown format, but in a simple console it will not support formatting.
![binaryart](https://user-images.githubusercontent.com/108028311/187042550-40b62fb6-5541-4e86-84d3-60cbb532f513.JPG)
### Symbols ASCII Art
The most widespread. Binary art is not widely used, because without formatting, it would be CHAOS. But we have one last chance: the ASCII symbol. Let's see an example of this refined drawing method.
```
 _______   ________  ___  ___  __       
|\  ___ \ |\   __  \|\  \|\  \|\  \     
\ \   __/|\ \  \|\  \ \  \ \  \/  /|_   
 \ \  \_|/_\ \   _  _\ \  \ \   ___  \  
  \ \  \_|\ \ \  \\  \\ \  \ \  \\ \  \      ::   :::::         ::::::::  ::     ::  ::::::
   \ \_______\ \__\\ _\\ \__\ \__\\ \__\     ::   ::               ::     ::     ::  ::
    \|_______|\|__|\|__|\|__|\|__| \|__|     ::   :::::            ::     :::::::::  ::::::
                                             ::      ::            ::     ::     ::  ::
                                             ::   :::::            ::     ::     ::  ::::::
                                             
                                             ::::::    ::::::::     ::::::::   ::::::::    :::
                                             ::   ::   ::           :::           ::       :::
                                             ::::::    ::::::::     ::::::::      ::       :::
                                             ::   ::   ::                :::      ::
                                             ::::::    ::::::::     ::::::::      ::       :::
                                                  
```

"Erik is the best" example 

> IMPORTANT: Serif or sans-serif are not used because writing in ASCII would become unreadable for the font type. Instead, in order to become composed, that is, that each symbol agrees with the others and that they do the same thing, it is necessary to use a monospace. A monospaced serif or a monospaced sans-serif is also fine!

## Style
```
Style {
   webkit-opacity:none;
   position-absolute:none;
   font-family: Segoe UI;
   nofont: Segoe UI; sans-serif;
   set code
}
```
## Switchers
Set on
<!--Start switchers-->
<a href="#switcherdata='switchercontrolset=0'">
    <svg> 
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="148" height="20" role="img"><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="148" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="41" height="20" fill="#555"/><rect x="41" width="107" height="20" fill="#007ec6"/><rect width="148" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Segoe UI" text-rendering="geometricPrecision" font-size="110"><text aria-hidden="true" x="215" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="310">0</text><text x="215" y="140" transform="scale(.1)" fill="#fff" textLength="310">0</text><text aria-hidden="true" x="935" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="970">Switcher</text><text x="935" y="140" transform="scale(.1)" fill="#fff" textLength="970">Switcher</text></g></svg>
</svg>
  </a>
  <br>
  <a href="#switcherdata='switchercontrolset=0'">
    <svg> 
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="148" height="20" role="img"><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="148" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="41" height="20" fill="#555"/><rect x="41" width="107" height="20" fill="#007ec6"/><rect width="148" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Segoe UI" text-rendering="geometricPrecision" font-size="110"><text aria-hidden="true" x="215" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="310">1</text><text x="215" y="140" transform="scale(.1)" fill="#fff" textLength="310">1</text><text aria-hidden="true" x="935" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="970">Switcher</text><text x="935" y="140" transform="scale(.1)" fill="#fff" textLength="970">Switcher</text></g></svg>
</svg>
  </a>
<!--Ending switchers>

## The Archive Universe
### 1. Informations
In computing, a word is the natural unit of data used by a particular processor design. A word is a fixed-sized datum handled as a unit by the instruction set or the hardware of the processor. The number of bits or digits in a word (the word size, word width, or word length) is an important characteristic of any specific processor design or computer architecture.

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

Each unit of allocation, in ascending order, is multiplied by a thousand times. For example, 1 GB repeated 1000 times becomes 1 Terabyte (1000 GB)

### 2: Comparison

Scroll Until the End

