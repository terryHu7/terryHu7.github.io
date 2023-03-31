#!/bin/bash
count=2
for file in ./DP/*;do
	./markdown2html -inPath=${file} -outPath="dp${count}.md"
	((count++))
done

