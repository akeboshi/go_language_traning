#!/bin/bash

for ((i=0;i<10;i++)) do  
  echo "=============================="
  echo GOMAXPROCS=$i
  GOMAXPROCS=$i go run madelbrot.go > /dev/null
done

echo "=============end=============="

# GOMAXPROCS=4の時が一番早い。
# 結果:
# cpu: Core™ i5-4250U (2core 4thread)
#
# ==============================
# GOMAXPROCS=0
# normal time: 0.456060
# parallel time: 0.482701
# ==============================
# GOMAXPROCS=1
# normal time: 0.514039
# parallel time: 0.493674
# ==============================
# GOMAXPROCS=2
# normal time: 0.466160
# parallel time: 0.239063
# ==============================
# GOMAXPROCS=3
# normal time: 0.481297
# parallel time: 0.227681
# ==============================
# GOMAXPROCS=4
# normal time: 0.440086
# parallel time: 0.296678
# ==============================
# GOMAXPROCS=5
# normal time: 0.485423
# parallel time: 0.209527
# ==============================
# GOMAXPROCS=6
# normal time: 0.473535
# parallel time: 0.209039
# ==============================
# GOMAXPROCS=7
# normal time: 0.459513
# parallel time: 0.210833
# ==============================
# GOMAXPROCS=8
# normal time: 0.447171
# parallel time: 0.214628
# ==============================
# GOMAXPROCS=9
# normal time: 0.441440
# parallel time: 0.212162
# =============end==============
