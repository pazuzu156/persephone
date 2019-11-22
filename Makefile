INSHOME=${HOME}/.config/persephone

ifndef PREFIX
PREFIX=/usr/bin
endif

all: persephone

persephone: clean
	NOTAR=true ./build

install:
ifdef GOPATH
	go install
else
	sudo install -D -m644 persephone $(PREFIX)/persephone
	sudo chmod +x $(PREFIX)/persephone
endif

	mkdir -pv ${INSHOME}/temp
	mkdir -pv ${INSHOME}/static/fonts
	mkdir -pv ${INSHOME}/static/images
	if ! [ -f ${INSHOME}/config.json ]; then install -D -m644 config.example.json ${INSHOME}/config.json; fi
	install -D -m644 static/fonts/NotoSans-Bold.ttf ${INSHOME}/static/fonts/NotoSans-Bold.ttf
	install -D -m644 static/fonts/NotoSans-Regular.ttf ${INSHOME}/static/fonts/NotoSans-Regular.ttf
	install -D -m644 static/images/background.png ${INSHOME}/static/images/background.png
	install -D -m644 static/images/bm.png ${INSHOME}/static/images/bm.png
	install -D -m644 artists.json ${INSHOME}/artists.json

uninstall:
ifdef GOPATH
	rm -rf $(GOPATH)/bin/persephone
else
	sudo rm -rf $(PREFIX)/persephone
endif

	rm -rf ${INSHOME}/temp
	rm -rf ${INSHOME}/static

clean:
	if [ -f persephone ]; then rm persephone; fi

.PHONY: all persephone install uninstall clean
