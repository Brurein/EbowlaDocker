#!/bin/bash

echo $1

/usr/bin/python2.7 /3bowla/Ebowla/ebowla.py /3bowla/host/$@ /3bowla/host/genetic.config

/3bowla/Ebowla/build_x64_go.sh /3bowla/output/go_symmetric_$@.go enc_$@

mv /3bowla/output/enc_$@ /3bowla/host/enc_$@