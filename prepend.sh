#!/bin/bash

tmp=`mktemp`
cat LICENSE > $tmp
cat "$1" >> $tmp
mv $tmp "$1"

