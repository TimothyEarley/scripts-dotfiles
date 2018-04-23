#! /bin/bash

amixer sget Master | grep -Po '\[.*%\]' | head -n 1 | grep -Po '\d+%'