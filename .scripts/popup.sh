#! /bin/bash

h=50
w=200

# dim=$( xdpyinfo | grep -Po '(?<=dimensions:    )\d+x\d+' )
dim=$( xrandr | grep -Po '\d+x\d+'  | head -n 1 )
sw=${dim%x*}
sh=${dim#*x}

middleY=$(( $sh / 2 - ( $h / 2 ) ))
middleX=$(( $sw / 2 - ( $w / 2 ) ))

echo $1 | dzen2 -p 1 -y $middleY -x $middleX -h $h -w $w 


