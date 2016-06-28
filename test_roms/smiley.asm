;; This is a test program that draws a smiley face in the middle of the screen
;; and a pixel at each corner of the screen. Useful for testing that your
;; renderer and coordinates are correct.

             CLR
smileypos    LOAD    r0,$1B
             LOAD    r1,$D
origin       LOAD    r2,$0
width        LOAD    r3,$38
height       LOAD    r4,$1F
drawsmiley   LOADI   smiley
             DRAW    r0,r1,$5
drawtl       LOADI   leftpixel
             DRAW    r2,r2,$1
drawtr       LOADI   rightpixel
             DRAW    r3,r2,$1
drawbl       LOADI   leftpixel
             DRAW    r2,r4,$1
drawbr       LOADI   rightpixel
             DRAW    r3,r4,$1
loop         JUMP    loop
smiley       FCB     $66
             FCB     $66
             FCB     $18
             FCB     $C3
             FCB     $3C
leftpixel    FCB     $80
rightpixel   FCB     $1
