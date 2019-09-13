INSHOME=${HOME}/.config/persephone

all: persephone

persephone: clean
	NOTAR=true ./build

install:
	sudo install -D -m644 persephone /usr/bin/persephone
	sudo chmod +x /usr/bin/persephone
	mkdir -pv ${INSHOME}/temp
	mkdir -pv ${INSHOME}/static/fonts
	mkdir -pv ${INSHOME}/static/images
	if ! [ -f ${INSHOME}/config.json ]; then install -D -m644 config.example.json ${INSHOME}/config.json; fi
	install -D -m644 static/fonts/NotoSans-Bold.ttf ${INSHOME}/static/fonts/NotoSans-Bold.ttf
	install -D -m644 static/fonts/NotoSans-Regular.ttf ${INSHOME}/static/fonts/NotoSans-Regular.ttf
	install -D -m644 static/images/background.png ${INSHOME}/static/images/background.png
	install -D -m644 artists.json ${INSHOME}/artists.json

uninstall:
	sudo rm -rf /usr/bin/persephone
	rm -rf ${INSHOME}/temp
	rm -rf ${INSHOME}/static

clean:
	if [ -f persephone ]; then rm persephone; fi

.PHONY: all persephone install uninstall clean
