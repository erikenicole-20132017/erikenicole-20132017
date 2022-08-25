# Welcome in the Erik Operating System
Here there will be only the source folder of the operating system and some information about it.

## Binary utility file

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
   dir.save
   when i receive delete.
   permission denied
   reccomend set(default)
   RECOGNIZE ACTIONS
   add for locale (unit.receive)
   EDIT FILE TO ALL USER
   username filter
   date filter
   subdate filter
   data filter
   action filter
   memory filter
   removable add only actions [private]
}
```
