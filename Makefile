mydir:=mydir

mydir2:=`mkdir mydir2`

mkdir:
	echo "123" 
	@if [ ! -d $(mydir) ]; then \
		echo "no" ;\
		$$$(shell mkdir $(mydir));\
	fi
