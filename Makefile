clean:
	rm *~

all:
	tinygo flash -target=pico -stack-size=8kb -monitor .
